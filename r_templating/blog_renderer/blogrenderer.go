package blog_renderer

import (
	"embed"
	"html/template"
	"io"
)

type Post struct {
	Title, Description, Body string
	Tags                     []string
}

var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

type PostRenderer struct {
	templ *template.Template
}

func NewPostRenderer() (*PostRenderer, error){
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	return &PostRenderer{templ}, err
}

func (r *PostRenderer) Render(w io.Writer, p Post) error {
	err := r.templ.ExecuteTemplate(w, "blog.gohtml",p);
	return err
}
