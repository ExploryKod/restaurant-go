package web

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// Function to upload a file to Amazon S3 and return the link
func uploadToS3(file multipart.File, filename string) (string, error) {
	// Configure AWS session
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("eu-west-3"),
		// You can also specify other AWS credentials here if needed
	})
	if err != nil {
		return "", err
	}

	// Initialize S3 client
	svc := s3.New(sess)

	// Prepare parameters for upload
	params := &s3.PutObjectInput{
		Bucket: aws.String("restaurant-go"),
		Key:    aws.String(filename),
		Body:   file,
		ACL:    aws.String("public-read"), // Make the file publicly accessible
	}

	// Upload the file to S3
	_, err = svc.PutObject(params)
	if err != nil {
		return "", err
	}

	// Build and return the link to the uploaded file
	fileURL := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", "restaurant-go", filename)
	return fileURL, nil
}

// HTTP handler to upload the image
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve the file from the form
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Failed to retrieve file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Generate a unique filename to avoid collisions
	filename := handler.Filename
	filename = strings.ReplaceAll(filename, " ", "_") // Replace spaces with underscores
	filename = fmt.Sprintf("%s%s", "UNIQUE_PREFIX_", filename)

	// Upload the file to Amazon S3
	fileURL, err := uploadToS3(file, filename)
	if err != nil {
		http.Error(w, "Failed to upload file to S3", http.StatusInternalServerError)
		return
	}

	// Respond with the link to the file in S3
	fmt.Fprintf(w, "File uploaded successfully. URL: %s", fileURL)
}
