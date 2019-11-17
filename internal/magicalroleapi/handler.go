package magicalroleapi

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// NewHandler creates an intance of Handler.
func NewHandler() Handler {
	// In the absence of a properly-wired DI layer, just manually wire up the
	// dependencies here for this simple example.
	// TODO(mrsheepuk) Refactor this using Wire
	clientSrc := k8sClientSource{
		mode: InCluster,
	}
	client, err := clientSrc.client()
	if err != nil {
		panic("Could not create k8s client")
	}
	rg := roleGetter{
		clientset: client,
	}
	api := roleAPI{
		roleGetter: &rg,
	}
	handler := Handler{
		api: &api,
	}
	return handler
}

// Handler is the top-level entry point to setup and manage the API.
type Handler struct {
	router *mux.Router
	api    *roleAPI
}

// Run sets up and runs the HTTP server, waiting for it to exit.
func (h *Handler) Run(port int) {
	h.router = mux.NewRouter().StrictSlash(true)
	h.setupRoutes()
	h.runServer(port)
}

func (h *Handler) setupRoutes() {
	h.router.HandleFunc("/", h.api.home)
	h.router.HandleFunc("/magicalroleapi/v1", h.api.getSubjectRoles)
}

func (h *Handler) runServer(port int) {
	srv := &http.Server{
		Handler:      h.router,
		Addr:         fmt.Sprintf(":%d", port),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	log.Println("Starting server on port", port)

	// TODO(mrsheepuk) Change this to start the server then have a separate
	// function to wait for exit and observe signals to terminate.
	log.Fatal(srv.ListenAndServe())
}
