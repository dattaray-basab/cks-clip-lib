package common

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/dattaray-basab/cks-clip-lib/globals"
)

func RunTemplate(data map[string]map[string][]string, templateText string, substitutionTemplate globals.SubstitionTemplateT) (string, error) {
	var buf bytes.Buffer

	template.Must(
		template.New("run").Parse(templateText),
	).Execute(&buf, Data)
	fmt.Println(buf.String())

	return buf.String(), nil

}
