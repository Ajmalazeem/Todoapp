package api

import (
	"context"
	"encoding/json"
	"net/http"
	"github.com/Ajmalazeem/models"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func decodePostTodoRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req models.PostTodoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func makePostTodoEndpoint(svc Todo) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {

		req := request.(models.PostTodoRequest)
		svc.PostTodo(req)
		return nil, nil

	}
}

func decodeGetTodoRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req models.GetTodoRequest
	vars := mux.Vars(r)
	id := vars["id"]
	//var err error
	req.Id = id
	return req, nil
}


func makeGetTodoEndpoint(svc Todo) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(models.GetTodoRequest)
		return svc.GetTodo(req)
	}
}

func decodePutTodoRequest(_ context.Context,r *http.Request) (interface{},error){
	var req models.PutTodoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	vars := mux.Vars(r)
	id := vars["id"]
	req.Id = id
	return req, nil
}

func makePutTodoEndpoint(svc Todo) endpoint.Endpoint{
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(models.PutTodoRequest)
		svc.PutTodo(req)
		return nil, nil
	}
}

func decodeDeleteTodoRequest(_ context.Context,r *http.Request)(interface{},error){
	var req models.DeleteTodoRequest
	vars := mux.Vars(r)
	id := vars["id"]
	req.Id = id 
	return req ,nil
}

func makeDeleteTodoEndpoint(svc Todo) endpoint.Endpoint{
	return func(_ context.Context, request interface{}) (response interface{}, err error) {
		req := request.(models.DeleteTodoRequest)
		svc.DeleteTodo(req)
		return nil,nil
	}
}


func MakeHandler(svc Todo) http.Handler {
	r := mux.NewRouter()
	PostHandler := httptransport.NewServer(
		makePostTodoEndpoint(svc),
		decodePostTodoRequest,
		encodeResponse,
	)
	GetHandler := httptransport.NewServer(
		makeGetTodoEndpoint(svc),
		decodeGetTodoRequest,
		encodeResponse,
	)
	PutHandler := httptransport.NewServer(
		makePutTodoEndpoint(svc),
		decodePutTodoRequest,
		encodeResponse,
	)
	DeleteHandler := httptransport.NewServer(
		makeDeleteTodoEndpoint(svc),
		decodeDeleteTodoRequest,
		encodeResponse,
	)

	r.Methods(http.MethodPost).Path("/todoer").Handler(PostHandler)
	r.Methods(http.MethodGet).Path("/todoer/{id}").Handler(GetHandler)
	r.Methods(http.MethodPut).Path("/todoer/{id}").Handler(PutHandler)
	r.Methods(http.MethodDelete).Path("todoer/{id}").Handler(DeleteHandler)

	return r
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
