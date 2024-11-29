package blog_renderer_test

import (
	"blog_renderer"
	"bytes"
	"testing"
	"github.com/approvals/go-approval-tests"
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

	t.Run("it converts a single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}
		err := blog_renderer.Render(&buf, aPost)

		if err != nil{
			t.Fatal(err)
		}
		approvals.VerifyString(t, buf.String())
	})
}
