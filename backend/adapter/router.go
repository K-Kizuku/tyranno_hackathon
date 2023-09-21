package adapter

import (
	"net/http"
	"time"
	mymiddleware "tyranno/backend/adapter/middleware"
	"tyranno/backend/adapter/websocket"
	"tyranno/backend/domain/service"
	"tyranno/backend/utils/config"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	Router *chi.Mux
}

func New() *Server {
	return &Server{
		Router: chi.NewRouter(),
	}
}

func (s *Server) Init() {
	// r := chi.NewRouter()
	config.LoadEnv()
	// http.ListenAndServe(":8080", r)
}

func (s *Server) InitRouter() {
	hub := service.NewHubModel()
	go hub.Run()

	s.Router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello tyranno"))
	})
	s.Router.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {

		websocket.InitWS(hub, w, r)
	})
	s.Router.Handle("/protected", mymiddleware.JWTMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("This is a protected route."))
	})))

}

func (s *Server) Middleware() {
	s.Router.Use(middleware.RequestID)
	s.Router.Use(middleware.Logger)
	s.Router.Use(middleware.Recoverer)

	// s.Router.Use(middleware.CloseNotify)
	s.Router.Use(middleware.Timeout(time.Second * 60))
}
