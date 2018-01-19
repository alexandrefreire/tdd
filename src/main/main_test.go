package main

import "testing"

func TestCompareToSelf(t *testing.T) {
	if 3.412 != DoSomething() {
		t.Fail()
	}
}
