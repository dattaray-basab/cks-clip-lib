package templates
var (
  PickControlTemplate = `
[
  {
	"op": "recast",
	"directives": {
    "token_id": "{{.FullQueryId}}",
        "options": [
        {{- range $k, $v := .MoveItemsInfo }}
            {
              "rel_paths": [
                "{{ $k }}"
              ],
              "sift": "{{ $v.Key }}"
            }
            {{- if $v.IsLastItem }}{{ else }}, {{ end -}}
        {{- end }}
        ]
    }
  }
]
`
)