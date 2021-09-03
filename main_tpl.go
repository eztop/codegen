package main

const (
	mainTpl = `
	package main

	import (
		"context"
		"flag"
		"github.com/eztop/go_common/ezrpc"
		"github.com/eztop/go_common/app"
		"github.com/eztop/go_common/log"
		"github.com/eztop/go_common/tracing"
		"{{$.ProjectPkg}}/pkg/conf"
		"{{.APIDocPkg}}"
	)
	var cfgPath = flag.String("config", "./conf/conf.toml", "-conf=./conf/conf.toml")
	
	func main() {
		flag.Parse()
		conf := conf.MustParse(*cfgPath)
		cleanup := tracing.MustInitGlobalTracer("{{$.ProjectName}}", conf.Tracing)
		defer cleanup()
		server := ezrpc.NewServer("{{$.ProjectName}}",conf.Server)
		{{range $idx,$each := .Services}}{{$each.Package}}Service:=Create{{$each.Package}}Service(conf)
		{{$.ProjectName}}.Register{{$each.Name}}RPCServer(server, {{$each.Package}}Service)
		{{end}}
		app.GoAsync(func(){
			if err := server.Start(); err != nil {
				log.WithError(err).
					Fatal("failed to Start")
			}
		})
		app.AddShutdownHook(func() {
			server.Stop(context.TODO())
		})
		app.Wait()
	}
	`
)
