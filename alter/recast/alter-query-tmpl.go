package recast
var (
  QueryTemplate = `
{
  "__CONTENT": [
      {
          "id": "{{.ShortQueryId}}",
          "kind": "multiselect",
          "prompt": "enter ...",
          "selector": [
        {{- range $k, $v := .MoveItemsInfo }}
            {{- if $v.IsFirstItem }}{{- $v.Index }}{{ else }}{{ end -}}
        {{- end -}}
          ],
          "children": {
            "kind": "literal",
              "value": [
        {{- range $k, $v := .MoveItemsInfo }}
            {{- if $v.IsFirstItem }}"{{- $v.Key }}"{{ else }}{{ end -}}
        {{- end -}}
              ]
            }
      }
  ]
}
`
)