package files

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"io"
	"net/url"
)

type s3Storage struct {
	endpointUrl     string
	bucket          string
	accessKeyId     string
	accessKeySecret string
	urlPrefix       string

	cachedAwsConfig *aws.Config
}

func NewS3Storage(endpointUrl, bucket, accessKeyId, accessKeySecret, urlPrefix string) Storage {
	return &s3Storage{
		endpointUrl:     endpointUrl,
		bucket:          bucket,
		accessKeyId:     accessKeyId,
		accessKeySecret: accessKeySecret,
		urlPrefix:       urlPrefix,
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
	cfg, err := s.awsConfig(ctx)
	if err != nil {
		return nil, err
	}
	client := s3.NewFromConfig(cfg)
	uploader := manager.NewUploader(client)

	output, err := uploader.Upload(ctx, &s3.PutObjectInput{
		Bucket:             aws.String(s.bucket),
		Key:                aws.String(path),
		ContentDisposition: aws.String("inline"),
		Body:               r,
		ACL:                types.ObjectCannedACLPublicRead,
	})
	if err != nil {
		return nil, err
	}

	uploadUrl := output.Location

	if s.urlPrefix != "" {
		urlPrefix, err := url.Parse(s.urlPrefix)
		if err != nil {
			return nil, err
		}
		s3Url, err := urlPrefix.Parse(output.Location)
		if err != nil {
			return nil, err
		}
		uploadUrl = urlPrefix.JoinPath(s3Url.Path).String()
	}

	return &UploadResult{
		Id:  *output.Key,
		Url: uploadUrl,
	}, nil
}
