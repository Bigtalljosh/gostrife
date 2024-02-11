package web

import (
	"net/http"

	"github.com/gorilla/mux"
)

type LandingController struct {
}

func NewLandingController() *LandingController {
	return &LandingController{}
}

func (c *LandingController) RegisterRoutes(router *mux.Router) {
    router.HandleFunc("/", c.GetCustomer).Methods("GET")
}

func (c *LandingController) GetCustomer(w http.ResponseWriter, r *http.Request) {
	message := "API Is up"

    w.Write([]byte(message))
}
