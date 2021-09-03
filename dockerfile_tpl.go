package main

const (
	dockerfileTpl = `
FROM alpine:latest
WORKDIR /app
RUN sed -i s/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g /etc/apk/repositories
RUN apk add tzdata
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && echo "Asia/Shanghai" > /etc/timezone
	
COPY {{.ProjectName}} /app/
RUN chmod +x /app/{{.ProjectName}}
	
CMD ["/app/{{.ProjectName}}"]
`
)
