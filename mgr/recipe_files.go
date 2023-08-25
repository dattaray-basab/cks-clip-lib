package mgr

import (


	"github.com/otiai10/copy"
)

func SetupRecipeFiles(templateMap map[string]string, dst_recipe_dirpath string, src_recipe_dirpath string) error {
	// query_token_path := filepath.Join(dst_recipe_dirpath, "__BLUEPRINTS", "__TOKENS", "__QUERY", "base.json")
	// fmt.Println(query_token_path)


	// response_token_path := filepath.Join(dst_recipe_dirpath, "__BLUEPRINTS", "__TOKENS", "0_default_context", "base.json")
	// fmt.Println(response_token_path)

	// directives_path := filepath.Join(dst_recipe_dirpath, "__BLUEPRINTS", "{{target}}", "__MISC", "directives.json")
	// fmt.Println(directives_path)

	// phase_path := filepath.Join(dst_recipe_dirpath, "__BLUEPRINTS", "{{target}}", "__PHASES", "{{phase_name}}.json")
	// fmt.Println(phase_path)

	// run_path := filepath.Join(dst_recipe_dirpath, "__BLUEPRINTS", "{{target}}", "run.py")
	// fmt.Println(run_path)

	err := copy.Copy(src_recipe_dirpath, dst_recipe_dirpath)
	if err != nil {
		return err
	}
	return nil
}
