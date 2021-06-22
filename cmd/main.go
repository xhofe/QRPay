package main

import (
	"github.com/Xhofe/QRPay/public"
	"github.com/Xhofe/QRPay/views"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/template/html"
	"net/http"
)

func main() {
	InitModel()
	//engine := html.New("./views",".html")
	engine := html.NewFileSystem(http.FS(views.Views),".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Use("/",filesystem.New(filesystem.Config{
		Root: http.FS(public.Public),
	}))
	InitView(app)
	InitApiRouter(app)
	_ = app.Listen(":3000")
}