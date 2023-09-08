package templates

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"
)

type Inventory struct {
	Material string
	Count    uint
}

func Junk(templateMap map[string]string, moveMap map[string]string) {
	var buf bytes.Buffer
	sweaters := Inventory{"wool", 17}

	// tmpl_, err := template.ParseFiles("x1.tmpl")
	// if err != nil {
	// 	panic(err)
	// }
	// logger.Log.Debug(tmpl_)

	tmplStr := "{{.Count}} items are made of {{.Material}}\n"
	tmpl := template.Must(template.New("test").Parse(tmplStr))
	err := tmpl.Execute(&buf, sweaters)
	if err != nil {
		panic(err)
	}
	tmplStr2 := "{{.Count}} --- {{.Material}}\n"
	tmpl = template.Must(template.New("test").Parse(tmplStr2))
	_ = tmpl.Execute(&buf, sweaters)
	s := buf.String()
	fmt.Println(s)

	cond := true
	tmplStr3 := "{{if .}}true{{else}}false{{end}}\n"
	tmpl = template.Must(template.New("test").Parse(tmplStr3))
	_ = tmpl.Execute(&buf, cond)
	s = buf.String()
	fmt.Println(s)

	names := []string{"Bart", "Lisa", "Maggie"}
	tmplStr4 := "Greetings by name:\n {{range .}}\tHowdy {{.}}!\n{{end}}Byes YouAlls.\n"
	tmpl = template.Must(template.New("test").Parse(tmplStr4))
	_ = tmpl.Execute(&buf, names)
	s = buf.String()
	fmt.Println(s)

	tmplStr5 := "Greetings by name:\n {{range .}}\tHowdy {{toupper .}}!\n{{end}}Byes YouAlls.\n"
	funcMap := template.FuncMap{
		"toupper": strings.ToUpper,
        "tolower": strings.ToLower,
	}
	tmpl = template.Must(template.New("func").Funcs(funcMap).Parse(tmplStr5))
	_ = tmpl.Execute(&buf, names)
	s = buf.String()
	fmt.Println(s)

}
