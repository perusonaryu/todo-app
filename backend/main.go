package main

import (
	"backend/models"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/google/uuid"
	postgresGorm "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func initDB() {
	dsn := os.Getenv("DATABASE_URL")
	var err error

	db, err = gorm.Open(postgresGorm.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	fmt.Println("Database connected")

	runMigrations()
}

func runMigrations() {
	dsn := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("failed to create migration driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file:///app/go/src/migrations",
		"postgres", driver)
	if err != nil {
		log.Fatalf("failed to create migrate instance: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("failed to run migrations: %v", err)
	}
	fmt.Println("Migrations ran successfully")
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1"})
	// CORSの設定を追加
	r.Use(cors.Default())

	// 初期化時にDB接続とマイグレーションを実行
	initDB()

	r.GET("api/tasks", func(c *gin.Context) {
		// データの取得
		var tasks []models.Task
		result := db.Find(&tasks)
		if result.Error != nil {
			log.Fatalf("failed to retrieve tasks: %v", result.Error)
		}

		// 取得したデータを表示
		for _, task := range tasks {
			fmt.Printf("ID: %s, Name: %s, Email: %s\n", task.ID, task.Title, task.Detail)
		}

		c.JSON(200, gin.H{
			"tasks": tasks,
		})
	})

	r.GET("/create", func(ctx *gin.Context) {

		newTask := models.Task{
			ID:     uuid.New(),
			Title:  "Finish project",
			Detail: "aaaaaaaaaaaaaaa",
			Status: "wait",
		}

		result := db.Create(&newTask)
		if result.Error != nil {
			log.Fatalf("failed to create task: %v", result.Error)
		}
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
