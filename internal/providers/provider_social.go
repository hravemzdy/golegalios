package providers

import (
	"github.com/hravemzdy/golegalios/internal/props"
	"github.com/hravemzdy/golegalios/internal/types"
	. "github.com/shopspring/decimal"
)

type IProviderSocial interface {
	ipropsProvider
	GetProps(period types.IPeriod) props.IPropsSocial

	MaxAnnualsBasis(period types.IPeriod) int32
	FactorEmployer(period types.IPeriod) Decimal
	FactorEmployerHigher(period types.IPeriod) Decimal
	FactorEmployee(period types.IPeriod) Decimal
	FactorEmployeeGarant(period types.IPeriod) Decimal
	FactorEmployeeReduce(period types.IPeriod) Decimal
	MarginIncomeEmp(period types.IPeriod) int32
	MarginIncomeAgr(period types.IPeriod) int32
}
