package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"
)

var (
	templatePath = flag.String("t", "", "path to the template file")
	target       = flag.String("f", "", "path to the file to be created or overwritten; write to stdout if not provided")
)

func main() {
	flag.Parse()
	if err := Run(*templatePath, *target); err != nil {
		log.Fatal(err)
	}
}

// Run parses template found at input path and writes the executed template to
// output path. If no output path is provided, write to stdout.
func Run(input, output string) error {

	tmpl, err := template.ParseFiles(input)
	if err != nil {
		return err
	}

	env := LoadEnv()
	if output == "" {
		return tmpl.Execute(os.Stdout, env)
	}
	f, err := os.Create(output)
	if err != nil {
		return fmt.Errorf("could not open outputfile %v", err)
	}
	defer f.Close()

	return tmpl.Execute(f, env)
}

// LoadEnv reads every environment variable into an Env struct's embedded map.
func LoadEnv() Env {

	var env []string
	env = os.Environ()

	mapped := make(map[string]string)

	for _, pair := range env {
		// os.Environ returns a list of KEY=VALUE pairs
		splitted := strings.Split(pair, "=")
		if len(splitted) != 2 {
			continue
		}
		k, v := splitted[0], splitted[1]
		mapped[k] = v
	}

	return Env{mapped}
}

// Env wraps a map[string]string. This type is convenient for passing to templates.
type Env struct {
	Vars map[string]string
}
