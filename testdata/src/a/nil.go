package a

import "fmt"

func nilBasicComparison() {
	if greeter == nil { // want "comparing interface to nil"
		fmt.Println("greeter is nil")
	}

	if greeter != nil { // want "comparing interface to nil"
		fmt.Println("greeter is not nil")
	}
}

func nilComplexComparison() {
	if now < 1 || greeter == nil { // want "comparing interface to nil"
		fmt.Println("greeter is nil")
	}

	if greeter != nil && now > 1 { // want "comparing interface to nil"
		fmt.Println("greeter is not nil")
	}
}

func nilBasicAssignment() {
	isNil := greeter == nil // want "comparing interface to nil"
	if isNil {
		fmt.Println("greeter is nil")
	}

	isNotNil := greeter != nil // want "comparing interface to nil"
	if isNotNil {
		fmt.Println("greeter is not nil")
	}
}

func nilComplexAssignment() {
	isNil := now < 1 || greeter == nil // want "comparing interface to nil"
	if isNil {
		fmt.Println("greeter is nil")
	}

	isNotNil := greeter != nil && now > 1 // want "comparing interface to nil"
	if isNotNil {
		fmt.Println("greeter is not nil")
	}
}

func nilBasicInlineAssignment() {
	if isNil := greeter == nil; isNil { // want "comparing interface to nil"
		fmt.Println("greeter is nil")
	}

	if isNotNil := greeter != nil; isNotNil { // want "comparing interface to nil"
		fmt.Println("greeter is not nil")
	}
}

func nilComplexInlineAssignment() {
	if isNil := now < 1 || greeter == nil; isNil { // want "comparing interface to nil"
		fmt.Println("greeter is nil")
	}

	if isNotNil := greeter != nil && now > 1; isNotNil { // want "comparing interface to nil"
		fmt.Println("greeter is not nil")
	}
}

func nilBasicMultiAssignment() {
	x, isNil := now, greeter == nil // want "comparing interface to nil"
	if isNil && x > 1 {
		fmt.Println("greeter is nil")
	}

	y, isNotNil := now, greeter != nil // want "comparing interface to nil"
	if isNotNil && y > 1 {
		fmt.Println("greeter is not nil")
	}
}

func nilComplexMultiAssignment() {
	x, isNil := now, now < 1 || greeter == nil // want "comparing interface to nil"
	if isNil && x > 1 {
		fmt.Println("greeter is nil")
	}

	y, isNotNil := now, greeter != nil && now > 1 // want "comparing interface to nil"
	if isNotNil && y > 1 {
		fmt.Println("greeter is not nil")
	}
}

func nilBasicInlineMultiAssignment() {
	if x, isNil := now, greeter == nil; isNil && x > 1 { // want "comparing interface to nil"
		fmt.Println("greeter is nil")
	}

	if y, isNotNil := now, greeter != nil; isNotNil && y > 1 { // want "comparing interface to nil"
		fmt.Println("greeter is not nil")
	}
}

func nilComplexInlineMultiAssignment() {
	if x, isNil := now, now < 1 || greeter == nil; isNil && x > 1 { // want "comparing interface to nil"
		fmt.Println("greeter is nil")
	}

	if y, isNotNil := now, greeter != nil && now > 1; isNotNil && y > 1 { // want "comparing interface to nil"
		fmt.Println("greeter is not nil")
	}
}

func nilSimpleForLoop() {
	for greeter == nil { // want "comparing interface to nil"
		fmt.Println("greeter is nil")
		greeter = costcoGreeter
	}

	for greeter != nil { // want "comparing interface to nil"
		fmt.Println("greeter is not nil")
		greeter = nil
	}
}

func nilComplexForLoop() {
	for now < 1 || greeter == nil { // want "comparing interface to nil"
		fmt.Println("greeter is nil")
		greeter = costcoGreeter
	}

	for greeter != nil && now > 1 { // want "comparing interface to nil"
		fmt.Println("greeter is not nil")
		greeter = nil
	}
}

func nilBasicInlineForLoop() {
	for isNil := greeter == nil; isNil; { // want "comparing interface to nil"
		fmt.Println("greeter is nil")
		isNil = false
	}

	for isNotNil := greeter != nil; isNotNil; { // want "comparing interface to nil"
		fmt.Println("greeter is not nil")
		isNotNil = false
	}
}

func nilComplexInlineForLoop() {
	for isNil := now < 1 || greeter == nil; isNil; { // want "comparing interface to nil"
		fmt.Println("greeter is nil")
		isNil = false
	}

	for isNotNil := greeter != nil && now < 1; isNotNil; { // want "comparing interface to nil"
		fmt.Println("greeter is not nil")
		isNotNil = false
	}
}

func nilBasicForRangeLoop() {
	for _, nextGreeter := range greeters {
		if nextGreeter == nil { // want "comparing interface to nil"
			fmt.Println("greeter is nil")
		}
	}

	for _, nextGreeter := range greeters {
		if nextGreeter != nil { // want "comparing interface to nil"
			fmt.Println("greeter is not nil")
		}
	}
}

func nilComplexForRangeLoop() {
	for _, nextGreeter := range greeters {
		if now < 1 || nextGreeter == nil { // want "comparing interface to nil"
			fmt.Println("greeter is nil")
		}
	}

	for _, nextGreeter := range greeters {
		if nextGreeter != nil && now > 1 { // want "comparing interface to nil"
			fmt.Println("greeter is not nil")
		}
	}
}

func nilBasicSwitch() {
	switch greeter == nil { // want "comparing interface to nil"
	case true:
		fmt.Println("greeter is nil")
	}

	switch greeter != nil { // want "comparing interface to nil"
	case true:
		fmt.Println("greeter is not nil")
	}
}

func nilComplexSwitch() {
	switch now < 1 || greeter == nil { // want "comparing interface to nil"
	case true:
		fmt.Println("greeter is nil")
	}

	switch greeter != nil && now > 1 { // want "comparing interface to nil"
	case true:
		fmt.Println("greeter is not nil")
	}
}

func nilBasicSwitchCase() {
	switch {
	case greeter == nil: // want "comparing interface to nil"
		fmt.Println("greeter is nil")
	}

	switch {
	case greeter != nil: // want "comparing interface to nil"
		fmt.Println("greeter not nil")
	}
}

func nilComplexSwitchCase() {
	switch {
	case now < 1 || greeter == nil: // want "comparing interface to nil"
		fmt.Println("greeter is nil")
	}

	switch {
	case greeter != nil && now > 1: // want "comparing interface to nil"
		fmt.Println("greeter not nil")
	}
}

func nilBasicSliceAssignment() {
	isNil := []bool{
		greeter == nil, // want "comparing interface to nil"
	}

	isNil[0] = greeter == nil // want "comparing interface to nil"
	if isNil[0] {
		fmt.Println("greeter is nil")
	}

	if isNil[0] = greeter == nil; isNil[0] { // want "comparing interface to nil"
		fmt.Println("greeter is nil")
	}

	isNotNil := []bool{
		greeter != nil, // want "comparing interface to nil"
	}

	isNotNil[0] = greeter != nil // want "comparing interface to nil"
	if isNotNil[0] {
		fmt.Println("greeter is not nil")
	}

	if isNotNil[0] = greeter != nil; isNotNil[0] { // want "comparing interface to nil"
		fmt.Println("greeter is not nil")
	}
}

func nilComplexSliceAssignment() {
	isNil := []bool{
		now < 1 || greeter == nil, // want "comparing interface to nil"
	}

	isNil[0] = now < 1 || greeter == nil // want "comparing interface to nil"
	if isNil[0] {
		fmt.Println("greeter is nil")
	}

	if isNil[0] = now < 1 || greeter == nil; isNil[0] { // want "comparing interface to nil"
		fmt.Println("greeter is nil")
	}

	isNotNil := []bool{
		greeter != nil && now > 1, // want "comparing interface to nil"
	}

	isNotNil[0] = greeter != nil && now > 1 // want "comparing interface to nil"
	if isNotNil[0] {
		fmt.Println("greeter is not nil")
	}

	if isNotNil[0] = greeter != nil && now > 1; isNotNil[0] { // want "comparing interface to nil"
		fmt.Println("greeter is not nil")
	}
}

func nilBasicMapAssignment() {
	isNil := map[string]bool{
		"a": greeter == nil, // want "comparing interface to nil"
	}

	isNil["a"] = greeter == nil // want "comparing interface to nil"
	if isNil["a"] {
		fmt.Println("greeter is nil")
	}

	if isNil["a"] = greeter == nil; isNil["a"] { // want "comparing interface to nil"
		fmt.Println("greeter is nil")
	}

	isNotNil := map[string]bool{
		"a": greeter != nil, // want "comparing interface to nil"
	}

	isNotNil["a"] = greeter != nil // want "comparing interface to nil"
	if isNotNil["a"] {
		fmt.Println("greeter is not nil")
	}

	if isNotNil["a"] = greeter != nil; isNotNil["a"] { // want "comparing interface to nil"
		fmt.Println("greeter is not nil")
	}
}

func nilComplexMapAssignment() {
	isNil := map[string]bool{
		"a": now < 1 || greeter == nil, // want "comparing interface to nil"
	}

	isNil["a"] = now < 1 || greeter == nil // want "comparing interface to nil"
	if isNil["a"] {
		fmt.Println("greeter is nil")
	}

	if isNil["a"] = now < 1 || greeter == nil; isNil["a"] { // want "comparing interface to nil"
		fmt.Println("greeter is nil")
	}

	isNotNil := map[string]bool{
		"a": greeter != nil && now > 1, // want "comparing interface to nil"
	}

	isNotNil["a"] = greeter != nil && now > 1 // want "comparing interface to nil"
	if isNotNil["a"] {
		fmt.Println("greeter is not nil")
	}

	if isNotNil["a"] = greeter != nil && now > 1; isNotNil["a"] { // want "comparing interface to nil"
		fmt.Println("greeter is not nil")
	}
}

func nilBasicStructAssignment() {
	isNil := struct {
		b bool
	}{
		b: greeter == nil, // want "comparing interface to nil"
	}

	isNil.b = greeter == nil // want "comparing interface to nil"
	if isNil.b {
		fmt.Println("greeter is nil")
	}

	if isNil.b = greeter == nil; isNil.b { // want "comparing interface to nil"
		fmt.Println("greeter is nil")
	}

	isNotNil := struct {
		b bool
	}{
		b: greeter != nil, // want "comparing interface to nil"
	}

	isNotNil.b = greeter != nil // want "comparing interface to nil"
	if isNotNil.b {
		fmt.Println("greeter is not nil")
	}

	if isNotNil.b = greeter != nil; isNotNil.b { // want "comparing interface to nil"
		fmt.Println("greeter is not nil")
	}
}

func nilComplexStructAssignment() {
	isNil := struct {
		b bool
	}{
		b: now < 1 || greeter == nil, // want "comparing interface to nil"
	}

	isNil.b = now < 1 || greeter == nil // want "comparing interface to nil"
	if isNil.b {
		fmt.Println("greeter is nil")
	}

	if isNil.b = now < 1 || greeter == nil; isNil.b { // want "comparing interface to nil"
		fmt.Println("greeter is nil")
	}

	isNotNil := struct {
		b bool
	}{
		b: greeter != nil && now > 1, // want "comparing interface to nil"
	}

	isNotNil.b = greeter != nil && now > 1 // want "comparing interface to nil"
	if isNotNil.b {
		fmt.Println("greeter is not nil")
	}

	if isNotNil.b = greeter != nil && now > 1; isNotNil.b { // want "comparing interface to nil"
		fmt.Println("greeter is not nil")
	}
}

func nilBasicChanSend() {
	boolChan <- greeter == nil // want "comparing interface to nil"
	if <-boolChan {
		fmt.Println("greeter is nil")
	}

	if boolChan <- greeter == nil; <-boolChan { // want "comparing interface to nil"
		fmt.Println("greeter is nil")
	}

	boolChan <- greeter != nil // want "comparing interface to nil"
	if <-boolChan {
		fmt.Println("greeter is not nil")
	}

	if boolChan <- greeter != nil; <-boolChan { // want "comparing interface to nil"
		fmt.Println("greeter is not nil")
	}
}

func nilComplexChanSend() {
	boolChan <- now < 1 || greeter == nil // want "comparing interface to nil"
	if <-boolChan {
		fmt.Println("greeter is nil")
	}

	if boolChan <- now < 1 || greeter == nil; <-boolChan { // want "comparing interface to nil"
		fmt.Println("greeter is nil")
	}

	boolChan <- greeter != nil && now > 1 // want "comparing interface to nil"
	if <-boolChan {
		fmt.Println("greeter is not nil")
	}

	if boolChan <- greeter != nil && now > 1; <-boolChan { // want "comparing interface to nil"
		fmt.Println("greeter is not nil")
	}
}

func nilBasicFuncReturn() {
	isNil := func() bool {
		return greeter == nil // want "comparing interface to nil"
	}

	if isNil() {
		fmt.Println("greeter is nil")
	}

	isNotNil := func() bool {
		return greeter != nil // want "comparing interface to nil"
	}

	if isNotNil() {
		fmt.Println("greeter is not nil")
	}
}

func nilComplexFuncReturn() {
	isNil := func() bool {
		return now < 1 || greeter == nil // want "comparing interface to nil"
	}

	if isNil() {
		fmt.Println("greeter is nil")
	}

	isNotNil := func() bool {
		return greeter != nil && now > 1 // want "comparing interface to nil"
	}

	if isNotNil() {
		fmt.Println("greeter is not nil")
	}
}

func nilBasicInlineFuncReturn() {
	if func() bool { return greeter == nil }() { // want "comparing interface to nil"
		fmt.Println("greeter is nil")
	}

	if func() bool { return greeter != nil }() { // want "comparing interface to nil"
		fmt.Println("greeter is not nil")
	}
}

func nilComplexInlineFuncReturn() {
	if func() bool { return now < 1 || greeter == nil }() { // want "comparing interface to nil"
		fmt.Println("greeter is nil")
	}

	if func() bool { return greeter != nil && now > 1 }() { // want "comparing interface to nil"
		fmt.Println("greeter is not nil")
	}
}

func nilBasicFuncMultiReturn() {
	isNilFn := func() (int64, bool) {
		return now, greeter == nil // want "comparing interface to nil"
	}

	if _, isNil := isNilFn(); isNil {
		fmt.Println("greeter is nil")
	}

	isNotNilFn := func() (bool, int64) {
		return greeter != nil, now // want "comparing interface to nil"
	}

	if isNotNil, _ := isNotNilFn(); isNotNil {
		fmt.Println("greeter is not nil")
	}
}

func nilComplexFuncMultiReturn() {
	isNilFn := func() (int64, bool) {
		return now, now < 1 || greeter == nil // want "comparing interface to nil"
	}

	if _, isNil := isNilFn(); isNil {
		fmt.Println("greeter is nil")
	}

	isNotNilFn := func() (bool, int64) {
		return greeter != nil && now > 1, now // want "comparing interface to nil"
	}

	if isNotNil, _ := isNotNilFn(); isNotNil {
		fmt.Println("greeter is not nil")
	}
}

func nilBasicInlineFuncMultiReturn() {
	if _, isNil := func() (int64, bool) { return now, greeter == nil }(); isNil { // want "comparing interface to nil"
		fmt.Println("greeter is nil")
	}

	if isNotNil, _ := func() (bool, int64) { return greeter != nil, now }(); isNotNil { // want "comparing interface to nil"
		fmt.Println("greeter is not nil")
	}
}

func nilComplexInlineFuncMultiReturn() {
	if _, isNil := func() (int64, bool) { return now, now < 1 || greeter == nil }(); isNil { // want "comparing interface to nil"
		fmt.Println("greeter is nil")
	}

	if isNotNil, _ := func() (bool, int64) { return greeter != nil && now > 1, now }(); isNotNil { // want "comparing interface to nil"
		fmt.Println("greeter is not nil")
	}
}

func nilTypeCast() {
	var something any = greeter

	if someGreeter, ok := something.(Greeter); ok && someGreeter == nil { // want "comparing interface to nil"
		fmt.Println("greeter is nil")
	}

	if someGreeter, ok := something.(Greeter); ok && someGreeter != nil { // want "comparing interface to nil"
		fmt.Println("greeter is not nil")
	}
}

func nilTypeSwitch() {
	switch greeter.(type) {
	case nil: // want "type switch case checking interface against nil"
		fmt.Println("greeter is nil")
	}
}

func nilSkipErrors() {
	var err error
	if err == nil {
		fmt.Println("err is nil") // OK, returning a `nil` error is idiomatic
	}

	if err != nil { // OK, returning a `nil` error is idiomatic
		fmt.Println("err is not nil")
	}
}
