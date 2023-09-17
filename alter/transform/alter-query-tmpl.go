package transform
var (
  TransformQueryTemplate = `
{
  "__CONTENT": [
  	{
      "id": "{{.ShortQueryId}}",
      "kind": "text",
      "prompt": "enter ...:",
      "value": "{{.FirstWordInFirstFile}}",
    }
  ]
}
`
)