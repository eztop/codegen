package main

const (
	confTpl = "[server]\naddr= \"127.0.0.1:8080\"\n[tracing]\nendpoint=\"http://127.0.0.1:14268/api/traces?format=jaeger.thrift\"\nsample_rate=1.0\n[db]\nuser=\"root\"\npassword=\"root\"\nhost=\"127.0.0.1\"\nport=3306\ndb_name=\"db\"\ncharset=\"utf8mb4\"\nmax_open_conns=10\nmax_idle_conns=10\nconn_max_lifetime=3600\n"
)
