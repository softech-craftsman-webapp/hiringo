package crypto

import (
    "fmt"
    "io/ioutil"
)

func ReadFile(filePath string) string {
	data, err := ioutil.ReadFile(filePath)

    if err != nil {
        fmt.Println("File reading error", err)
        return ""
    }
    
	return string(data)
}