package expt3

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/dattaray-basab/cks-clip-lib/globals"
	"github.com/dattaray-basab/cks-clip-lib/templates"
)

func run(text string)(string, error) {
	var buf bytes.Buffer
	//   fmt.Printf("Template:\n%s\n", text)
	//   fmt.Printf("Output:\n\n'''\n")
	template.Must(
		template.New("run").Parse(text),
	).Execute(&buf, Data)
	fmt.Println(buf.String())
	// fmt.Printf("”'\n\n")
	return buf.String(), nil
}

func Expt3(templateMap map[string]string, substitutionTemplate globals.SubstitionTemplateT) (string, error) {
	result, err := run(templates.T3)
	return result, err
}


