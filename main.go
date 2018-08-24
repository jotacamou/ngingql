package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/jotacamou/ngingql/db"
	"github.com/jotacamou/ngingql/resolver"

	"io/ioutil"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

var schema *graphql.Schema

func init() {
	schemaFile, err := ioutil.ReadFile("schema.graphql")
	if err != nil {
		panic(err)
	}

	schema = graphql.MustParseSchema(string(schemaFile), &resolver.Resolver{})

}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write(page)
	})

	r.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		gql := &relay.Handler{Schema: schema}
		ctx := context.Background()
		gql.ServeHTTP(w, r.WithContext(ctx))
	})

	routeLogger := handlers.LoggingHandler(os.Stdout, r)

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", routeLogger))

}

var page = []byte(`
<!DOCTYPE html>
<html>
	<head>
		<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/graphiql/0.10.2/graphiql.css" />
		<script src="https://cdnjs.cloudflare.com/ajax/libs/fetch/1.1.0/fetch.min.js"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/react/15.5.4/react.min.js"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/react/15.5.4/react-dom.min.js"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/graphiql/0.10.2/graphiql.js"></script>
	</head>
	<body style="width: 100%; height: 100%; margin: 0; overflow: hidden;">
		<div id="graphiql" style="height: 100vh;">Loading...</div>
		<script>
			function graphQLFetcher(graphQLParams) {
				return fetch("/graphql", {
					method: "post",
					headers: {
						'Accept': 'application/json',
						'Content-Type': 'application/json',
						'Authorization': 'Bearer f66aa7b3-add4-4fa2-88c8-534b2905c6ca'
					},
					body: JSON.stringify(graphQLParams),
					credentials: "include",
				}).then(function (response) {
					return response.text();
				}).then(function (responseBody) {
					try {
						return JSON.parse(responseBody);
					} catch (error) {
						return responseBody;
					}
				});
			}
			ReactDOM.render(
				React.createElement(GraphiQL, {fetcher: graphQLFetcher}),
				document.getElementById("graphiql")
			);
		</script>
	</body>
</html>
`)
