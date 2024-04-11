package main

import (
	"github.com/carmel/go-pwa/pkg/analytics"
	"github.com/carmel/go-pwa/pkg/app"
)

type migratePage struct {
	app.Compo
}

func newMigratePage() *migratePage {
	return &migratePage{}
}

func (p *migratePage) OnNav(ctx app.Context) {
	p.initPage(ctx)
}

func (p *migratePage) initPage(ctx app.Context) {
	ctx.Page().SetTitle("Migrate Codebase From go-pwa v8 To v9")
	ctx.Page().SetDescription("Documentation about what changed between go-pwa v8 and v9.")
	analytics.Page("migrate", nil)
}

func (p *migratePage) Render() app.UI {
	return newPage().
		Title("Migrate From v8 to v9").
		Icon(swapSVG).
		Index(
			newIndexLink().Title("Intro"),
			newIndexLink().Title("Changes"),
			newIndexLink().Title("    General"),
			newIndexLink().Title("    Components"),
			newIndexLink().Title("    Context"),
			newIndexLink().Title("API Design Decisions"),

			app.Div().Class("separator"),
		).
		Content(
			newRemoteMarkdownDoc().Src("/web/documents/migrate.md"),
		)
}