package models

type Todo struct {
	Description string `json: "description" xml: "description" form: "description"`
}
