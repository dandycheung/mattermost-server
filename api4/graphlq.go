// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package api4

import (
	"net/http"

	"github.com/mattermost/mattermost-server/v5/graphql"

	"github.com/99designs/gqlgen/handler"
)

func (api *API) InitGraphql() {
	api.BaseRoutes.Graphql.Handle("", api.ApiHandlerTrustRequester(graphqlHandler))
	api.BaseRoutes.Graphql.Handle("/playground", handler.Playground("Mattermost GraphQL endpoint", "/api/v4/graphql"))
}

func graphqlHandler(c *Context, w http.ResponseWriter, r *http.Request) {
	handler.GraphQL(graphql.NewExecutableSchema(graphql.Config{Resolvers: &graphql.Resolver{c.App}})).ServeHTTP(w, r)
}
