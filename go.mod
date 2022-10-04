module app

go 1.17

replace gopkg.in/urfave/cli.v2 => github.com/urfave/cli/v2 v2.1.1

require gorm.io/driver/sqlite v1.3.1

require (
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.4 // indirect
	github.com/mattn/go-sqlite3 v1.14.9 // indirect
	gorm.io/gorm v1.23.1 // indirect
)
