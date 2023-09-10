package common

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/dattaray-basab/cks-clip-lib/globals"
)

func RunTemplate(templateText string, tmplRootData globals.SubstitionTemplateT) (string, error) {
	var buf bytes.Buffer

	template.Must(
		template.New("run").Parse(templateText),
	).Execute(&buf, tmplRootData)
	fmt.Println(buf.String())

	return buf.String(), nil

}
