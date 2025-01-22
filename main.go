package main

import (
	"fiber-poc-api/database/entity"
	internal "fiber-poc-api/routes"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	jwtware "github.com/gofiber/jwt/v3"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"strings"
	"time"
)

func main() {

	// ==> Get config from config.yaml
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	err := viper.ReadInConfig()
	if err != nil {
		log.Errorf("fatal error config file: %+v \n", err)
	}

	// ==> Connect database mysql
	db := databaseConnection()
	generateTable := true
	log.Infof("config generateTable: %t", generateTable)
	if generateTable {
		db.AutoMigrate(&entity.User{})
		db.AutoMigrate(&entity.Role{})
		db.AutoMigrate(&entity.Privilege{})
		db.AutoMigrate(&entity.UserRole{})
		db.AutoMigrate(&entity.RolePrivilege{})
		db.AutoMigrate(&entity.LoginHistory{})
		log.Infof("Generate Tables success")
	}

	app := fiber.New()

	// ==> cors
	app.Use(cors.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// ==> JWT Middleware configuration
	jwtMiddleware := jwtware.New(jwtware.Config{
		SigningKey:   []byte(viper.GetString("jwt.secret")),
		ErrorHandler: jwtError,
	})

	// ==> routes
	internal.Router(app, jwtMiddleware, db)

	// ==> server start
	port := viper.GetString("server.port")
	err = app.Listen(fmt.Sprintf(":%s", port))
	if err != nil {
		log.Errorf("error server: %+v \n", err.Error())
		return
	}
}

func jwtError(c *fiber.Ctx, err error) error {
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
	}
	return c.Next()
}

func databaseConnection() *gorm.DB {

	newLogger := logger.New(nil,
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,          // Don't include params in the SQL log
			Colorful:                  false,         // Disable color
		},
	)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Bangkok",
		viper.GetString("database.host"),
		viper.GetString("database.user"),
		viper.GetString("database.password"),
		viper.GetString("database.name"),
		viper.GetString("database.port"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "",
			NoLowerCase: true,
		},
	})
	if err != nil {
		log.Errorf("error database: %+v \n", err.Error())
		return nil
	}

	fmt.Println("Successfully connected to the database")
	return db
}
