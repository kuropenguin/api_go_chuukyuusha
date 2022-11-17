package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kuropenguin/api_go_chuukyuusha/models"
)

func main() {

	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(db-for-go)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		fmt.Println("err")
		fmt.Println(err)
	}
	defer db.Close()

	articleID := 10
	const sqlStr = `
	select *
	from articles
	where article_id = ?
	;
	`
	row := db.QueryRow(sqlStr, articleID)
	// if err := row.Err(); err != nil {
	// fmt.Println(err)
	// return
	// }

	var createdTime sql.NullTime
	var article models.Article
	err = row.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &createdTime)
	if err != nil {
		fmt.Println("fmt")
		fmt.Println(err)
		return
	}

	if createdTime.Valid {
		article.CreatedAt = createdTime.Time
	}

	fmt.Printf("%+v\n", article)

}
