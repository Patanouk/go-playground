package templating

import (
	"embed"
	"go-playground/blog"
	"html/template"
	"io"
)

var (
	//go:embed "templates/*.gohtml"
	blogTemplate embed.FS
)

type blogRenderer struct {
	template *template.Template
}

func NewBlogRenderer() (*blogRenderer, error) {
	blogTemplate, err := template.ParseFS(blogTemplate, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}
	return &blogRenderer{template: blogTemplate}, nil
}

func (r blogRenderer) Render(w io.Writer, p *blog.Post) error {

	if err := r.template.Execute(w, p); err != nil {
		return err
	}

	return nil
}
