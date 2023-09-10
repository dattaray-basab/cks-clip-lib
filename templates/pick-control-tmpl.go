package templates
var (
  PickQueryTemplate = `
{
  "__CONTENT": [

      {
          "id": {{.ShortQueryId}},
          "kind": "multiselect",
          "prompt": "enter ...",
          "selector": [
        {{ range $k, $v := .MoveItemsInfo -}}
            {{ $v.Index }}
            {{- if $v.IsLastItem -}}{{else}},{{- end }}
        {{- end }}
          ],
          "children": {
            "kind": "literal",
              "value": [
        {{ range $k, $v := .MoveItemsInfo -}}
                  {{ $v.Key }},
        {{- end }}
              ]
            },
      }
  ]
}
`
)