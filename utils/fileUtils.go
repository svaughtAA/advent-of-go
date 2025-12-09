package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

func GetFileContents(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("Error getting the contents of file %s: %w", path, err)
	}
	return string(data), nil
}

func PrettyPrint(obj any) {
	jsonData, err := json.MarshalIndent(obj, "", "\t")
	if err == nil {
		fmt.Println(string(jsonData))
	}
}
