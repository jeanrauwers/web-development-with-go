package views

import (
	"html/template"
	"path/filepath"
)

func NewView(layout string, files ...string) *View {
	files = append(files, listLayoutFiles()...)

	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &View{
		Template: t,
		Layout:   layout,
	}
}

type View struct {
	Template *template.Template
	Layout   string
}

func listLayoutFiles() []string {
	files, err := filepath.Glob("views/layouts/*.gohtml")
	if err != nil {
		panic(err)
	}
	return files
}
