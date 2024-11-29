package blog_renderer_test

import (
	"blog_renderer"
	"bytes"
	"io"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
)

func TestRender(t *testing.T) {
	var (
		aPost = blog_renderer.Post{
			Title: "Hello World",
			Body: "This is a post",
			Description: "This is a description",
			Tags: []string{"go", "tdd"},
		}
	)

	postRenderer, err := blog_renderer.NewPostRenderer();

	if err != nil {
		t.Fatal(err)
	}

	t.Run("it converts a single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}
		err := postRenderer.Render(&buf, aPost)

		if err != nil{
			t.Fatal(err)
		}
		approvals.VerifyString(t, buf.String())
	})
}


func BenchmarkRender(b *testing.B) {
	var (
		aPost = blog_renderer.Post{
			Title: "Hello World",
			Body: "This is a post",
			Description: "This is a description",
			Tags: []string{"go", "tdd"},
		}
	)

	postRenderer, err := blog_renderer.NewPostRenderer();
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i:=0; i<b.N; i++{
		postRenderer.Render(io.Discard, aPost)
	}
}