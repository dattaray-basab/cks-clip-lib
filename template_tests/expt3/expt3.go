package expt3

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/dattaray-basab/cks-clip-lib/templates"
)

func run(text string) {
	var buf bytes.Buffer
	//   fmt.Printf("Template:\n%s\n", text)
	//   fmt.Printf("Output:\n\n'''\n")
	template.Must(
		template.New("run").Parse(text),
	).Execute(&buf, Data)
	fmt.Println(buf.String())
	// fmt.Printf("”'\n\n")
}

func Expt3(templateMap map[string]string, moveMap map[string]string) {
	run(templates.T3)
}


