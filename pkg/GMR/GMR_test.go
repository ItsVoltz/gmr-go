package GMR

import (
	"bytes"
	"os"
	"testing"
)

func TestEncrypt(t *testing.T) {
	script := []byte("print('Hello World')")
	expected := []byte("dHF86X8zqTduHvNkp9cmtQ+SH7rlsL+b0fSS6IpzYXA=")
	encrypted, err := Encrypt(script)
	if err != nil {
		t.Error(err)
	}

	if encrypted == nil {
		t.Error("Encrypted script is nil")
	}

	if !bytes.Equal(encrypted, expected) {
		t.Errorf("Expected %s, got %s", expected, encrypted)
	}
}

func TestEncryptFile(t *testing.T) {
	src, dst := "../hello_world.lua", "../hello_world_encrypted.lua"
	expected := []byte("GMR.RunEncryptedScript(\"dHF86X8zqTduHvNkp9cmtQ+SH7rlsL+b0fSS6IpzYXA=\")")
	err := EncryptFile(src, dst)
	if err != nil {
		t.Error(err)
	}

	encryptedBytes, err := os.ReadFile(dst)
	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal(encryptedBytes, expected) {
		t.Errorf("Expected %s, got %s", expected, encryptedBytes)
	}

	err = os.Remove(dst) // remove encrypted test file
	if err != nil {
		t.Error(err)
	}
}
