package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/daybaryour/golang-starter/04-mongo-golang/models"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserController struct {
	session *mgo.Session
}

func NewUserController(mongo_session *mgo.Session) *UserController { //as sent from the
	return &UserController{mongo_session} //return an address of user controller
}

func (uc UserController) GetUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	users := models.User{}
	if err := uc.session.DB("golang-learning").C("users").Find().All(users); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	usersToJson, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", usersToJson)
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) { //Declaring the GetUser Struct Method
	id := params.ByName("id") //Get the ID from the url

	if !bson.IsObjectIdHex(id) { //check if the id is an hexadecimal
		w.WriteHeader(http.StatusNotFound) //send 404 response to postman
	}

	oid := bson.ObjectIdHex(id) //all is well set the id as oid while converting it to hexadecimal (mongo db object id)

	user := models.User{} //assign the user to a variable a struct of type model
	if err := uc.session.DB("golang-learning").C("users").FindId(oid).One(&user); err != nil {
		w.WriteHeader(404)
		return
	}

	userToJson, err := json.Marshal(user) //marshall the response into json
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", userToJson)
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	user := models.User{} //assign the user to a variable a struct of type model
	//get the payload from postman and decode to json
	_ = json.NewDecoder(r.Body).Decode(&user)
	user.Id = bson.NewObjectId()

	if err := uc.session.DB("golang-learning").C("users").Insert(user); err != nil {
		w.WriteHeader(404)
		return
	}

	uj, err := json.Marshal(user) //marshall the response into json
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := params.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
	}

	oid := bson.ObjectIdHex(id)

	if err := uc.session.DB("golang-learning").C("users").RemoveId(oid); err != nil {
		w.WriteHeader(404)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Deleted User", oid)
}
