package period2019

import (
	"github.com/hravemzdy/golegalios/internal/providers/period2018"
)

const (
	HEALTH_VERSION_CODE      int16 = 2019
	HEALTH_MIN_MONTHLY_BASIS       = SALARY_MIN_MONTHLY_WAGE
	HEALTH_MAX_ANNUALS_BASIS       = period2018.HEALTH_MAX_ANNUALS_BASIS
	HEALTH_LIM_MONTHLY_STATE       = period2018.HEALTH_LIM_MONTHLY_STATE
	HEALTH_LIM_MONTHLY_DIS50       = 7540
	HEALTH_FACTOR_COMPOUND         = period2018.HEALTH_FACTOR_COMPOUND
	HEALTH_FACTOR_EMPLOYEE         = period2018.HEALTH_FACTOR_EMPLOYEE
	HEALTH_MARGIN_INCOME_EMP       = 3000
	HEALTH_MARGIN_INCOME_AGR       = period2018.HEALTH_MARGIN_INCOME_AGR
)
