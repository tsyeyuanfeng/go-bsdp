# go-bsdp
An binary diff&amp;patch library based on bsdiff algorithm(v4.3) for Go

## Installation

```bash
go get -v github.com/tsyeyuanfeng/go-bsdp/...
```

## API

```go
package main

import (
    "os"
    "github.com/tsyeyuanfeng/go-bsdp/bsdp"
)

const (
    oldFilePath = "your old file path"
    newFilePath = "your new file path"
    patchFilePath = "your patch file path"
)

func main() {
    // Generate patch
    bsdp.Diff(oldFilePath, newFilePath, patchFilePath)

    // Apply patch
    bsdp.patch(oldFilePath, newFilePath, patchFilePath)
}
```

## Command line tool
```bash
# Generate patch
go-bsdp diff "old file" "new file" "patch file"

# Apply patch
go-bsdp patch "old file" "new file" "patch file"
```