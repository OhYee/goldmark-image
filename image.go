// package image is a extension for the goldmark(http://github.com/yuin/goldmark).
//
// This extension adds a link to the image

package dot

import (
	"bytes"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

// RenderWrapperFunc render wrapper function
type RenderWrapperFunc func(args ImgArgs, class string, renderImg RenderImgFunc) string

// RenderImgFunc render img function
type RenderImgFunc func(args ImgArgs) string

// Img extension
type Img struct {
	className     string
	renderWrapper RenderWrapperFunc
}

// NewImg initial a Img extension
func NewImg(className string, renderWrapper RenderWrapperFunc) goldmark.Extender {
	return &Img{className: className, renderWrapper: renderWrapper}
}

// Extend implements goldmark.Extender.
func (img *Img) Extend(m goldmark.Markdown) {
	m.Renderer().AddOptions(renderer.WithNodeRenderers(
		util.Prioritized(img, 0),
	))
}

// RegisterFuncs implements NodeRenderer.RegisterFuncs.
func (img *Img) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(ast.KindImage, img.render)
}

// ImgArgs arguments of Img
type ImgArgs struct {
	Title string
	Alt   string
	Src   string
}

func (img *Img) wrapper(args ImgArgs, class string, renderImg RenderImgFunc) string {
	buf := bytes.NewBuffer([]byte{})
	buf.WriteString(`<a`)
	if img.className != "" {
		buf.WriteString(` class="`)
		buf.WriteString(img.className)
		buf.WriteString(`"`)
	}
	buf.WriteString(` target="_blank" rel="noopener noreferrer" href="`)

	buf.WriteString(args.Src)
	buf.WriteString(`">`)
	buf.WriteString(renderImg(args))
	buf.WriteString(`</a>`)
	return buf.String()
}

func (img *Img) renderImg(args ImgArgs) string {
	buf := bytes.NewBuffer([]byte{})
	buf.WriteString(`<img src="`)
	buf.WriteString(args.Src)
	buf.WriteString(`"`)
	if len(args.Alt) != 0 {
		buf.WriteString(` alt="`)
		buf.WriteString(args.Alt)
		buf.WriteString(`"`)
	}
	if len(args.Title) != 0 {
		buf.WriteString(` title="`)
		buf.WriteString(args.Title)
		buf.WriteString(`"`)
	}
	buf.WriteString(` />`)
	return buf.String()
}

func (img *Img) render(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if !entering {
		return ast.WalkContinue, nil
	}
	n := node.(*ast.Image)
	if img.renderWrapper == nil {
		w.WriteString(img.wrapper(ImgArgs{
			Title: string(n.Title),
			Alt:   string(n.Text(source)),
			Src:   string(n.Destination),
		}, img.className, img.renderImg))
	} else {
		w.WriteString(img.renderWrapper(ImgArgs{
			Title: string(n.Title),
			Alt:   string(n.Text(source)),
			Src:   string(n.Destination),
		}, img.className, img.renderImg))
	}
	return ast.WalkSkipChildren, nil
}
