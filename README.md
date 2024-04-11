## go-pwa

> Notice: this is a fork of [maxence-charriere/go-app](https://github.com/maxence-charriere/go-app).

go-pwa is a package for **building progressive web apps (PWA)** with the [Go programming language (Golang)](https://golang.org) and [WebAssembly (Wasm)](https://webassembly.org).

Shaping a UI is done by using a **[declarative syntax](https://go-pwa.dev/declarative-syntax) that creates and compose HTML elements only by using the Go programing language**.

It **uses [Go HTTP standard](https://golang.org/pkg/net/http) model**.

An app created with go-pwa can out of the box **runs in its own window**, **supports offline mode**, and is **SEO friendly**.

## Install

**go-pwa** requirements:

- [Go 1.18](https://golang.org/doc/go1.17) or newer
- [Go module](https://github.com/golang/go/wiki/Modules)

```sh
go mod init
go get -u github.com/carmel/go-pwa
```

## Declarative syntax

go-pwa uses a `declarative syntax` so you can **write reusable component-based UI elements** just by using the Go programming language.

Here is a Hello World component that takes an input and displays its value in its title:

```go
type hello struct {
	app.Compo

	name string
}

func (h *hello) Render() app.UI {
	return app.Div().Body(
		app.H1().Body(
			app.Text("Hello, "),
			app.If(h.name != "",
				app.Text(h.name),
			).Else(
				app.Text("World!"),
			),
		),
		app.P().Body(
			app.Input().
				Type("text").
				Value(h.name).
				Placeholder("What is your name?").
				AutoFocus(true).
				OnChange(h.ValueTo(&h.name)),
		),
	)
}
```

## Standard HTTP

Apps created with go-pwa complies with [Go standard HTTP](https://golang.org/pkg/net/http) package interfaces.

```go
func main() {
	// Components routing:
	app.Route("/", &hello{})
	app.Route("/hello", &hello{})
	app.RunWhenOnBrowser()

	// HTTP routing:
	http.Handle("/", &app.Handler{
		Name:        "Hello",
		Description: "An Hello World! example",
	})

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
```

## Sample

- [Lofimusic.app](https://lofimusic.app/collegemusic-lonely)
- [Murlok.io](https://murlok.io/)
- [liwasc](https://pojntfx.github.io/liwasc/)
- [go-pwa Docs](https://go-pwa.dev/)
