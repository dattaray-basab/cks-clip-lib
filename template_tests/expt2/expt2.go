package expt2

import (
	"bytes"
	"fmt"
	"go/format"
	"html/template"
	"log"

	"github.com/dattaray-basab/cks-clip-lib/templates"
)



func Expt2(templateMap map[string]string, moveMap map[string]string) {

	tmplStr := templates.T2

	tmpl := template.Must(template.New("t2").Parse(tmplStr))

	var processed bytes.Buffer
	err := tmpl.Execute(&processed, DataVal)
	if err != nil {
		log.Fatalf("unable to parse data into template: %v\n", err)
	}
	formatted, err := format.Source(processed.Bytes())
	if err != nil {
		log.Fatalf("Could not format processed template: %v\n", err)
	}
	fmt.Println(string(formatted))
}