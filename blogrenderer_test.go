package blogrenderer_test

import (
	"bytes"
	"testing"

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

	t.Run("Convert post to HTML", func(t *testing.T) {
		buf := bytes.Buffer{}
		err := blogrenderer.Render(&buf, aPost)

		if err != nil {
			t.Fatal(err)
		}

		got := buf.String()
		want := `<h1>hello world</h1>
<p>This is a description</p>
Tags: <ul><li>go</li><li>tdd</li></ul>
`
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
