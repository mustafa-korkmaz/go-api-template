check-swagger:
	which swagger || (GO111MODULE=off go get -u github.com/go-swagger/go-swagger/cmd/swagger)

swagger: check-swagger
	C:\Users\mustafa.korkmaz\.gopath\bin\swagger.exe generate spec -o .\third_party\swagger\swaggerui\swagger.yaml --scan-models

serve-swagger: check-swagger
	C:\Users\mustafa.korkmaz\.gopath\bin\swagger.exe serve -F=swagger swagger.yaml
