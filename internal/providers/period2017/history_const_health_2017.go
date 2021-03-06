package period2017

import (
	"github.com/hravemzdy/golegalios/internal/providers/period2016"
)

const (
	HEALTH_VERSION_CODE      int16 = 2017
	HEALTH_MIN_MONTHLY_BASIS       = SALARY_MIN_MONTHLY_WAGE
	HEALTH_MAX_ANNUALS_BASIS       = period2016.HEALTH_MAX_ANNUALS_BASIS
	HEALTH_LIM_MONTHLY_STATE       = period2016.HEALTH_LIM_MONTHLY_STATE
	HEALTH_LIM_MONTHLY_DIS50       = 6814
	HEALTH_FACTOR_COMPOUND         = period2016.HEALTH_FACTOR_COMPOUND
	HEALTH_FACTOR_EMPLOYEE         = period2016.HEALTH_FACTOR_EMPLOYEE
	HEALTH_MARGIN_INCOME_EMP       = period2016.HEALTH_MARGIN_INCOME_EMP
	HEALTH_MARGIN_INCOME_AGR       = period2016.HEALTH_MARGIN_INCOME_AGR
)
