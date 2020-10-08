package rest

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/gorilla/mux"
)

// RegisterRoutes registers phraseservice-related REST handlers to a router
func RegisterRoutes(ctx context.CLIContext, r *mux.Router) {
	root := r.PathPrefix("/phraseservice/phrase").Subrouter()

	root.HandleFunc("", createPhraseHandler(ctx)).Methods("POST")
	root.HandleFunc("/{owner}", listPhraseHandler(ctx, "phraseservice")).Methods("GET")
	root.HandleFunc("", setPhraseHandler(ctx)).Methods("PUT")
	root.HandleFunc("", deletePhraseHandler(ctx)).Methods("DELETE")
}
