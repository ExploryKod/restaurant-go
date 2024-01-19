package demoHTTP

type User struct {
	id            int    `json:"id"`
	name          string `json:"name"`
	firstname     string `json:"firstname"`
	mail          string `json:"mail"`
	phone         string `json:"phone"`
	is_superadmin bool   `json:"is_superadmin"`
	birthday      string `json:"birthday"`
}

type TodoStoreInterface interface {
	GetTodos() ([]TodoItem, error)
	AddTodo(item TodoItem) (int, error)
	DeleteTodo(id int) error
	ToggleTodo(id int) error
}
