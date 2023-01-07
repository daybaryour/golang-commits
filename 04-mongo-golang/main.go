package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"

	"github.com/daybaryour/golang-starter/04-mongo-golang/controllers"
)

func main() {
	router := httprouter.New() //declaring r as the router variable so we don't have to prepend each codebase with r
	user_controller := controllers.NewUserController(getSession())

	router.GET("/users", user_controller.GetUsers)
	router.GET("/users/:id", user_controller.GetUser)
	router.POST("/users", user_controller.CreateUser)
	router.DELETE("/users/:id", user_controller.DeleteUser)

	http.ListenAndServe("localhost:8000", router)
}

func getSession() *mgo.Session { //return pointer to mongo db session
	mongodb_session, err := mgo.Dial("mongodb+srv://debaryour:qeQ0Va4ZZj49HTAw@cluster0.1nbznun.mongodb.net/?retryWrites=true&w=majority") //dial the mongo db connecton string
	if err != nil {
		panic(err) //stops execution and output error message
	}

	return mongodb_session
}
