package period2017

import (
	"github.com/mzdyhrave/legaliosgo/internal/props"
	"github.com/mzdyhrave/legaliosgo/internal/providers"
	"github.com/mzdyhrave/legaliosgo/internal/types"
	. "github.com/shopspring/decimal"
)

type providerTaxing2017 struct {
	providers.ProviderBase
}

func NewProviderTaxing2017() providers.IProviderTaxing {
	return providerTaxing2017{
		ProviderBase: providers.ProviderBase{Version: types.GetVersionId(2017)},
	}
}

func (b providerTaxing2017) GetProps(period types.IPeriod) props.IPropsTaxing {
	return props.NewPropsTaxing2014(b.Version,
		b.AllowancePayer(period),
		b.AllowanceDisab1st(period),
		b.AllowanceDisab2nd(period),
		b.AllowanceDisab3rd(period),
		b.AllowanceStudy(period),
		b.AllowanceChild1st(period),
		b.AllowanceChild2nd(period),
		b.AllowanceChild3rd(period),
		b.FactorAdvances(period),
		b.FactorWithhold(period),
		b.FactorSolidary(period),
		b.FactorTaxRate2(period),
		b.MinAmountOfTaxBonus(period),
		b.MaxAmountOfTaxBonus(period),
		b.MarginIncomeOfTaxBonus(period),
		b.MarginIncomeOfRounding(period),
		b.MarginIncomeOfWithhold(period),
		b.MarginIncomeOfSolidary(period),
		b.MarginIncomeOfTaxRate2(period),
		b.MarginIncomeOfWthEmp(period),
		b.MarginIncomeOfWthAgr(period))
}

func (b providerTaxing2017) AllowancePayer(period types.IPeriod) int32 {
	return TAXING_ALLOWANCE_PAYER
}

func (b providerTaxing2017) AllowanceDisab1st(period types.IPeriod) int32 {
	return TAXING_ALLOWANCE_DISAB_1ST
}

func (b providerTaxing2017) AllowanceDisab2nd(period types.IPeriod) int32 {
	return TAXING_ALLOWANCE_DISAB_2ND
}

func (b providerTaxing2017) AllowanceDisab3rd(period types.IPeriod) int32 {
	return TAXING_ALLOWANCE_DISAB_3RD
}

func (b providerTaxing2017) AllowanceStudy(period types.IPeriod) int32 {
	return TAXING_ALLOWANCE_STUDY
}

func (b providerTaxing2017) AllowanceChild1st(period types.IPeriod) int32 {
	return TAXING_ALLOWANCE_CHILD_1ST
}

func (b providerTaxing2017) AllowanceChild2nd(period types.IPeriod) int32 {
	if b.ProviderBase.IsPeriodGreaterOrEqualThan(period, 2017, 7) {
		return VAR07_TAXING_ALLOWANCE_CHILD_2ND
	}
	return TAXING_ALLOWANCE_CHILD_2ND
}

func (b providerTaxing2017) AllowanceChild3rd(period types.IPeriod) int32 {
	if b.ProviderBase.IsPeriodGreaterOrEqualThan(period, 2017, 7) {
		return VAR07_TAXING_ALLOWANCE_CHILD_3RD
	}
	return TAXING_ALLOWANCE_CHILD_3RD
}

func (b providerTaxing2017) FactorAdvances(period types.IPeriod) Decimal {
	return NewFromFloat(TAXING_FACTOR_ADVANCES)
}

func (b providerTaxing2017) FactorWithhold(period types.IPeriod) Decimal {
	return NewFromFloat(TAXING_FACTOR_WITHHOLD)
}

func (b providerTaxing2017) FactorSolidary(period types.IPeriod) Decimal {
	return NewFromFloat(TAXING_FACTOR_SOLIDARY)
}

func (b providerTaxing2017) FactorTaxRate2(period types.IPeriod) Decimal {
	return NewFromFloat(TAXING_FACTOR_TAXRATE2)
}

func (b providerTaxing2017) MinAmountOfTaxBonus(period types.IPeriod) int32 {
	return TAXING_MIN_AMOUNT_OF_TAXBONUS
}

func (b providerTaxing2017) MaxAmountOfTaxBonus(period types.IPeriod) int32 {
	return TAXING_MAX_AMOUNT_OF_TAXBONUS
}

func (b providerTaxing2017) MarginIncomeOfTaxBonus(period types.IPeriod) int32 {
	return TAXING_MARGIN_INCOME_OF_TAXBONUS
}

func (b providerTaxing2017) MarginIncomeOfRounding(period types.IPeriod) int32 {
	return TAXING_MARGIN_INCOME_OF_ROUNDING
}

func (b providerTaxing2017) MarginIncomeOfWithhold(period types.IPeriod) int32 {
	return TAXING_MARGIN_INCOME_OF_WITHHOLD
}

func (b providerTaxing2017) MarginIncomeOfSolidary(period types.IPeriod) int32 {
	return TAXING_MARGIN_INCOME_OF_SOLIDARY
}

func (b providerTaxing2017) MarginIncomeOfTaxRate2(period types.IPeriod) int32 {
	return TAXING_MARGIN_INCOME_OF_TAXRATE2
}

func (b providerTaxing2017) MarginIncomeOfWthEmp(period types.IPeriod) int32 {
	return TAXING_MARGIN_INCOME_OF_WHT_EMP
}

func (b providerTaxing2017) MarginIncomeOfWthAgr(period types.IPeriod) int32 {
	return TAXING_MARGIN_INCOME_OF_WHT_AGR
}

