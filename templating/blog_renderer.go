package templating

import (
	"embed"
	"github.com/gomarkdown/markdown"
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

func (r blogRenderer) RenderIndex(w io.Writer, posts []blog.Post) error {

	if err := r.template.ExecuteTemplate(w, "index.gohtml", posts); err != nil {
		return err
	}

	return nil
}

func (r blogRenderer) RenderPost(w io.Writer, p *blog.Post) error {
	html := markdown.ToHTML([]byte(p.Body), nil, nil)

	if err := r.template.ExecuteTemplate(w, "blog.gohtml", blog.Post{
		Title:       p.Title,
		Description: p.Description,
		Body:        string(html),
		Tags:        p.Tags,
	}); err != nil {
		return err
	}

	return nil
}
