package models

type Todo struct {
	Id          uint   `json: "id" xml: "id" form: "id"`
	Description string `json: "description" xml: "description" form: "description"`
}
