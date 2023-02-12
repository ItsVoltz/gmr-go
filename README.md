
# gmr-go

A Go implementation of the GMR Lua encryption API


## Installation
```bash
go get github.com/ItsVoltz/gmr-go
```

## Usage/Examples

```go
import "github.com/ItsVoltz/gmr-go/pkg/GMR"

function main() {
    // Encrypt, the raw encryption string
    data, err := GMR.Encrypt([]byte("print(\"Hello, World!\")")
    if err != nil {
        panic(err)
    }

    // EncryptFile, encrypts the source file and outputs executable GMR.RunEncryptedScript(...)
    err := GMR.EncryptFile("my_file.lua", "my_encrypted_file.lua")
    if err != nil {
        panic(err)
    }
}
```