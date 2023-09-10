package templates
var (
  T3 = `
{{- range $name, $map := . -}}
  {{ $name }}:
    {{- range $name2, $values := $map }}
      {{ $name2 }}:
        {{- range $value := $values }}
          {{ $value }}
        {{- end }}
    {{- end }}
{{ end -}}
`
)