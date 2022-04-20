package providers

import (
	"github.com/hravemzdy/golegalios/internal/props"
	"github.com/hravemzdy/golegalios/internal/types"
)

type IProviderSalary interface {
	ipropsProvider
	GetProps(period types.IPeriod) props.IPropsSalary

	WorkingShiftWeek(period types.IPeriod) int32
	WorkingShiftTime(period types.IPeriod) int32
	MinMonthlyWage(period types.IPeriod) int32
	MinHourlyWage(period types.IPeriod) int32
}
