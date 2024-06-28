package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
	"os"
)

type ModelGenerator struct {
}

type DsnConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func (m *ModelGenerator) Generate(outputDir string) {
	config := ResolveDBConfig()
	g := gen.NewGenerator(gen.Config{
		OutPath: outputDir,
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})
	gormdb, _ := gorm.Open(postgres.Open(config.ToString()))
	g.UseDB(gormdb) // reuse your gorm db

	//// Generate basic type-safe DAO API for struct `model.User` following conventions
	//
	//g.ApplyBasic(
	//	// Generate struct `User` based on table `users`
	//	g.GenerateModel("users"),
	//
	//	// Generate struct `Employee` based on table `users`
	//	g.GenerateModelAs("users", "Employee"),
	//
	//	// Generate struct `User` based on table `users` and generating options
	//	g.GenerateModel("users", gen.FieldIgnore("address"), gen.FieldType("id", "int64")),
	//
	//	// Generate struct `Customer` based on table `customer` and generating options
	//	// customer table may have a tags column, it can be JSON type, gorm/gen tool can generate for your JSON data type
	//	g.GenerateModel("customer", gen.FieldType("tags", "datatypes.JSON")),
	//)
	g.ApplyBasic(
		// Generate structs from all tables of current database
		g.GenerateAllTable()...,
	)
	// Generate the code
	g.Execute()
}

func ResolveDBConfig() *DsnConfig {
	return &DsnConfig{
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	}
}

func (c DsnConfig) ToString() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Moscow",
		c.Host, c.User, c.Password, c.Database, c.Port,
	)
}
