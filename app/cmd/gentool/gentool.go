package main

import (
	"denny/gin_gorm_practise_dining/app/db"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	var d db.IDb

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if os.Getenv("APP_ENV") == "testing" {
		d = &db.Sqlite{}
	} else {
		d = &db.Mysql{}
	}

	conn := db.InitDb(d.NewConnInfo())

	g := gen.NewGenerator(gen.Config{
		OutPath:       "./app/query",
		Mode:          gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
		FieldNullable: true,
	})

	g.UseDB(conn)

	dataMap := map[string]func(detailType gorm.ColumnType) (dataType string){
		"tinyint":   func(detailType gorm.ColumnType) (dataType string) { return "int64" },
		"smallint":  func(detailType gorm.ColumnType) (dataType string) { return "int64" },
		"mediumint": func(detailType gorm.ColumnType) (dataType string) { return "int64" },
		"bigint":    func(detailType gorm.ColumnType) (dataType string) { return "int64" },
		"int":       func(detailType gorm.ColumnType) (dataType string) { return "int64" },
		"time":      func(columnType gorm.ColumnType) (dataType string) { return "datatypes.Time" },
	}

	// 要先于`ApplyBasic`执行
	g.WithDataTypeMap(dataMap)

	// g.GenerateAllTable()
	g.ApplyBasic(g.GenerateAllTable()...)

	g.Execute()

}
