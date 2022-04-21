package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/SargntSprinkles/lion-turtle-api/config"
	"github.com/sirupsen/logrus"
)

var server *http.Server

func Serve() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/ping", ping)
	mux.HandleFunc("/gql", gql)
	mux.Handle("/ui", playground.Handler("playground", "/gql"))

	port := config.Port()
	server = &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}
	logrus.Infof("listening on %s", port)
	logrus.Fatal(server.ListenAndServe())
}

func index(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		writer.WriteHeader(http.StatusOK)
		fmt.Fprint(writer, "Hello! You have found the Lion Turtle API for the Avatar Legends RPG. Please use /gql to make requests or /ui for API docs and testing. Thanks! -SargntSprinkles")
	default:
		http.NotFound(writer, request)
	}
}

func ping(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		writer.WriteHeader(http.StatusOK)
		fmt.Fprint(writer, "pong")
	default:
		http.NotFound(writer, request)
	}
}

func gql(w http.ResponseWriter, r *http.Request) {
	resolverToolChain := &Resolver{}
	opts := []graphql.SchemaOpt{}

	schema := graphql.MustParseSchema(resolverToolChain.schema(), resolverToolChain, opts...)
	gqlHandler := &relay.Handler{Schema: schema}
	ctx := r.Context()
	ctx = context.WithValue(ctx, "request", r)
	ctx = context.WithValue(ctx, "response", w)
	r = r.WithContext(ctx)
	gqlHandler.ServeHTTP(w, r)
}
