package common

import "os"

func ReadJsonFile(filepath string) (map[string]interface{}, error) {
	jsonMap := make(map[string]interface{})
	jsonFile, err := os.Open(filepath)
	if err != nil {
		return jsonMap, err
	}
	defer jsonFile.Close()
	return jsonMap, nil
}