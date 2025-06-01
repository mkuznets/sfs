package filestore

import (
	"context"
	"html/template"
	"io"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3FileStore struct {
	endpointUrl     string
	bucket          string
	accessKeyId     string
	accessKeySecret string
	urlTemplate     string
	usePathStyle    bool
}

type Object struct {
	Bucket string
	Key    string
}

func NewS3FileStore(endpointUrl, bucket, accessKeyId, accessKeySecret, urlTemplate string, usePathStyle bool) *S3FileStore {
	return &S3FileStore{
		endpointUrl:     endpointUrl,
		bucket:          bucket,
		accessKeyId:     accessKeyId,
		accessKeySecret: accessKeySecret,
		urlTemplate:     urlTemplate,
		usePathStyle:    usePathStyle,
	}
}

func (s *S3FileStore) Upload(ctx context.Context, path string, body io.Reader) (*UploadResult, error) {
	s3client := s3.New(s3.Options{
		Credentials:      credentials.NewStaticCredentialsProvider(s.accessKeyId, s.accessKeySecret, ""),
		EndpointResolver: s3.EndpointResolverFromURL(s.endpointUrl),
		UsePathStyle:     s.usePathStyle,
	})

	urlTpl, err := template.New("url").Parse(s.urlTemplate)
	if err != nil {
		return nil, err
	}

	_, err = s3client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:             aws.String(s.bucket),
		Key:                aws.String(path),
		ContentDisposition: aws.String("inline"),
		Body:               body,
	})
	if err != nil {
		return nil, err
	}

	var urlB strings.Builder
	if err := urlTpl.Execute(&urlB, Object{Bucket: s.bucket, Key: path}); err != nil {
		return nil, err
	}

	return &UploadResult{
		ID:  path,
		URL: urlB.String(),
	}, nil
}
