package main

import (
	"flag"
	"log"
	"os"

	"github.com/mitchellh/cli"
)

func main() {
	c := cli.NewCLI("cli", "0.1.0")
	c.Args = os.Args[1:]

	//register subcommand
	c.Commands = map[string]cli.CommandFactory{
		"add": func() (cli.Command, error) {
			return &AddCommand{
				UI: generateUI(),
			}, nil
		},
		"nest": func() (cli.Command, error) {
			return &NestCommand{
				UI: generateUI(),
			}, nil
		},
	}

	code, err := c.Run()
	if err != nil {
		log.Printf("fail to call Run() %v\n", err)
	}
	os.Exit(code)
}

func generateUI() cli.Ui {
	return &cli.ColoredUi{
		InfoColor:  cli.UiColorBlue,
		ErrorColor: cli.UiColorRed,
		Ui: &cli.BasicUi{
			Writer:      os.Stdout,
			ErrorWriter: os.Stderr,
			Reader:      os.Stdin,
		},
	}
}

//add subcommand
type AddCommand struct {
	UI cli.Ui
}

func (c *AddCommand) Synopsis() string {
	return "add subcommand"
}

func (c *AddCommand) Help() string {
	return "Usage: cli add "
}

func (c *AddCommand) Run(args []string) int {
	var debug bool
	flags := flag.NewFlagSet("add", flag.ContinueOnError)
	flags.BoolVar(&debug, "debug", false, "Run as DEBUG mode")

	if err := flags.Parse(args); err != nil {
		return 1
	}
	c.UI.Output("Normal Message")
	c.UI.Error("Error Message")
	log.Println(debug)

	return 0
}

//nest subcommand
type NestCommand struct {
	UI cli.Ui
}

func (c *NestCommand) Synopsis() string {
	return "nest subcommand"
}

func (c *NestCommand) Help() string {
	return "Usage: cli nest"
}

func (c *NestCommand) Run(args []string) int {
	flags := flag.NewFlagSet("nest", flag.ContinueOnError)

	if err := flags.Parse(args); err != nil {
		return 1
	}

	//create further subcommand
	cl := cli.NewCLI("nest", "0.1.0")
	cl.Args = args

	//register subcommand
	cl.Commands = map[string]cli.CommandFactory{
		"other": func() (cli.Command, error) {
			return &OtherCommand{
				UI: generateUI(),
			}, nil
		},
	}

	code, err := cl.Run()
	if err != nil {
		log.Printf("fail to call Run() %v\n", err)
	}
	return code
}

//other subcommand
type OtherCommand struct {
	UI cli.Ui
}

func (c *OtherCommand) Synopsis() string {
	return "other subcommand"
}

func (c *OtherCommand) Help() string {
	return "Usage: cli nest other "
}

func (c *OtherCommand) Run(args []string) int {
	var debug bool
	flags := flag.NewFlagSet("other", flag.ContinueOnError)
	flags.BoolVar(&debug, "debug", false, "Run as DEBUG mode")

	if err := flags.Parse(args); err != nil {
		return 1
	}
	log.Println(c.Help())
	log.Println(c.Synopsis())
	log.Println(debug)

	return 0
}
