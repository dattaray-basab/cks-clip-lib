package filegen

import (
	"os"
	"path/filepath"

	"github.com/dattaray-basab/cks-clip-lib/common"
	"github.com/dattaray-basab/cks-clip-lib/globals"
)

func CreateRecipeFiles(baseDirpath string) error {
	err := os.RemoveAll(baseDirpath)
	if err != nil {
		return err
	}

	recipeScaffold := globals.ScaffoldInfoTListT{
		{
			Filepath: filepath.Join(baseDirpath, "__BLUEPRINTS", "__TOKENS", "__QUERY", "base.json"),
			Content: `
{
  "__SCHEMA": "s_context",
  "__CONTENT": [
	{
	  "id": "id_name1",
	  "kind": "text",
	  "prompt": "please enter your name:",
	  "value": "Ashwini"
	},
	{
	  "id": "id_name2",
	  "kind": "text",
	  "prompt": "please enter your friend's name:",
	  "value": "Basab"
	},
	{
	  "id": "id_closed",
	  "kind": "text",
	  "prompt": "please enter if weekend closings are effective:",
	  "value": "0"
	},
	{
	  "id": "id_port",
	  "kind": "select",
	  "prompt": "enter ...",
	  "selector": 1,
	  "children": {
		"kind": "literal",
		"value": [
		  "3200",
		  "3201",
		  "3202",
		  "3203"
		]
	  }
	},
	{
	  "id": "id_fixed_pages",
	  "kind": "multiselect",
	  "prompt": "enter ...",
	  "selector": [
		1,
		0
	  ],
	  "children": {
		"kind": "literal",
		"value": [
		  "contact",
		  "about",
		  "posts"
		]
	  }
	},
	{
	  "id": "id_var_pages",
	  "kind": "multiselect",
	  "prompt": "enter ...",
	  "selector": [
		4,
		2,
		1
	  ],
	  "children": {
		"kind": "literal",
		"value": [
		  "hotels",
		  "restaurants",
		  "attractions",
		  "events",
		  "shopping",
		  "nightlife"
		]
	  }
	}
  ]
}
		`,
		},
		{
			Filepath: filepath.Join(baseDirpath, "__BLUEPRINTS", "{{target}}", "__MISC", "directives.json"),
			Content: `
{
  "LOG_LEVEL": 10,
  "MODE_LOG": true,
  "ACTIVE_PHASE_LIST": null,
  "CURRENT_CONTEXT": "1_response_context",
  "OPERATION_FOLDER_PREFIX": "__"
}
		`,
		},
		{
			Filepath: filepath.Join(baseDirpath, "__BLUEPRINTS", "{{target}}", "__PHASES", "{{phase_name}}.json"),
			Content: `
{
  "__CODE_BLOCK": "{{code_block}}",
  "ops_pipeline": [
	{
	  "remove": {
		"filter": "[*]"
	  }
	},
	{
	  "copy": {
		"filter": "~[__*]|~[.DS_Store, .gitignore]"
	  }
	}
  ]
}
		`,
		},
		{
			Filepath: filepath.Join(baseDirpath, "__BLUEPRINTS", "{{target}}", "run.py"),
			Content: `
from code_transformer.src.main._main_generator import fn_start

error, _ = fn_start(__file__)
if error is not None:
    exit(1)
		`,
		},
		{
			Filepath: filepath.Join(baseDirpath, "__RECIPE_CONFIG.json"),
			Content: `
[
  "../../__BLUEPRINTS"
]
		`,
		},
	}

	err = common.CreateFiles(recipeScaffold)
	return err
}
