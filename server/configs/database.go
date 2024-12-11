package configs

import (
	"log"
	"os"
	"server/models"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Database connection variables
var DB *gorm.DB;
var err error;

func InitDatabase() {
	DB_HOST := os.Getenv("DB_HOST");
	DB_PORT := os.Getenv("DB_PORT");
	DB_USER := os.Getenv("DB_USER");
	DB_NAME := os.Getenv("DB_NAME");
	DB_PASSWORD := os.Getenv("DB_PASSWORD");
	DB_SSLMODE := os.Getenv("DB_SSLMODE");

	// dsn := "host=localhost user=admin password=123456 dbname=jira_clone port=5432 sslmode=disable";

	var builder strings.Builder

	builder.WriteString("host=");
	builder.WriteString(DB_HOST);
	builder.WriteString(" ");
	builder.WriteString("user=");
	builder.WriteString(DB_USER);
	builder.WriteString(" ");
	builder.WriteString("password=");
	builder.WriteString(DB_PASSWORD);
	builder.WriteString(" ");
	builder.WriteString("dbname=");
	builder.WriteString(DB_NAME);
	builder.WriteString(" ");
	builder.WriteString("port=");
	builder.WriteString(DB_PORT);
	builder.WriteString(" ");
	builder.WriteString("sslmode=");
	builder.WriteString(DB_SSLMODE);

	dsn := builder.String();

	println(Yellow, ">>>>> DSN: " + dsn, Reset);

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{});

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err);
		println(Red, ">>>>> Failed to connect to database", Reset);
	}

	// Enable UUID extension
	db, err := DB.DB();

	if err == nil {
		_, err = db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";");

		if err != nil {
			log.Fatalf("Failed to create uuid extension: %v", err);
			println(Red, ">>>>> Failed to create uuid extension", Reset);
		}
	}

	// Migrate the schema in the correct order
	err = DB.AutoMigrate(&models.User{}); // Migrate User first

	if err != nil {
		log.Fatalf("Failed to migrate 'User': %v", err);
		println(Red, ">>>>> Failed to migrate 'User'", Reset);
	}

	err = DB.AutoMigrate(&models.Workspace{}); // Then migrate Workspace

	if err != nil {
		log.Fatalf("Failed to migrate 'Workspace': %v", err);
		println(Red, ">>>>> Failed to migrate 'Workspace'", Reset);
	}

	err = DB.AutoMigrate(&models.Project{}); // Then migrate Project

	if err != nil {
		log.Fatalf("Failed to migrate 'Project': %v", err);
		println(Red, ">>>>> Failed to migrate 'Project'", Reset);
	}

	err = DB.AutoMigrate(&models.Task{}); // Finally migrate Task

	if err != nil {
		log.Fatalf("Failed to migrate 'Task': %v", err);
		println(Red, ">>>>> Failed to migrate 'Task'", Reset);
	}

	println(Green, ">>>>> Connect database successfully", Reset);
}