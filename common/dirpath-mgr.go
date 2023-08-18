package common

import (
	"fmt"
	"strings"
)

func GetNamedPath(dirPath string, searchPathFragment string) (string, error)  {
	dirPathParts := strings.Split(dirPath, "/")
	newPath := ""
	for i, dirPathPart := range dirPathParts {
		newPath = newPath + dirPathPart + "/"
		if dirPathPart == searchPathFragment {
			fmt.Println("misc/envmgr.go::SetRecipePath: index, dirPathPart::", i, dirPathPart)
			return newPath, nil
		}
	}
	return "", fmt.Errorf("misc/envmgr.go::SetRecipePath: %s not found in %s", searchPathFragment, dirPath)
}