package commands

import "flag"

type SalaryOptions struct {
	Output
}

func (s SalaryOptions) AddFlags(cmd *flag.FlagSet) {
	s.Output.AddFlags(cmd)
}

func Salary(opt SalaryOptions) (err error) {
	panic("unimplemented")
}
