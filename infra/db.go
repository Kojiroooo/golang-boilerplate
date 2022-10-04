package infra

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"app/domain/model"
)

// 同時コネクションとかを考えた方がいい
// コネクションがどれだけ貼られるか不明なのでシングルトンにした
// そもそもポインタ参照だからシングルトンにする必要はないかもしれない
// Db.poolのアクセスを防げてないしあまり機能してない
type dbConnection struct {
	pool *gorm.DB
}

var Db dbConnection

func (db dbConnection) GetConnection() *gorm.DB {
	if db.pool == nil {
		_pool, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
			PrepareStmt: true,
		})
		db.pool = _pool
		if err != nil {
			panic("failed to connect database")
		}
	}
	return db.pool
}

func (db dbConnection) Close() {
	if db.pool != nil {
		_pool := db.pool
		_db, _ := _pool.DB()
		_db.Close()
	}
}

// マイグレートは別で管理したい
func DbInit() {
	db := Db.GetConnection()
	db.AutoMigrate(&model.Post{})
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Auth{})
	defer Db.Close()
}
