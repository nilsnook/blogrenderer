package blogrenderer

import (
	"fmt"
	"io"
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

func Render(w io.Writer, p Post) error {
	fmt.Fprintf(w, "<h1>%s</h1>", p.Title)
	return nil
}
