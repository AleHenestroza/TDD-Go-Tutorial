package blogrenderer_test

import (
	"bytes"
	"testing"

	approvals "github.com/approvals/go-approval-tests"

	blogrenderer "go-tdd-tutorial/18-Templating"
)

func TestRender(t *testing.T) {
	var (
		aPost = blogrenderer.Post{
			Title:       "Hello, world",
			Description: "This is the description",
			Body:        "This is a post",
			Tags:        []string{"go", "tdd"},
		}
	)

	t.Run("convert a single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}
		err := blogrenderer.Render(&buf, aPost)

		if err != nil {
			t.Fatalf(err.Error())
		}

		approvals.VerifyString(t, buf.String())
	})
}
