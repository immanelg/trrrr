package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

type config struct {
	source string
	target string

	backend string

	lingvaDomain string
}

func defaults(c *config) {
	c.source = "auto"
	c.target = "en"
	c.backend = "google"
	c.lingvaDomain = "translate.plausibility.cloud"
}

func readCommand(c *config) {
}

func main() {
	var c config

	defaults(&c)

	text := ""

	// CLI
	var source string
	var target string
	var backend string
	var configPath string
	flag.StringVar(&source, "s", "", "source language")
	flag.StringVar(&target, "t", "", "target language")
	flag.StringVar(&backend, "b", "", "backend")
	flag.StringVar(&configPath, "C", "", "config file")
	flag.StringVar(&text, "S", text, "translate from string instead of stdin")
	flag.Parse()

	if configPath == "" { configPath = findConfigPath() }
	// override defaults by config file
	_ = readConfig(configPath, &c)

	// override config file by CLI
	if source != "" { c.source = source }
	if target != "" { c.target = target }
	if backend != "" { c.backend = backend }

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
