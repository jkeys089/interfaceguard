package a

import "fmt"

func interfaceBasicComparison() {
	if greeter == otherGreeter { // want "comparing two interfaces"
		fmt.Println("greeter == otherGreeter")
	}

	if greeter != otherGreeter { // want "comparing two interfaces"
		fmt.Println("greeter != otherGreeter")
	}
}

func interfaceComplexComparison() {
	if now < 1 || greeter == otherGreeter { // want "comparing two interfaces"
		fmt.Println("greeter == otherGreeter")
	}

	if greeter != otherGreeter && now > 1 { // want "comparing two interfaces"
		fmt.Println("greeter != otherGreeter")
	}
}

func interfaceBasicAssignment() {
	isNil := greeter == otherGreeter // want "comparing two interfaces"
	if isNil {
		fmt.Println("greeter == otherGreeter")
	}

	isNotNil := greeter != otherGreeter // want "comparing two interfaces"
	if isNotNil {
		fmt.Println("greeter != otherGreeter")
	}
}

func interfaceComplexAssignment() {
	isNil := now < 1 || greeter == otherGreeter // want "comparing two interfaces"
	if isNil {
		fmt.Println("greeter == otherGreeter")
	}

	isNotNil := greeter != otherGreeter && now > 1 // want "comparing two interfaces"
	if isNotNil {
		fmt.Println("greeter != otherGreeter")
	}
}

func interfaceBasicInlineAssignment() {
	if isNil := greeter == otherGreeter; isNil { // want "comparing two interfaces"
		fmt.Println("greeter == otherGreeter")
	}

	if isNotNil := greeter != otherGreeter; isNotNil { // want "comparing two interfaces"
		fmt.Println("greeter != otherGreeter")
	}
}

func interfaceComplexInlineAssignment() {
	if isNil := now < 1 || greeter == otherGreeter; isNil { // want "comparing two interfaces"
		fmt.Println("greeter == otherGreeter")
	}

	if isNotNil := greeter != otherGreeter && now > 1; isNotNil { // want "comparing two interfaces"
		fmt.Println("greeter != otherGreeter")
	}
}

func interfaceBasicMultiAssignment() {
	x, isNil := now, greeter == otherGreeter // want "comparing two interfaces"
	if isNil && x > 1 {
		fmt.Println("greeter == otherGreeter")
	}

	y, isNotNil := now, greeter != otherGreeter // want "comparing two interfaces"
	if isNotNil && y > 1 {
		fmt.Println("greeter != otherGreeter")
	}
}

func interfaceComplexMultiAssignment() {
	x, isNil := now, now < 1 || greeter == otherGreeter // want "comparing two interfaces"
	if isNil && x > 1 {
		fmt.Println("greeter == otherGreeter")
	}

	y, isNotNil := now, greeter != otherGreeter && now > 1 // want "comparing two interfaces"
	if isNotNil && y > 1 {
		fmt.Println("greeter != otherGreeter")
	}
}

func interfaceBasicInlineMultiAssignment() {
	if x, isNil := now, greeter == otherGreeter; isNil && x > 1 { // want "comparing two interfaces"
		fmt.Println("greeter == otherGreeter")
	}

	if y, isNotNil := now, greeter != otherGreeter; isNotNil && y > 1 { // want "comparing two interfaces"
		fmt.Println("greeter != otherGreeter")
	}
}

func interfaceComplexInlineMultiAssignment() {
	if x, isNil := now, now < 1 || greeter == otherGreeter; isNil && x > 1 { // want "comparing two interfaces"
		fmt.Println("greeter == otherGreeter")
	}

	if y, isNotNil := now, greeter != otherGreeter && now > 1; isNotNil && y > 1 { // want "comparing two interfaces"
		fmt.Println("greeter != otherGreeter")
	}
}

func interfaceSimpleForLoop() {
	for greeter == otherGreeter { // want "comparing two interfaces"
		fmt.Println("greeter == otherGreeter")
		greeter = costcoGreeter
	}

	for greeter != otherGreeter { // want "comparing two interfaces"
		fmt.Println("greeter != otherGreeter")
		greeter = nil
	}
}

func interfaceComplexForLoop() {
	for now < 1 || greeter == otherGreeter { // want "comparing two interfaces"
		fmt.Println("greeter == otherGreeter")
		greeter = costcoGreeter
	}

	for greeter != otherGreeter && now > 1 { // want "comparing two interfaces"
		fmt.Println("greeter != otherGreeter")
		greeter = nil
	}
}

func interfaceBasicInlineForLoop() {
	for isNil := greeter == otherGreeter; isNil; { // want "comparing two interfaces"
		fmt.Println("greeter == otherGreeter")
		isNil = false
	}

	for isNotNil := greeter != otherGreeter; isNotNil; { // want "comparing two interfaces"
		fmt.Println("greeter != otherGreeter")
		isNotNil = false
	}
}

func interfaceComplexInlineForLoop() {
	for isNil := now < 1 || greeter == otherGreeter; isNil; { // want "comparing two interfaces"
		fmt.Println("greeter == otherGreeter")
		isNil = false
	}

	for isNotNil := greeter != otherGreeter && now < 1; isNotNil; { // want "comparing two interfaces"
		fmt.Println("greeter != otherGreeter")
		isNotNil = false
	}
}

func interfaceBasicForRangeLoop() {
	for _, nextGreeter := range greeters {
		if nextGreeter == otherGreeter { // want "comparing two interfaces"
			fmt.Println("greeter == otherGreeter")
		}
	}

	for _, nextGreeter := range greeters {
		if nextGreeter != otherGreeter { // want "comparing two interfaces"
			fmt.Println("greeter != otherGreeter")
		}
	}
}

func interfaceComplexForRangeLoop() {
	for _, nextGreeter := range greeters {
		if now < 1 || nextGreeter == otherGreeter { // want "comparing two interfaces"
			fmt.Println("greeter == otherGreeter")
		}
	}

	for _, nextGreeter := range greeters {
		if nextGreeter != otherGreeter && now > 1 { // want "comparing two interfaces"
			fmt.Println("greeter != otherGreeter")
		}
	}
}

func interfaceBasicSwitch() {
	switch greeter == otherGreeter { // want "comparing two interfaces"
	case true:
		fmt.Println("greeter == otherGreeter")
	}

	switch greeter != otherGreeter { // want "comparing two interfaces"
	case true:
		fmt.Println("greeter != otherGreeter")
	}
}

func interfaceComplexSwitch() {
	switch now < 1 || greeter == otherGreeter { // want "comparing two interfaces"
	case true:
		fmt.Println("greeter == otherGreeter")
	}

	switch greeter != otherGreeter && now > 1 { // want "comparing two interfaces"
	case true:
		fmt.Println("greeter != otherGreeter")
	}
}

func interfaceBasicSwitchCase() {
	switch {
	case greeter == otherGreeter: // want "comparing two interfaces"
		fmt.Println("greeter == otherGreeter")
	}

	switch {
	case greeter != otherGreeter: // want "comparing two interfaces"
		fmt.Println("greeter not nil")
	}
}

func interfaceComplexSwitchCase() {
	switch {
	case now < 1 || greeter == otherGreeter: // want "comparing two interfaces"
		fmt.Println("greeter == otherGreeter")
	}

	switch {
	case greeter != otherGreeter && now > 1: // want "comparing two interfaces"
		fmt.Println("greeter not nil")
	}
}

func interfaceBasicSliceAssignment() {
	isNil := []bool{
		greeter == otherGreeter, // want "comparing two interfaces"
	}

	isNil[0] = greeter == otherGreeter // want "comparing two interfaces"
	if isNil[0] {
		fmt.Println("greeter == otherGreeter")
	}

	if isNil[0] = greeter == otherGreeter; isNil[0] { // want "comparing two interfaces"
		fmt.Println("greeter == otherGreeter")
	}

	isNotNil := []bool{
		greeter != otherGreeter, // want "comparing two interfaces"
	}

	isNotNil[0] = greeter != otherGreeter // want "comparing two interfaces"
	if isNotNil[0] {
		fmt.Println("greeter != otherGreeter")
	}

	if isNotNil[0] = greeter != otherGreeter; isNotNil[0] { // want "comparing two interfaces"
		fmt.Println("greeter != otherGreeter")
	}
}

func interfaceComplexSliceAssignment() {
	isNil := []bool{
		now < 1 || greeter == otherGreeter, // want "comparing two interfaces"
	}

	isNil[0] = now < 1 || greeter == otherGreeter // want "comparing two interfaces"
	if isNil[0] {
		fmt.Println("greeter == otherGreeter")
	}

	if isNil[0] = now < 1 || greeter == otherGreeter; isNil[0] { // want "comparing two interfaces"
		fmt.Println("greeter == otherGreeter")
	}

	isNotNil := []bool{
		greeter != otherGreeter && now > 1, // want "comparing two interfaces"
	}

	isNotNil[0] = greeter != otherGreeter && now > 1 // want "comparing two interfaces"
	if isNotNil[0] {
		fmt.Println("greeter != otherGreeter")
	}

	if isNotNil[0] = greeter != otherGreeter && now > 1; isNotNil[0] { // want "comparing two interfaces"
		fmt.Println("greeter != otherGreeter")
	}
}

func interfaceBasicMapAssignment() {
	isNil := map[string]bool{
		"a": greeter == otherGreeter, // want "comparing two interfaces"
	}

	isNil["a"] = greeter == otherGreeter // want "comparing two interfaces"
	if isNil["a"] {
		fmt.Println("greeter == otherGreeter")
	}

	if isNil["a"] = greeter == otherGreeter; isNil["a"] { // want "comparing two interfaces"
		fmt.Println("greeter == otherGreeter")
	}

	isNotNil := map[string]bool{
		"a": greeter != otherGreeter, // want "comparing two interfaces"
	}

	isNotNil["a"] = greeter != otherGreeter // want "comparing two interfaces"
	if isNotNil["a"] {
		fmt.Println("greeter != otherGreeter")
	}

	if isNotNil["a"] = greeter != otherGreeter; isNotNil["a"] { // want "comparing two interfaces"
		fmt.Println("greeter != otherGreeter")
	}
}

func interfaceComplexMapAssignment() {
	isNil := map[string]bool{
		"a": now < 1 || greeter == otherGreeter, // want "comparing two interfaces"
	}

	isNil["a"] = now < 1 || greeter == otherGreeter // want "comparing two interfaces"
	if isNil["a"] {
		fmt.Println("greeter == otherGreeter")
	}

	if isNil["a"] = now < 1 || greeter == otherGreeter; isNil["a"] { // want "comparing two interfaces"
		fmt.Println("greeter == otherGreeter")
	}

	isNotNil := map[string]bool{
		"a": greeter != otherGreeter && now > 1, // want "comparing two interfaces"
	}

	isNotNil["a"] = greeter != otherGreeter && now > 1 // want "comparing two interfaces"
	if isNotNil["a"] {
		fmt.Println("greeter != otherGreeter")
	}

	if isNotNil["a"] = greeter != otherGreeter && now > 1; isNotNil["a"] { // want "comparing two interfaces"
		fmt.Println("greeter != otherGreeter")
	}
}

func interfaceBasicStructAssignment() {
	isNil := struct {
		b bool
	}{
		b: greeter == otherGreeter, // want "comparing two interfaces"
	}

	isNil.b = greeter == otherGreeter // want "comparing two interfaces"
	if isNil.b {
		fmt.Println("greeter == otherGreeter")
	}

	if isNil.b = greeter == otherGreeter; isNil.b { // want "comparing two interfaces"
		fmt.Println("greeter == otherGreeter")
	}

	isNotNil := struct {
		b bool
	}{
		b: greeter != otherGreeter, // want "comparing two interfaces"
	}

	isNotNil.b = greeter != otherGreeter // want "comparing two interfaces"
	if isNotNil.b {
		fmt.Println("greeter != otherGreeter")
	}

	if isNotNil.b = greeter != otherGreeter; isNotNil.b { // want "comparing two interfaces"
		fmt.Println("greeter != otherGreeter")
	}
}

func interfaceComplexStructAssignment() {
	isNil := struct {
		b bool
	}{
		b: now < 1 || greeter == otherGreeter, // want "comparing two interfaces"
	}

	isNil.b = now < 1 || greeter == otherGreeter // want "comparing two interfaces"
	if isNil.b {
		fmt.Println("greeter == otherGreeter")
	}

	if isNil.b = now < 1 || greeter == otherGreeter; isNil.b { // want "comparing two interfaces"
		fmt.Println("greeter == otherGreeter")
	}

	isNotNil := struct {
		b bool
	}{
		b: greeter != otherGreeter && now > 1, // want "comparing two interfaces"
	}

	isNotNil.b = greeter != otherGreeter && now > 1 // want "comparing two interfaces"
	if isNotNil.b {
		fmt.Println("greeter != otherGreeter")
	}

	if isNotNil.b = greeter != otherGreeter && now > 1; isNotNil.b { // want "comparing two interfaces"
		fmt.Println("greeter != otherGreeter")
	}
}

func interfaceBasicChanSend() {
	boolChan <- greeter == otherGreeter // want "comparing two interfaces"
	if <-boolChan {
		fmt.Println("greeter == otherGreeter")
	}

	if boolChan <- greeter == otherGreeter; <-boolChan { // want "comparing two interfaces"
		fmt.Println("greeter == otherGreeter")
	}

	boolChan <- greeter != otherGreeter // want "comparing two interfaces"
	if <-boolChan {
		fmt.Println("greeter != otherGreeter")
	}

	if boolChan <- greeter != otherGreeter; <-boolChan { // want "comparing two interfaces"
		fmt.Println("greeter != otherGreeter")
	}
}

func interfaceComplexChanSend() {
	boolChan <- now < 1 || greeter == otherGreeter // want "comparing two interfaces"
	if <-boolChan {
		fmt.Println("greeter == otherGreeter")
	}

	if boolChan <- now < 1 || greeter == otherGreeter; <-boolChan { // want "comparing two interfaces"
		fmt.Println("greeter == otherGreeter")
	}

	boolChan <- greeter != otherGreeter && now > 1 // want "comparing two interfaces"
	if <-boolChan {
		fmt.Println("greeter != otherGreeter")
	}

	if boolChan <- greeter != otherGreeter && now > 1; <-boolChan { // want "comparing two interfaces"
		fmt.Println("greeter != otherGreeter")
	}
}

func interfaceBasicFuncReturn() {
	isNil := func() bool {
		return greeter == otherGreeter // want "comparing two interfaces"
	}

	if isNil() {
		fmt.Println("greeter == otherGreeter")
	}

	isNotNil := func() bool {
		return greeter != otherGreeter // want "comparing two interfaces"
	}

	if isNotNil() {
		fmt.Println("greeter != otherGreeter")
	}
}

func interfaceComplexFuncReturn() {
	isNil := func() bool {
		return now < 1 || greeter == otherGreeter // want "comparing two interfaces"
	}

	if isNil() {
		fmt.Println("greeter == otherGreeter")
	}

	isNotNil := func() bool {
		return greeter != otherGreeter && now > 1 // want "comparing two interfaces"
	}

	if isNotNil() {
		fmt.Println("greeter != otherGreeter")
	}
}

func interfaceBasicInlineFuncReturn() {
	if func() bool { return greeter == otherGreeter }() { // want "comparing two interfaces"
		fmt.Println("greeter == otherGreeter")
	}

	if func() bool { return greeter != otherGreeter }() { // want "comparing two interfaces"
		fmt.Println("greeter != otherGreeter")
	}
}

func interfaceComplexInlineFuncReturn() {
	if func() bool { return now < 1 || greeter == otherGreeter }() { // want "comparing two interfaces"
		fmt.Println("greeter == otherGreeter")
	}

	if func() bool { return greeter != otherGreeter && now > 1 }() { // want "comparing two interfaces"
		fmt.Println("greeter != otherGreeter")
	}
}

func interfaceBasicFuncMultiReturn() {
	isNilFn := func() (int64, bool) {
		return now, greeter == otherGreeter // want "comparing two interfaces"
	}

	if _, isNil := isNilFn(); isNil {
		fmt.Println("greeter == otherGreeter")
	}

	isNotNilFn := func() (bool, int64) {
		return greeter != otherGreeter, now // want "comparing two interfaces"
	}

	if isNotNil, _ := isNotNilFn(); isNotNil {
		fmt.Println("greeter != otherGreeter")
	}
}

func interfaceComplexFuncMultiReturn() {
	isNilFn := func() (int64, bool) {
		return now, now < 1 || greeter == otherGreeter // want "comparing two interfaces"
	}

	if _, isNil := isNilFn(); isNil {
		fmt.Println("greeter == otherGreeter")
	}

	isNotNilFn := func() (bool, int64) {
		return greeter != otherGreeter && now > 1, now // want "comparing two interfaces"
	}

	if isNotNil, _ := isNotNilFn(); isNotNil {
		fmt.Println("greeter != otherGreeter")
	}
}

func interfaceBasicInlineFuncMultiReturn() {
	if _, isNil := func() (int64, bool) { return now, greeter == otherGreeter }(); isNil { // want "comparing two interfaces"
		fmt.Println("greeter == otherGreeter")
	}

	if isNotNil, _ := func() (bool, int64) { return greeter != otherGreeter, now }(); isNotNil { // want "comparing two interfaces"
		fmt.Println("greeter != otherGreeter")
	}
}

func interfaceComplexInlineFuncMultiReturn() {
	if _, isNil := func() (int64, bool) { return now, now < 1 || greeter == otherGreeter }(); isNil { // want "comparing two interfaces"
		fmt.Println("greeter == otherGreeter")
	}

	if isNotNil, _ := func() (bool, int64) { return greeter != otherGreeter && now > 1, now }(); isNotNil { // want "comparing two interfaces"
		fmt.Println("greeter != otherGreeter")
	}
}

func interfaceTypeCast() {
	var something any = greeter

	if someGreeter, ok := something.(Greeter); ok && someGreeter == otherGreeter { // want "comparing two interfaces"
		fmt.Println("greeter == otherGreeter")
	}

	if someGreeter, ok := something.(Greeter); ok && someGreeter != otherGreeter { // want "comparing two interfaces"
		fmt.Println("greeter != otherGreeter")
	}
}
