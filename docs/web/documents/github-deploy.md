## Intro

[GitHub Pages](https://pages.github.com) is a pretty easy way to have a running and available website. The limitation is that only static websites can be hosted.

PWAs created with go-pwa are usually served by a [standard Go HTTP server](https://pkg.go.dev/net/http#Server). That said, **it is possible to generate them as completely static websites in order to have them running from hosts like GitHub Pages**.

## Generate a Static Website

A static website can be generated with the help of the [GenerateStaticWebsite()](/reference#GenerateStaticWebsite) function:

```go
func main() {
	app.Route("/", &hello{})
	app.Route("/hello"), &hello{})
	app.RunWhenOnBrowser()

	err := app.GenerateStaticWebsite(".", &app.Handler{
		Name:        "Hello",
		Description: "An Hello World! example",
    })

    if err != nil {
        log.Fatal(err)
    }
}
```

It can then be generated by building and running the go executable:

```sh
# Build app.wasm:
GOARCH=wasm GOOS=js go build -o /test-app/web/app.wasm

# Build and generate static website:
go build
./hello
```

The current directory will then have all the required files to run the PWA as a static website:

```bash
.                        # Current dir.
├── app-worker.js        # Service-worker file (Generated).
├── app.js               # Js support file (Generated).
├── index.html           # Index page (Generated from "/").
├── hello.html           # hello page (Generated from "/hello").
├── manifest.webmanifest # PWA manifest (Generated).
├── wasm_exec.js         # Wasm support file (Generated).
└── web                  # Web directory.
    └── app.wasm         # Wasm app (Manually built).
```

## Deployment

Once generated, the static website can directly be dropped in a GitHub repository, either in the root or in the `docs` directory depending on how [GitHub Pages](https://pages.github.com) is configured.

### Domainless Repository

By default, no domain is associated with a GitHub repository. The website is accessed from a GitHub address that looks like this:

```sh
https://USERNAME.github.io/REPOSITORY_NAME
```

In that scenario, the [Handler](/reference#Handler) resource provider must be changed to be a [GitHubPages](/reference#GitHubPages) provider:

```go
func main() {
	app.Route("/", &hello{})
	app.Route("/hello"), &hello{})
	app.RunWhenOnBrowser()

	err := app.GenerateStaticWebsite(".", &app.Handler{
		Name:        "Hello",
		Description: "An Hello World! example",
		Resources:   app.GitHubPages("REPOSITORY_NAME"),
    })

    if err != nil {
        log.Fatal(err)
    }
}
```

This will fix static resources issues caused by GitHub adding the repository name as an URL path prefix.