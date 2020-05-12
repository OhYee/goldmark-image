# goldmark-image

[![Sync to Gitee](https://github.com/OhYee/goldmark-image/workflows/Sync%20to%20Gitee/badge.svg)](https://gitee.com/OhYee/goldmark-image) [![w
orkflow state](https://github.com/OhYee/goldmark-image/workflows/test/badge.svg)](https://github.com/OhYee/goldmark-image/actions) [![codecov](https://codecov.io/gh/OhYee/goldmark-image/branch/master/graph/badge.svg)](https://codecov.io/gh/OhYee/goldmark-image) [![version](https://img.shields.io/github/v/tag/OhYee/goldmark-image)](https://github.com/OhYee/goldmark-image/tags)bn

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
	img "github.com/OhYee/goldmark-image"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
)

func main() {
	var buf bytes.Buffer
	source := []byte(`![Logo](https://www.oyohyee.com/static/img/logo.svg "title")`)

	md := goldmark.New(
	    goldmark.WithExtensions(
			img.NewImg("image", nil),
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
