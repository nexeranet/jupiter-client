generate-client:
	oapi-codegen -package client -generate client,types ./client/api.yaml > ./client/client.gen.go
