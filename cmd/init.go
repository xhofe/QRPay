package main

import (
	"github.com/Xhofe/QRPay/controllers"
	"github.com/Xhofe/QRPay/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

func InitModel() {
	db, err := gorm.Open(sqlite.Open("QRPay.db"),&gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "qr_code",
		},
	})
	models.DB = db
	if err != nil {
		log.Fatalln("数据库连接失败")
	}
	err = models.DB.AutoMigrate(&models.QRCode{})
	if err != nil {
		log.Fatalln("数据库迁移失败")
	}
}

func InitApiRouter(app *fiber.App)  {
	api := app.Group("/api")
	{
		api.Use(cors.New())
		api.Post("/create", controllers.Create)
		api.Get("/get/:link", controllers.Get)
		api.Get("/themes",controllers.GetThemes)
	}
}

func InitView(app *fiber.App) {
	//app.Static("/","./public")
	//app.Static("/themes","./views/themes")
	app.Get("/themes/:theme",controllers.Theme)
	app.Get("/:link", controllers.View)
}