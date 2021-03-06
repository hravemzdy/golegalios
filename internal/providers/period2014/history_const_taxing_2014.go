package period2014

import (
	"github.com/hravemzdy/golegalios/internal/providers/period2013"
)

const (
	TAXING_VERSION_CODE              int16   = 2014
	TAXING_ALLOWANCE_PAYER                   = period2013.TAXING_ALLOWANCE_PAYER
	TAXING_ALLOWANCE_DISAB_1ST               = period2013.TAXING_ALLOWANCE_DISAB_1ST
	TAXING_ALLOWANCE_DISAB_2ND               = period2013.TAXING_ALLOWANCE_DISAB_2ND
	TAXING_ALLOWANCE_DISAB_3RD               = period2013.TAXING_ALLOWANCE_DISAB_3RD
	TAXING_ALLOWANCE_STUDY                   = period2013.TAXING_ALLOWANCE_STUDY
	TAXING_ALLOWANCE_CHILD_1ST               = period2013.TAXING_ALLOWANCE_CHILD_1ST
	TAXING_ALLOWANCE_CHILD_2ND               = period2013.TAXING_ALLOWANCE_CHILD_2ND
	TAXING_ALLOWANCE_CHILD_3RD               = period2013.TAXING_ALLOWANCE_CHILD_3RD
	TAXING_FACTOR_ADVANCES           float64 = period2013.TAXING_FACTOR_ADVANCES
	TAXING_FACTOR_WITHHOLD           float64 = period2013.TAXING_FACTOR_WITHHOLD
	TAXING_FACTOR_SOLIDARY           float64 = period2013.TAXING_FACTOR_SOLIDARY
	TAXING_FACTOR_TAXRATE2           float64 = period2013.TAXING_FACTOR_TAXRATE2
	TAXING_MIN_AMOUNT_OF_TAXBONUS            = period2013.TAXING_MIN_AMOUNT_OF_TAXBONUS
	TAXING_MAX_AMOUNT_OF_TAXBONUS            = period2013.TAXING_MAX_AMOUNT_OF_TAXBONUS
	TAXING_MARGIN_INCOME_OF_TAXBONUS         = (SALARY_MIN_MONTHLY_WAGE / 2)
	TAXING_MARGIN_INCOME_OF_ROUNDING         = period2013.TAXING_MARGIN_INCOME_OF_ROUNDING
	TAXING_MARGIN_INCOME_OF_WITHHOLD         = 0
	TAXING_MARGIN_INCOME_OF_SOLIDARY         = (4 * 25942)
	TAXING_MARGIN_INCOME_OF_TAXRATE2         = period2013.TAXING_MARGIN_INCOME_OF_TAXRATE2
	TAXING_MARGIN_INCOME_OF_WHT_EMP          = period2013.TAXING_MARGIN_INCOME_OF_WHT_EMP
	TAXING_MARGIN_INCOME_OF_WHT_AGR          = 10000
)
