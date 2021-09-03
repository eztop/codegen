package main

const (
	serviceTpl = `
	package service
	
	import (
		"context"
		"github.com/google/wire"
		"{{.ProjectPkg}}/pkg/{{.Package}}/repo"
		"{{.ProjectPkg}}/pkg/conf"
		"{{.APIDocPkg}}"
	)
	// Interface Interface
	type Interface interface {
		{{range $idx,$each := .APIs}}{{$each.Name}}(ctx context.Context, req *{{$.ProjectName}}.{{$each.Req}}) (*{{$.ProjectName}}.{{$each.Resp}}, error)
		{{end}}
	}

	type service struct {
		repo repo.Interface
		conf *conf.Conf
	}

	// SuperSet SuperSet
	var SuperSet = wire.NewSet(MustNew, repo.SuperSet)

	// MustNew MustNew
	func MustNew(repo repo.Interface, conf *conf.Conf) Interface{
		return &service{
			repo:repo,
			conf:conf,
		}
	}
	{{range $idx,$each := .APIs}}
	func(s *service) {{$each.Name}}(ctx context.Context, req *{{$.ProjectName}}.{{$each.Req}}) (*{{$.ProjectName}}.{{$each.Resp}}, error){
		return &{{$.ProjectName}}.{{$each.Resp}}{},nil
	}
	{{end}}
	`
)
