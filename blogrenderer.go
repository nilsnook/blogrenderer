package blogrenderer

import (
	"embed"
	"html/template"
	"io"
)

var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

type PostRenderer struct {
	tmpl *template.Template
}

func NewPostRenderer() (*PostRenderer, error) {
	t, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}

	return &PostRenderer{tmpl: t}, nil
}

func (r *PostRenderer) Render(w io.Writer, p Post) error {
	if err := r.tmpl.Execute(w, p); err != nil {
		return err
	}
	return nil
}

// func Render(w io.Writer, p Post) error {
// 	// tmpl, err := template.New("blog").Parse(postTemplate)
// 	tmpl, err := template.ParseFS(postTemplates, "templates/*.gohtml")
// 	if err != nil {
// 		return err
// 	}
//
// 	if err := tmpl.Execute(w, p); err != nil {
// 		return err
// 	}
//
// 	return nil
// }
