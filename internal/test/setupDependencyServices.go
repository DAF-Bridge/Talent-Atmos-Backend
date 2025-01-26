//go:builds integration

package test

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/models"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/logs"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/opensearch-project/opensearch-go"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB_TEST *gorm.DB
var ESClient_TEST *opensearch.Client
var s3Client *minio.Client

// setupTestDB initializes the test database before running tests
func setupTestDB() {
	var err error
	dsn := os.Getenv("DATABASE_URL_TEST")
	DB_TEST, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	sqlDB, err := DB_TEST.DB()
	if err != nil {
		log.Fatal("Failed to get database instance: ", err)
	}
	if err := sqlDB.Ping(); err != nil {
		log.Fatal("Failed to ping database: ", err)
	}
	logs.Info("Connected to test database successfully.")

	err = DB_TEST.AutoMigrate(&models.Event{})
	if err != nil {
		log.Fatal("Failed to create test table:", err)
	}

	if err := DB_TEST.Exec(`INSERT INTO events (id, organization_id, category_id, name, pic_url, start_date, end_date, start_time, end_time, description, highlight, requirement, key_takeaway, location_name, latitude, longitude, province, location_type, audience, price_type) VALUES (1, 1, 1, 'Builds Renewable Energy Summit', 'https://drive.google.com/uc?export=view&id=1-wqxOT_uo1pE_mEPHbJVoirMMH2Be3Ks', '2025-01-15', '2025-01-16', '09:00:00', '17:00:00', 'Explore advancements in renewable energy technologies.', 'Top speakers from the renewable energy sector.', 'Open to professionals in the energy sector.', 'Learn about the latest trends in solar and wind energy.', 'Conference Hall A', 13.7563, 100.5018, 'Bangkok', 'onsite', 'students', 'free');`).Error; err != nil {
		log.Fatal("Failed to insert test data:", err)
	}

	fmt.Println("Test database initialized.")
}

// teardownTestDB cleans up the database after tests finish
func teardownTestDB() {
	err := DB_TEST.Exec("DELETE FROM events;")
	if err != nil {
		log.Fatal("Failed to clean up test data:", err)
	}

	fmt.Println("Test database cleaned up.")
}

// setupOpenSearch initializes OpenSearch client
func setupOpenSearch() {
	var err error
	ESClient_TEST, err = opensearch.NewClient(opensearch.Config{
		Addresses: []string{
			os.Getenv("OPENSEARCH_URL_TEST"),
		},
	})
	if err != nil {
		log.Fatal("Failed to create OpenSearch client:", err)
	}

	fmt.Println("OpenSearch client initialized.")
}

// teardownOpenSearch cleans up OpenSearch resources after tests
func teardownOpenSearch() {
	_, err := ESClient_TEST.Indices.Delete([]string{"events"})
	if err != nil {
		log.Fatal("Failed to delete OpenSearch index:", err)
	}

	fmt.Println("OpenSearch resources cleaned up.")
}

// setupS3 initializes MinIO client to simulate S3
func setupS3() {
	endpoint := os.Getenv("S3_ENDPOINT")
	accessKeyID := os.Getenv("S3_ACCESS_KEY")
	secretAccessKey := os.Getenv("S3_SECRET_KEY")
	useSSL := false

	var err error
	s3Client, err = minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatal("Failed to create MinIO client:", err)
	}

	// Create a bucket
	bucketName := os.Getenv("S3_BUCKET_NAME")
	location := os.Getenv("S3_BUCKET_LOCATION")

	err = s3Client.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		exists, errBucketExists := s3Client.BucketExists(context.Background(), bucketName)
		if errBucketExists == nil && exists {
			log.Printf("Bucket '%s' already exists.", bucketName)
		} else {
			log.Fatal("Failed to create S3 bucket:", err)
		}

		log.Fatal("Failed to create S3 bucket:", err)
	} else {
		logs.Info("S3 (MinIO) client initialized.")
	}

	// Upload a test object
	filePath := "internal\\test\\metadata\\Google.png"
	objectName := "events/banners/1"
	contentType := "image/jpeg"

	_, err = s3Client.FPutObject(context.Background(), bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Fatal("Failed to upload test object to S3:", err)
	}

	logs.Info("Successfully uploaded test object to S3.")
}

// teardownS3 cleans up S3 resources after tests
func teardownS3() {
	err := s3Client.RemoveObject(context.Background(), os.Getenv("S3_BUCKET_NAME"), "events/banners/1", minio.RemoveObjectOptions{})
	if err != nil {
		log.Fatal("Failed to remove S3 object:", err)
	}

	fmt.Println("S3 resources cleaned up.")
}

// TestMain is the entry point for all tests in this package
func TestMain(m *testing.M) {
	setupTestDB()
	setupOpenSearch()
	setupS3()

	defer func() {
		teardownS3()
		teardownOpenSearch()
		teardownTestDB()
	}() // Ensure cleanup after tests finish

	exitCode := m.Run()
	os.Exit(exitCode)
}
