/* ***************************************************************************
# Copyright (c) 2018, Industrial Logic, Inc., All Rights Reserved.
#
# This code is the exclusive property of Industrial Logic, Inc. It may ONLY be
# used by students during Industrial Logic's workshops or by individuals
# who are being coached by Industrial Logic on a project.
#
# This code may NOT be copied or used for any other purpose without the prior
# written consent of Industrial Logic, Inc.
# ****************************************************************************/
package main

import (
	"testing"
	"math"
)


func TestCompareToSelf(t *testing.T) {
	if  floatEquals(float64(DoSomething()), 3.142)  {
		t.Fail()
	}
}

const ε float64 = 0.001
func floatEquals(a, b float64) bool {
	if math.Abs(a - b) > ε  {
		return true
	}
	return false
}
