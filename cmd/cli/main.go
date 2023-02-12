package main

import (
	"fmt"
	"github.com/ItsVoltz/gmr-go/pkg/GMR"
	"log"
	"os"
	"path"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Fatal("Please provide a source file.")
	}

	src := args[1]
	var dst string
	ext := path.Ext(src)

	if len(args) > 2 {
		dst = args[2]
	} else {
		dst = src[:len(src)-len(ext)] + fmt.Sprintf("_encrypted%s", ext)
	}

	err := GMR.EncryptFile(src, dst)
	if err != nil {
		log.Fatalf("Failed to encrypt file: %s", err.Error())
	}

	println("Encrypted file successfully.\nPress Enter to close.")
	fmt.Scanln()
}
