package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath:       "../dal/query",
		ModelPkgPath:  "../dal/model",
		FieldNullable: true,
	})

	db, _ := gorm.Open(mysql.Open("root:123456@(127.0.0.1:3306)/dev?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})

	g.UseDB(db)

	g.ApplyBasic(g.GenerateModel("user_info"),
		g.GenerateModel("credits_record"),
		g.GenerateModel("props_info"),
		g.GenerateModel("user_props"),
		g.GenerateModel("props_record"),
		g.GenerateModel("virtual_currency_record"))

	g.Execute()
}
