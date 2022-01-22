package api

import (
	"net/http"
)

type SchabloneServer struct{}

// Make sure we conform to ServerInterface
// var _ ServerInterface = (*SchabloneServer)(nil)

// Create new server
func NewSchabloneServer() *SchabloneServer {
	return &SchabloneServer{}
}

func (*SchabloneServer) GetCategoryCreate(w http.ResponseWriter, r *http.Request) {
}
