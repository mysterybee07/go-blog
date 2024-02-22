package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mysterybee07/blogbackend/controller"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controller.Register)
	app.Post("/api/login", controller.Login)

	// app.Use("middleware.IsAuthenticate")
	app.Post("/api/create-blog", controller.CreatePost)
	app.Get("/api/allposts", controller.AllPost)
	app.Get("/api/post/:id", controller.DetailPost)
	app.Put("/api/updatepost/:id", controller.UpdatePost)
	app.Delete("/api/deletepost/:id", controller.DeletePost)
	app.Get("/api/uniquepost", controller.UniquePost)

}
