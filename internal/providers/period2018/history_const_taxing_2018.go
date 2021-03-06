package period2018

import (
	"github.com/hravemzdy/golegalios/internal/providers/period2017"
)

const (
	TAXING_VERSION_CODE              int16   = 2018
	TAXING_ALLOWANCE_PAYER                   = period2017.TAXING_ALLOWANCE_PAYER
	TAXING_ALLOWANCE_DISAB_1ST               = period2017.TAXING_ALLOWANCE_DISAB_1ST
	TAXING_ALLOWANCE_DISAB_2ND               = period2017.TAXING_ALLOWANCE_DISAB_2ND
	TAXING_ALLOWANCE_DISAB_3RD               = period2017.TAXING_ALLOWANCE_DISAB_3RD
	TAXING_ALLOWANCE_STUDY                   = period2017.TAXING_ALLOWANCE_STUDY
	TAXING_ALLOWANCE_CHILD_1ST               = 1267
	TAXING_ALLOWANCE_CHILD_2ND               = period2017.VAR07_TAXING_ALLOWANCE_CHILD_2ND
	TAXING_ALLOWANCE_CHILD_3RD               = period2017.VAR07_TAXING_ALLOWANCE_CHILD_3RD
	TAXING_FACTOR_ADVANCES           float64 = period2017.TAXING_FACTOR_ADVANCES
	TAXING_FACTOR_WITHHOLD           float64 = period2017.TAXING_FACTOR_WITHHOLD
	TAXING_FACTOR_SOLIDARY           float64 = period2017.TAXING_FACTOR_SOLIDARY
	TAXING_FACTOR_TAXRATE2           float64 = period2017.TAXING_FACTOR_TAXRATE2
	TAXING_MIN_AMOUNT_OF_TAXBONUS            = period2017.TAXING_MIN_AMOUNT_OF_TAXBONUS
	TAXING_MAX_AMOUNT_OF_TAXBONUS            = period2017.TAXING_MAX_AMOUNT_OF_TAXBONUS
	TAXING_MARGIN_INCOME_OF_TAXBONUS         = (SALARY_MIN_MONTHLY_WAGE / 2)
	TAXING_MARGIN_INCOME_OF_ROUNDING         = period2017.TAXING_MARGIN_INCOME_OF_ROUNDING
	TAXING_MARGIN_INCOME_OF_WITHHOLD         = period2017.TAXING_MARGIN_INCOME_OF_WITHHOLD
	TAXING_MARGIN_INCOME_OF_SOLIDARY         = (4 * 29979)
	TAXING_MARGIN_INCOME_OF_TAXRATE2         = period2017.TAXING_MARGIN_INCOME_OF_TAXRATE2
	TAXING_MARGIN_INCOME_OF_WHT_EMP          = 2500
	TAXING_MARGIN_INCOME_OF_WHT_AGR          = period2017.TAXING_MARGIN_INCOME_OF_WHT_AGR
)
