package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ahmedsat/erp-reports-cli/commands"
	"github.com/ahmedsat/erp-reports-cli/utils"
)

var subcommands = []string{"totals", "farms", "training"}

func usage() {
	fmt.Printf("Usage: %s subcommand [options]\n", os.Args[0])
	fmt.Println("subcommands:")
	for _, subcommand := range subcommands {
		fmt.Printf("  %s\n", subcommand)
	}
}

func main() {

	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "totals":
		cmd := flag.NewFlagSet("totals", flag.ExitOnError)
		opt := commands.TotalsOptions{}
		opt.AddFlags(cmd)

		utils.HandelErr(cmd.Parse(os.Args[2:]))

		utils.HandelErr(commands.Totals(opt))

	case "farms":
		cmd := flag.NewFlagSet("farms", flag.ExitOnError)
		opt := commands.FarmsOptions{}
		opt.AddFlags(cmd)

		utils.HandelErr(cmd.Parse(os.Args[2:]))

		utils.HandelErr(commands.Farms(opt))
	case "training":
		cmd := flag.NewFlagSet("training", flag.ExitOnError)
		opt := commands.TrainingOptions{}
		opt.AddFlags(cmd)

		utils.HandelErr(cmd.Parse(os.Args[2:]))

		utils.HandelErr(commands.Training(opt))

	default:
		usage()
	}
}
