package blogrenderer_test

import (
	"bytes"
	"io"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
	"github.com/nilsnook/blogrenderer"
)

func TestRenderer(t *testing.T) {
	var (
		aPost = blogrenderer.Post{
			Title:       "hello world",
			Description: "This is a description",
			Body:        "This is a post",
			Tags:        []string{"go", "tdd"},
		}
	)

	postRenderer, err := blogrenderer.NewPostRenderer()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("Convert single post to HTML", func(t *testing.T) {
		buf := bytes.Buffer{}

		if err := postRenderer.Render(&buf, aPost); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})

	t.Run("Renders an index of posts", func(t *testing.T) {
		buf := bytes.Buffer{}
		posts := []blogrenderer.Post{{Title: "Untitled Player"}, {Title: "Little Sunshine"}}

		if err := postRenderer.RenderIndex(&buf, posts); err != nil {
			t.Fatal(err)
		}

		got := buf.String()
		want := `<ol><li><a href="/post/untitled-player">Untitled Player</a></li><li><a href="/post/little-sunshine">Little Sunshine</a></li></ol>`

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

func BenchmarkRender(b *testing.B) {
	var (
		aPost = blogrenderer.Post{
			Title:       "hello world",
			Description: "This is a description",
			Body:        "This is a post",
			Tags:        []string{"go", "tdd"},
		}
	)

	postRenderer, err := blogrenderer.NewPostRenderer()
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		postRenderer.Render(io.Discard, aPost)
	}
}
