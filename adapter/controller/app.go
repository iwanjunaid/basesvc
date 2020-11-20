package controller

type AppController struct {
	Author interface{ AuthorController }
}
