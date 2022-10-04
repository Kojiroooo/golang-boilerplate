package main

import (
	"fmt"
	"net/http"

	"app/domain/model"
	"app/infra"
)

// https://dev.classmethod.jp/articles/go-sample-rest-api/
// Response構造体を作ってない
// dbがsqlite
// package分割できてない
// testコード書いてない(go-sqlmockとか使うといいのかな？)

func main() {
	infra.DbInit()
	HandleRequest()
	testQuery()
	startServer()

}

func startServer() {
	fmt.Println("[INFO] Server listening")
	http.ListenAndServe(":3000", nil)
}

func testQuery() {
	db := infra.Db.GetConnection()
	var user model.User
	db.First(&user, 1)
	fmt.Println(user)
}
