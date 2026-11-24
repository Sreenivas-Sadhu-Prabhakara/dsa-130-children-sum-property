package main

import "testing"

func TestHasChildrenSum(t *testing.T) {
	ok := tn(10, tn(8, tn(3, nil, nil), tn(5, nil, nil)), tn(2, tn(2, nil, nil), nil))
	if !HasChildrenSum(ok) {
		t.Error("expected property to hold")
	}
	bad := tn(10, tn(8, nil, nil), tn(3, nil, nil)) // 8+3 != 10
	if HasChildrenSum(bad) {
		t.Error("expected property to fail at the root")
	}
	deepBad := tn(10, tn(8, tn(4, nil, nil), tn(5, nil, nil)), tn(2, nil, nil)) // 4+5 != 8
	if HasChildrenSum(deepBad) {
		t.Error("expected property to fail below the root")
	}
	if !HasChildrenSum(nil) {
		t.Error("empty tree holds trivially")
	}
	if !HasChildrenSum(tn(7, nil, nil)) {
		t.Error("a leaf holds trivially")
	}
	oneChild := tn(5, tn(5, nil, nil), nil)
	if !HasChildrenSum(oneChild) {
		t.Error("single child equal to parent holds")
	}
}
