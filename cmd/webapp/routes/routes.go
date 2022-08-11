package routes

import (
	"encoding/json"
	"fmt"
	"gosample/internal/gosample/models"
	"io/ioutil"
	"net/http"

	"github.com/MakMoinee/go-mith/pkg/goserve"
)

type routesHandler struct {
}

func Set(httpService *goserve.Service) {
	handler := routesHandler{}

	httpService.Router.Get("/get/users", handler.GetUsers)
	httpService.Router.Post("/post/users", handler.StoreUser)
}

func (svc *routesHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	fmt.Println("id ===>>", id)
	w.Write([]byte("Hello World"))
}

func (svc *routesHandler) StoreUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	user := models.User{}

	err = json.Unmarshal(body, &user)
	if err != nil {
		panic(err)
	}

	fmt.Print(fmt.Sprintf("%#v", user))
	w.Write([]byte("You have just send a post request"))
}
