package services

import (
	"article-content-analysis/internal/models"
	"bytes"
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/pkg/errors"
)

const BucketName = "gu-article-content-analysis" //TODO - config?

// Returns error if object is not in s3
func GetContentAnalysisFromS3(path string) (*models.ContentAnalysis, error) {
	var contentAnalysis *models.ContentAnalysis = nil

	sess, err := GetAwsSession("developerPlayground", "eu-west-1")
	if err != nil {
		return contentAnalysis, errors.Wrap(err, "failed to create aws session")
	}

	downloader := s3manager.NewDownloader(sess)

	buffer := aws.NewWriteAtBuffer([]byte{})

	_, err = downloader.Download(buffer, &s3.GetObjectInput{
		Bucket: aws.String(BucketName),
		Key:    aws.String(path),
	})
	if err != nil {
		return contentAnalysis, errors.Wrap(err, "failed to download file")
	}

	unmarshallError := json.Unmarshal(buffer.Bytes(), &contentAnalysis)
	if unmarshallError != nil {
		return contentAnalysis, errors.Wrap(err, "failed to unmarshall s3 data")
	}

	return contentAnalysis, nil
}

func StoreContentAnalysisInS3(contentAnalysis *models.ContentAnalysis) error {
	sess, err := GetAwsSession("developerPlayground", "eu-west-1")
	uploader := s3manager.NewUploader(sess)

	marshalled, err := json.Marshal(contentAnalysis)
	if err != nil {
		return errors.Wrap(err, "failed to marshall data for S3 upload")
	}

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(BucketName),
		Key:    aws.String(contentAnalysis.Path),
		Body:   bytes.NewReader(marshalled),
	})

	return err
}
