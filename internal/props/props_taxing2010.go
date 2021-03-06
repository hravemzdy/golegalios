package props

import (
	"github.com/hravemzdy/golegalios/internal/types"
	. "github.com/shopspring/decimal"
)

type PropsTaxing2010 struct {
	propsTaxingBase
}

func (p PropsTaxing2010) AllowancePayer() int32 {
	return p.propsTaxingBase.AllowancePayer()
}

func (p PropsTaxing2010) AllowanceDisab1st() int32 {
	return p.propsTaxingBase.AllowanceDisab1st()
}

func (p PropsTaxing2010) AllowanceDisab2nd() int32 {
	return p.propsTaxingBase.AllowanceDisab2nd()
}

func (p PropsTaxing2010) AllowanceDisab3rd() int32 {
	return p.propsTaxingBase.AllowanceDisab3rd()
}

func (p PropsTaxing2010) AllowanceStudy() int32 {
	return p.propsTaxingBase.AllowanceStudy()
}

func (p PropsTaxing2010) AllowanceChild1st() int32 {
	return p.propsTaxingBase.AllowanceChild1st()
}

func (p PropsTaxing2010) AllowanceChild2nd() int32 {
	return p.propsTaxingBase.AllowanceChild2nd()
}

func (p PropsTaxing2010) AllowanceChild3rd() int32 {
	return p.propsTaxingBase.AllowanceChild3rd()
}

func (p PropsTaxing2010) FactorAdvances() Decimal {
	return p.propsTaxingBase.FactorAdvances()
}

func (p PropsTaxing2010) FactorWithhold() Decimal {
	return p.propsTaxingBase.FactorWithhold()
}

func (p PropsTaxing2010) FactorSolidary() Decimal {
	return p.propsTaxingBase.FactorSolidary()
}

func (p PropsTaxing2010) FactorTaxRate2() Decimal {
	return p.propsTaxingBase.FactorTaxRate2()
}

func (p PropsTaxing2010) MinAmountOfTaxBonus() int32 {
	return p.propsTaxingBase.MinAmountOfTaxBonus()
}

func (p PropsTaxing2010) MaxAmountOfTaxBonus() int32 {
	return p.propsTaxingBase.MaxAmountOfTaxBonus()
}

func (p PropsTaxing2010) MarginIncomeOfTaxBonus() int32 {
	return p.propsTaxingBase.MarginIncomeOfTaxBonus()
}

func (p PropsTaxing2010) MarginIncomeOfRounding() int32 {
	return p.propsTaxingBase.MarginIncomeOfRounding()
}

func (p PropsTaxing2010) MarginIncomeOfWithhold() int32 {
	return p.propsTaxingBase.MarginIncomeOfWithhold()
}

func (p PropsTaxing2010) MarginIncomeOfSolidary() int32 {
	return p.propsTaxingBase.MarginIncomeOfSolidary()
}

func (p PropsTaxing2010) MarginIncomeOfTaxRate2() int32 {
	return p.propsTaxingBase.MarginIncomeOfTaxRate2()
}

func (p PropsTaxing2010) MarginIncomeOfWthEmp() int32 {
	return p.propsTaxingBase.MarginIncomeOfWthEmp()
}

func (p PropsTaxing2010) MarginIncomeOfWthAgr() int32 {
	return p.propsTaxingBase.MarginIncomeOfWthAgr()
}

func (p PropsTaxing2010) ValueEquals(otherTaxing IPropsTaxing) bool {
	return p.propsTaxingBase.ValueEquals(otherTaxing)
}

func (p PropsTaxing2010) HasWithholdIncome(termOpt types.WorkTaxingTerms, signOpt types.TaxDeclSignOption, noneOpt types.TaxNoneSignOption, incomeSum int32) bool {
	//*****************************************************************************
	// Tax income for advance from Year 2008 to Year 2013
	//*****************************************************************************
	// - withhold tax (non-signed declaration) and income is less than X CZK
	//*****************************************************************************

	var withholdIncome bool = false
	if signOpt != types.DECL_TAX_NO_SIGNED {
		return false
	}
	if noneOpt != types.NOSIGN_TAX_WITHHOLD {
		return false
	}
	if p.MarginIncomeOfWithhold() == 0 || incomeSum <= p.MarginIncomeOfWithhold() {
		if incomeSum > 0 {
			withholdIncome = true
		}
	}
	return withholdIncome
}

func (p PropsTaxing2010) BenefitAllowancePayer(signOpts types.TaxDeclSignOption, benefitOpts types.TaxDeclBenfOption) int32 {
	return p.propsTaxingBase.BenefitAllowancePayer(signOpts, benefitOpts)
}

func (p PropsTaxing2010) BenefitAllowanceDisab(signOpts types.TaxDeclSignOption, benefitOpts types.TaxDeclDisabOption) int32 {
	return p.propsTaxingBase.BenefitAllowanceDisab(signOpts, benefitOpts)
}

func (p PropsTaxing2010) BenefitAllowanceStudy(signOpts types.TaxDeclSignOption, benefitOpts types.TaxDeclBenfOption) int32 {
	return p.propsTaxingBase.BenefitAllowanceStudy(signOpts, benefitOpts)
}

func (p PropsTaxing2010) BenefitAllowanceChild(signOpts types.TaxDeclSignOption, benefitOpts types.TaxDeclBenfOption, benefitOrds int32, disabelOpts int32) int32 {
	return p.propsTaxingBase.BenefitAllowanceChild(signOpts, benefitOpts, benefitOrds, disabelOpts)
}

func (p PropsTaxing2010) BonusChildRaw(income int32, benefit int32, rebated int32) int32 {
	return p.propsTaxingBase.BonusChildRaw(income, benefit, rebated)
}

func (p PropsTaxing2010) BonusChildFix(income int32, benefit int32, rebated int32) int32 {
	return p.propsTaxingBase.BonusChildFix(income, benefit, rebated)
}

func (p PropsTaxing2010) TaxableIncomeSupers(incomeResult int32, healthResult int32, socialResult int32) int32 {
	return p.propsTaxingBase.TaxableIncomeSupers(incomeResult, healthResult, socialResult)
}

func (p PropsTaxing2010) TaxableIncomeBasis(incomeResult int32) int32 {
	return p.propsTaxingBase.TaxableIncomeBasis(incomeResult)
}

func (p PropsTaxing2010) RoundedRawBaseAdvances(incomeResult int32) int32 {
	return p.propsTaxingBase.RoundedRawBaseAdvances(incomeResult)
}

func (p PropsTaxing2010) RoundedBaseAdvances(incomeResult int32, healthResult int32, socialResult int32) int32 {
	return p.propsTaxingBase.RoundedBaseAdvances(incomeResult, healthResult, socialResult)
}

func (p PropsTaxing2010) RoundedBaseSolidary(incomeResult int32) int32 {
	return p.propsTaxingBase.RoundedBaseSolidary(incomeResult)
}

func (p PropsTaxing2010) RoundedAdvancesPaym(supersResult int32, basisResult int32) int32 {
	factorAdvances := types.Divide(p.FactorAdvances(), NewFromInt32(100))

	var advanceTaxing int32 = 0
	if basisResult <= p.MarginIncomeOfRounding() {
		advanceTaxing = p.propsTaxingBase.intTaxRoundUp(types.Multiply(NewFromInt32(supersResult), factorAdvances))
		return advanceTaxing
	}
	advanceTaxing = p.propsTaxingBase.intTaxRoundUp(types.Multiply(NewFromInt32(supersResult), factorAdvances))
	return advanceTaxing
}

func (p PropsTaxing2010) RoundedSolidaryPaym(basisResult int32) int32 {
	return p.propsTaxingBase.RoundedSolidaryPaym(basisResult)
}

func (p PropsTaxing2010) RoundedBaseWithhold(incomeResult int32) int32 {
	return p.propsTaxingBase.RoundedBaseWithhold(incomeResult)
}

func (p PropsTaxing2010) RoundedWithholdPaym(supersResult int32, basisResult int32) int32 {
	return p.propsTaxingBase.RoundedWithholdPaym(supersResult, basisResult)
}

func NewPropsTaxing2010(versionId types.IVersionId,
	allowancePayer int32,
	allowanceDisab1st int32,
	allowanceDisab2nd int32,
	allowanceDisab3rd int32,
	allowanceStudy int32,
	allowanceChild1st int32,
	allowanceChild2nd int32,
	allowanceChild3rd int32,
	factorAdvances Decimal,
	factorWithhold Decimal,
	factorSolidary Decimal,
	factorTaxRate2 Decimal,
	minAmountOfTaxBonus int32,
	maxAmountOfTaxBonus int32,
	marginIncomeOfTaxBonus int32,
	marginIncomeOfRounding int32,
	marginIncomeOfWithhold int32,
	marginIncomeOfSolidary int32,
	marginIncomeOfTaxRate2 int32,
	marginIncomeOfWthEmp int32,
	marginIncomeOfWthAgr int32) IPropsTaxing {
	return PropsTaxing2010{
		propsTaxingBase: propsTaxingBase{
			propsBase:              propsBase{Version: versionId},
			allowancePayer:         allowancePayer,
			allowanceDisab1st:      allowanceDisab1st,
			allowanceDisab2nd:      allowanceDisab2nd,
			allowanceDisab3rd:      allowanceDisab3rd,
			allowanceStudy:         allowanceStudy,
			allowanceChild1st:      allowanceChild1st,
			allowanceChild2nd:      allowanceChild2nd,
			allowanceChild3rd:      allowanceChild3rd,
			factorAdvances:         factorAdvances,
			factorWithhold:         factorWithhold,
			factorSolidary:         factorSolidary,
			factorTaxRate2:         factorTaxRate2,
			minAmountOfTaxBonus:    minAmountOfTaxBonus,
			maxAmountOfTaxBonus:    maxAmountOfTaxBonus,
			marginIncomeOfTaxBonus: marginIncomeOfTaxBonus,
			marginIncomeOfRounding: marginIncomeOfRounding,
			marginIncomeOfWithhold: marginIncomeOfWithhold,
			marginIncomeOfSolidary: marginIncomeOfSolidary,
			marginIncomeOfTaxRate2: marginIncomeOfTaxRate2,
			marginIncomeOfWthEmp:   marginIncomeOfWthEmp,
			marginIncomeOfWthAgr:   marginIncomeOfWthAgr,
		},
	}
}

func EmptyPropsTaxing2010() IPropsTaxing {
	return PropsTaxing2010{
		propsTaxingBase: propsTaxingBase{
			propsBase:              propsBase{Version: types.GetVersionId(types.VERSION_ZERO)},
			allowancePayer:         0,
			allowanceDisab1st:      0,
			allowanceDisab2nd:      0,
			allowanceDisab3rd:      0,
			allowanceStudy:         0,
			allowanceChild1st:      0,
			allowanceChild2nd:      0,
			allowanceChild3rd:      0,
			factorAdvances:         NewFromFloat(0),
			factorWithhold:         NewFromFloat(0),
			factorSolidary:         NewFromFloat(0),
			factorTaxRate2:         NewFromFloat(0),
			minAmountOfTaxBonus:    0,
			maxAmountOfTaxBonus:    0,
			marginIncomeOfTaxBonus: 0,
			marginIncomeOfRounding: 0,
			marginIncomeOfWithhold: 0,
			marginIncomeOfSolidary: 0,
			marginIncomeOfTaxRate2: 0,
			marginIncomeOfWthEmp:   0,
			marginIncomeOfWthAgr:   0,
		},
	}
}
