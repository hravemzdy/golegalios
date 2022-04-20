package period2019

import (
	"github.com/hravemzdy/golegalios/internal/props"
	"github.com/hravemzdy/golegalios/internal/providers"
	"github.com/hravemzdy/golegalios/internal/types"
	. "github.com/shopspring/decimal"
)

type providerSocial2019 struct {
	providers.ProviderBase
}

func NewProviderSocial2019() providers.IProviderSocial {
	return providerSocial2019{
		ProviderBase: providers.ProviderBase{Version: types.GetVersionId(2019)},
	}
}

func (b providerSocial2019) GetProps(period types.IPeriod) props.IPropsSocial {
	return props.NewPropsSocial(b.Version,
		b.MaxAnnualsBasis(period),
		b.FactorEmployer(period),
		b.FactorEmployerHigher(period),
		b.FactorEmployee(period),
		b.FactorEmployeeGarant(period),
		b.FactorEmployeeReduce(period),
		b.MarginIncomeEmp(period),
		b.MarginIncomeAgr(period))
}

func (b providerSocial2019) MaxAnnualsBasis(period types.IPeriod) int32 {
	return SOCIAL_MAX_ANNUALS_BASIS
}

func (b providerSocial2019) FactorEmployer(period types.IPeriod) Decimal {
	if b.ProviderBase.IsPeriodGreaterOrEqualThan(period, 2019, 7) {
		return NewFromFloat(VAR07_SOCIAL_FACTOR_EMPLOYER)
	}
	return NewFromFloat(SOCIAL_FACTOR_EMPLOYER)
}

func (b providerSocial2019) FactorEmployerHigher(period types.IPeriod) Decimal {
	return NewFromFloat(SOCIAL_FACTOR_EMPLOYER_HIGHER)
}

func (b providerSocial2019) FactorEmployee(period types.IPeriod) Decimal {
	return NewFromFloat(SOCIAL_FACTOR_EMPLOYEE)
}

func (b providerSocial2019) FactorEmployeeGarant(period types.IPeriod) Decimal {
	return NewFromFloat(SOCIAL_FACTOR_EMPLOYEE_GARANT)
}

func (b providerSocial2019) FactorEmployeeReduce(period types.IPeriod) Decimal {
	return NewFromFloat(SOCIAL_FACTOR_EMPLOYEE_REDUCE)
}

func (b providerSocial2019) MarginIncomeEmp(period types.IPeriod) int32 {
	return SOCIAL_MARGIN_INCOME_EMP
}

func (b providerSocial2019) MarginIncomeAgr(period types.IPeriod) int32 {
	return SOCIAL_MARGIN_INCOME_AGR
}
