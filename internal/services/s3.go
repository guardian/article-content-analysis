package services

import (
	"article-content-analysis/internal/models"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

const BucketName = "article-content-analysis"	//TODO - config?

func GetContentAnalysisFromS3(path string) (*models.ContentAnalysis, error) {
	var contentAnalysis *models.ContentAnalysis = nil

	sess, err := GetAwsSession("membership", "eu-west-1")

	downloader := s3manager.NewDownloader(sess)

	buffer := aws.NewWriteAtBuffer([]byte{})

	_, err = downloader.Download(buffer, &s3.GetObjectInput{
		Bucket: aws.String(BucketName),
		Key:    aws.String(path),
	})
	if err != nil {
		return contentAnalysis, fmt.Errorf("failed to download file, %v", err)
	}

	unmarshallError := json.Unmarshal(buffer.Bytes(), &contentAnalysis)
	if unmarshallError != nil {
		return contentAnalysis, fmt.Errorf("failed to unmarshall s3 data, %v", err)
	}

	return contentAnalysis, nil
}

func StoreContentAnalysisInS3(contentAnalysis *models.ContentAnalysis) error {
	sess, err := GetAwsSession("membership", "eu-west-1")
	uploader := s3manager.NewUploader(sess)

	marshalled, err := json.Marshal(contentAnalysis)
	if err != nil {
		return fmt.Errorf("failed to marshall data for S3 upload, %v", err)
	}

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(BucketName),
		Key:    aws.String(contentAnalysis.Path),
		Body:   bytes.NewReader(marshalled),
	})

	return err
}
