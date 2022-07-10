package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)


func IndexHandler(w http.ResponseWriter,r *http.Request, params httprouter.Params){
	name := params.ByName("name")
	w.Write([]byte(fmt.Sprintf("Hello %s",name)))
}
func main(){
	router := httprouter.New()
	router.GET("/",IndexHandler)

	listener,err := net.Listen("tcp","0.0.0.0:1234")
	if err != nil{
		 panic(err)
	}
	server := &http.Server{
		Handler: router,
		WriteTimeout:  15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}
	log.Fatalln(server.Serve(listener))
}