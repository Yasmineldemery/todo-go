package main

import (
"fmt"
"github.com/julienschmidt/httprouter"
"gopkg.in/mgo.v2"
"net/http"
)

func main(){

	router := httprouter.New()
	todoController := controllers.NewTodoController(getSession())
	router.GET("/todo", todoController.getTodos)
	router.GET("/todo/:id", todoController.getTodo)
	router.POST("/todo", todoController.createTodo)
	router.DELETE("/todo/:id", todoController.deleteTodo)
	router.PUT("/todo/:id",todoController.updateTodo)
	http.ListenAndServe("localhost:3000",router)

}

func getSession() *mgo.Session{
	s,err := mgo.Dial("mongodb://localhost:27017")
	if err!= nil{
		fmt.Println(err)
		panic(err)
	} 
	return s
}