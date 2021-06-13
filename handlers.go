package main

import (
	"github.com/unrolled/render"
	"net/http"
)

type Handlers struct {
	Render       *render.Render
	LayoutRender *render.Render
}

func BaseHandlers() *Handlers {
	lr := render.New(render.Options{
		Layout:     "layout",
		Extensions: []string{".tmpl", ".gohtml"}, // Sp
	})

	r := render.New(render.Options{})

	return &Handlers{LayoutRender: lr, Render: r}
}

func (h Handlers) Landing(w http.ResponseWriter, req *http.Request) {
	h.Render.HTML(w, http.StatusOK, "landing", nil)
}
