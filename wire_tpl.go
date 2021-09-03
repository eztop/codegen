package main

const (
	wireTpl = `
	//+build wireinject

	package main
	
	import (
		"github.com/google/wire"
		"{{.ProjectPkg}}/pkg/conf"
		{{range $idx,$each := .Services}}{{$each.Package}} "{{$.ProjectPkg}}/pkg/{{$each.Package}}/service"
		{{end}}
	)
	
	{{range $idx,$each := .Services}}
	// Create{{$each.Package}}Service Create{{$each.Package}}Service
	func Create{{$each.Package}}Service(conf *conf.Conf) {{$each.Package}}.Interface{
		panic(wire.Build({{$each.Package}}.SuperSet))
	}
	{{end}}
	`
)
