default: build

vet:
	go fmt github.com/adamdecaf/terraform-provider-nearlyfreespeech/nearlyfreespeech
	go tool vet .

dev: vet
	go build -o terraform-provider-nearlyfreespeech github.com/adamdecaf/terraform-provider-nearlyfreespeech/nearlyfreespeech

build: vet
	GOOS=darwin GOARCH=amd64 go build -o terraform-provider-nearlyfreespeech-osx github.com/adamdecaf/terraform-provider-nearlyfreespeech/nearlyfreespeech
	GOOS=linux GOARCH=amd64 go build -o terraform-provider-nearlyfreespeech-linux github.com/adamdecaf/terraform-provider-nearlyfreespeech/nearlyfreespeech
