package main

import (
	"fmt"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"go.bbkane.com/warg"
	"go.bbkane.com/warg/value/scalar"
)

var casing = []string{"lower", "upper", "title"}

func app() *warg.App {
	app := warg.New(
		"convert_case",
		"v0.1.0",
		warg.NewSection(
			"String case conversion utility tool.",
			warg.NewSubCmd(
				"convert",
				"Convert an input string to a different casing.",
				convert,
				warg.NewCmdFlag(
					"--input",
					"The string to convert.",
					scalar.String(),
					warg.Alias("-i"),
					warg.EnvVars("CONVERT_CASE_INPUT"),
					warg.Required(),
				),
				warg.NewCmdFlag(
					"--casing",
					"The casing to convert to.",
					scalar.String(
						scalar.Choices(casing...),
					),
					warg.Alias("-c"),
					warg.EnvVars("CONVERT_CASE_CASING"),
					warg.Required(),
				),
			),
		),
	)
	return &app
}

func convert(ctx warg.CmdContext) error {
	input := ctx.Flags["--input"].(string)
	casing := ctx.Flags["--casing"].(string)

	caser := cases.Caser{}

	switch casing {
	case "lower":
		caser = cases.Lower(language.Und)
	case "upper":
		caser = cases.Upper(language.Und)
	case "title":
		caser = cases.Title(language.Und)
	}

	fmt.Fprintln(ctx.Stdout, caser.String(input))

	return nil
}

func main() {
	app := app()
	app.MustRun()
}
