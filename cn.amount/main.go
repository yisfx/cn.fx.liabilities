package main

import (
	"fmt"
	"time"

	"cn.amount/utils"
)

func main() {
	layout := "2006-01-02 15:04:05"

	zz := &utils.BaseInfo{}
	zz.Rate = 0.033

	timeStr := "2048-07-02 01:00:01"

	EndTime, _ := time.Parse(layout, timeStr)
	zz.EndTime = EndTime
	zz.TotalMonth = utils.MonthsBetween(time.Now(), zz.EndTime)

	sh1 := &utils.BaseInfo{}
	sh1.Rate = 0.0285
	sh1.TotalAmount = 635439.56
	timeStr = "2039-08-07 01:00:01"
	EndTime, _ = time.Parse(layout, timeStr)
	sh1.EndTime = EndTime
	sh1.TotalMonth = utils.MonthsBetween(time.Now(), sh1.EndTime)

	sh2 := &utils.BaseInfo{}
	sh2.Rate = 0.0315
	sh2.TotalAmount = 178550.48
	timeStr = "2054-08-06 01:00:01"
	EndTime, _ = time.Parse(layout, timeStr)
	sh2.EndTime = EndTime
	sh2.TotalMonth = utils.MonthsBetween(time.Now(), sh2.EndTime)

	zz.TotalAmount = 5000000.00
	sh1.TotalAmount = 5000000.00
	sh2.TotalAmount = 5000000.00

	currMonth := &utils.MonthDetail{}
	currMonth.ZZ = utils.CalcFirthMonth(zz)
	currMonth.SH1 = utils.CalcFirthMonth(sh1)
	currMonth.SH2 = utils.CalcFirthMonth(sh2)
	currMonth.MonthNumber = 1
	currMonth = utils.CalcSummary(currMonth)

	utils.PrintMonth(currMonth, true)

	for currMonth.MonthNumber < 5*12 {
		if currMonth.MonthNumber%12 == 0 {

			fmt.Println(utils.ColorGreen, currMonth.MonthNumber/12, "次提前还款10万", utils.ColorReset)
			zz.TotalAmount = currMonth.ZZ.LeftAmount - 100000
			zz.TotalMonth = currMonth.ZZ.LeftMonth

			currMonth.ZZ = utils.CalcFirthMonth(zz)
		}

		currMonth.ZZ = utils.CalcMonthlyDetail(currMonth.ZZ)
		currMonth.SH1 = utils.CalcMonthlyDetail(currMonth.SH1)
		currMonth.SH2 = utils.CalcMonthlyDetail(currMonth.SH2)
		currMonth.MonthNumber += 1
		currMonth = utils.CalcSummary(currMonth)

		utils.PrintMonth(currMonth, false)

	}
	fmt.Println(utils.ColorBlue, "当前在第", currMonth.MonthNumber, "月", utils.ColorReset)
	for currMonth.MonthNumber < 6*12 {
		if currMonth.MonthNumber%12 == 0 {

			fmt.Println(utils.ColorGreen, currMonth.MonthNumber/12, "次提前还款", utils.ColorReset)
			sh2.TotalAmount = currMonth.SH2.LeftAmount - 30000
			sh2.TotalMonth = currMonth.SH2.LeftMonth
			currMonth.SH2 = utils.CalcFirthMonth(sh2)

			zz.TotalAmount = currMonth.ZZ.LeftAmount - 70000
			zz.TotalMonth = currMonth.ZZ.LeftMonth

			currMonth.ZZ = utils.CalcFirthMonth(zz)
		}

		currMonth.ZZ = utils.CalcMonthlyDetail(currMonth.ZZ)
		currMonth.SH1 = utils.CalcMonthlyDetail(currMonth.SH1)
		currMonth.SH2 = utils.CalcMonthlyDetail(currMonth.SH2)
		currMonth.MonthNumber += 1
		currMonth = utils.CalcSummary(currMonth)
		utils.PrintMonth(currMonth, false)
	}
	fmt.Println(utils.ColorBlue, "当前在第", currMonth.MonthNumber, "月", utils.ColorReset)
	for currMonth.MonthNumber < 7*12 {
		if currMonth.MonthNumber%12 == 0 {

			fmt.Println(utils.ColorGreen, currMonth.MonthNumber/12, "次提前还款", utils.ColorReset)
			sh2.TotalAmount = currMonth.SH2.LeftAmount - 100000
			sh2.TotalMonth = currMonth.SH2.LeftMonth
			currMonth.SH2 = utils.CalcFirthMonth(sh2)
		}

		currMonth.ZZ = utils.CalcMonthlyDetail(currMonth.ZZ)
		currMonth.SH1 = utils.CalcMonthlyDetail(currMonth.SH1)
		currMonth.SH2 = utils.CalcMonthlyDetail(currMonth.SH2)
		currMonth.MonthNumber += 1
		currMonth = utils.CalcSummary(currMonth)
		utils.PrintMonth(currMonth, false)
	}
	fmt.Println(utils.ColorBlue, "当前在第", currMonth.MonthNumber, "月", utils.ColorReset)
	for currMonth.MonthNumber < 8*12 {
		if currMonth.MonthNumber%12 == 0 {

			fmt.Println(utils.ColorGreen, currMonth.MonthNumber/12, "次提前还款", utils.ColorReset)
			sh2.TotalAmount = currMonth.SH2.LeftAmount - 20000
			sh2.TotalMonth = currMonth.SH2.LeftMonth
			currMonth.SH2 = utils.CalcFirthMonth(sh2)

			sh1.TotalAmount = currMonth.SH1.LeftAmount - 80000
			sh1.TotalMonth = currMonth.SH1.LeftMonth
			currMonth.SH1 = utils.CalcFirthMonth(sh1)
		}

		currMonth.ZZ = utils.CalcMonthlyDetail(currMonth.ZZ)
		currMonth.SH1 = utils.CalcMonthlyDetail(currMonth.SH1)
		currMonth.SH2 = utils.CalcMonthlyDetail(currMonth.SH2)
		currMonth.MonthNumber += 1
		currMonth = utils.CalcSummary(currMonth)
		utils.PrintMonth(currMonth, false)
	}

	fmt.Println(utils.ColorBlue, "当前在第", currMonth.MonthNumber, "月", utils.ColorReset)
	for currMonth.MonthNumber < 10*12 {
		if currMonth.MonthNumber%12 == 0 {

			fmt.Println(utils.ColorGreen, currMonth.MonthNumber/12, "次提前还款", utils.ColorReset)

			sh1.TotalAmount = currMonth.SH1.LeftAmount - 100000
			sh1.TotalMonth = currMonth.SH1.LeftMonth
			currMonth.SH1 = utils.CalcFirthMonth(sh1)
		}

		currMonth.ZZ = utils.CalcMonthlyDetail(currMonth.ZZ)
		currMonth.SH1 = utils.CalcMonthlyDetail(currMonth.SH1)
		currMonth.SH2 = utils.CalcMonthlyDetail(currMonth.SH2)
		currMonth.MonthNumber += 1
		currMonth = utils.CalcSummary(currMonth)
		utils.PrintMonth(currMonth, false)
	}

	return
	for currMonth.MonthNumber < 11*12 {
		if currMonth.MonthNumber%12 == 0 {

			fmt.Println(utils.ColorGreen, currMonth.MonthNumber/12, "次提前还款", utils.ColorReset)

			zz.TotalAmount = currMonth.ZZ.LeftAmount - 6000
			zz.TotalMonth = currMonth.ZZ.LeftMonth
			currMonth.ZZ = utils.CalcFirthMonth(zz)

			sh1.TotalAmount = currMonth.SH1.LeftAmount - 17000
			sh1.TotalMonth = currMonth.SH1.LeftMonth
			currMonth.SH1 = utils.CalcFirthMonth(sh1)

			sh2.TotalAmount = currMonth.SH2.LeftAmount - 3600
			sh2.TotalMonth = currMonth.SH2.LeftMonth
			currMonth.SH2 = utils.CalcFirthMonth(sh2)

		}

		currMonth.ZZ = utils.CalcMonthlyDetail(currMonth.ZZ)
		currMonth.SH1 = utils.CalcMonthlyDetail(currMonth.SH1)
		currMonth.SH2 = utils.CalcMonthlyDetail(currMonth.SH2)
		currMonth.MonthNumber += 1
		currMonth = utils.CalcSummary(currMonth)
		utils.PrintMonth(currMonth, false)
	}

}
