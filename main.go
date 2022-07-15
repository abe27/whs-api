package main

import (
	"os"
	"strings"
	"time"

	"github.com/abe27/bugtracker/api/controllers"
	"github.com/abe27/bugtracker/api/models"
	"github.com/abe27/bugtracker/api/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", controllers.Hello)
	r := app.Group("/api/v1")
	r.Get("/hello", controllers.Hello)
	u := app.Group("/api/v1/member")
	u.Post("/register", controllers.Register)
	u.Post("/login", controllers.Login)
}

func init() {
	var err error
	err = godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}

	// dns := "host=" + os.Getenv("DBHOST") + " user=" + os.Getenv("DBUSER") + " password=" + os.Getenv("DBPASSWORD") + " dbname=" + os.Getenv("DBNAME") + " port=" + os.Getenv("DBPORT") + " sslmode=" + os.Getenv("SSLMODE") + " TimeZone=" + os.Getenv("TZNAME") + ""
	dns := "host=" + os.Getenv("DBHOST") +
		" user=" + os.Getenv("DBUSER") +
		" password=" + os.Getenv("DBPASSWORD") +
		" dbname=" + os.Getenv("DBNAME") +
		" port=" + os.Getenv("DBPORT") +
		" sslmode=" + os.Getenv("SSLMODE") +
		" TimeZone=" + os.Getenv("TZNAME") + ""
	services.DBConn, err = gorm.Open(postgres.Open(dns), &gorm.Config{
		SkipDefaultTransaction: true,
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "tbt_", // table name prefix, table for `User` would be `t_users`
			SingularTable: false,  // use singular table name, table for `User` would be `user` with this option enabled
			NoLowerCase:   false,  // skip the snake_casing of names
			NameReplacer:  strings.NewReplacer("CID", "Cid"),
		},
	})

	if err != nil {
		panic("Failed to connect to database")
	}
	services.DBConn.AutoMigrate(&models.User{})
}

func main() {
	config := fiber.Config{
		Prefork:                      true,
		ServerHeader:                 "Taweechai Yuenyang API Server",
		AppName:                      "API Service",
		StreamRequestBody:            false,
		DisablePreParseMultipartForm: false,
		ReduceMemoryUsage:            false,
	}
	// Initialize Fiber Framework
	app := fiber.New(config)
	app.Use(requestid.New())
	app.Use(logger.New())
	SetupRoutes(app)
	app.Listen(":3000")
}
