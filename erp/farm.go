package erp

import (
	"time"

	"github.com/ahmedsat/erp-reports-cli/utils"
)

type FarmsOptions struct {
	From            time.Time
	To              time.Time
	Fields          utils.List
	IncludeCanceled bool
}

func GetFarms[T any](opt FarmsOptions) (result []T, err error) {

	filters := utils.Filters{
		utils.NewFilter("type", utils.Eq, "farm"),
		utils.NewFilter("creation_date", utils.Gte, opt.From.Format("2006-01-02")),
		// ? to is excluded, so we have to add 1 day
		utils.NewFilter("creation_date", utils.Lte, opt.To.AddDate(0, 0, 1).Format("2006-01-02")),
	}

	if !opt.IncludeCanceled {
		filters = append(filters, utils.NewFilter("farm_status", utils.Neq, "Cancelled"))
	}

	result, err = Get[T]("Farm", "", filters, opt.Fields)
	if err != nil {
		return
	}

	return
}
