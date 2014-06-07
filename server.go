package main

import (
  "github.com/go-martini/martini"
  "github.com/codegangsta/martini-contrib/render"
)

func main() {
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
    data := struct {
      URL_string string
      JS_string string
    }{
      params["_1"],
      "function collide(e){var t=e.radius+16,n=e.x-t,r=e.x+t,i=e.y-t,s=e.y+t;return function(t,o,u,a,f){if(t.point&&t.point!==e){var l=e.x-t.point.x,c=e.y-t.point.y,h=Math.sqrt(l*l+c*c),p=e.radius+t.point.radius;if(h<p){h=(h-p)/h*.5;e.x-=l*=h;e.y-=c*=h;t.point.x+=l;t.point.y+=c}}return o>r||a<n||u>s||f<i}}var w=940,h=600;var nodes=d3.range(200).map(function(){return{radius:Math.random()*12+4}}),color=d3.scale.category20();var force=d3.layout.force().gravity(.05).charge(function(e,t){return t?0:-2e3}).nodes(nodes).size([w,h]);var root=nodes[0];root.radius=0;root.fixed=true;force.start();var svg=d3.select('#screen').append('svg:svg').attr('width',w).attr('height',h);svg.selectAll('circle').data(nodes.slice(1)).enter().append('svg:circle').attr('r',function(e){return e.radius-2}).style('fill',function(e,t){return color(t%4)});force.on('tick',function(e){var t=d3.geom.quadtree(nodes),n=0,r=nodes.length;while(++n<r){t.visit(collide(nodes[n]))}svg.selectAll('circle').attr('cx',function(e){return e.x}).attr('cy',function(e){return e.y})});svg.on('mousemove',function(){var e=d3.mouse(this);root.px=e[0];root.py=e[1];force.resume()})",
    }
    r.HTML(200, "d", data)
  })

  m.Run()
}
