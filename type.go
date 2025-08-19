package main

type RootResponse struct {
	Success bool `json:"success"`
	Message string `json:"message"`
}

type Product struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	ImgUrl string `json:"imgUrl"`
	Status bool `json:"status"`
}

type User struct {
	Name string
	Email string
	Role string
}