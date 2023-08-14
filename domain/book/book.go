package book

type Book struct {
	ID     string `json:"id"`
	Author string `json:"author"`
	Title  string `json:"title"`
	Pages  int    `json:"pages"`
}
