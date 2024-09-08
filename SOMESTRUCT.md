<pre><code>
package main

import (
"database/sql"
"fmt"
"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
// Чтение переменных окружения
serverAddress := os.Getenv("SERVER_ADDRESS")
postgresConnURL := os.Getenv("POSTGRES_CONN")

	// Подключение к базе данных PostgreSQL
	db, err := sql.Open("postgres", postgresConnURL)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Создание экземпляра Gin
	r := gin.Default()

	// Маршруты API
	r.GET("/api/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})

	tenders := r.Group("/api/tenders")
	{
		tenders.GET("", listTenders)
		tenders.POST("/new", createTender)
		tenders.GET("/my", getTendersByUser)
		tenders.PATCH("/:id/edit", editTender)
		tenders.PUT("/:id/rollback/:version", rollbackTender)
	}

	bids := r.Group("/api/bids")
	{
		bids.POST("/new", createBid)
		bids.GET("/my", getBidsByUser)
		bids.GET("/:tenderId/list", getBidsForTender)
		bids.PATCH("/:id/edit", editBid)
		bids.PUT("/:id/rollback/:version", rollbackBid)
	}

	// Запуск сервера
	r.Run(serverAddress)
}

// Обработчики маршрутов
func listTenders(c *gin.Context) {
// Логика получения списка тендеров
}

func createTender(c *gin.Context) {
// Логика создания нового тендера
}

func getTendersByUser(c *gin.Context) {
// Логика получения тендеров пользователя
}

func editTender(c *gin.Context) {
// Логика редактирования тендера
}

func rollbackTender(c *gin.Context) {
// Логика отката версии тендера
}

func createBid(c *gin.Context) {
// Логика создания нового предложения
}

func getBidsByUser(c *gin.Context) {
// Логика получения предложений пользователя
}

func getBidsForTender(c *gin.Context) {
// Логика получения предложений для тендера
}

func editBid(c *gin.Context) {
// Логика редактирования предложения
}

func rollbackBid(c *gin.Context) {
// Логика отката версии предложения
}

</code></pre>