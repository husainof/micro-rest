package user

import (
	handlers "micro/internal/handlers"
	"micro/pkg/logging"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const (
	usersURL = "/users"
	userURL  = "/user/:uuid"
)

type handler struct {
	logger *logging.Logger
}

var _ handlers.Handler = handler{}

func NewHandler(logger *logging.Logger) handlers.Handler {
	return &handler{logger}
}

func (h handler) Register(router *httprouter.Router) {
	router.GET(usersURL, h.GetUserList)
	router.GET(userURL, h.GetUserByUUID)
	router.POST(userURL, h.CreateUser)
	router.PUT(userURL, h.UpdateUser)
	router.DELETE(userURL, h.DeleteUser)
}

func (h *handler) GetUserList(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	h.logger.Info("get user list method")
	w.Write([]byte("hello world!"))
}

func (h *handler) GetUserByUUID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

}
