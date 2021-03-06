package props

import (
	"github.com/hravemzdy/golegalios/internal/types"
	. "github.com/shopspring/decimal"
)

type PropsSocial2010 struct {
	propsSocialBase
}

func (p PropsSocial2010) MaxAnnualsBasis() int32 {
	return p.propsSocialBase.MaxAnnualsBasis()
}

func (p PropsSocial2010) FactorEmployer() Decimal {
	return p.propsSocialBase.factorEmployer
}

func (p PropsSocial2010) FactorEmployerHigher() Decimal {
	return p.propsSocialBase.factorEmployerHigher
}

func (p PropsSocial2010) FactorEmployee() Decimal {
	return p.propsSocialBase.factorEmployee
}

func (p PropsSocial2010) FactorEmployeeGarant() Decimal {
	return p.propsSocialBase.factorEmployeeGarant
}

func (p PropsSocial2010) FactorEmployeeReduce() Decimal {
	return p.propsSocialBase.factorEmployeeReduce
}

func (p PropsSocial2010) MarginIncomeEmp() int32 {
	return p.propsSocialBase.marginIncomeEmp
}

func (p PropsSocial2010) MarginIncomeAgr() int32 {
	return p.propsSocialBase.marginIncomeAgr
}

func (p PropsSocial2010) ValueEquals(otherSocial IPropsSocial) bool {
	return p.propsSocialBase.ValueEquals(otherSocial)
}

func (p PropsSocial2010) HasParticy(term types.WorkSocialTerms, incomeTerm int32, incomeSpec int32) bool {
	return p.propsSocialBase.HasParticyWithAdapters(term, incomeTerm, incomeSpec,
		p.hasTermExemptionParticy,
		p.hasIncomeBasedEmploymentParticy,
		p.hasIncomeBasedAgreementsParticy,
		p.hasIncomeCumulatedParticy)
}

func (p PropsSocial2010) hasTermExemptionParticy(_term types.WorkSocialTerms) bool {
	return false
}
func (p PropsSocial2010) hasIncomeBasedEmploymentParticy(_term types.WorkSocialTerms) bool {
	return _term == types.SOCIAL_TERM_SMALLS_EMPL

}
func (p PropsSocial2010) hasIncomeBasedAgreementsParticy(_term types.WorkSocialTerms) bool {
	return false
}
func (p PropsSocial2010) hasIncomeCumulatedParticy(_term types.WorkSocialTerms) bool {
	var particy bool = false
	switch _term {
	case types.SOCIAL_TERM_EMPLOYMENTS:
		particy = false
	case types.SOCIAL_TERM_AGREEM_TASK:
		particy = false
	case types.SOCIAL_TERM_SMALLS_EMPL:
		particy = false
	case types.SOCIAL_TERM_SHORTS_MEET:
		particy = false
	case types.SOCIAL_TERM_SHORTS_DENY:
		particy = false
	case types.SOCIAL_TERM_BY_CONTRACT:
		particy = false
	default:
		particy = false
	}
	return particy
}

func (p PropsSocial2010) RoundedEmployeePaym(basisResult int32) int32 {
	factorEmployee := types.Divide(p.FactorEmployee(), NewFromInt32(100))
	return p.propsSocialBase.intInsuranceRoundUp(types.Multiply(NewFromInt32(basisResult), factorEmployee))
}

func (p PropsSocial2010) RoundedEmployerPaym(basisResult int32) int32 {
	factorEmployer := types.Divide(p.FactorEmployer(), NewFromInt32(100))
	return p.propsSocialBase.intInsuranceRoundUp(types.Multiply(NewFromInt32(basisResult), factorEmployer))
}

func (p PropsSocial2010) ResultOvercaps(baseSuma int32, overCaps int32) OvercapsResultPair {
	maxBaseEmployee := max32(0, baseSuma-overCaps)
	empBaseOvercaps := max32(0, baseSuma-maxBaseEmployee)
	valBaseOvercaps := max32(0, overCaps-empBaseOvercaps)
	return OvercapsResultPair{maxBaseEmployee, valBaseOvercaps}
}

func (p PropsSocial2010) AnnualsBasisCut(incomeList []ParticySocialTarget, annuityBasis int32) ParticySocialResultTriple {
	return p.propsSocialBase.AnnualsBasisCut(incomeList, annuityBasis)
}

func NewPropsSocial2010(versionId types.IVersionId,
	maxAnnualsBasis int32,
	factorEmployer Decimal,
	factorEmployerHigher Decimal,
	factorEmployee Decimal,
	factorEmployeeGarant Decimal,
	factorEmployeeReduce Decimal,
	marginIncomeEmp int32,
	marginIncomeAgr int32) IPropsSocial {
	return PropsSocial2010{
		propsSocialBase: propsSocialBase{
			propsBase:            propsBase{Version: versionId},
			maxAnnualsBasis:      maxAnnualsBasis,
			factorEmployer:       factorEmployer,
			factorEmployerHigher: factorEmployerHigher,
			factorEmployee:       factorEmployee,
			factorEmployeeGarant: factorEmployeeGarant,
			factorEmployeeReduce: factorEmployeeReduce,
			marginIncomeEmp:      marginIncomeEmp,
			marginIncomeAgr:      marginIncomeAgr,
		},
	}
}

func EmptyPropsSocial2010() IPropsSocial {
	return PropsSocial2010{
		propsSocialBase: propsSocialBase{
			propsBase:            propsBase{Version: types.GetVersionId(types.VERSION_ZERO)},
			maxAnnualsBasis:      0,
			factorEmployer:       NewFromFloat(0),
			factorEmployerHigher: NewFromFloat(0),
			factorEmployee:       NewFromFloat(0),
			factorEmployeeGarant: NewFromFloat(0),
			factorEmployeeReduce: NewFromFloat(0),
			marginIncomeEmp:      0,
			marginIncomeAgr:      0,
		},
	}
}
