package templates
var (
  T3 = `
{{- range $name, $map := . -}}
{{ $name }}:
  {{- range $name, $values := $map }}
  {{ $name }}:
    {{- range $value := $values }}
    {{ $value }}
    {{- end }}
  {{- end }}
{{ end -}}
`
)