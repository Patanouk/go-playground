package templating

import (
	"bytes"
	"github.com/approvals/go-approval-tests"
	"go-playground/blog"
	"io"
	"testing"
)

func TestRender(t *testing.T) {
	var (
		aPost = blog.Post{
			Title:       "hello world",
			Body:        "#This is a test  `code`",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)

	renderer, err := NewBlogRenderer()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("it converts a single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}

		if err := renderer.Render(&buf, &aPost); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})
}

func BenchmarkRender(b *testing.B) {
	var (
		aPost = blog.Post{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)

	renderer, err := NewBlogRenderer()
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		renderer.Render(io.Discard, &aPost)
	}
}
