package todo

type Todo struct {
	Title string `json:"title"`
	Description string `json:"description"`
}

type TodoHandler struct {
	db *goarm.DB
}
