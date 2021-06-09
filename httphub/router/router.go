package router

import (
	"net/http"

	"github.com/ElMehdi19/httphub/httphub/helpers"
	"github.com/ElMehdi19/httphub/httphub/middlewares"
	"github.com/gorilla/mux"
)

func New() http.Handler {
	handler := mux.NewRouter()
	handler.Use(middlewares.Recover, middlewares.Logger, middlewares.JSONContent, middlewares.CORS)

	if helpers.IsDevMode() {
		handler.HandleFunc("/debug", SourceCodeHandler)
	}

	// HTTP Methods handlers
	handler.HandleFunc("/get", ViewGet).Methods("GET")
	handler.HandleFunc("/put", ViewPut).Methods("PUT")
	handler.HandleFunc("/post", ViewPost).Methods("POST")
	handler.HandleFunc("/patch", ViewPatch).Methods("patch")
	handler.HandleFunc("/delete", ViewDelete).Methods("delete")
	handler.HandleFunc("/any", ViewAny)

	// Request inspection handlers
	handler.HandleFunc("/request", ViewRequest).Methods("GET")
	handler.HandleFunc("/ip", ViewIP).Methods("GET")
	handler.HandleFunc("/user-agent", ViewUserAgent).Methods("GET")
	handler.HandleFunc("/headers", ViewHeaders).Methods("GET")

	// Status codes handlers
	handler.HandleFunc("/status/{code}", ViewStatusCodes)

	// Auth handlers
	handler.HandleFunc("/auth/basic/{user}/{passwd}", ViewBasicAuth).Methods("GET")
	handler.HandleFunc("/auth/basic-hidden/{user}/{passwd}", ViewBasicAuthHidden).Methods("GET")
	handler.HandleFunc("/auth/bearer", ViewBearerAuth).Methods("GET")

	// Response inspection handlers
	handler.HandleFunc("/response-headers", ViewResponseHeader).Methods("GET")
	handler.HandleFunc("/cache", ViewCache).Methods("GET")

	return handler
}
