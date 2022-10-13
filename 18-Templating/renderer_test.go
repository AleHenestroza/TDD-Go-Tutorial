package blogrenderer_test

import (
	"bytes"
	// "io"
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

	postRenderer, err := blogrenderer.NewPostRenderer()
	if err != nil {
		t.Fatalf(err.Error())
	}

	t.Run("convert a single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}

		if err := postRenderer.Render(&buf, aPost); err != nil {
			t.Fatalf(err.Error())
		}

		approvals.VerifyString(t, buf.String())
	})

	t.Run("render the index with all posts", func(t *testing.T) {
		buf := bytes.Buffer{}
		posts := []blogrenderer.Post{{Title: "Hello World"}, {Title: "Hello World 2"}}

		if err := postRenderer.RenderIndex(&buf, posts); err != nil {
			t.Fatalf(err.Error())
		}

		approvals.VerifyString(t, buf.String())
	})
}

// func BenchmarkRender(b *testing.B) {
// 	var (
// 		aPost = blogrenderer.Post{
// 			Title:       "hello world",
// 			Body:        "This is a post",
// 			Description: "This is a description",
// 			Tags:        []string{"go", "tdd"},
// 		}
// 	)

// 	postRenderer, err := blogrenderer.NewPostRenderer()

// 	if err != nil {
// 		b.Fatal(err)
// 	}

// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		postRenderer.Render(io.Discard, aPost)
// 	}
// }
