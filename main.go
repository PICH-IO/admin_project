package main

import (
	"log"

	"github.com/PICH-IO/admin-api/configs"
	"github.com/PICH-IO/admin-api/configs/database"
	util_common "github.com/PICH-IO/admin-api/pkg/utils/common"
	"github.com/PICH-IO/admin-api/routes"

	"github.com/gofiber/contrib/fiberi18n/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	_ "github.com/lib/pq"
	"golang.org/x/text/language"
)

func main() {
	configs.InitConfig()                       // Initialize configuration
	if err := util_common.Init(); err != nil { // Initialize i18n bundle
		log.Fatalf("Failed to initialize i18n: %v", err)
	}
	db := database.GetDB()

	app := fiber.New()
	app.Use(logger.New())
	app.Use(recover.New())

	app.Use(cors.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// i18n middleware setup
	app.Use(
		fiberi18n.New(&fiberi18n.Config{
			RootPath: "pkg/translate",
			AcceptLanguages: []language.Tag{
				language.English,
				language.MustParse("km"),
			},
			DefaultLanguage: language.English,
		}),
	)
	// CORS setup
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, DELETE, PATCH",
	}))
	// fmt.Println("helllo")

	routes.SetupRoutes(app, db)
	log.Fatal(app.Listen(":" + configs.PORT))
}
