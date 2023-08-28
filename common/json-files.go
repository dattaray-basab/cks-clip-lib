package common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func ReadJsonFile(filePath string) (map[string]interface{}, error) {
// Open our jsonFile
    jsonFile, err := os.Open(filePath)
    // if we os.Open returns an error then handle it
    if err != nil {
        return nil, err
    }
    fmt.Println("Successfully Opened users.json")    // defer the closing of our jsonFile so that we can parse it later on
    defer jsonFile.Close()

    byteValue, _ := ioutil.ReadAll(jsonFile)

    var result map[string]interface{}
    json.Unmarshal([]byte(byteValue), &result)

	return result, nil


}
