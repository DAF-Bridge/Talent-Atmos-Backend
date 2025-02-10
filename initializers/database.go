package initializers

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/infrastructure"
	"github.com/go-redis/redis/v8"
	"github.com/opensearch-project/opensearch-go"

	"github.com/DAF-Bridge/Talent-Atmos-Backend/logs"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// var ESClient *elasticsearch.Client
var ESClient *opensearch.Client

var S3 *infrastructure.S3Uploader

var ctx = context.Background()

func ConnectToDB() {
	// Define the PostgreSQL connection details
	dsn := os.Getenv("DATABASE_URL")

	// Load the configuration file
	utils.InitConfig()

	// Using Viper to load the configuration file
	// dsn := fmt.Sprintf("%v://%v:%v@%v:%v/%v?sslmode=disable&TimeZone=Asia/Bangkok",
	// 	viper.GetString("db.driver"),
	// 	viper.GetString("db.user"),
	// 	viper.GetString("db.password"),
	// 	viper.GetString("db.host"),
	// 	viper.GetInt("db.port"),
	// 	viper.GetString("db.database"),
	// )

	// Initialize GORM with the PostgreSQL driver
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	// Assign the db instance to the global DB variable
	DB = db

	// Test the connection
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get database instance: ", err)
	}
	if err := sqlDB.Ping(); err != nil {
		log.Fatal("Failed to ping database: ", err)
	}

	// fmt.Println("Successfully connected to PostgreSQL!")
	logs.Info("Successfully connected to PostgreSQL!")
}

func ConnectToS3() {
	// Define the S3 bucket name
	bucketName := os.Getenv("S3_BUCKET_NAME")

	// Initialize the S3 uploader
	s3, err := infrastructure.NewS3Uploader(bucketName)
	if err != nil {
		log.Fatalf("Failed to connect to S3: %s", err)
	}

	// Assign the S3 instance to the global S3 variable
	S3 = s3
	logs.Info("Successfully connected to S3!")
}

func ConnectToElasticSearch() *opensearch.Client {
	bonsaiURL := os.Getenv("ELASTICSEARCH_URL")
	username := os.Getenv("ELASTICSEARCH_USERNAME")
	password := os.Getenv("ELASTICSEARCH_PASSWORD")

	cfg := opensearch.Config{
		Addresses: []string{bonsaiURL},
		Username:  username,
		Password:  password,
	}

	// Create the client
	var err error
	ESClient, err = opensearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the Elasticsearch client: %s", err)
	}

	// Test connection
	res, err := ESClient.Info()
	if err != nil {
		log.Fatalf("Error getting info from Elasticsearch: %s", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalf("Error closing the response body: %s", err)
		}
	}(res.Body)

	status := res.Status()
	logs.Info(fmt.Sprintf("Successfully connected to Elasticsearch!, %s", status))

	return ESClient
}

func ConnectToRedis() *redis.Client {
	// Connect to Redis
	rs := os.Getenv("REDIS_URL")

	opt, _ := redis.ParseURL(rs)
	client := redis.NewClient(opt)

	res, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to ping to Redis: %v!\n", err)
	}

	// fmt.Println("Successfully connected to Redis!")
	logs.Info(fmt.Sprintf("Successfully connected to Redis!, %s", res))

	return client
}
