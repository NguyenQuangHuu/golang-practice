package main

import (
	"awesomeProject/internal/config"
	"awesomeProject/internal/handler"
	"awesomeProject/internal/middleware"
	"awesomeProject/internal/repository"
	"awesomeProject/internal/service"
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "nguyenthikimyen"
	password = "12345678"
	dbname   = "chitchat"
	sslmode  = "disable"
)

//TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.
// biến toàn cục để chạy kết nối đến toàn hệ thống

func main() {
	connection := &config.PSQLConnection{Host: host, Port: port, User: user, Password: password, Name: dbname, SslMode: sslmode}
	db := connection.PSQLConnection()
	//Connect to PostgresSQL
	//db = pgDBConnection()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)
	//DI Word

	wordRepository := repository.NewWordRepository(db)
	wordService := service.NewWordService(wordRepository)
	wordHandle := handler.NewWordHandleRequest(wordService)
	//DI User
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandle := handler.NewUserHandle(userService)
	//Router
	gin.SetMode(gin.DebugMode)
	//TIP Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined or highlighted text
	// to see how GoLand suggests fixing it.
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())
	wordHandler := router.Group("/api/vocabulary")
	wordHandler.GET("/search", wordHandle.FindByWord)
	wordHandler.Use(middleware.RequireAuthentication)
	wordHandler.GET("/words", middleware.RoleRequired("ADMIN", "VIP_USER"), wordHandle.GetAllWords)
	wordHandler.GET("/words/:id", wordHandle.GetWordByID)
	wordHandler.POST("/words", middleware.RoleRequired("ADMIN"), wordHandle.AddWord)
	wordHandler.PUT("/words/:id", wordHandle.UpdateWordByID)

	userHandler := router.Group("/api/user")
	userHandler.POST("/login", userHandle.Login)
	userHandler.POST("/register", userHandle.Register)
	userHandler.GET("/logout", userHandle.Logout)
	router.GET("/ws", handler.HandleWebsocket)
	err := router.Run()
	if err != nil {
		return
	}
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
