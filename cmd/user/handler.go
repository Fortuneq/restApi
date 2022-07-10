package user

import ("github.com/julienschmidt/httprouter"
"net/http")
type handler struct {

}

func (h * handler) Register(router *httprouter.Router){
	router.GET("/users", h.Getlist)
}

func (h * handler)Getlist(w http.ResponseWriter, r * http.Request,params){
	w.Write([]byte("this is list of users"))
}