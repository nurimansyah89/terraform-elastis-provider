default: install

generate:
	go generate ./...

install:
	go install .

build:
	go build -o terraform-provider-elastis
	mkdir -p $(HOME)/.terraform.d/plugins/elastis.id/provider/elastis/$(VER)/$(ARCH)
	mv ./terraform-provider-elastis $(HOME)/.terraform.d/plugins/elastis.id/provider/elastis/$(VER)/$(ARCH)

test:
	go test -count=1 -parallel=4 ./...

testacc:
	TF_ACC=1 go test -count=1 -parallel=4 -timeout 10m -v ./...
