
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

## CLI Tool
#### Alternatively you can use the CLI tool to encrypt your GMR scripts. ([Download release here](https://github.com/ItsVoltz/gmr-go/releases/tag/CLI-Tool))
Via command line:
```bash
./gmr path/to/my/script.lua
```
Via command line with custom destination:
```bash
./gmr path/to/my/script.lua path/to/destination/script_encrypted.lua
```

If you are on windows you can also drag and drop a script onto the executable and it will automatically encrypt that file for you.