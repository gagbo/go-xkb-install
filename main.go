package main

import (
	"github.com/gagbo/go-xkb-install/internal/rules"
	"github.com/gagbo/go-xkb-install/internal/utils"

	"github.com/jessevdk/go-flags"

	"fmt"
	"log"
	"os"
)

type Options struct {
	Verbose        []bool `short:"v" long:"verbose" description:"Verbose output"`
	XkbSymbol      string `short:"S" long:"symbol" required:"yes" description:"XkbSymbol name associated with the layout"`
	XkbVariant     string `short:"V" long:"variant" required:"yes" description:"XkbVariant name to be associated with the layout"`
	XkbDescription string `short:"d" long:"description" required:"yes" description:"Description of the variant"`
	XkbComposePath string `long:"compose" description:"Path to an optional XCompose file to use."`
	Positional     struct {
		Path string `positional-arg-name:"path" required:"1" description:"Path to the symbols definition file"`
	} `positional-args:"yes"`
}

var options Options

var parser = flags.NewParser(&options, flags.Default)

func main() {
	if _, err := parser.Parse(); err != nil {
		switch flagsErr := err.(type) {
		case flags.ErrorType:
			if flagsErr == flags.ErrHelp {
				os.Exit(0)
			}
			os.Exit(1)
		default:
			os.Exit(1)
		}
	}

	fmt.Printf("Path to file that's supposed to have the symbols data: %s", options.Positional.Path)

	err := rules.AddLstVariant(options.XkbSymbol, options.XkbVariant, options.XkbDescription)
	if err != nil {
		log.Fatalf("Error dealing adding rule to lst file: %s", err)
	}

	err = rules.UpdateVariant(options.XkbSymbol, options.XkbVariant, options.XkbDescription)
	if err != nil {
		log.Fatalf("Error dealing with rules xml file: %s", err)
	}


	if options.XkbComposePath != "" {
		if _, err := utils.Copy(options.XkbComposePath, "~/.Xcompose"); err != nil {
			log.Fatalf("Error copying XCompose: %s.", err)
		}
	}
}
