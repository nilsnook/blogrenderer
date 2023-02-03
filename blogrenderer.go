package blogrenderer

import (
	"embed"
	"html/template"
	"io"
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

func Render(w io.Writer, p Post) error {
	// tmpl, err := template.New("blog").Parse(postTemplate)
	tmpl, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return err
	}

	if err := tmpl.Execute(w, p); err != nil {
		return err
	}

	return nil
}
