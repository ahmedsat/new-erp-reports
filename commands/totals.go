package commands

import (
	"fmt"

	"github.com/ahmedsat/erp-reports-cli/erp"
	"github.com/ahmedsat/erp-reports-cli/utils"
)

type TotalsOptions struct {
	GlobalOptions
}

func Totals(opt TotalsOptions) (err error) {

	err = opt.Validate()
	if err != nil {
		return
	}

	type Farm struct {
		Name         string  `json:"name"`
		Region       string  `json:"region"`
		TotalFarmers int     `json:"total_farmers"`
		Area         float64 `json:"farm_area__feddan"`
	}

	farms, err := erp.GetFarms[Farm](erp.FarmsOptions{
		From:            opt.From,
		To:              opt.To,
		Fields:          utils.List{"name", "region", "total_farmers", "farm_area__feddan"},
		IncludeCanceled: opt.IncludeCanceled,
	})
	if err != nil {
		return
	}

	regions := make(map[string]struct {
		count        int
		area         float64
		totalFarmers int
	})
	for _, farm := range farms {
		region := regions[farm.Region]
		region.count++
		region.area += farm.Area
		region.totalFarmers += farm.TotalFarmers
		regions[farm.Region] = region

		total := regions["total"]
		total.count++
		total.area += farm.Area
		total.totalFarmers += farm.TotalFarmers
		regions["total"] = total
	}

	t := utils.TableBase{}

	t.SetHeader("region", "total_farmers", "count of farms", "area")
	for k, region := range regions {
		if k == "total" {
			continue
		}

		t.AppendRow(
			k,
			fmt.Sprintf("%d", region.totalFarmers),
			fmt.Sprintf("%d", region.count),
			fmt.Sprintf("%.2f", region.area),
		)
	}

	t.AppendRow("total", fmt.Sprintf("%d", regions["total"].totalFarmers), fmt.Sprintf("%d", regions["total"].count), fmt.Sprintf("%.2f", regions["total"].area))

	data := opt.TablePrinter(&t)
	opt.Print(data)

	return nil
}
