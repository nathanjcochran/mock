package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"text/template"

	"github.com/nicheinc/mock/iface"
	"golang.org/x/tools/imports"
)

func main() {
	var (
		dir     = flag.String("d", ".", "Directory to search for interface in")
		outFile = flag.String("o", "", "Output file (default stdout)")
	)
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [options] interface\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "Options:\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	// First argument is interface name
	args := flag.Args()
	if len(args) < 1 {
		log.Fatal("Not enough args")
	}
	ifaceName := args[0]

	// Parse the package and get info about the interface
	iface, err := iface.GetInterface(*dir, ifaceName)
	if err != nil {
		log.Fatalf("Error getting interface information: %s", err)
	}

	// Parse the template
	tmpl, err := template.New("default").Parse(tmpl)
	if err != nil {
		log.Fatalf("Error parsing template: %s", err)
	}

	// Execute/output the template
	buf := &bytes.Buffer{}
	if err := tmpl.Execute(buf, &iface); err != nil {
		log.Fatalf("Error executing template: %s", err)
	}

	// Format it with go imports
	formatted, err := imports.Process(*outFile, buf.Bytes(), nil)
	if err != nil {
		log.Fatalf("Error formating output: %s", err)
	}

	// Open the file, if provided, or use stdout
	out := os.Stdout
	if *outFile != "" {
		out, err = os.Create(*outFile)
		if err != nil {
			log.Fatalf("Error creating output file: %s", err)
		}
		defer out.Close()
	}

	// Write the formatted output to the file
	if _, err := out.Write(formatted); err != nil {
		log.Fatalf("Error writing to file: %s", err)
	}
}
