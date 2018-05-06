package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"log"
	"os"
	"text/template"

	"github.com/nathanjcochran/mock/iface"
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

	// First argument is interface name:
	args := flag.Args()
	if len(args) < 1 {
		log.Fatal("Not enough args")
	}
	intfName := args[0]

	// Parse the package and get info about the interface:
	intf, err := iface.GetInterface(*dir, intfName)
	if err != nil {
		log.Fatalf("Error getting interface information: %s", err)
	}

	// Parse the template:
	tmpl, err := template.New("default").Parse(defaultTmpl)
	if err != nil {
		log.Fatalf("Error parsing template: %s", err)
	}

	// Execute/output the template:
	buf := &bytes.Buffer{}
	if err := tmpl.Execute(buf, &intf); err != nil {
		log.Fatalf("Error executing template: %s", err)
	}

	// Format it with go fmt:
	result, err := format.Source(buf.Bytes())
	if err != nil {
		log.Fatalf("Error formating output: %s", err)
	}

	// Open the file, if provided, or use stdout:
	out := os.Stdout
	if *outFile != "" {
		out, err = os.Create(*outFile)
		if err != nil {
			log.Fatalf("Error creating output file: %s", err)
		}
		defer out.Close()
	}

	// Write the formatted output to the file:
	if _, err := out.Write(result); err != nil {
		log.Fatalf("Error writing to file: %s", err)
	}
}
