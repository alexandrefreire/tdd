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

import "math"

func CalculateBonus(sales, quota, commission, tax int) float64 {
	taxPercentage := percentage(100-tax)
	return bonusWithCommission(commission, sales, quota) * taxPercentage
}

func bonusWithCommission(commission int, sales int, quota int) float64 {
	commissionPercentage := percentage(commission)
	bonus := baseBonus(sales, quota)
	bonusWithCommission := bonus * commissionPercentage
	return bonusWithCommission
}


func baseBonus(sales int, quota int) float64 {
	return math.Max(float64(sales-quota), 0)
}

func percentage(value int) float64 {
	return float64(value) / 100.0
}

func CalculateTeamBonus(sales, quota, members, commission int) float64 {
	if members <= 0 {
		return 0
	}
	return bonusWithCommission( commission, sales, quota )/ float64(members)
}