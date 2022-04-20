package props

import (
	"github.com/hravemzdy/golegalios/internal/types"
	. "github.com/shopspring/decimal"
)

type PropsTaxing struct {
	propsTaxingBase
}

func (p PropsTaxing) AllowancePayer() int32 {
	return p.propsTaxingBase.AllowancePayer()
}

func (p PropsTaxing) AllowanceDisab1st() int32 {
	return p.propsTaxingBase.AllowanceDisab1st()
}

func (p PropsTaxing) AllowanceDisab2nd() int32 {
	return p.propsTaxingBase.AllowanceDisab2nd()
}

func (p PropsTaxing) AllowanceDisab3rd() int32 {
	return p.propsTaxingBase.AllowanceDisab3rd()
}

func (p PropsTaxing) AllowanceStudy() int32 {
	return p.propsTaxingBase.AllowanceStudy()
}

func (p PropsTaxing) AllowanceChild1st() int32 {
	return p.propsTaxingBase.AllowanceChild1st()
}

func (p PropsTaxing) AllowanceChild2nd() int32 {
	return p.propsTaxingBase.AllowanceChild2nd()
}

func (p PropsTaxing) AllowanceChild3rd() int32 {
	return p.propsTaxingBase.AllowanceChild3rd()
}

func (p PropsTaxing) FactorAdvances() Decimal {
	return p.propsTaxingBase.FactorAdvances()
}

func (p PropsTaxing) FactorWithhold() Decimal {
	return p.propsTaxingBase.FactorWithhold()
}

func (p PropsTaxing) FactorSolidary() Decimal {
	return p.propsTaxingBase.FactorSolidary()
}

func (p PropsTaxing) FactorTaxRate2() Decimal {
	return p.propsTaxingBase.FactorTaxRate2()
}

func (p PropsTaxing) MinAmountOfTaxBonus() int32 {
	return p.propsTaxingBase.MinAmountOfTaxBonus()
}

func (p PropsTaxing) MaxAmountOfTaxBonus() int32 {
	return p.propsTaxingBase.MaxAmountOfTaxBonus()
}

func (p PropsTaxing) MarginIncomeOfTaxBonus() int32 {
	return p.propsTaxingBase.MarginIncomeOfTaxBonus()
}

func (p PropsTaxing) MarginIncomeOfRounding() int32 {
	return p.propsTaxingBase.MarginIncomeOfRounding()
}

func (p PropsTaxing) MarginIncomeOfWithhold() int32 {
	return p.propsTaxingBase.MarginIncomeOfWithhold()
}

func (p PropsTaxing) MarginIncomeOfSolidary() int32 {
	return p.propsTaxingBase.MarginIncomeOfSolidary()
}

func (p PropsTaxing) MarginIncomeOfTaxRate2() int32 {
	return p.propsTaxingBase.MarginIncomeOfTaxRate2()
}

func (p PropsTaxing) MarginIncomeOfWthEmp() int32 {
	return p.propsTaxingBase.MarginIncomeOfWthEmp()
}

func (p PropsTaxing) MarginIncomeOfWthAgr() int32 {
	return p.propsTaxingBase.MarginIncomeOfWthAgr()
}

func (p PropsTaxing) ValueEquals(otherTaxing IPropsTaxing) bool {
	return p.propsTaxingBase.ValueEquals(otherTaxing)
}

func (p PropsTaxing) HasWithholdIncome(termOpt types.WorkTaxingTerms, signOpt types.TaxDeclSignOption, noneOpt types.TaxNoneSignOption, incomeSum int32) bool {
	//*****************************************************************************
	// Tax income for advance from Year 2014 to Year 2017
	//*****************************************************************************
	// - withhold tax (non-signed declaration) and income
	// if (period.Year >= 2018)
	// -- income from DPP is less than X CZK
	// -- income from low-income employment is less than X CZK
	// -- income from statutory employment and non-resident is always withhold tax

	var withholdIncome bool = false
	if signOpt != types.DECL_TAX_NO_SIGNED {
		return false
	}
	if noneOpt != types.NOSIGN_TAX_WITHHOLD {
		return false
	}
	if termOpt == types.TAXING_TERM_AGREEM_TASK {
		if p.MarginIncomeOfWthAgr() == 0 || incomeSum <= p.MarginIncomeOfWthAgr() {
			if incomeSum > 0 {
				withholdIncome = true
			}
		}
		return withholdIncome
	}
	if termOpt == types.TAXING_TERM_EMPLOYMENTS {
		if p.MarginIncomeOfWthEmp() == 0 || incomeSum <= p.MarginIncomeOfWthEmp() {
			if incomeSum > 0 {
				withholdIncome = true
			}
		}
		return withholdIncome
	}
	if termOpt == types.TAXING_TERM_STATUT_PART {
		if incomeSum > 0 {
			withholdIncome = true
		}
		return withholdIncome
	}
	return withholdIncome
}

func (p PropsTaxing) BenefitAllowancePayer(signOpts types.TaxDeclSignOption, benefitOpts types.TaxDeclBenfOption) int32 {
	return p.propsTaxingBase.BenefitAllowancePayer(signOpts, benefitOpts)
}

func (p PropsTaxing) BenefitAllowanceDisab(signOpts types.TaxDeclSignOption, benefitOpts types.TaxDeclDisabOption) int32 {
	return p.propsTaxingBase.BenefitAllowanceDisab(signOpts, benefitOpts)
}

func (p PropsTaxing) BenefitAllowanceStudy(signOpts types.TaxDeclSignOption, benefitOpts types.TaxDeclBenfOption) int32 {
	return p.propsTaxingBase.BenefitAllowanceStudy(signOpts, benefitOpts)
}

func (p PropsTaxing) BenefitAllowanceChild(signOpts types.TaxDeclSignOption, benefitOpts types.TaxDeclBenfOption, benefitOrds int32, disabelOpts int32) int32 {
	return p.propsTaxingBase.BenefitAllowanceChild(signOpts, benefitOpts, benefitOrds, disabelOpts)
}

func (p PropsTaxing) BonusChildRaw(income int32, benefit int32, rebated int32) int32 {
	return p.propsTaxingBase.BonusChildRaw(income, benefit, rebated)
}

func (p PropsTaxing) BonusChildFix(income int32, benefit int32, rebated int32) int32 {
	return p.propsTaxingBase.BonusChildFix(income, benefit, rebated)
}

func (p PropsTaxing) TaxableIncomeSupers(incomeResult int32, healthResult int32, socialResult int32) int32 {
	return p.propsTaxingBase.TaxableIncomeSupers(incomeResult, healthResult, socialResult)
}

func (p PropsTaxing) TaxableIncomeBasis(incomeResult int32) int32 {
	return p.propsTaxingBase.TaxableIncomeBasis(incomeResult)
}

func (p PropsTaxing) RoundedRawBaseAdvances(incomeResult int32) int32 {
	return p.propsTaxingBase.RoundedRawBaseAdvances(incomeResult)
}

func (p PropsTaxing) RoundedBaseAdvances(incomeResult int32, healthResult int32, socialResult int32) int32 {
	return p.propsTaxingBase.RoundedBaseAdvances(incomeResult, healthResult, socialResult)
}

func (p PropsTaxing) RoundedBaseSolidary(incomeResult int32) int32 {
	return p.propsTaxingBase.RoundedBaseSolidary(incomeResult)
}

func (p PropsTaxing) RoundedAdvancesPaym(supersResult int32, basisResult int32) int32 {
	factorAdvances := types.Divide(p.FactorAdvances(), NewFromInt32(100))
	factorTaxRate2 := types.Divide(p.FactorTaxRate2(), NewFromInt32(100))

	var taxRate1Basis int32 = basisResult
	var taxRate2Basis int32 = 0
	if p.MarginIncomeOfTaxRate2() != 0 {
		taxRate1Basis = min32(basisResult, p.MarginIncomeOfTaxRate2())
		taxRate2Basis = max32(0, basisResult-p.MarginIncomeOfTaxRate2())
	}
	var taxRate1Taxing Decimal = Zero
	if basisResult <= p.MarginIncomeOfRounding() {
		taxRate1Taxing = types.Multiply(NewFromInt32(taxRate1Basis), factorAdvances)
	} else {
		taxRate1Taxing = types.Multiply(NewFromInt32(taxRate1Basis), factorAdvances)
	}
	var taxRate2Taxing Decimal = Zero
	if p.MarginIncomeOfTaxRate2() != 0 {
		taxRate2Taxing = types.Multiply(NewFromInt32(taxRate2Basis), factorTaxRate2)
	}
	return p.intTaxRoundUp(taxRate1Taxing.Add(taxRate2Taxing))
}

func (p PropsTaxing) RoundedSolidaryPaym(basisResult int32) int32 {
	return p.propsTaxingBase.RoundedSolidaryPaym(basisResult)
}

func (p PropsTaxing) RoundedBaseWithhold(incomeResult int32) int32 {
	return p.propsTaxingBase.RoundedBaseWithhold(incomeResult)
}

func (p PropsTaxing) RoundedWithholdPaym(supersResult int32, basisResult int32) int32 {
	return p.propsTaxingBase.RoundedWithholdPaym(supersResult, basisResult)
}

func NewPropsTaxing(versionId types.IVersionId,
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
	return PropsTaxing{
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

func EmptyPropsTaxing() IPropsTaxing {
	return PropsTaxing{
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
