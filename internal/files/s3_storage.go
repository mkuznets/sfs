package files

import (
	"context"
	"html/template"
	"io"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type s3Storage struct {
	endpointUrl     string
	bucket          string
	accessKeyId     string
	accessKeySecret string
	urlTemplate     string

	// cachedAwsConfig *aws.Config
}

type Object struct {
	Bucket string
	Key    string
}

func NewS3Storage(endpointUrl, bucket, accessKeyId, accessKeySecret, urlTemplate string) Storage {
	return &s3Storage{
		endpointUrl:     endpointUrl,
		bucket:          bucket,
		accessKeyId:     accessKeyId,
		accessKeySecret: accessKeySecret,
		urlTemplate:     urlTemplate,
	}
}

func (s *s3Storage) Upload(ctx context.Context, path string, r io.Reader) (*UploadResult, error) {
	s3client := s3.New(s3.Options{
		Credentials:      credentials.NewStaticCredentialsProvider(s.accessKeyId, s.accessKeySecret, ""),
		EndpointResolver: s3.EndpointResolverFromURL(s.endpointUrl),
	})

	urlTpl, err := template.New("url").Parse(s.urlTemplate)
	if err != nil {
		return nil, err
	}

	_, err = s3client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:             aws.String(s.bucket),
		Key:                aws.String(path),
		ContentDisposition: aws.String("inline"),
		Body:               r,
	})
	if err != nil {
		return nil, err
	}

	var urlB strings.Builder
	if err := urlTpl.Execute(&urlB, Object{Bucket: s.bucket, Key: path}); err != nil {
		return nil, err
	}

	return &UploadResult{
		Id:  path,
		Url: urlB.String(),
	}, nil
}
