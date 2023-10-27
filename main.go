package main

import (
"github.com/julienschmidt/httprouter"
"gopkg.in/mgo.v2"
"net/http"
"todo/controllers"
)

func main(){

	router := httprouter.New()
	todoController := controllers.NewTodoController(getSession())
	router.GET("/todo/:id", todoController.GetTodo)
	router.POST("/todo", todoController.CreateTodo)
	router.DELETE("/todo/:id", todoController.DeleteTodo)
	http.ListenAndServe("localhost:3000",router)

}

func getSession() *mgo.Session{
	s, err := mgo.Dial("mongodb://localhost:27017")
	if err!= nil{
		panic(err)
	} 
	return s
}