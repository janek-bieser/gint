# GinT

GinT is an HTML template plug-in for the great
[Gin](https://github.com/gin-gonic/gin) framework. It is somewhat
opinionated about how templates are organized and rendered, but in
exchange you wont't have to write a lot of boilerplate code.

## How to use it

```go
import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default()

    // create the renderer and plug it into the gin framework
    htmlRender := gint.NewHTMLRender()
    r.HTMLRender = htmlRender

    r.GET("/", func(c *gin.Context){
        c.HTML(http.StatusOK, "home")
    })
    
    r.Run()
}
```

## How it works

GinT expects you to organize all your HTML templates in a single
folder. Your `templateRoot` needs to contain a `layout.*` file
which defines your application layout. The file extension of your
templates is configurable, so it could be called
`layout.html, layout.tpl...`. Every time you call `.HTML(status, templateName)` on
the `gin.Context` object, GinT will look inside the templates folder and
render `templateName` inside your `layout`.
