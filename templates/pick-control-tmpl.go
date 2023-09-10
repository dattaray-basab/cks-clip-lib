package templates
var (
  PickControlTemplate = `
[
  {
	"op": "pick",
	"directives": {
      "token_id": "{{.FullQueryId}}"}}",
      {{.MoveItemsInfo}}": {
        {{- range $k, $v := . -}}
          {{ $v }}
            "options": [
              {
                "rel_paths": [
                  "{{ $k }}"
                ],
                "sift": "{{.Index}}}}"
              }
            ]
        {{- end }}
    }
  }
]
`
)