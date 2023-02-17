package entity

type Book struct {
	Id        string    `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Pubs      Publisher `json:"publishers"`
	Authors   []Author  `json:"authors"`
	Slug      string    `json:"slug"`
	ImageName string    `json:"image_name"`
}

type Publisher struct {
	Title string `json:"title"`
}

type Author struct {
	Name string `json:"name"`
}

/*
	type structs for parsing result of Fidibo API
*/

type FidiboBook struct {
	AllBook Books `json:"books"`
}

type Books struct {
	AllHit Hits `json:"hits"`
}
type Hits struct {
	AllHit []Source `json:"hits"`
}

type Source struct {
	Boo Book `json:"_source"`
}
