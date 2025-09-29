package commands

import (
	"errors"
	"flag"
	"fmt"

	"github.com/ahmedsat/erp-reports-cli/utils"
)

type TrainingOptions struct {
	GlobalOptions
}

func (t *TrainingOptions) AddFlags(cmd *flag.FlagSet) {
	t.GlobalOptions.AddFlags(cmd)
	// ? this is a placeholder
}

func (t *TrainingOptions) Validate() (err error) {
	err = t.GlobalOptions.Validate()
	if err != nil {
		return errors.Join(err, fmt.Errorf("%s : failed to validate options", utils.WhereAmI()))
	}

	// ? this is a placeholder

	return
}

func Training(opt TrainingOptions) (err error) {
	err = opt.Validate()
	if err != nil {
		return errors.Join(err, fmt.Errorf("%s : failed to validate options", utils.WhereAmI()))
	}

	fmt.Println("Training")

	return
}
