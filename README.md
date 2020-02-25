# goldmark-image

[![master_test](https://github.com/OhYee/goldmark-image/workflows/master_test/badge.svg)](https://github.com/OhYee/goldmark-image/actions?workflow=master_test)

goldmark-image is an extension for [goldmark](https://github.com/yuin/goldmark).  

## Installation

```bash
go get -u github.com/OhYee/goldmark-image
```

## Usage

See `image_test.go`

```go
package main

import (
	"bytes"
	"fmt"
	"github.com/OhYee/goldmark-dot"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
)

func main() {
	var buf bytes.Buffer
	source := []byte(`![Logo](https://www.oyohyee.com/static/img/logo.svg "title")`)

	md := goldmark.New(
	    goldmark.WithExtensions(
			NewImg("image", nil),
		),
	)

	if err := md.Convert(source, &buf); err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", buf.Bytes())
}
```

## License

[MIT](LICENSE)
