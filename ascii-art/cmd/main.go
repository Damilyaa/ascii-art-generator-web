package main

import (
	"fmt"
	"os"

	"ascii-art-web/ascii-art/banners"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run() error {
	args := os.Args[1:]
	if len(args) < 1 || len(args) > 2 {
		fmt.Println(help)
		return nil
	}

	in := args[0]

	banner := "standard"
	if len(args) == 2 {
		banner = args[1]
	}

	templates, err := banners.ParseTemplates()
	if err != nil {
		return err
	}

	tmpl, ok := templates[banner]
	if !ok {
		return fmt.Errorf("banner %q does not exist", banner)
	}

	out, err := tmpl.Execute(in)
	if err != nil {
		return err
	}

	if out == "" {
		return nil
	}

	if empty(out) {
		fmt.Print(out)
	} else {
		fmt.Println(out)
	}

	return nil
}

func empty(s string) bool {
	for _, r := range s {
		if r != '\n' {
			return false
		}
	}
	return true
}

const help = `Usage: go run . [STRING] [BANNER]

EX: go run . something standard`
