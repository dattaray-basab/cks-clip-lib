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
              {{ $v.Index }},
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