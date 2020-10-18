package apiserver

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ktroitskiy/http-rest-api/internal/store"
	"github.com/sirupsen/logrus"
)

// APIServer ...
type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

// New ...
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Start ...
func (server *APIServer) Start() error {
	if err := server.configureLogger(); err != nil {
		return err
	}

	server.configureRouter()

	if err := server.configureStore(); err != nil {
		return err
	}

	server.logger.Info("starting api server")

	return http.ListenAndServe(server.config.BindAddress, server.router)
}

func (server *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(server.config.LogLevel)

	if err != nil {
		return err
	}

	server.logger.SetLevel(level)

	return nil
}

func (server *APIServer) configureRouter() {
	server.router.HandleFunc("/hello", server.handleHello())
}

func (server *APIServer) configureStore() error {
	store := store.New(server.config.Store)

	if err := store.Open(); err != nil {
		return err
	}

	server.store = store

	return nil
}

func (server *APIServer) handleHello() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		io.WriteString(writer, "hello")
	}
}
