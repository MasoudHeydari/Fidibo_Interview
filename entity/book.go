package entity

type Book struct {
	Id         string    `json:"id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	Publishers Publisher `json:"publishers"`
	Authors    Author    `json:"authors"`
	Slug       string    `json:"slug"`
	ImageName  string    `json:"image_name"`
}

type Publisher struct {
	Title string `json:"title"`
}

type Author struct {
	Name string `json:"name"`
}
