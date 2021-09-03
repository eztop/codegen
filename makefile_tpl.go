package main

const (
	makefielTpl = `tag?=v1
dev: clean build
	@./{{.ProjectName}}
build: fmt vet
	@cd cmd/{{.ProjectName}} && wire && go build -o ../../ && cd ../..
build-dev: build
	@docker build -t {{.DockerRegistryDEV}}/{{.ProjectName}} .
	@docker push {{.DockerRegistryDEV}}/{{.ProjectName}}
build-qa: build
	@docker build -t {{.DockerRegistryQA}}/{{.ProjectName}} .
	@docker push {{.DockerRegistryQA}}/{{.ProjectName}}
build-pl: build
	@docker build -t {{.DockerRegistryPL}}/{{.ProjectName}}:${tag} .
	@docker push {{.DockerRegistryPL}}/{{.ProjectName}}:${tag}
build-ol: build
	@docker build -t {{.DockerRegistryOL}}/{{.ProjectName}}:${tag} .
	@docker push {{.DockerRegistryOL}}/{{.ProjectName}}:${tag}
clean:
	@rm -rf {{.ProjectName}}
fmt:
	@go fmt ./...
vet:
	@go vet ./...
`
)