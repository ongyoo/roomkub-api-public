package awss3

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	// "github.com/aws/aws-sdk-go-v2/config"
	// "github.com/aws/aws-sdk-go-v2/credentials"
	// "github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	// "github.com/aws/aws-sdk-go-v2/service/s3"
	// "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

const (
	ProfileKey = "profile"
	ProductKey = "product"
)

type Storage interface {
	PutObject(ctx context.Context, key string, body io.Reader, metadata map[string]string) (*string, error) //manager.UploadOutput
	GetObject(ctx context.Context, key, fileName string) (*strings.Reader, error)
}

type storage struct {
	awsS3Client     *s3.Client
	bucketName      string
	accountID       string
	accessKeyID     string
	accessKeySecret string
}

func NewStorage() Storage {
	/*
		creds := credentials.NewStaticCredentialsProvider(os.Getenv("AWS_ACCESS_KEY"), os.Getenv("AWS_SECRET_KEY"), "")
		cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithCredentialsProvider(creds), config.WithRegion(os.Getenv("AWS_S3_REGION")))
		if err != nil {
			log.Printf("error: %v", err)
		}

		awsS3Client := s3.NewFromConfig(cfg)
		return storage{awsS3Client}
	*/
	bucketName := os.Getenv("CF_BUCKET")
	accountID := os.Getenv("CF_ID")
	accessKeyID := os.Getenv("CF_ACCESS_KEY")
	accessKeySecret := os.Getenv("CF_SECRET_KEY")

	r2Resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			//https://a3c9c45f3cbf131b32dba7f814ef6951.r2.cloudflarestorage.com/roomkub
			URL: fmt.Sprintf("https://%s.r2.cloudflarestorage.com", accountID),
		}, nil
	})

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithEndpointResolverWithOptions(r2Resolver),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKeyID, accessKeySecret, "")),
	)
	if err != nil {
		log.Fatal(err)
	}

	client := s3.NewFromConfig(cfg)
	return storage{client, bucketName, accountID, accessKeyID, accessKeySecret}
}

func (s storage) PutObject(ctx context.Context, key string, body io.Reader, metadata map[string]string) (*string, error) {
	uploader := manager.NewUploader(s.awsS3Client)
	res, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket:   aws.String(s.bucketName),
		Key:      aws.String(key),
		Body:     body,
		Metadata: metadata,
	})

	return res.Key, err
}

func (s storage) GetObject(ctx context.Context, key, fileName string) (*strings.Reader, error) {
	// Name of the file where you want to save the downloaded file
	// Create the file
	newFile, err := os.Create(fileName)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer newFile.Close()

	downloader := manager.NewDownloader(s.awsS3Client)
	numBytes, err := downloader.Download(context.TODO(), newFile, &s3.GetObjectInput{
		Bucket: aws.String(s.bucketName),
		Key:    aws.String(key),
	})

	if numBytes <= 0 {
		return nil, err
	}

	buf := bytes.NewBuffer(nil)
	_, err = io.Copy(buf, newFile)
	if err != nil {
		return nil, err
	}

	bufStr := strings.NewReader(buf.String())
	return bufStr, nil
}
