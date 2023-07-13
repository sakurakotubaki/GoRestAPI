package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)
// Bookという構造体を定義
type Book struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}
// Book型のスライスを定義
var books []Book

func main() {
	// Book型のスライスに本の情報を追加
	books = append(books, Book{ID: 1, Title: "Go言語"})
	books = append(books, Book{ID: 2, Title: "Rails"})
	books = append(books, Book{ID: 3, Title: "Flutter"})
  // ginのルーターを作成
	r := gin.Default()
  // GETリクエストを受け取った時の処理
	r.GET("/books", func(c *gin.Context) {
		// 本の情報をJSONで返す
		c.JSON(http.StatusOK, books)
	})
  // POSTリクエストを受け取った時の処理
	r.POST("/books", func(c *gin.Context) {
		// Book型の変数を定義
		var book Book
		// リクエストボディをJSONにバインド
		if err := c.BindJSON(&book); err != nil {
			// エラーがあればエラーを返す
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// 本の情報を追加
		book.ID = books[len(books)-1].ID + 1
		// appendでスライスに追加
		books = append(books, book)
		// c.JSONでJSONを返す
		c.JSON(http.StatusOK, book)
	})
  // r.Run()でサーバーを起動
	r.Run() // listen and serve on 0.0.0.0:8080
}
