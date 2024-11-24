package routes

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/benleem/show-beam/internal/templates"
	"github.com/benleem/show-beam/internal/templates/pages"
)

type IndexHandler struct{}

func NewIndexHandler() *IndexHandler {
	return &IndexHandler{}
}

func (h *IndexHandler) Get(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	var page templ.Component
	// hxReq := c.Request().Header.Get("Hx-Request")
	// if hxReq != "" {
	// 	page = pages.Home(true)
	// 	c.Response().Header().Set(echo.HeaderVary, "Hx-Request")
	// 	return page.Render(context.Background(), c.Response().Writer)
	// }
	page = pages.Home()
	templates.Layout(page, "showbeam - home").Render(r.Context(), w)
}
