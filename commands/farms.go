package commands

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"slices"
	"strings"
	"time"

	"github.com/ahmedsat/erp-reports-cli/erp"
	"github.com/ahmedsat/erp-reports-cli/utils"
)

type FarmsOptions struct {
	GlobalOptions
	Fields                 ListFlagString `json:"fields"`
	FarmFields             []string
	FarmApplicationsFields []string
}

func (f *FarmsOptions) AddFlags(cmd *flag.FlagSet) {
	f.GlobalOptions.AddFlags(cmd)

	cmd.Var(&f.Fields, "fields", "Fields to get")
}

var validFarmsFields = []string{
	"arabic_name",
	"name",
	"region",
	"total_farmers",
	"farm_area__feddan",
	"farm_id",
	"farm_application",
	"creation_date",
}

var validFarmApplicationsFields = []string{
	"name",
	"engineer_name",
}

func tryCorrectField(field string) string {
	switch field {
	case "ar_name":
		return "arabic_name"
	case "a_engineer":
		return "a_engineer_name"
	case "engineer":
		return "engineer_name"
	case "code":
		return "farm_id"
	case "area":
		return "farm_area__feddan"
	case "total":
		return "total_farmers"
	case "application":
		return "farm_application"
	default:
		return field
	}
}

func (f *FarmsOptions) Validate() (err error) {
	err = f.GlobalOptions.Validate()
	if err != nil {
		return
	}

	if len(f.Fields) == 0 {
		return fmt.Errorf("no fields set in --fields\n available fields are %v", slices.Concat(validFarmsFields, validFarmApplicationsFields))
	}

	for i := range f.Fields {

		f.Fields[i] = tryCorrectField(f.Fields[i])

		if after, ok := strings.CutPrefix(f.Fields[i], "a_"); ok {
			withoutPrefix := after
			// check if the field is valid
			if !slices.Contains(validFarmApplicationsFields, withoutPrefix) {
				return fmt.Errorf("invalid field a_%s", withoutPrefix)
			}
			f.FarmApplicationsFields = append(f.FarmApplicationsFields, withoutPrefix)
		} else {
			// check if the field is valid
			if !slices.Contains(validFarmsFields, f.Fields[i]) {
				return fmt.Errorf("invalid field %s", f.Fields[i])
			}
			f.FarmFields = append(f.FarmFields, f.Fields[i])
		}
	}

	if len(f.FarmFields) == 0 {
		return errors.New("no farm fields set")
	}

	return
}

type FarmApplication struct {
	Name     string `json:"name"`
	Engineer string `json:"engineer_name"`
}

func (f FarmApplication) GetField(field string) string {

	switch field {
	case "name":
		return f.Name
	case "engineer_name":
		return f.Engineer
	default:
		fmt.Fprintf(os.Stderr, "invalid field %s\n", field)
		return ""
	}
}

type Farm struct {
	Name            string  `json:"name"`
	ArabicName      string  `json:"arabic_name"`
	Region          string  `json:"region"`
	TotalFarmers    int     `json:"total_farmers"`
	Area            float64 `json:"farm_area__feddan"`
	Code            string  `json:"farm_id"`
	Application     string  `json:"farm_application"`
	CreationDateStr string  `json:"creation_date"`

	FarmApplication FarmApplication `json:"-"`

	// parsed
	CreationDate time.Time `json:"-"`
}

func (f Farm) GetField(field string) string {

	if after, ok := strings.CutPrefix(field, "a_"); ok {
		return f.FarmApplication.GetField(after)
	}

	switch field {
	case "arabic_name":
		return f.ArabicName
	case "name":
		return f.Name
	case "region":
		return f.Region
	case "total_farmers":
		return fmt.Sprintf("%d", f.TotalFarmers)
	case "farm_area__feddan":
		return fmt.Sprintf("%.2f", f.Area)
	case "farm_id":
		return f.Code
	case "farm_application":
		return f.Application
	case "creation_date":
		return f.CreationDate.Format("01-02-2006")
	default:
		fmt.Fprintf(os.Stderr, "invalid field %s\n", field)
		return "UNKNOWN: " + field
	}
}

func Farms(opt FarmsOptions) (err error) {

	err = opt.Validate()
	if err != nil {
		return
	}

	farms, err := erp.GetFarms[Farm](erp.FarmsOptions{
		From: opt.From,
		To:   opt.To,
		Fields: utils.Ternary(
			len(opt.FarmApplicationsFields) > 0,
			append(opt.FarmFields, "farm_application"),
			opt.FarmFields),
		IncludeCanceled: opt.IncludeCanceled,
	})
	if err != nil {
		return
	}

	var farmApplications = []FarmApplication{}

	if len(opt.FarmApplicationsFields) > 0 {
		farmApplications, err = erp.GetFarmApplicants[FarmApplication](erp.FarmApplicantsOptions{
			Fields: append(opt.FarmApplicationsFields, "name"),
		})
		if err != nil {
			return
		}

		for i := range farms {
			farms[i].FarmApplication = utils.FindF(farmApplications, func(application FarmApplication) bool {
				return application.Name == farms[i].Application
			})
		}
	}

	t := utils.TableBase{}
	t.SetHeader(opt.Fields...)

	for _, farm := range farms {
		farm.CreationDate, err = time.Parse("2006-01-02 15:04:05", farm.CreationDateStr)
		var row []string
		for _, field := range opt.Fields {
			row = append(row, farm.GetField(field))
		}
		t.AppendRow(row...)
	}

	data := opt.TablePrinter(&t)
	opt.Print(data)

	return
}
