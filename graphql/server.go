package graphql

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/websocket"
	"github.com/tsisar/extended-log-go/log"
	"github.com/tsisar/starknet-indexer/generated/ent"
	"github.com/tsisar/starknet-indexer/internal/config"
	"github.com/vektah/gqlparser/v2/ast"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"gorm.io/gorm"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
)

const defaultPort = "8080"

func Start(_ context.Context, db *gorm.DB, client *ent.Client) error {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	query := "/query"
	indexerName := config.App.IndexerName
	if indexerName != "" {
		query = fmt.Sprintf("/%s/query", indexerName)
	}

	schema := NewExecutableSchema(Config{Resolvers: &Resolver{
		Client: client,
		DB:     db,
	}})
	srv := handler.New(schema)
	srv.AddTransport(&transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true // Allow all origins for WebSocket connections
			},
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
		KeepAlivePingInterval: 10 * time.Second,
	})
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{})

	srv.Use(extension.Introspection{})
	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	//srv.AroundOperations(func(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
	//	op := graphql.GetOperationContext(ctx)
	//	log.Debugf("Operation: %s\nQuery:\n%s\n", op.OperationName, op.RawQuery)
	//	return next(ctx)
	//})
	srv.SetRecoverFunc(func(ctx context.Context, err interface{}) error {
		log.Debugf("Panic occurred: %v", err)
		return gqlerror.Errorf("Internal server error")
	})

	// === GraphQL IDEs ===
	http.Handle("/", playground.Handler(
		"GraphiQL Playground",
		query,
		playground.WithGraphiqlEnablePluginExplorer(true),
	))
	http.Handle("/apollo", playground.ApolloSandboxHandler(
		"Apollo Sandbox",
		query,
	))
	http.Handle("/altair", playground.AltairHandler(
		"Altair GraphQL",
		query,
		map[string]any{},
	))

	// === GraphQL endpoint ===
	http.Handle(query, srv)

	log.Debugf("GraphQL server ready at:\n"+
		"- http://localhost:%s%s/		 (GraphiQL with Explorer)\n"+
		"- http://localhost:%s%s/apollo  (Apollo Sandbox)\n"+
		"- http://localhost:%s%s/altair  (Altair)\n", port, query, port, query, port, query)

	return http.ListenAndServe(":"+port, nil)
}
