// Code generated automatically. DO NOT EDIT.
package disableall

import "fmt"

func nilBasicComparison() {
	if greeter == nil { 
		fmt.Println("greeter is nil")
	}

	if greeter != nil { 
		fmt.Println("greeter is not nil")
	}
}

func nilComplexComparison() {
	if now < 1 || greeter == nil { 
		fmt.Println("greeter is nil")
	}

	if greeter != nil && now > 1 { 
		fmt.Println("greeter is not nil")
	}
}

func nilBasicAssignment() {
	isNil := greeter == nil 
	if isNil {
		fmt.Println("greeter is nil")
	}

	isNotNil := greeter != nil 
	if isNotNil {
		fmt.Println("greeter is not nil")
	}
}

func nilComplexAssignment() {
	isNil := now < 1 || greeter == nil 
	if isNil {
		fmt.Println("greeter is nil")
	}

	isNotNil := greeter != nil && now > 1 
	if isNotNil {
		fmt.Println("greeter is not nil")
	}
}

func nilBasicInlineAssignment() {
	if isNil := greeter == nil; isNil { 
		fmt.Println("greeter is nil")
	}

	if isNotNil := greeter != nil; isNotNil { 
		fmt.Println("greeter is not nil")
	}
}

func nilComplexInlineAssignment() {
	if isNil := now < 1 || greeter == nil; isNil { 
		fmt.Println("greeter is nil")
	}

	if isNotNil := greeter != nil && now > 1; isNotNil { 
		fmt.Println("greeter is not nil")
	}
}

func nilBasicMultiAssignment() {
	x, isNil := now, greeter == nil 
	if isNil && x > 1 {
		fmt.Println("greeter is nil")
	}

	y, isNotNil := now, greeter != nil 
	if isNotNil && y > 1 {
		fmt.Println("greeter is not nil")
	}
}

func nilComplexMultiAssignment() {
	x, isNil := now, now < 1 || greeter == nil 
	if isNil && x > 1 {
		fmt.Println("greeter is nil")
	}

	y, isNotNil := now, greeter != nil && now > 1 
	if isNotNil && y > 1 {
		fmt.Println("greeter is not nil")
	}
}

func nilBasicInlineMultiAssignment() {
	if x, isNil := now, greeter == nil; isNil && x > 1 { 
		fmt.Println("greeter is nil")
	}

	if y, isNotNil := now, greeter != nil; isNotNil && y > 1 { 
		fmt.Println("greeter is not nil")
	}
}

func nilComplexInlineMultiAssignment() {
	if x, isNil := now, now < 1 || greeter == nil; isNil && x > 1 { 
		fmt.Println("greeter is nil")
	}

	if y, isNotNil := now, greeter != nil && now > 1; isNotNil && y > 1 { 
		fmt.Println("greeter is not nil")
	}
}

func nilSimpleForLoop() {
	for greeter == nil { 
		fmt.Println("greeter is nil")
		greeter = costcoGreeter
	}

	for greeter != nil { 
		fmt.Println("greeter is not nil")
		greeter = nil
	}
}

func nilComplexForLoop() {
	for now < 1 || greeter == nil { 
		fmt.Println("greeter is nil")
		greeter = costcoGreeter
	}

	for greeter != nil && now > 1 { 
		fmt.Println("greeter is not nil")
		greeter = nil
	}
}

func nilBasicInlineForLoop() {
	for isNil := greeter == nil; isNil; { 
		fmt.Println("greeter is nil")
		isNil = false
	}

	for isNotNil := greeter != nil; isNotNil; { 
		fmt.Println("greeter is not nil")
		isNotNil = false
	}
}

func nilComplexInlineForLoop() {
	for isNil := now < 1 || greeter == nil; isNil; { 
		fmt.Println("greeter is nil")
		isNil = false
	}

	for isNotNil := greeter != nil && now < 1; isNotNil; { 
		fmt.Println("greeter is not nil")
		isNotNil = false
	}
}

func nilBasicForRangeLoop() {
	for _, nextGreeter := range greeters {
		if nextGreeter == nil { 
			fmt.Println("greeter is nil")
		}
	}

	for _, nextGreeter := range greeters {
		if nextGreeter != nil { 
			fmt.Println("greeter is not nil")
		}
	}
}

func nilComplexForRangeLoop() {
	for _, nextGreeter := range greeters {
		if now < 1 || nextGreeter == nil { 
			fmt.Println("greeter is nil")
		}
	}

	for _, nextGreeter := range greeters {
		if nextGreeter != nil && now > 1 { 
			fmt.Println("greeter is not nil")
		}
	}
}

func nilBasicSwitch() {
	switch greeter == nil { 
	case true:
		fmt.Println("greeter is nil")
	}

	switch greeter != nil { 
	case true:
		fmt.Println("greeter is not nil")
	}
}

func nilComplexSwitch() {
	switch now < 1 || greeter == nil { 
	case true:
		fmt.Println("greeter is nil")
	}

	switch greeter != nil && now > 1 { 
	case true:
		fmt.Println("greeter is not nil")
	}
}

func nilBasicSwitchCase() {
	switch {
	case greeter == nil: 
		fmt.Println("greeter is nil")
	}

	switch {
	case greeter != nil: 
		fmt.Println("greeter not nil")
	}
}

func nilComplexSwitchCase() {
	switch {
	case now < 1 || greeter == nil: 
		fmt.Println("greeter is nil")
	}

	switch {
	case greeter != nil && now > 1: 
		fmt.Println("greeter not nil")
	}
}

func nilBasicSliceAssignment() {
	isNil := []bool{
		greeter == nil, 
	}

	isNil[0] = greeter == nil 
	if isNil[0] {
		fmt.Println("greeter is nil")
	}

	if isNil[0] = greeter == nil; isNil[0] { 
		fmt.Println("greeter is nil")
	}

	isNotNil := []bool{
		greeter != nil, 
	}

	isNotNil[0] = greeter != nil 
	if isNotNil[0] {
		fmt.Println("greeter is not nil")
	}

	if isNotNil[0] = greeter != nil; isNotNil[0] { 
		fmt.Println("greeter is not nil")
	}
}

func nilComplexSliceAssignment() {
	isNil := []bool{
		now < 1 || greeter == nil, 
	}

	isNil[0] = now < 1 || greeter == nil 
	if isNil[0] {
		fmt.Println("greeter is nil")
	}

	if isNil[0] = now < 1 || greeter == nil; isNil[0] { 
		fmt.Println("greeter is nil")
	}

	isNotNil := []bool{
		greeter != nil && now > 1, 
	}

	isNotNil[0] = greeter != nil && now > 1 
	if isNotNil[0] {
		fmt.Println("greeter is not nil")
	}

	if isNotNil[0] = greeter != nil && now > 1; isNotNil[0] { 
		fmt.Println("greeter is not nil")
	}
}

func nilBasicMapAssignment() {
	isNil := map[string]bool{
		"a": greeter == nil, 
	}

	isNil["a"] = greeter == nil 
	if isNil["a"] {
		fmt.Println("greeter is nil")
	}

	if isNil["a"] = greeter == nil; isNil["a"] { 
		fmt.Println("greeter is nil")
	}

	isNotNil := map[string]bool{
		"a": greeter != nil, 
	}

	isNotNil["a"] = greeter != nil 
	if isNotNil["a"] {
		fmt.Println("greeter is not nil")
	}

	if isNotNil["a"] = greeter != nil; isNotNil["a"] { 
		fmt.Println("greeter is not nil")
	}
}

func nilComplexMapAssignment() {
	isNil := map[string]bool{
		"a": now < 1 || greeter == nil, 
	}

	isNil["a"] = now < 1 || greeter == nil 
	if isNil["a"] {
		fmt.Println("greeter is nil")
	}

	if isNil["a"] = now < 1 || greeter == nil; isNil["a"] { 
		fmt.Println("greeter is nil")
	}

	isNotNil := map[string]bool{
		"a": greeter != nil && now > 1, 
	}

	isNotNil["a"] = greeter != nil && now > 1 
	if isNotNil["a"] {
		fmt.Println("greeter is not nil")
	}

	if isNotNil["a"] = greeter != nil && now > 1; isNotNil["a"] { 
		fmt.Println("greeter is not nil")
	}
}

func nilBasicStructAssignment() {
	isNil := struct {
		b bool
	}{
		b: greeter == nil, 
	}

	isNil.b = greeter == nil 
	if isNil.b {
		fmt.Println("greeter is nil")
	}

	if isNil.b = greeter == nil; isNil.b { 
		fmt.Println("greeter is nil")
	}

	isNotNil := struct {
		b bool
	}{
		b: greeter != nil, 
	}

	isNotNil.b = greeter != nil 
	if isNotNil.b {
		fmt.Println("greeter is not nil")
	}

	if isNotNil.b = greeter != nil; isNotNil.b { 
		fmt.Println("greeter is not nil")
	}
}

func nilComplexStructAssignment() {
	isNil := struct {
		b bool
	}{
		b: now < 1 || greeter == nil, 
	}

	isNil.b = now < 1 || greeter == nil 
	if isNil.b {
		fmt.Println("greeter is nil")
	}

	if isNil.b = now < 1 || greeter == nil; isNil.b { 
		fmt.Println("greeter is nil")
	}

	isNotNil := struct {
		b bool
	}{
		b: greeter != nil && now > 1, 
	}

	isNotNil.b = greeter != nil && now > 1 
	if isNotNil.b {
		fmt.Println("greeter is not nil")
	}

	if isNotNil.b = greeter != nil && now > 1; isNotNil.b { 
		fmt.Println("greeter is not nil")
	}
}

func nilBasicChanSend() {
	boolChan <- greeter == nil 
	if <-boolChan {
		fmt.Println("greeter is nil")
	}

	if boolChan <- greeter == nil; <-boolChan { 
		fmt.Println("greeter is nil")
	}

	boolChan <- greeter != nil 
	if <-boolChan {
		fmt.Println("greeter is not nil")
	}

	if boolChan <- greeter != nil; <-boolChan { 
		fmt.Println("greeter is not nil")
	}
}

func nilComplexChanSend() {
	boolChan <- now < 1 || greeter == nil 
	if <-boolChan {
		fmt.Println("greeter is nil")
	}

	if boolChan <- now < 1 || greeter == nil; <-boolChan { 
		fmt.Println("greeter is nil")
	}

	boolChan <- greeter != nil && now > 1 
	if <-boolChan {
		fmt.Println("greeter is not nil")
	}

	if boolChan <- greeter != nil && now > 1; <-boolChan { 
		fmt.Println("greeter is not nil")
	}
}

func nilBasicFuncReturn() {
	isNil := func() bool {
		return greeter == nil 
	}

	if isNil() {
		fmt.Println("greeter is nil")
	}

	isNotNil := func() bool {
		return greeter != nil 
	}

	if isNotNil() {
		fmt.Println("greeter is not nil")
	}
}

func nilComplexFuncReturn() {
	isNil := func() bool {
		return now < 1 || greeter == nil 
	}

	if isNil() {
		fmt.Println("greeter is nil")
	}

	isNotNil := func() bool {
		return greeter != nil && now > 1 
	}

	if isNotNil() {
		fmt.Println("greeter is not nil")
	}
}

func nilBasicInlineFuncReturn() {
	if func() bool { return greeter == nil }() { 
		fmt.Println("greeter is nil")
	}

	if func() bool { return greeter != nil }() { 
		fmt.Println("greeter is not nil")
	}
}

func nilComplexInlineFuncReturn() {
	if func() bool { return now < 1 || greeter == nil }() { 
		fmt.Println("greeter is nil")
	}

	if func() bool { return greeter != nil && now > 1 }() { 
		fmt.Println("greeter is not nil")
	}
}

func nilBasicFuncMultiReturn() {
	isNilFn := func() (int64, bool) {
		return now, greeter == nil 
	}

	if _, isNil := isNilFn(); isNil {
		fmt.Println("greeter is nil")
	}

	isNotNilFn := func() (bool, int64) {
		return greeter != nil, now 
	}

	if isNotNil, _ := isNotNilFn(); isNotNil {
		fmt.Println("greeter is not nil")
	}
}

func nilComplexFuncMultiReturn() {
	isNilFn := func() (int64, bool) {
		return now, now < 1 || greeter == nil 
	}

	if _, isNil := isNilFn(); isNil {
		fmt.Println("greeter is nil")
	}

	isNotNilFn := func() (bool, int64) {
		return greeter != nil && now > 1, now 
	}

	if isNotNil, _ := isNotNilFn(); isNotNil {
		fmt.Println("greeter is not nil")
	}
}

func nilBasicInlineFuncMultiReturn() {
	if _, isNil := func() (int64, bool) { return now, greeter == nil }(); isNil { 
		fmt.Println("greeter is nil")
	}

	if isNotNil, _ := func() (bool, int64) { return greeter != nil, now }(); isNotNil { 
		fmt.Println("greeter is not nil")
	}
}

func nilComplexInlineFuncMultiReturn() {
	if _, isNil := func() (int64, bool) { return now, now < 1 || greeter == nil }(); isNil { 
		fmt.Println("greeter is nil")
	}

	if isNotNil, _ := func() (bool, int64) { return greeter != nil && now > 1, now }(); isNotNil { 
		fmt.Println("greeter is not nil")
	}
}

func nilTypeCast() {
	var something any = greeter

	if someGreeter, ok := something.(Greeter); ok && someGreeter == nil { 
		fmt.Println("greeter is nil")
	}

	if someGreeter, ok := something.(Greeter); ok && someGreeter != nil { 
		fmt.Println("greeter is not nil")
	}
}

func nilTypeSwitch() {
	switch greeter.(type) {
	case nil: 
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
