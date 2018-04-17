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


func TestGotMyBonus(t *testing.T) {
	if  floatEquals(CalculateBonus(200, 100, 100, 0), 100.0)  {
		t.Fail()
	}
}

func TestGotNoBonusMatchQuota(t *testing.T) {
	if floatEquals(CalculateBonus(200, 200, 100, 0), 0.0) {
		t.Fail()
	}
}

func TestGotNoBonusBelowQuota(t *testing.T) {
    if  floatEquals(CalculateBonus(300, 400, 100, 0), 0.0)  {
		t.Fail()
	}
}

func TestBonusWithCommission(t *testing.T) {
	if  floatEquals(CalculateBonus(200, 100, 50, 0), 50.0)  {
		t.Fail()
	}
}

func TestBonusWithCommissionAndTax(t *testing.T) {
	if  floatEquals(CalculateBonus(200, 100, 50, 10), 45.0)  {
		t.Fail()
	}
}

func TestNoBonusWithCommission(t *testing.T) {
	if  floatEquals(CalculateBonus(200, 300, -50, 0), 0.0)  {
		t.Fail()
	}
}

func TestGotTeamBonus(t *testing.T) {
	if  floatEquals(CalculateTeamBonus(200, 100, 4,100), 25.0)  {
		t.Fail()
	}
}

func TestNoTeamBonus(t *testing.T) {
	if  floatEquals(CalculateTeamBonus(200, 100, 0, 100), 0)  {
		t.Fail()
	}
}

func TestTeamBonusWithCommission(t *testing.T) {
	if  floatEquals(CalculateTeamBonus(300, 100, 2, 80), 80)  {
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
