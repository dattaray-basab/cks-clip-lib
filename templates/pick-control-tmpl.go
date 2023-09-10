package templates
var (
  PickQueryTemplate = `
{
  "__CONTENT": [

      {
          "id": {{.ShortQueryId}},
          "kind": "multiselect",
          "prompt": "enter ...",
        {{ range $k, $v := .MoveItemsInfo -}}
          "selector": [
            {{ $v.Index }}
            {{if $v.IsLastItem}},{{else}}{{end}}
          ],
          "children": {
            "kind": "literal",
              "value": [
                  {{ $v.Key }},
              ]
            },
          {{- end }}
      }
  ]
}
`
)