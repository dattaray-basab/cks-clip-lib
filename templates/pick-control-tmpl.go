package templates
var (
  PickControlTemplate = `
[
  {
	"op": "pick",
	"directives": {
    "token_id": {{.FullQueryId}},
      {{- range $k, $v := .MoveItemsInfo -}}
          "options": [
            {
              "rel_paths": [
                "{{ $k }}"
              ],
              "sift": "{{ $v.Key }}"
            }
          ]
      {{- end }}
    }
  }
]
`
)