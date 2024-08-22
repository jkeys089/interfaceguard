// Code generated automatically. DO NOT EDIT.
package disableiface

import "fmt"

func interfaceBasicComparison() {
	if greeter == otherGreeter { 
		fmt.Println("greeter == otherGreeter")
	}

	if greeter != otherGreeter { 
		fmt.Println("greeter != otherGreeter")
	}
}

func interfaceComplexComparison() {
	if now < 1 || greeter == otherGreeter { 
		fmt.Println("greeter == otherGreeter")
	}

	if greeter != otherGreeter && now > 1 { 
		fmt.Println("greeter != otherGreeter")
	}
}

func interfaceBasicAssignment() {
	isNil := greeter == otherGreeter 
	if isNil {
		fmt.Println("greeter == otherGreeter")
	}

	isNotNil := greeter != otherGreeter 
	if isNotNil {
		fmt.Println("greeter != otherGreeter")
	}
}

func interfaceComplexAssignment() {
	isNil := now < 1 || greeter == otherGreeter 
	if isNil {
		fmt.Println("greeter == otherGreeter")
	}

	isNotNil := greeter != otherGreeter && now > 1 
	if isNotNil {
		fmt.Println("greeter != otherGreeter")
	}
}

func interfaceBasicInlineAssignment() {
	if isNil := greeter == otherGreeter; isNil { 
		fmt.Println("greeter == otherGreeter")
	}

	if isNotNil := greeter != otherGreeter; isNotNil { 
		fmt.Println("greeter != otherGreeter")
	}
}

func interfaceComplexInlineAssignment() {
	if isNil := now < 1 || greeter == otherGreeter; isNil { 
		fmt.Println("greeter == otherGreeter")
	}

	if isNotNil := greeter != otherGreeter && now > 1; isNotNil { 
		fmt.Println("greeter != otherGreeter")
	}
}

func interfaceBasicMultiAssignment() {
	x, isNil := now, greeter == otherGreeter 
	if isNil && x > 1 {
		fmt.Println("greeter == otherGreeter")
	}

	y, isNotNil := now, greeter != otherGreeter 
	if isNotNil && y > 1 {
		fmt.Println("greeter != otherGreeter")
	}
}

func interfaceComplexMultiAssignment() {
	x, isNil := now, now < 1 || greeter == otherGreeter 
	if isNil && x > 1 {
		fmt.Println("greeter == otherGreeter")
	}

	y, isNotNil := now, greeter != otherGreeter && now > 1 
	if isNotNil && y > 1 {
		fmt.Println("greeter != otherGreeter")
	}
}

func interfaceBasicInlineMultiAssignment() {
	if x, isNil := now, greeter == otherGreeter; isNil && x > 1 { 
		fmt.Println("greeter == otherGreeter")
	}

	if y, isNotNil := now, greeter != otherGreeter; isNotNil && y > 1 { 
		fmt.Println("greeter != otherGreeter")
	}
}

func interfaceComplexInlineMultiAssignment() {
	if x, isNil := now, now < 1 || greeter == otherGreeter; isNil && x > 1 { 
		fmt.Println("greeter == otherGreeter")
	}

	if y, isNotNil := now, greeter != otherGreeter && now > 1; isNotNil && y > 1 { 
		fmt.Println("greeter != otherGreeter")
	}
}

func interfaceSimpleForLoop() {
	for greeter == otherGreeter { 
		fmt.Println("greeter == otherGreeter")
		greeter = costcoGreeter
	}

	for greeter != otherGreeter { 
		fmt.Println("greeter != otherGreeter")
		greeter = nil
	}
}

func interfaceComplexForLoop() {
	for now < 1 || greeter == otherGreeter { 
		fmt.Println("greeter == otherGreeter")
		greeter = costcoGreeter
	}

	for greeter != otherGreeter && now > 1 { 
		fmt.Println("greeter != otherGreeter")
		greeter = nil
	}
}

func interfaceBasicInlineForLoop() {
	for isNil := greeter == otherGreeter; isNil; { 
		fmt.Println("greeter == otherGreeter")
		isNil = false
	}

	for isNotNil := greeter != otherGreeter; isNotNil; { 
		fmt.Println("greeter != otherGreeter")
		isNotNil = false
	}
}

func interfaceComplexInlineForLoop() {
	for isNil := now < 1 || greeter == otherGreeter; isNil; { 
		fmt.Println("greeter == otherGreeter")
		isNil = false
	}

	for isNotNil := greeter != otherGreeter && now < 1; isNotNil; { 
		fmt.Println("greeter != otherGreeter")
		isNotNil = false
	}
}

func interfaceBasicForRangeLoop() {
	for _, nextGreeter := range greeters {
		if nextGreeter == otherGreeter { 
			fmt.Println("greeter == otherGreeter")
		}
	}

	for _, nextGreeter := range greeters {
		if nextGreeter != otherGreeter { 
			fmt.Println("greeter != otherGreeter")
		}
	}
}

func interfaceComplexForRangeLoop() {
	for _, nextGreeter := range greeters {
		if now < 1 || nextGreeter == otherGreeter { 
			fmt.Println("greeter == otherGreeter")
		}
	}

	for _, nextGreeter := range greeters {
		if nextGreeter != otherGreeter && now > 1 { 
			fmt.Println("greeter != otherGreeter")
		}
	}
}

func interfaceBasicSwitch() {
	switch greeter == otherGreeter { 
	case true:
		fmt.Println("greeter == otherGreeter")
	}

	switch greeter != otherGreeter { 
	case true:
		fmt.Println("greeter != otherGreeter")
	}
}

func interfaceComplexSwitch() {
	switch now < 1 || greeter == otherGreeter { 
	case true:
		fmt.Println("greeter == otherGreeter")
	}

	switch greeter != otherGreeter && now > 1 { 
	case true:
		fmt.Println("greeter != otherGreeter")
	}
}

func interfaceBasicSwitchCase() {
	switch {
	case greeter == otherGreeter: 
		fmt.Println("greeter == otherGreeter")
	}

	switch {
	case greeter != otherGreeter: 
		fmt.Println("greeter not nil")
	}
}

func interfaceComplexSwitchCase() {
	switch {
	case now < 1 || greeter == otherGreeter: 
		fmt.Println("greeter == otherGreeter")
	}

	switch {
	case greeter != otherGreeter && now > 1: 
		fmt.Println("greeter not nil")
	}
}

func interfaceBasicSliceAssignment() {
	isNil := []bool{
		greeter == otherGreeter, 
	}

	isNil[0] = greeter == otherGreeter 
	if isNil[0] {
		fmt.Println("greeter == otherGreeter")
	}

	if isNil[0] = greeter == otherGreeter; isNil[0] { 
		fmt.Println("greeter == otherGreeter")
	}

	isNotNil := []bool{
		greeter != otherGreeter, 
	}

	isNotNil[0] = greeter != otherGreeter 
	if isNotNil[0] {
		fmt.Println("greeter != otherGreeter")
	}

	if isNotNil[0] = greeter != otherGreeter; isNotNil[0] { 
		fmt.Println("greeter != otherGreeter")
	}
}

func interfaceComplexSliceAssignment() {
	isNil := []bool{
		now < 1 || greeter == otherGreeter, 
	}

	isNil[0] = now < 1 || greeter == otherGreeter 
	if isNil[0] {
		fmt.Println("greeter == otherGreeter")
	}

	if isNil[0] = now < 1 || greeter == otherGreeter; isNil[0] { 
		fmt.Println("greeter == otherGreeter")
	}

	isNotNil := []bool{
		greeter != otherGreeter && now > 1, 
	}

	isNotNil[0] = greeter != otherGreeter && now > 1 
	if isNotNil[0] {
		fmt.Println("greeter != otherGreeter")
	}

	if isNotNil[0] = greeter != otherGreeter && now > 1; isNotNil[0] { 
		fmt.Println("greeter != otherGreeter")
	}
}

func interfaceBasicMapAssignment() {
	isNil := map[string]bool{
		"a": greeter == otherGreeter, 
	}

	isNil["a"] = greeter == otherGreeter 
	if isNil["a"] {
		fmt.Println("greeter == otherGreeter")
	}

	if isNil["a"] = greeter == otherGreeter; isNil["a"] { 
		fmt.Println("greeter == otherGreeter")
	}

	isNotNil := map[string]bool{
		"a": greeter != otherGreeter, 
	}

	isNotNil["a"] = greeter != otherGreeter 
	if isNotNil["a"] {
		fmt.Println("greeter != otherGreeter")
	}

	if isNotNil["a"] = greeter != otherGreeter; isNotNil["a"] { 
		fmt.Println("greeter != otherGreeter")
	}
}

func interfaceComplexMapAssignment() {
	isNil := map[string]bool{
		"a": now < 1 || greeter == otherGreeter, 
	}

	isNil["a"] = now < 1 || greeter == otherGreeter 
	if isNil["a"] {
		fmt.Println("greeter == otherGreeter")
	}

	if isNil["a"] = now < 1 || greeter == otherGreeter; isNil["a"] { 
		fmt.Println("greeter == otherGreeter")
	}

	isNotNil := map[string]bool{
		"a": greeter != otherGreeter && now > 1, 
	}

	isNotNil["a"] = greeter != otherGreeter && now > 1 
	if isNotNil["a"] {
		fmt.Println("greeter != otherGreeter")
	}

	if isNotNil["a"] = greeter != otherGreeter && now > 1; isNotNil["a"] { 
		fmt.Println("greeter != otherGreeter")
	}
}

func interfaceBasicStructAssignment() {
	isNil := struct {
		b bool
	}{
		b: greeter == otherGreeter, 
	}

	isNil.b = greeter == otherGreeter 
	if isNil.b {
		fmt.Println("greeter == otherGreeter")
	}

	if isNil.b = greeter == otherGreeter; isNil.b { 
		fmt.Println("greeter == otherGreeter")
	}

	isNotNil := struct {
		b bool
	}{
		b: greeter != otherGreeter, 
	}

	isNotNil.b = greeter != otherGreeter 
	if isNotNil.b {
		fmt.Println("greeter != otherGreeter")
	}

	if isNotNil.b = greeter != otherGreeter; isNotNil.b { 
		fmt.Println("greeter != otherGreeter")
	}
}

func interfaceComplexStructAssignment() {
	isNil := struct {
		b bool
	}{
		b: now < 1 || greeter == otherGreeter, 
	}

	isNil.b = now < 1 || greeter == otherGreeter 
	if isNil.b {
		fmt.Println("greeter == otherGreeter")
	}

	if isNil.b = now < 1 || greeter == otherGreeter; isNil.b { 
		fmt.Println("greeter == otherGreeter")
	}

	isNotNil := struct {
		b bool
	}{
		b: greeter != otherGreeter && now > 1, 
	}

	isNotNil.b = greeter != otherGreeter && now > 1 
	if isNotNil.b {
		fmt.Println("greeter != otherGreeter")
	}

	if isNotNil.b = greeter != otherGreeter && now > 1; isNotNil.b { 
		fmt.Println("greeter != otherGreeter")
	}
}

func interfaceBasicChanSend() {
	boolChan <- greeter == otherGreeter 
	if <-boolChan {
		fmt.Println("greeter == otherGreeter")
	}

	if boolChan <- greeter == otherGreeter; <-boolChan { 
		fmt.Println("greeter == otherGreeter")
	}

	boolChan <- greeter != otherGreeter 
	if <-boolChan {
		fmt.Println("greeter != otherGreeter")
	}

	if boolChan <- greeter != otherGreeter; <-boolChan { 
		fmt.Println("greeter != otherGreeter")
	}
}

func interfaceComplexChanSend() {
	boolChan <- now < 1 || greeter == otherGreeter 
	if <-boolChan {
		fmt.Println("greeter == otherGreeter")
	}

	if boolChan <- now < 1 || greeter == otherGreeter; <-boolChan { 
		fmt.Println("greeter == otherGreeter")
	}

	boolChan <- greeter != otherGreeter && now > 1 
	if <-boolChan {
		fmt.Println("greeter != otherGreeter")
	}

	if boolChan <- greeter != otherGreeter && now > 1; <-boolChan { 
		fmt.Println("greeter != otherGreeter")
	}
}

func interfaceBasicFuncReturn() {
	isNil := func() bool {
		return greeter == otherGreeter 
	}

	if isNil() {
		fmt.Println("greeter == otherGreeter")
	}

	isNotNil := func() bool {
		return greeter != otherGreeter 
	}

	if isNotNil() {
		fmt.Println("greeter != otherGreeter")
	}
}

func interfaceComplexFuncReturn() {
	isNil := func() bool {
		return now < 1 || greeter == otherGreeter 
	}

	if isNil() {
		fmt.Println("greeter == otherGreeter")
	}

	isNotNil := func() bool {
		return greeter != otherGreeter && now > 1 
	}

	if isNotNil() {
		fmt.Println("greeter != otherGreeter")
	}
}

func interfaceBasicInlineFuncReturn() {
	if func() bool { return greeter == otherGreeter }() { 
		fmt.Println("greeter == otherGreeter")
	}

	if func() bool { return greeter != otherGreeter }() { 
		fmt.Println("greeter != otherGreeter")
	}
}

func interfaceComplexInlineFuncReturn() {
	if func() bool { return now < 1 || greeter == otherGreeter }() { 
		fmt.Println("greeter == otherGreeter")
	}

	if func() bool { return greeter != otherGreeter && now > 1 }() { 
		fmt.Println("greeter != otherGreeter")
	}
}

func interfaceBasicFuncMultiReturn() {
	isNilFn := func() (int64, bool) {
		return now, greeter == otherGreeter 
	}

	if _, isNil := isNilFn(); isNil {
		fmt.Println("greeter == otherGreeter")
	}

	isNotNilFn := func() (bool, int64) {
		return greeter != otherGreeter, now 
	}

	if isNotNil, _ := isNotNilFn(); isNotNil {
		fmt.Println("greeter != otherGreeter")
	}
}

func interfaceComplexFuncMultiReturn() {
	isNilFn := func() (int64, bool) {
		return now, now < 1 || greeter == otherGreeter 
	}

	if _, isNil := isNilFn(); isNil {
		fmt.Println("greeter == otherGreeter")
	}

	isNotNilFn := func() (bool, int64) {
		return greeter != otherGreeter && now > 1, now 
	}

	if isNotNil, _ := isNotNilFn(); isNotNil {
		fmt.Println("greeter != otherGreeter")
	}
}

func interfaceBasicInlineFuncMultiReturn() {
	if _, isNil := func() (int64, bool) { return now, greeter == otherGreeter }(); isNil { 
		fmt.Println("greeter == otherGreeter")
	}

	if isNotNil, _ := func() (bool, int64) { return greeter != otherGreeter, now }(); isNotNil { 
		fmt.Println("greeter != otherGreeter")
	}
}

func interfaceComplexInlineFuncMultiReturn() {
	if _, isNil := func() (int64, bool) { return now, now < 1 || greeter == otherGreeter }(); isNil { 
		fmt.Println("greeter == otherGreeter")
	}

	if isNotNil, _ := func() (bool, int64) { return greeter != otherGreeter && now > 1, now }(); isNotNil { 
		fmt.Println("greeter != otherGreeter")
	}
}

func interfaceTypeCast() {
	var something any = greeter

	if someGreeter, ok := something.(Greeter); ok && someGreeter == otherGreeter { 
		fmt.Println("greeter == otherGreeter")
	}

	if someGreeter, ok := something.(Greeter); ok && someGreeter != otherGreeter { 
		fmt.Println("greeter != otherGreeter")
	}
}
