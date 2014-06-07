package main

import (
  "github.com/go-martini/martini"
  "github.com/codegangsta/martini-contrib/render"
  "os/exec"
  "fmt"
)

func main() {
  fmt.Println("fmt included for debug. remove for production.")
  m := martini.Classic()
  m.Use(render.Renderer(render.Options{
    Directory: "template",
    Layout: "layout",
    Extensions: []string{".html"},
  }))

  m.Get("/", func(r render.Render) {
    random_string := RandomString(10, 30)
    r.HTML(200, "main", "/d/" + random_string)
  })

  m.Get("/d/**", func(params martini.Params, r render.Render) {
    cmd := exec.Command("ruby", "./interpreter/interpreter.rb", string(params["_1"]))
    translation, _ := cmd.Output()
    data := struct {
      URL_string string
      JS_string string
    }{
      params["_1"],
      string(translation),
    }
    r.HTML(200, "d", data)
  })

  m.Run()
}
