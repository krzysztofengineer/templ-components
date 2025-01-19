package handler

import (
	"net/http"
	"templ-components/pages"
)

type Home struct {
}

func NewHome() *Home {
	return &Home{}
}

func (h *Home) Index(w http.ResponseWriter, r *http.Request) {
	pages.Home().Render(r.Context(), w)
}
