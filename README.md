# schablone-server

## Generate OpenApi code
~/go/bin/oapi-codegen -generate types schablone-api.yaml > api/types.gen.go
~/go/bin/oapi-codegen -generate server swagger.yaml > api/schablone.gen.go
