package pkg

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/kurin/blazer/b2"
)

func UploadToB2(file []byte, filename string) (string, error) {
	// Load Env
	keyId := os.Getenv("BLACKBLAZE_ID")
	bucketName := os.Getenv("BLACKBLAZE_NAME")
	applicationKey := os.Getenv("BLACKBLAZE_KEY")

	ctx := context.Background()

	// Initialize B2 client
	client, err := b2.NewClient(ctx, keyId, applicationKey)
	if err != nil {
		return "", err
	}

	// Get Bucket
	bucket, err := client.Bucket(ctx, bucketName)
	if err != nil {
		return "", err
	}

	// Upload file
	unixName := strconv.FormatInt(time.Now().Unix(), 10)
	objectPath := "uploads/" + filename + unixName
	obj := bucket.Object(objectPath)
	w := obj.NewWriter(ctx)
	defer w.Close()
	if _, err := w.Write(file); err != nil {
		return "", err
	}
	if err := w.Close(); err != nil {
		return "", err
	}

	// Return Public URL
	return "https://f002.backblazeb2.com/file/" + bucketName + "/" + objectPath, nil
}
