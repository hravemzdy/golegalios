package period2022

import (
	"github.com/hravemzdy/golegalios/internal/props"
	"github.com/hravemzdy/golegalios/internal/providers"
	"github.com/hravemzdy/golegalios/internal/types"
	. "github.com/shopspring/decimal"
)

type providerTaxing2022 struct {
	providers.ProviderBase
}

func NewProviderTaxing2022() providers.IProviderTaxing {
	return providerTaxing2022{
		ProviderBase: providers.ProviderBase{Version: types.GetVersionId(2022)},
	}
}

func (b providerTaxing2022) GetProps(period types.IPeriod) props.IPropsTaxing {
	return props.NewPropsTaxing(b.Version,
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

func (b providerTaxing2022) AllowancePayer(period types.IPeriod) int32 {
	return TAXING_ALLOWANCE_PAYER
}

func (b providerTaxing2022) AllowanceDisab1st(period types.IPeriod) int32 {
	return TAXING_ALLOWANCE_DISAB_1ST
}

func (b providerTaxing2022) AllowanceDisab2nd(period types.IPeriod) int32 {
	return TAXING_ALLOWANCE_DISAB_2ND
}

func (b providerTaxing2022) AllowanceDisab3rd(period types.IPeriod) int32 {
	return TAXING_ALLOWANCE_DISAB_3RD
}

func (b providerTaxing2022) AllowanceStudy(period types.IPeriod) int32 {
	return TAXING_ALLOWANCE_STUDY
}

func (b providerTaxing2022) AllowanceChild1st(period types.IPeriod) int32 {
	return TAXING_ALLOWANCE_CHILD_1ST
}

func (b providerTaxing2022) AllowanceChild2nd(period types.IPeriod) int32 {
	return TAXING_ALLOWANCE_CHILD_2ND
}

func (b providerTaxing2022) AllowanceChild3rd(period types.IPeriod) int32 {
	return TAXING_ALLOWANCE_CHILD_3RD
}

func (b providerTaxing2022) FactorAdvances(period types.IPeriod) Decimal {
	return NewFromFloat(TAXING_FACTOR_ADVANCES)
}

func (b providerTaxing2022) FactorWithhold(period types.IPeriod) Decimal {
	return NewFromFloat(TAXING_FACTOR_WITHHOLD)
}

func (b providerTaxing2022) FactorSolidary(period types.IPeriod) Decimal {
	return NewFromFloat(TAXING_FACTOR_SOLIDARY)
}

func (b providerTaxing2022) FactorTaxRate2(period types.IPeriod) Decimal {
	return NewFromFloat(TAXING_FACTOR_TAXRATE2)
}

func (b providerTaxing2022) MinAmountOfTaxBonus(period types.IPeriod) int32 {
	return TAXING_MIN_AMOUNT_OF_TAXBONUS
}

func (b providerTaxing2022) MaxAmountOfTaxBonus(period types.IPeriod) int32 {
	return TAXING_MAX_AMOUNT_OF_TAXBONUS
}

func (b providerTaxing2022) MarginIncomeOfTaxBonus(period types.IPeriod) int32 {
	return TAXING_MARGIN_INCOME_OF_TAXBONUS
}

func (b providerTaxing2022) MarginIncomeOfRounding(period types.IPeriod) int32 {
	return TAXING_MARGIN_INCOME_OF_ROUNDING
}

func (b providerTaxing2022) MarginIncomeOfWithhold(period types.IPeriod) int32 {
	return TAXING_MARGIN_INCOME_OF_WITHHOLD
}

func (b providerTaxing2022) MarginIncomeOfSolidary(period types.IPeriod) int32 {
	return TAXING_MARGIN_INCOME_OF_SOLIDARY
}

func (b providerTaxing2022) MarginIncomeOfTaxRate2(period types.IPeriod) int32 {
	return TAXING_MARGIN_INCOME_OF_TAXRATE2
}

func (b providerTaxing2022) MarginIncomeOfWthEmp(period types.IPeriod) int32 {
	return TAXING_MARGIN_INCOME_OF_WHT_EMP
}

func (b providerTaxing2022) MarginIncomeOfWthAgr(period types.IPeriod) int32 {
	return TAXING_MARGIN_INCOME_OF_WHT_AGR
}
