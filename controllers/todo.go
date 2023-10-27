package controllers

import(
	"fmt"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"todo/models"
)

type TodoController struct{
	session *mgo.Session
}

func NewTodoController(s *mgo.Session) *TodoController{
return &TodoController{s}
}

func(todoController TodoController) GetTodo(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id){
		w.WriteHeader(http.StatusNotFound)
	}

	oid := bson.ObjectIdHex(id)

	t := models.Todo{}

	if err := todoController.session.DB("golang").C("todos").FindId(oid).One(&t); err != nil{
		w.WriteHeader(404)
		return
	}
	tjson , err := json.Marshal(t)

	if err!=nil{
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", tjson)
}

func(todoController TodoController) CreateTodo(w http.ResponseWriter, r *http.Request, _ httprouter.Params){

	t := models.Todo{}

	json.NewDecoder(r.Body).Decode(&t)

	t.Id = bson.NewObjectId()

	todoController.session.DB("golang").C("todos").Insert(t)
	
	tjson, err := json.Marshal(t)

	if err!=nil{
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", tjson)
}

func(todoController TodoController) DeleteTodo(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id){
		w.WriteHeader(http.StatusNotFound)
	}

	oid := bson.ObjectIdHex(id)


	if err := todoController.session.DB("golang").C("todos").RemoveId(oid); err != nil{
		w.WriteHeader(404)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Deleted todo", oid, "\n")
}