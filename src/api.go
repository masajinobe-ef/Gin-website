package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Создаем новый роутер
	r := gin.Default()

	// Указываем директорию для статических файлов
	r.Static("/static", "./static")

	// Загружаем HTML-шаблон
	r.LoadHTMLFiles("static/index.html")

	// Определяем маршрут для HTML-страницы
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// Определяем маршрут для API
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Привет, мир!",
		})
	})

	// Запускаем сервер на порту 8080
	r.Run(":8080")
}
