// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.0 DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"//+dL591fjsAAA==",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}

