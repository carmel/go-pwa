package main

import (
	"github.com/carmel/go-pwa/pkg/analytics"
	"github.com/carmel/go-pwa/pkg/app"
)

type architecturePage struct {
	app.Compo
}

func newArchitecturePage() *architecturePage {
	return &architecturePage{}
}

func (p *architecturePage) OnNav(ctx app.Context) {
	p.initPage(ctx)
}

func (p *architecturePage) initPage(ctx app.Context) {
	ctx.Page().SetTitle("Understanding go-pwa Architecture")
	ctx.Page().SetDescription("Documentation about how go-pwa parts are working together to form a Progressive Web App (PWA).")
	analytics.Page("architecture", nil)
}

func (p *architecturePage) Render() app.UI {
	return newPage().
		Title("Architecture").
		Icon(fileTreeSVG).
		Index(
			newIndexLink().Title("Overview"),
			newIndexLink().Title("Web Browser"),
			newIndexLink().Title("Server"),
			newIndexLink().Title("HTML Pages"),
			newIndexLink().Title("Package Resources"),
			newIndexLink().Title("app.wasm"),
			newIndexLink().Title("Static Resources"),

			app.Div().Class("separator"),

			newIndexLink().Title("Next"),
		).
		Content(
			newRemoteMarkdownDoc().Src("/web/documents/architecture.md"),
		)
}
