default: build

vet:
	@go fmt github.com/adamdecaf/terraform-provider-nearlyfreespeech
	go fmt github.com/adamdecaf/terraform-provider-nearlyfreespeech/nearlyfreespeech
	go tool vet .

dev: vet
	@rm -f terraform-provider-nearlyfreespeech*
	go build -o terraform-provider-nearlyfreespeech github.com/adamdecaf/terraform-provider-nearlyfreespeech
	@chmod +x terraform-provider-nearlyfreespeech
	@ln -f -s $(shell pwd)/terraform-provider-nearlyfreespeech ${GOPATH}/bin/terraform-provider-nearlyfreespeech

build: vet
	@rm -f terraform-provider-nearlyfreespeech*
	GOOS=darwin GOARCH=amd64 go build -o terraform-provider-nearlyfreespeech-osx github.com/adamdecaf/terraform-provider-nearlyfreespeech
	GOOS=linux GOARCH=amd64 go build -o terraform-provider-nearlyfreespeech-linux github.com/adamdecaf/terraform-provider-nearlyfreespeech
	@chmod +x terraform-provider-nearlyfreespeech-*
