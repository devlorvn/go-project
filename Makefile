check_install_swagger:
	which swagger | GO111MODULE=off go get -u github.com/go-swagger/go-swagger/cmd/swagger@latest

swagger:
	GO111MODULE=off swagger generate spec -o ./swagger.yaml --scan-models