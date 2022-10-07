package main

import (
	"io"
	"log"
	"net/http"

	"github.com/kuropenguin/api_go_chuukyuusha/handlers"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {

	if req.Method == http.MethodGet {
		io.WriteString(w, "Hello, world\n")
	}
}

func main() {

	http.HandleFunc("/hello", handlers.HelloHandler)
	http.HandleFunc("/article", handlers.PostArticleHandler)
	http.HandleFunc("/article/list", handlers.ArticleListHandler)
	http.HandleFunc("/article/1", handlers.ArticleDetailHandler)
	http.HandleFunc("/article/nice", handlers.PostNiceHandler)
	http.HandleFunc("/comment", handlers.PostCommentHandler)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
