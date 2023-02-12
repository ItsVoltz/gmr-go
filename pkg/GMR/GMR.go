package GMR

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

const URI = "https://gmrwow.com/api" // The URI for GMR API
var client = &http.Client{}

// Encrypt sends script bytes to the GMR API for Lua encryption
func Encrypt(data []byte) ([]byte, error) {
	URL := URI + "/lua/encrypt"
	request, err := http.NewRequest(http.MethodPost, URL, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "text/plain")
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	return io.ReadAll(response.Body)
}

// EncryptFile encrypts a script and writes to a destination with GMR.RunEncryptedScript(...)
func EncryptFile(src, dst string) error {
	data, err := os.ReadFile(src)
	if err != nil {
		return err
	}

	encryptedBytes, err := Encrypt(data)
	if err != nil {
		return err
	}

	encryptedData := []byte(fmt.Sprintf("GMR.RunEncryptedScript(\"%s\")", encryptedBytes))
	return os.WriteFile(dst, encryptedData, os.ModePerm)
}
