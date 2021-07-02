package views

import (
	"fmt"
	"html/template"
	"path/filepath"
)

func NewView(layout string, files ...string) *View {
	files = append(files, layoutFiles()...)
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

/**
 * returns a slice of strings representing  layout of
 * files used in our application.
 */
func layoutFiles() []string {
	files, err := filepath.Glob("views/layouts/*.html")
	if err != nil {
		panic(err)
	}
	fmt.Print(files)
	return files
}
