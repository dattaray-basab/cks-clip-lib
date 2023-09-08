package templates

var T2 string = `
	{{ range . }}
// {{ .Desc }}
type {{ .Name }} struct {
    {{ range .Fields -}}
        {{ .Name }} {{ .TypeName }}
    {{ end }}
}
{{ end }}
	`

