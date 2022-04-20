//go:build protokolFile
// +build protokolFile

package protokol

import (
	"github.com/hravemzdy/golegalios/internal/factories"
	"github.com/hravemzdy/golegalios/internal/props"
	"github.com/hravemzdy/golegalios/internal/types"
	. "github.com/shopspring/decimal"
	"os"
	"testing"
)

type propsSalaryIntFunc func(props props.IPropsSalary) int32
type propsSalaryDecFunc func(props props.IPropsSalary) Decimal

func exportSalaryPropsIntLine(protokol *os.File, testYear int16, sut factories.IFactorySalary, function propsSalaryIntFunc) {
	exportPropsYear(protokol, testYear)

	var testMonth int16
	for testMonth = 1; testMonth <= 12; testMonth++ {
		testPeriod := types.GetPeriodWithYearMonth(testYear, testMonth)
		testResult, exists := sut.GetProps(testPeriod)
		var testValue int32 = 0
		if exists && testResult != nil {
			testValue = function(testResult)
		}
		exportPropsIntValue(protokol, testValue)
	}
	exportPropsEnd(protokol)
}

func exportSalaryPropsDecLine(protokol *os.File, testYear int16, sut factories.IFactorySalary, function propsSalaryDecFunc) {
	exportPropsYear(protokol, testYear)

	var testMonth int16
	for testMonth = 1; testMonth <= 12; testMonth++ {
		testPeriod := types.GetPeriodWithYearMonth(testYear, testMonth)
		testResult, exists := sut.GetProps(testPeriod)
		var testValue Decimal = NewFromFloat(0)
		if exists && testResult != nil {
			testValue = function(testResult)
		}
		exportPropsDecValue(protokol, testValue)
	}
	exportPropsEnd(protokol)
}

func exportSalaryPropsIntFile(t *testing.T, fileName string, minYear int16, maxYear int16, sut factories.IFactorySalary, function propsSalaryIntFunc) {
	testProtokol, err := createProtokolFile(fileName)

	if err != nil {
		t.Errorf("Error creating file %s", err)
		return
	}

	defer testProtokol.Close()

	exportPropsStart(testProtokol)

	for testYear := minYear; testYear <= maxYear; testYear++ {
		exportSalaryPropsIntLine(testProtokol, testYear, sut, function)
	}
}

func exportSalaryPropsDecFile(t *testing.T, fileName string, minYear int16, maxYear int16, sut factories.IFactorySalary, function propsSalaryDecFunc) {
	testProtokol, err := createProtokolFile(fileName)

	if err != nil {
		t.Errorf("Error creating file %s", err)
		return
	}

	defer testProtokol.Close()

	exportPropsStart(testProtokol)

	for testYear := minYear; testYear <= maxYear; testYear++ {
		exportSalaryPropsDecLine(testProtokol, testYear, sut, function)
	}
}
