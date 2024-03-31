package main

import (
	"net/http"
	"errors"
	"github.com/gin-gonic/gin"
)

type todo struct {
	ID					string 	`json:"id"` // these are to help with mapping to JSON
	Item				string 	`json:"item"`
	Completed 	bool 		`json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "Clean Room", Completed: false},
	{ID: "2", Item: "Read Book", Completed: false},
	{ID: "3", Item: "Record Video", Completed: false},
}

func getTodos(context *gin.Context){ // context is the request object in Go (body, header, etc.)
	context.IndentedJSON(http.StatusOK, todos) // to convert the data (in this case todos) in the following response body in Go to JSON format
}

func addTodo(context *gin.Context){
	var newTodo todo

	// converts the request body from JSON format into the format stipulated in Go (in this case the newTodo type, todo)
	if err := context.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)

	context.IndentedJSON(http.StatusCreated, newTodo)
}

func getTodo(context *gin.Context){
	id := context.Param("id")
	todo, err := getTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, todo)

}

func toggleTodoStatus(context *gin.Context){
	id := context.Param("id")
	todo, err := getTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}

	todo.Completed = !todo.Completed

	context.IndentedJSON(http.StatusOK, todo)
}

// returns either a todo or an error
func getTodoById(id string) (*todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil // return the todo, return nil as an error
		}
	}
	return nil, errors.New("todo not found") // if todo not found, return nil for the todo and return an error
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodo)
	router.POST("/todos", addTodo)
	router.PATCH("/todos/:id", toggleTodoStatus)
	router.Run("localhost:9090")
}