package main

import (
	"net/http"

	"github.com/alvinliju/go-learning/tree/main/go-mongo-crud/controllers"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main(){
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())

	r.GET("/user/:id", uc.GetUser )
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8000", r)
}


func getSession() *mgo.Session{
	session, err := mgo.Dial("mongodb://localhost:27107")
	if err != nil {
		panic(err)
	}

	return session
}
