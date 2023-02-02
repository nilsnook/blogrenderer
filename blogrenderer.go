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
	// Title and Description
	_, err := fmt.Fprintf(w, "<h1>%s</h1>\n<p>%s</p>\n", p.Title, p.Description)
	if err != nil {
		return err
	}
	// Tags
	_, err = fmt.Fprint(w, "Tags: <ul>")
	if err != nil {
		return err
	}
	for _, t := range p.Tags {
		_, err = fmt.Fprintf(w, "<li>%s</li>", t)
		if err != nil {
			return err
		}
	}
	_, err = fmt.Fprint(w, "</ul>")
	if err != nil {
		return err
	}

	return nil
}
