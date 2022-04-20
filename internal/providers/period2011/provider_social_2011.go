package period2011

import (
	"github.com/hravemzdy/golegalios/internal/props"
	"github.com/hravemzdy/golegalios/internal/providers"
	"github.com/hravemzdy/golegalios/internal/types"
	. "github.com/shopspring/decimal"
)

type providerSocial2011 struct {
	providers.ProviderBase
}

func NewProviderSocial2011() providers.IProviderSocial {
	return providerSocial2011{
		ProviderBase: providers.ProviderBase{Version: types.GetVersionId(2011)},
	}
}

func (b providerSocial2011) GetProps(period types.IPeriod) props.IPropsSocial {
	return props.NewPropsSocial2010(b.Version,
		b.MaxAnnualsBasis(period),
		b.FactorEmployer(period),
		b.FactorEmployerHigher(period),
		b.FactorEmployee(period),
		b.FactorEmployeeGarant(period),
		b.FactorEmployeeReduce(period),
		b.MarginIncomeEmp(period),
		b.MarginIncomeAgr(period))
}

func (b providerSocial2011) MaxAnnualsBasis(period types.IPeriod) int32 {
	return SOCIAL_MAX_ANNUALS_BASIS
}

func (b providerSocial2011) FactorEmployer(period types.IPeriod) Decimal {
	return NewFromFloat(SOCIAL_FACTOR_EMPLOYER)
}

func (b providerSocial2011) FactorEmployerHigher(period types.IPeriod) Decimal {
	return NewFromFloat(SOCIAL_FACTOR_EMPLOYER_HIGHER)
}

func (b providerSocial2011) FactorEmployee(period types.IPeriod) Decimal {
	return NewFromFloat(SOCIAL_FACTOR_EMPLOYEE)
}

func (b providerSocial2011) FactorEmployeeGarant(period types.IPeriod) Decimal {
	return NewFromFloat(SOCIAL_FACTOR_EMPLOYEE_GARANT)
}

func (b providerSocial2011) FactorEmployeeReduce(period types.IPeriod) Decimal {
	return NewFromFloat(SOCIAL_FACTOR_EMPLOYEE_REDUCE)
}

func (b providerSocial2011) MarginIncomeEmp(period types.IPeriod) int32 {
	return SOCIAL_MARGIN_INCOME_EMP
}

func (b providerSocial2011) MarginIncomeAgr(period types.IPeriod) int32 {
	return SOCIAL_MARGIN_INCOME_AGR
}
