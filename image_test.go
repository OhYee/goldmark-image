package dot

import (
	"bytes"
	"strings"
	"testing"

	"github.com/yuin/goldmark"
)

func Test_default(t *testing.T) {

	testcases := []struct {
		name          string
		class         string
		renderWrapper RenderWrapperFunc
		source        string
		want          string
	}{
		{
			name:          "src only",
			class:         "",
			renderWrapper: nil,
			source:        `![](https://www.oyohyee.com/static/img/logo.svg)`,
			want:          `<p><a target="_blank" rel="noopener noreferrer" href="https://www.oyohyee.com/static/img/logo.svg"><img src="https://www.oyohyee.com/static/img/logo.svg" /></a></p>`,
		},
		{
			name:          "class only",
			class:         "image",
			renderWrapper: nil,
			source:        `![](https://www.oyohyee.com/static/img/logo.svg)`,
			want:          `<p><a class="image" target="_blank" rel="noopener noreferrer" href="https://www.oyohyee.com/static/img/logo.svg"><img src="https://www.oyohyee.com/static/img/logo.svg" /></a></p>`,
		},
		{
			name:          "alt only",
			class:         "",
			renderWrapper: nil,
			source:        `![Logo](https://www.oyohyee.com/static/img/logo.svg)`,
			want:          `<p><a target="_blank" rel="noopener noreferrer" href="https://www.oyohyee.com/static/img/logo.svg"><img src="https://www.oyohyee.com/static/img/logo.svg" alt="Logo" /></a></p>`,
		},
		{
			name:          "title only",
			class:         "",
			renderWrapper: nil,
			source:        `![](https://www.oyohyee.com/static/img/logo.svg "title")`,
			want:          `<p><a target="_blank" rel="noopener noreferrer" href="https://www.oyohyee.com/static/img/logo.svg"><img src="https://www.oyohyee.com/static/img/logo.svg" title="title" /></a></p>`,
		},
		{
			name:          "with others",
			class:         "",
			renderWrapper: nil,
			source:        `**Image**:![](https://www.oyohyee.com/static/img/logo.svg)`,
			want:          `<p><strong>Image</strong>:<a target="_blank" rel="noopener noreferrer" href="https://www.oyohyee.com/static/img/logo.svg"><img src="https://www.oyohyee.com/static/img/logo.svg" /></a></p>`,
		},
		{
			name:  "with wapper",
			class: "",
			renderWrapper: func(args ImgArgs, class string, r RenderImgFunc) string {
				return r(args)
			},
			source: `![](https://www.oyohyee.com/static/img/logo.svg)`,
			want:   `<p><img src="https://www.oyohyee.com/static/img/logo.svg" /></p>`,
		},
		{
			name:          "all",
			class:         "image",
			renderWrapper: nil,
			source:        `**Image**:![Logo](https://www.oyohyee.com/static/img/logo.svg "title")`,
			want:          `<p><strong>Image</strong>:<a class="image" target="_blank" rel="noopener noreferrer" href="https://www.oyohyee.com/static/img/logo.svg"><img src="https://www.oyohyee.com/static/img/logo.svg" alt="Logo" title="title" /></a></p>`,
		},
	}

	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			buf := bytes.NewBuffer([]byte{})
			md := goldmark.New(
				goldmark.WithExtensions(
					NewImg(tt.class, tt.renderWrapper),
				),
			)
			if err := md.Convert([]byte(tt.source), buf); err != nil {
				t.Error(err)
			}
			if n := strings.Compare(buf.String(), tt.want+"\n"); n != 0 {
				t.Errorf("Got:\n|>%s<|\nExcepted:\n|>%s<|\n", buf.String(), tt.want+"\n")
			}
		})
	}
}
