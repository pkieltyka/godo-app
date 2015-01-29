package data

import (
	"gopkg.in/mgo.v2/bson"
	"upper.io/db"
)

var TodoCollection db.Collection

type Todo struct {
	ID        bson.ObjectId `bson:"_id,omitempty" json:"id"`
	UserID    bson.ObjectId `bson:"user_id,omitempty" json:"user_id"`
	Title     string        `bson:"title" json:"title"`
	Completed bool          `bson:"completed" json:"completed"`
}

func NewTodo() *Todo {
	return &Todo{}
}

func FindTodo(cond db.Cond) (todo *Todo, err error) {
	res := TodoCollection.Find(cond)
	err = res.One(&todo)
	return
}

func FindAllTodos(conds ...db.Cond) ([]*Todo, error) {
	var cond db.Cond
	if len(conds) > 0 {
		cond = conds[0]
	}

	todos := make([]*Todo, 0)
	res := TodoCollection.Find(cond)
	err := res.All(&todos)
	if err != nil {
		return nil, err
	}

	return todos, nil
}

func FindTodoById(id bson.ObjectId) (todo *Todo, err error) {
	return FindTodo(db.Cond{"_id": id})
}
