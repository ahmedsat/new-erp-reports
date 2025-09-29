package erp

import (
	"errors"
	"fmt"

	"github.com/ahmedsat/erp-reports-cli/utils"
)

type FarmApplicantsOptions struct {
	Fields utils.List
}

func GetFarmApplicants[T any](opt FarmApplicantsOptions) (result []T, err error) {

	result, err = Get[T]("/api/resource/Farm Application", utils.Filters{}, opt.Fields)
	if err != nil {
		err = errors.Join(err, fmt.Errorf("%s : failed to prepare request", utils.WhereAmI()))
		return
	}

	return
}
