package utils

import (
	"fmt"
	"math"
	"time"
)

type MonthDetail struct {
	MonthNumber int

	ZZ  *Detail
	SH1 *Detail
	SH2 *Detail

	CurrentCapitalAmount float64
	Gearing              float64

	LeftTotal float64
	Total     float64
}

type Detail struct {
	LeftAmount float64
	LeftMonth  int

	MonthNumber int

	CurrentCapitalAmount float64

	CurrentPrincipalAmount float64

	CurrentInterestAmount float64
	CurrentRate           float64
}

type BaseInfo struct {
	Rate        float64
	TotalAmount float64
	EndTime     time.Time
	TotalMonth  int
}

const (
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorReset  = "\033[0m"
)

func CalcFirthMonth(baseInfo *BaseInfo) *Detail {
	detail := &Detail{}

	n := baseInfo.TotalMonth
	yr := baseInfo.Rate
	p := baseInfo.TotalAmount

	detail.LeftMonth = n - 1

	detail.MonthNumber = 1

	detail.CurrentRate = yr

	r := yr / 12

	detail.CurrentInterestAmount = p * r

	detail.CurrentCapitalAmount = (p * r * math.Pow((1+r), float64(n))) / (math.Pow(1+r, float64(n)) - 1)

	detail.CurrentPrincipalAmount = detail.CurrentCapitalAmount - detail.CurrentInterestAmount

	detail.LeftAmount = p - detail.CurrentPrincipalAmount

	return detail
}

func CalcMonthlyDetail(curr *Detail) *Detail {

	detail := &Detail{}

	n := curr.LeftMonth
	yr := curr.CurrentRate

	p := curr.LeftAmount

	detail.LeftMonth = n - 1

	detail.MonthNumber = curr.MonthNumber + 1

	detail.CurrentRate = yr

	r := yr / 12

	detail.CurrentInterestAmount = p * r

	detail.CurrentCapitalAmount = curr.CurrentCapitalAmount

	detail.CurrentPrincipalAmount = detail.CurrentCapitalAmount - detail.CurrentInterestAmount

	detail.LeftAmount = p - detail.CurrentPrincipalAmount

	return detail
}

func MonthsBetween(start, end time.Time) int {
	// 确保开始时间早于结束时间
	if start.After(end) {
		start, end = end, start
	}

	// 初始化月差计数器
	months := 0
	current := start

	// 循环直到当前时间超过结束时间
	for !current.After(end) {
		// 每个月加1
		months++
		// 当前时间加一个月
		current = current.AddDate(0, 1, 0)
	}

	return months
}
func ToDecimal(f float64) string {
	return fmt.Sprintf("%.2f", f)
}
func Print(detail *Detail) {
	fmt.Print("\t本金：", ToDecimal(detail.CurrentPrincipalAmount), "\t")
	fmt.Print("\t利息：", ToDecimal(detail.CurrentInterestAmount), "\t")

	fmt.Print("\t月供：", ToDecimal(detail.CurrentCapitalAmount), "\t")
	fmt.Print("\t利率：", ToDecimal(detail.CurrentRate*100), "\t")
	fmt.Print("\t剩余贷款：", ToDecimal(detail.LeftAmount), "\t")
	fmt.Print("\t剩余月数：", detail.LeftMonth, "\t\r\n")

}

func PrintMonth(month *MonthDetail, forcePrint bool) {

	if !forcePrint && month.MonthNumber%12 != 11 && month.MonthNumber%12 != 1 && month.MonthNumber%12 != 0 {
		return
	}

	fmt.Print(ColorRed)

	fmt.Print("第", month.MonthNumber, "个月\t\t")
	fmt.Print("总资产：", ToDecimal(month.Total))
	fmt.Print("\t总负债：", ToDecimal(month.LeftTotal))
	fmt.Print("\t负债率:", ToDecimal(month.Gearing), "%")
	fmt.Print("\t 本月供:", ToDecimal(month.CurrentCapitalAmount), "\r\n")
	fmt.Print(ColorReset)

	fmt.Print("\tZZ:\t")
	Print(month.ZZ)
	fmt.Print("\tSH1:\t")
	Print(month.SH1)
	fmt.Print("\tSH2:\t")
	Print(month.SH2)
}

func CalcSummary(month *MonthDetail) *MonthDetail {
	month.LeftTotal = month.ZZ.LeftAmount + month.SH1.LeftAmount + month.SH2.LeftAmount
	month.CurrentCapitalAmount = month.ZZ.CurrentCapitalAmount + month.SH1.CurrentCapitalAmount + month.SH2.CurrentCapitalAmount
	month.Gearing = month.CurrentCapitalAmount / (3000 + 0) * 100
	month.Total = (10000 + 10000) - month.LeftTotal
	return month
}
