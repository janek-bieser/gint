# GinT

GinT is an HTML template plug-in for the great
[Gin](https://github.com/gin-gonic/gin) framework. It is somewhat
opinionated about how templates are organized and rendered, but in
exchange you wont't have to write a lot of boilerplate code.

## Getting Started

```go
// example.go
import (
    "github.com/gin-gonic/gin"
    "github.com/janek-bieser/gint"
)

func main() {
    r := gin.Default()

    // create the renderer and plug it into the gin framework
    htmlRender := gint.NewHTMLRender()
    r.HTMLRender = htmlRender

    r.GET("/", func(c *gin.Context){
        c.HTML(http.StatusOK, "home", gin.H{"title": "Test"})
    })
    
    r.Run()
}
```
```html
<!-- templates/layout.tmpl -->
<html>
    <head>
        <title>{{ .title }}</title>
    </head>
    <body>
        <h1>Example</h1>
        {{ template "content" . }}
    <body>
</html>
```
```html
<!-- templates/home.tmpl -->
<h2>Home</h2>
```
The call to `c.HTML(http.StatusOK, "home", gin.H{"title": "Test"})` will render the following HTML:

```html
<html>
    <head>
        <title>Test</title>
    </head>
    <body>
        <h1>Example</h1>
        <h2>Home</h2>
    <body>
</html>
```

## How it works

GinT expects you to organize all your HTML templates in a single
folder. Your `TemplateDir` needs to contain a `layout.*` file
which defines your application layout. The file extension of your
templates is configurable (default is `tmpl`), so it could be called
`layout.html, layout.tpl...`. Every time you call `.HTML(status,
templateName, data)` on the `gin.Context` object, GinT will look
inside the templates folder and render `templateName` inside your
`layout`. In order for this to work, you have to render the content
template inside your layout using `{{ template "content" . }}` or `{{
block "content" . }}`
