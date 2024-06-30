package parser

import (
	"bytes"
	"html/template"

	katex "github.com/FurqanSoftware/goldmark-katex"
	callout "github.com/VojtaStruhar/goldmark-obsidian-callout"
	chromahtml "github.com/alecthomas/chroma/v2/formatters/html"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
	"go.abhg.dev/goldmark/mermaid"
	"go.abhg.dev/goldmark/wikilink"
)

type MdParser struct {
}

func (mdp MdParser) Convert(mdFile []byte) (template.HTML, error) {
	var buffer bytes.Buffer
	err := goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,
			&wikilink.Extender{},
			highlighting.NewHighlighting(
				highlighting.WithStyle("catppuccin-mocha"),
				highlighting.WithFormatOptions(
					chromahtml.WithLineNumbers(true),
				),
			),
			&mermaid.Extender{},
			&katex.Extender{},
			callout.ObsidianCallout,
		),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithXHTML(),
		),
	).Convert(mdFile, &buffer)

	if err != nil {
		return template.HTML(""), err
	}
	return template.HTML(buffer.String()), nil
}

func NewMdParser() MdParser {
	return MdParser{}
}
