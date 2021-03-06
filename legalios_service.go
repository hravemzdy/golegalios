package golegalios

import (
	"github.com/hravemzdy/golegalios/internal/factories"
	"github.com/hravemzdy/golegalios/internal/props"
	"github.com/hravemzdy/golegalios/internal/types"
)

type IPeriod = types.IPeriod
type IVersionId = types.IVersionId
type IBundleProps = factories.IBundleProps

func NewVersionId() IVersionId {
	return types.NewVersionId()
}

func GetVersionId(id int16) IVersionId {
	return types.GetVersionId(id)
}

func PeriodZero() IPeriod {
	return types.PeriodZero()
}

func NewPeriodWithCode(code int32) IPeriod {
	return types.GetPeriodWithCode(code)
}

func NewPeriodWithYearMonth(year int16, month int16) IPeriod {
	return types.GetPeriodWithYearMonth(year, month)
}

func NewBundleProps(period types.IPeriod,
	bundleSalary props.IPropsSalary,
	bundleHealth props.IPropsHealth,
	bundleSocial props.IPropsSocial,
	bundleTaxing props.IPropsTaxing) IBundleProps {
	return factories.NewBundleProps(period, bundleSalary, bundleHealth, bundleSocial, bundleTaxing)
}

func EmptyBundleProps(period IPeriod) IBundleProps {
	return factories.EmptyBundleProps(period)
}

type IServiceLegalios interface {
	GetBundle(period types.IPeriod) (provider factories.IBundleProps, err error)
}

type legaliosService struct {
	builder factories.IBundleBuilder
}

func (s legaliosService) GetBundle(period types.IPeriod) (provider factories.IBundleProps, err error) {
	bundle, exists := s.builder.GetBundle(period)
	if exists == false {
		err = newBundleNoneError()
		return nil, err
	}
	if bundle == nil {
		err = newBundleNullError()
		return nil, err
	}
	if bundle.GetPeriodProps().GetCode() == 0 {
		err = newBundleEmptyError()
		return nil, err
	}
	if bundle.GetPeriodProps().GetCode() != period.GetCode() {
		err = newBundleInvalidError()
		return nil, err
	}
	return bundle, nil
}

func NewLegaliosService() IServiceLegalios {
	return legaliosService{builder: factories.NewBundleBuilder()}
}
