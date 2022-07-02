# embedsfs

Golang embedded file system structured processing helper.

## Installation

```shell
go get -u github.com/coolstina/embedsfs
```

## Usage

### Specify embeds relative path

```go
package main

import (
	"embed"
	"fmt"
	"github.com/coolstina/embedsfs"
	"log"
)

//go:embed embeds
var embeds embed.FS

func main() {
	// Initialize embedsfs instance.
	ins := embedsfs.NewEmbedsFS(embeds, embedsfs.WithPath("embeds"))

	// Get golang internal structure embed.FS
	embeds := ins.Embeds()
	fmt.Println(embeds)

	// Get golang embeds filesystem filename
	name := ins.Filename("readme.txt")
	fmt.Println(name)

	// Get golang embeds filesystem filename
	data, err := ins.Content("readme.txt")
	if err != nil {
		log.Fatalf("get embeds file content to failure: %v", err)
	}

	fmt.Println(string(data))
}
```

### Not specify embeds relative path

```go
package main

import (
	"embed"
	"fmt"
	"github.com/coolstina/embedsfs"
	"log"
)

//go:embed embeds
var embeds embed.FS

func main() {
	// Initialize embedsfs instance.
	ins := embedsfs.NewEmbedsFS(embeds)

	// Get golang internal structure embed.FS
	embeds := ins.Embeds()
	fmt.Println(embeds)

	// Get golang embeds filesystem filename
	name := ins.Filename("embeds/readme.txt")
	fmt.Println(name)

	// Get golang embeds filesystem filename
	data, err := ins.Content("embeds/readme.txt")
	if err != nil {
		log.Fatalf("get embeds file content to failure: %v", err)
	}

	fmt.Println(string(data))
}
```

## LICENSE

- [Apache License](LICENSE)

## Author

- [@helloshaohua](https://github.com/helloshaohua)


