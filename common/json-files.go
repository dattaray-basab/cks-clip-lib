package common

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

func ReadJsonFile(filePath string) (map[string]interface{}, error) {
    jsonFile, err := os.Open(filePath)

    if err != nil {
        return nil, err
    }
    log.Println("Successfully Opened users.json")  
    defer jsonFile.Close()

    byteValue, _ := io.ReadAll(jsonFile)

    var result map[string]interface{}
    json.Unmarshal([]byte(byteValue), &result)

	return result, nil


}
