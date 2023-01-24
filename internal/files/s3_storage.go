package files

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"html/template"
	"io"
	"strings"
)

type s3Storage struct {
	endpointUrl     string
	bucket          string
	accessKeyId     string
	accessKeySecret string
	urlTemplate     string

	cachedAwsConfig *aws.Config
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

func (s *s3Storage) awsConfig(ctx context.Context) (aws.Config, error) {
	if s.cachedAwsConfig != nil {
		return *s.cachedAwsConfig, nil
	}

	resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL: s.endpointUrl,
		}, nil
	})

	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithEndpointResolverWithOptions(resolver),
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(s.accessKeyId, s.accessKeySecret, ""),
		),
	)
	if err != nil {
		return aws.Config{}, err
	}

	s.cachedAwsConfig = &cfg

	return cfg, nil
}

func (s *s3Storage) Upload(ctx context.Context, path string, r io.Reader) (*UploadResult, error) {
	urlTpl, err := template.New("url").Parse(s.urlTemplate)
	if err != nil {
		return nil, err
	}

	cfg, err := s.awsConfig(ctx)
	if err != nil {
		return nil, err
	}
	client := s3.NewFromConfig(cfg)

	_, err = client.PutObject(ctx, &s3.PutObjectInput{
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
