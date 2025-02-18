package functions

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func ReadJSON(filename string) string {
	_, currentFilePath, _, _ := runtime.Caller(0)
	mockDataDirectory := filepath.Join(currentFilePath, "../../mockData")
	jsonFileContent, err := os.ReadFile(fmt.Sprintf("%s/%s", mockDataDirectory, filename))
	if err != nil {
		fmt.Println("[ERROR] Error while reading a JSON file: " + err.Error())
		panic(err)
	}

	return string(jsonFileContent)
}
