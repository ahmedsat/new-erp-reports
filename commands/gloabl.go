package commands

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/ahmedsat/erp-reports-cli/utils"
	"github.com/atotto/clipboard"
)

type GlobalOptions struct {
	FromStr         string `json:"from"`
	ToStr           string `json:"to"`
	DateLayout      string `json:"date_layout"`
	Format          string `json:"format"`
	Copy            bool   `json:"copy"`
	OutputFile      string `json:"output"`
	IncludeCanceled bool   `json:"include_canceled"`

	// parsed
	From         time.Time          `json:"-"`
	To           time.Time          `json:"-"`
	TablePrinter utils.TablePrinter `json:"-"`
}

func (g *GlobalOptions) AddFlags(cmd *flag.FlagSet) {
	cmd.StringVar(&g.FromStr, "from", "1-1-2022", "Date part of ISO")
	cmd.StringVar(&g.ToStr, "to", time.Now().Format("2-1-2006"), "Date part of ISO")
	cmd.StringVar(&g.DateLayout, "date-layout", "2-1-2006", "Date layout")
	cmd.StringVar(&g.Format, "format", "csv", "Format of output [csv, tsv]")
	cmd.BoolVar(&g.Copy, "copy", false, "Copy to clipboard")
	cmd.StringVar(&g.OutputFile, "output", "", "Output file")
	cmd.BoolVar(&g.IncludeCanceled, "include-canceled", false, "Include canceled farms")
}

func (g *GlobalOptions) Validate() (err error) {
	if g.DateLayout == "" {
		return fmt.Errorf("%s : -date-layout not set", utils.WhereAmI())
	}

	// check if the layout is valid
	// bug: this does not check the layout
	_, err = time.Parse(g.DateLayout, "2-1-2006")
	if err != nil {
		return errors.Join(err, fmt.Errorf("%s : invalid date layout", utils.WhereAmI()))
	}

	if g.FromStr == "" {
		return fmt.Errorf("%s : -from not set", utils.WhereAmI())
	}

	if g.ToStr == "" {
		return fmt.Errorf("%s : -to not set", utils.WhereAmI())
	}

	g.From, err = time.Parse(g.DateLayout, g.FromStr)
	if err != nil {
		return errors.Join(err, fmt.Errorf("%s : failed to parse from", utils.WhereAmI()))
	}

	g.To, err = time.Parse(g.DateLayout, g.ToStr)
	if err != nil {
		return errors.Join(err, fmt.Errorf("%s : failed to parse to", utils.WhereAmI()))
	}

	if g.From.After(g.To) {
		return fmt.Errorf("%s : -from is after -to", utils.WhereAmI())
	}

	switch g.Format {
	case "csv":
		g.TablePrinter = utils.TablePrinterCsv
	case "tsv":
		g.TablePrinter = utils.TablePrinterTsv
	default:
		g.TablePrinter = utils.TablePrinterCsv
		fmt.Printf("unknown format: %s defaulting to csv\n", g.Format)
	}

	return nil
}

func (g *GlobalOptions) Print(data string) (err error) {

	if g.Copy {
		err = clipboard.WriteAll(data)
		if err != nil {
			return errors.Join(err, fmt.Errorf("%s : failed to copy to clipboard", utils.WhereAmI()))
		}
	}

	if g.OutputFile != "" {
		err = os.WriteFile(g.OutputFile, []byte(data), 0644)
		if err != nil {
			return errors.Join(err, fmt.Errorf("%s : failed to write file", utils.WhereAmI()))
		}
	}

	if g.OutputFile == "" && !g.Copy {
		fmt.Println(data)
	}

	return
}
