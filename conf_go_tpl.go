package main

const (
	confGoTpl = "package conf\n import (\n\"github.com/eztop/go_common/ezrpc\"\n\"github.com/eztop/go_common/tracing\"\n\"github.com/eztop/go_common/confx\"\n\"github.com/eztop/go_common/gormx\"\n)\n // Conf Conf\ntype Conf struct {\nServer *ezrpc.ServerConf `toml:\"server\"`\nTracing *tracing.Conf `toml:\"tracing\"`\nDB *gormx.Conf `toml:\"db\"`\n} \n//MustParse MustParse\nfunc MustParse(cfgPath string) *Conf {\nconf:=&Conf{}\nconfx.MustParseFile(cfgPath,conf)\nreturn conf\n}\n"
)
