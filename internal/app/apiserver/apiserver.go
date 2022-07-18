package apiserver

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type APIserver struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

func New(config *Config) *APIserver {
	return &APIserver{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (s *APIserver) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}
	s.configureRouter()
	s.logger.Info("starting api server")
	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *APIserver) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)

	return nil
}

func (s *APIserver) configureRouter() {
	s.router.HandleFunc("/hello", s.Handlehello())
}

func (s *APIserver) Handlehello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello")
	}
}
