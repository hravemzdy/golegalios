package period2013

import (
	"github.com/hravemzdy/golegalios/internal/providers/period2012"
)

const (
	HEALTH_VERSION_CODE      int16 = 2013
	HEALTH_MIN_MONTHLY_BASIS       = SALARY_MIN_MONTHLY_WAGE
	HEALTH_MAX_ANNUALS_BASIS       = 0
	HEALTH_LIM_MONTHLY_STATE       = period2012.HEALTH_LIM_MONTHLY_STATE
	HEALTH_LIM_MONTHLY_DIS50       = period2012.HEALTH_LIM_MONTHLY_DIS50
	HEALTH_FACTOR_COMPOUND         = period2012.HEALTH_FACTOR_COMPOUND
	HEALTH_FACTOR_EMPLOYEE         = period2012.HEALTH_FACTOR_EMPLOYEE
	HEALTH_MARGIN_INCOME_EMP       = period2012.HEALTH_MARGIN_INCOME_EMP
	HEALTH_MARGIN_INCOME_AGR       = period2012.HEALTH_MARGIN_INCOME_AGR

	VAR11_HEALTH_LIM_MONTHLY_DIS50 = 5829

	VAR08_HEALTH_MIN_MONTHLY_BASIS = VAR08_SALARY_MIN_MONTHLY_WAGE
)
