新建$HOME/.codegen,内容如下
```
{
    "git_prefix": "git@github.com:eztop",
    "pkg_prefix": "github.com/eztop",
    "apidoc": "apidoc",
    "dir": "golang",
    "docker_registry_dev": "myregistry/dev",
    "docker_registry_qa": "myregistry/qa",
    "docker_registry_pl": "myregistry/pl",
    "docker_registry_ol": "myregistry/ol"
}
```

编译二进制
```
git clone git@github.com:eztop
cd codegen && go install
```

使用
```
codegen demo

```
得到如下目录结构
```
.
├── Dockerfile
├── Makefile
├── cmd
│   └── demo
│       ├── main.go
│       ├── wire.go
│       └── wire_gen.go
├── conf
│   └── conf.toml
├── go.mod
├── go.sum
└── pkg
    ├── conf
    │   └── conf.go
    └── v1
        ├── model
        ├── repo
        │   └── repo.go
        └── service
            └── service.go
```