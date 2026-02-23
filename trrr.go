package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

const defaultSrc = "auto"
const defaultTgt = "en"
const defaultBackend = "google"

type config struct {
	source string
	target string

	backend string

	googleConfig struct{}
	lingvaConfig lingvaConfig
}
type lingvaConfig struct {
	domain string
}

func main() {
	var c config

	var text string
	var configFileLocation string

	flag.StringVar(&c.source, "s", defaultSrc, "source language")
	flag.StringVar(&c.target, "t", defaultTgt, "target language")
	flag.StringVar(&c.backend, "b", defaultBackend, "backend")
	flag.StringVar(&c.backend, "C", configFileLocation, "config file")
	flag.StringVar(&text, "S", "", "translate from string instead of stdin")
	flag.Parse()

	if text == "" {
		stat, _ := os.Stdin.Stat()

		if (stat.Mode() & os.ModeCharDevice) == 0 {
			c, err := io.ReadAll(os.Stdin)
			if err != nil && errors.Is(err, io.EOF) {
				fmt.Fprintln(os.Stderr, "failed to read stdin:", err)
				os.Exit(1)
			}
			text = string(c)
		} else {
			fmt.Fprintf(os.Stderr, "no input")
			return
		}
	}

	c.lingvaConfig.domain = "translate.plausibility.cloud"
	// if configFileLocation == "" {
	//     findConfigFile(&configFileLocation)
	// }
	// readConfig(configFileLocation, &c)

	switch c.backend {
	case "google":
		translation, err := google(c, text)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
		}
		fmt.Print(translation)
	case "lingva":
		translation, err := lingva(c, text)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
		}
		fmt.Print(translation)
	default:
		panic("unknown backend")
	}
}
