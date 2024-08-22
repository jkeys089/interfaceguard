package interfaceguard

import (
	"flag"
	"go/ast"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/analysis"
)

func NewAnalyzer(defaultNilCheckDisabled, defaultInterfaceCheckDisabled bool) *analysis.Analyzer {
	r := &runner{}
	flags := flag.NewFlagSet("interfaceguard", flag.ExitOnError)

	flags.BoolVar(
		&r.interfaceCheckDisabled, "i", defaultInterfaceCheckDisabled, "disable interface-to-interface comparison checks")

	flags.BoolVar(
		&r.nilCheckDisabled, "n", defaultNilCheckDisabled, "disable interface nil comparison checks")

	return &analysis.Analyzer{
		Name:  "interfaceguard",
		Doc:   "interfaceguard checks subtle interface-related issues, including improper comparisons and type mismatches.",
		Flags: *flags,
		Run:   r.run,
	}
}

type runner struct {
	nilCheckDisabled       bool
	interfaceCheckDisabled bool
}

func (r *runner) run(pass *analysis.Pass) (interface{}, error) {
	if r.nilCheckDisabled && r.interfaceCheckDisabled {
		return nil, nil
	}

	for _, file := range pass.Files {
		ast.Inspect(file, func(node ast.Node) bool {
			switch n := node.(type) {
			case *ast.IfStmt:
				r.checkCondition(n.Cond, pass)
			case *ast.SwitchStmt:
				if n.Tag != nil {
					r.checkCondition(n.Tag, pass)
				}
				for _, stmt := range n.Body.List {
					if caseClause, ok := stmt.(*ast.CaseClause); ok {
						for _, expr := range caseClause.List {
							r.checkCondition(expr, pass)
						}
					}
				}
			case *ast.TypeSwitchStmt:
				r.checkTypeSwitch(n, pass)
			case *ast.AssignStmt:
				for _, rhs := range n.Rhs {
					r.checkAssignment(rhs, pass)
				}
			case *ast.CompositeLit:
				for _, elt := range n.Elts {
					r.checkAssignment(elt, pass)
				}
			case *ast.KeyValueExpr:
				r.checkAssignment(n.Value, pass)
			case *ast.ForStmt:
				if n.Cond != nil {
					r.checkCondition(n.Cond, pass)
				}
			case *ast.SendStmt:
				r.checkAssignment(n.Value, pass)
			case *ast.ReturnStmt:
				for _, result := range n.Results {
					r.checkAssignment(result, pass)
				}
			}
			return true
		})
	}
	return nil, nil
}

func (r *runner) checkCondition(cond ast.Expr, pass *analysis.Pass) {
	switch expr := cond.(type) {
	case *ast.BinaryExpr:
		if expr.Op == token.EQL || expr.Op == token.NEQ {
			r.checkBinaryExpr(expr, pass)
		} else {
			r.checkCondition(expr.X, pass)
			r.checkCondition(expr.Y, pass)
		}
	case *ast.ParenExpr:
		r.checkCondition(expr.X, pass)
	case *ast.CallExpr:
		for _, arg := range expr.Args {
			r.checkCondition(arg, pass)
		}
	case *ast.UnaryExpr:
		r.checkCondition(expr.X, pass)
	}
}

func (r *runner) checkBinaryExpr(expr *ast.BinaryExpr, pass *analysis.Pass) {
	xType := pass.TypesInfo.TypeOf(expr.X)
	yType := pass.TypesInfo.TypeOf(expr.Y)

	if !r.nilCheckDisabled && (isNil(expr.X) || isNil(expr.Y)) {
		nonNilExpr := expr.Y
		if isNil(expr.Y) {
			nonNilExpr = expr.X
		}
		if isInterfaceType(pass.TypesInfo.TypeOf(nonNilExpr)) && !isErrorType(pass.TypesInfo.TypeOf(nonNilExpr)) {
			pass.Reportf(expr.Pos(), "comparing interface to nil")
		}
	} else if !r.interfaceCheckDisabled && isInterfaceType(xType) && isInterfaceType(yType) {
		if !isErrorType(xType) && !isErrorType(yType) {
			pass.Reportf(expr.Pos(), "comparing two interfaces")
		}
	}
}

func (r *runner) checkTypeSwitch(stmt *ast.TypeSwitchStmt, pass *analysis.Pass) {
	if r.nilCheckDisabled {
		return
	}

	for _, bodyStmt := range stmt.Body.List {
		if caseClause, ok := bodyStmt.(*ast.CaseClause); ok {
			for _, expr := range caseClause.List {
				var ident *ast.Ident
				if ident, ok = expr.(*ast.Ident); ok && ident.Name == "nil" {
					pass.Reportf(ident.Pos(), "type switch case checking interface against nil")
				}
			}
		}
	}
}

func (r *runner) checkAssignment(rhs ast.Expr, pass *analysis.Pass) {
	if binaryExpr, ok := rhs.(*ast.BinaryExpr); ok {
		if (binaryExpr.Op == token.EQL || binaryExpr.Op == token.NEQ) && (isNil(binaryExpr.X) || isNil(binaryExpr.Y)) {
			if r.nilCheckDisabled {
				return
			}

			var exprToCheck ast.Expr
			if isNil(binaryExpr.X) {
				exprToCheck = binaryExpr.Y
			} else {
				exprToCheck = binaryExpr.X
			}

			exprType := pass.TypesInfo.TypeOf(exprToCheck)
			if exprType != nil && isInterfaceType(exprType) && !isErrorType(exprType) {
				pass.Reportf(rhs.Pos(), "comparing interface to nil")
			}
		} else if isInterfaceType(pass.TypesInfo.TypeOf(binaryExpr.X)) && isInterfaceType(pass.TypesInfo.TypeOf(binaryExpr.Y)) {
			if r.interfaceCheckDisabled {
				return
			}

			if !isErrorType(pass.TypesInfo.TypeOf(binaryExpr.X)) && !isErrorType(pass.TypesInfo.TypeOf(binaryExpr.Y)) {
				pass.Reportf(rhs.Pos(), "comparing two interfaces")
			}
		} else {
			r.checkCondition(binaryExpr.X, pass)
			r.checkCondition(binaryExpr.Y, pass)
		}
	} else if _, ok = rhs.(*ast.BasicLit); !ok {
		r.checkCondition(rhs, pass)
	}
}

func isInterfaceType(typ types.Type) bool {
	_, ok := typ.Underlying().(*types.Interface)
	return ok
}

func isErrorType(typ types.Type) bool {
	return typ.String() == "error"
}

func isNil(expr ast.Expr) bool {
	ident, ok := expr.(*ast.Ident)
	return ok && ident.Name == "nil"
}
