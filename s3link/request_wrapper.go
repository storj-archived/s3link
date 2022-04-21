package s3link

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/s3"
)

func (s *S3Link) AbortMultipartUploadWithContext(ctx aws.Context, input *s3.AbortMultipartUploadInput, opts ...request.Option) (*s3.AbortMultipartUploadOutput, error) {
	req, out := s.AbortMultipartUploadRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) CompleteMultipartUploadWithContext(ctx aws.Context, input *s3.CompleteMultipartUploadInput, opts ...request.Option) (*s3.CompleteMultipartUploadOutput, error) {
	req, out := s.CompleteMultipartUploadRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) CopyObjectWithContext(ctx aws.Context, input *s3.CopyObjectInput, opts ...request.Option) (*s3.CopyObjectOutput, error) {
	req, out := s.CopyObjectRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) CreateBucketWithContext(ctx aws.Context, input *s3.CreateBucketInput, opts ...request.Option) (*s3.CreateBucketOutput, error) {
	req, out := s.CreateBucketRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) CreateMultipartUploadWithContext(ctx aws.Context, input *s3.CreateMultipartUploadInput, opts ...request.Option) (*s3.CreateMultipartUploadOutput, error) {
	req, out := s.CreateMultipartUploadRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) DeleteBucketWithContext(ctx aws.Context, input *s3.DeleteBucketInput, opts ...request.Option) (*s3.DeleteBucketOutput, error) {
	req, out := s.DeleteBucketRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) DeleteBucketAnalyticsConfigurationWithContext(ctx aws.Context, input *s3.DeleteBucketAnalyticsConfigurationInput, opts ...request.Option) (*s3.DeleteBucketAnalyticsConfigurationOutput, error) {
	req, out := s.DeleteBucketAnalyticsConfigurationRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) DeleteBucketCorsWithContext(ctx aws.Context, input *s3.DeleteBucketCorsInput, opts ...request.Option) (*s3.DeleteBucketCorsOutput, error) {
	req, out := s.DeleteBucketCorsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) DeleteBucketEncryptionWithContext(ctx aws.Context, input *s3.DeleteBucketEncryptionInput, opts ...request.Option) (*s3.DeleteBucketEncryptionOutput, error) {
	req, out := s.DeleteBucketEncryptionRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) DeleteBucketIntelligentTieringConfigurationWithContext(ctx aws.Context, input *s3.DeleteBucketIntelligentTieringConfigurationInput, opts ...request.Option) (*s3.DeleteBucketIntelligentTieringConfigurationOutput, error) {
	req, out := s.DeleteBucketIntelligentTieringConfigurationRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) DeleteBucketInventoryConfigurationWithContext(ctx aws.Context, input *s3.DeleteBucketInventoryConfigurationInput, opts ...request.Option) (*s3.DeleteBucketInventoryConfigurationOutput, error) {
	req, out := s.DeleteBucketInventoryConfigurationRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) DeleteBucketLifecycleWithContext(ctx aws.Context, input *s3.DeleteBucketLifecycleInput, opts ...request.Option) (*s3.DeleteBucketLifecycleOutput, error) {
	req, out := s.DeleteBucketLifecycleRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) DeleteBucketMetricsConfigurationWithContext(ctx aws.Context, input *s3.DeleteBucketMetricsConfigurationInput, opts ...request.Option) (*s3.DeleteBucketMetricsConfigurationOutput, error) {
	req, out := s.DeleteBucketMetricsConfigurationRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) DeleteBucketOwnershipControlsWithContext(ctx aws.Context, input *s3.DeleteBucketOwnershipControlsInput, opts ...request.Option) (*s3.DeleteBucketOwnershipControlsOutput, error) {
	req, out := s.DeleteBucketOwnershipControlsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) DeleteBucketPolicyWithContext(ctx aws.Context, input *s3.DeleteBucketPolicyInput, opts ...request.Option) (*s3.DeleteBucketPolicyOutput, error) {
	req, out := s.DeleteBucketPolicyRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) DeleteBucketReplicationWithContext(ctx aws.Context, input *s3.DeleteBucketReplicationInput, opts ...request.Option) (*s3.DeleteBucketReplicationOutput, error) {
	req, out := s.DeleteBucketReplicationRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) DeleteBucketTaggingWithContext(ctx aws.Context, input *s3.DeleteBucketTaggingInput, opts ...request.Option) (*s3.DeleteBucketTaggingOutput, error) {
	req, out := s.DeleteBucketTaggingRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) DeleteBucketWebsiteWithContext(ctx aws.Context, input *s3.DeleteBucketWebsiteInput, opts ...request.Option) (*s3.DeleteBucketWebsiteOutput, error) {
	req, out := s.DeleteBucketWebsiteRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) DeleteObjectWithContext(ctx aws.Context, input *s3.DeleteObjectInput, opts ...request.Option) (*s3.DeleteObjectOutput, error) {
	req, out := s.DeleteObjectRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) DeleteObjectTaggingWithContext(ctx aws.Context, input *s3.DeleteObjectTaggingInput, opts ...request.Option) (*s3.DeleteObjectTaggingOutput, error) {
	req, out := s.DeleteObjectTaggingRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) DeleteObjectsWithContext(ctx aws.Context, input *s3.DeleteObjectsInput, opts ...request.Option) (*s3.DeleteObjectsOutput, error) {
	req, out := s.DeleteObjectsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) DeletePublicAccessBlockWithContext(ctx aws.Context, input *s3.DeletePublicAccessBlockInput, opts ...request.Option) (*s3.DeletePublicAccessBlockOutput, error) {
	req, out := s.DeletePublicAccessBlockRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) GetBucketAccelerateConfigurationWithContext(ctx aws.Context, input *s3.GetBucketAccelerateConfigurationInput, opts ...request.Option) (*s3.GetBucketAccelerateConfigurationOutput, error) {
	req, out := s.GetBucketAccelerateConfigurationRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) GetBucketAclWithContext(ctx aws.Context, input *s3.GetBucketAclInput, opts ...request.Option) (*s3.GetBucketAclOutput, error) {
	req, out := s.GetBucketAclRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) GetBucketAnalyticsConfigurationWithContext(ctx aws.Context, input *s3.GetBucketAnalyticsConfigurationInput, opts ...request.Option) (*s3.GetBucketAnalyticsConfigurationOutput, error) {
	req, out := s.GetBucketAnalyticsConfigurationRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) GetBucketCorsWithContext(ctx aws.Context, input *s3.GetBucketCorsInput, opts ...request.Option) (*s3.GetBucketCorsOutput, error) {
	req, out := s.GetBucketCorsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) GetBucketEncryptionWithContext(ctx aws.Context, input *s3.GetBucketEncryptionInput, opts ...request.Option) (*s3.GetBucketEncryptionOutput, error) {
	req, out := s.GetBucketEncryptionRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) GetBucketIntelligentTieringConfigurationWithContext(ctx aws.Context, input *s3.GetBucketIntelligentTieringConfigurationInput, opts ...request.Option) (*s3.GetBucketIntelligentTieringConfigurationOutput, error) {
	req, out := s.GetBucketIntelligentTieringConfigurationRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) GetBucketInventoryConfigurationWithContext(ctx aws.Context, input *s3.GetBucketInventoryConfigurationInput, opts ...request.Option) (*s3.GetBucketInventoryConfigurationOutput, error) {
	req, out := s.GetBucketInventoryConfigurationRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) GetBucketLifecycleWithContext(ctx aws.Context, input *s3.GetBucketLifecycleInput, opts ...request.Option) (*s3.GetBucketLifecycleOutput, error) {
	req, out := s.GetBucketLifecycleRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) GetBucketLifecycleConfigurationWithContext(ctx aws.Context, input *s3.GetBucketLifecycleConfigurationInput, opts ...request.Option) (*s3.GetBucketLifecycleConfigurationOutput, error) {
	req, out := s.GetBucketLifecycleConfigurationRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) GetBucketLocationWithContext(ctx aws.Context, input *s3.GetBucketLocationInput, opts ...request.Option) (*s3.GetBucketLocationOutput, error) {
	req, out := s.GetBucketLocationRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) GetBucketLoggingWithContext(ctx aws.Context, input *s3.GetBucketLoggingInput, opts ...request.Option) (*s3.GetBucketLoggingOutput, error) {
	req, out := s.GetBucketLoggingRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) GetBucketMetricsConfigurationWithContext(ctx aws.Context, input *s3.GetBucketMetricsConfigurationInput, opts ...request.Option) (*s3.GetBucketMetricsConfigurationOutput, error) {
	req, out := s.GetBucketMetricsConfigurationRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) GetBucketNotificationWithContext(ctx aws.Context, input *s3.GetBucketNotificationConfigurationRequest, opts ...request.Option) (*s3.NotificationConfigurationDeprecated, error) {
	req, out := s.GetBucketNotificationRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) GetBucketNotificationConfigurationWithContext(ctx aws.Context, input *s3.GetBucketNotificationConfigurationRequest, opts ...request.Option) (*s3.NotificationConfiguration, error) {
	req, out := s.GetBucketNotificationConfigurationRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) GetBucketOwnershipControlsWithContext(ctx aws.Context, input *s3.GetBucketOwnershipControlsInput, opts ...request.Option) (*s3.GetBucketOwnershipControlsOutput, error) {
	req, out := s.GetBucketOwnershipControlsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) GetBucketPolicyWithContext(ctx aws.Context, input *s3.GetBucketPolicyInput, opts ...request.Option) (*s3.GetBucketPolicyOutput, error) {
	req, out := s.GetBucketPolicyRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) GetBucketPolicyStatusWithContext(ctx aws.Context, input *s3.GetBucketPolicyStatusInput, opts ...request.Option) (*s3.GetBucketPolicyStatusOutput, error) {
	req, out := s.GetBucketPolicyStatusRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) GetBucketReplicationWithContext(ctx aws.Context, input *s3.GetBucketReplicationInput, opts ...request.Option) (*s3.GetBucketReplicationOutput, error) {
	req, out := s.GetBucketReplicationRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) GetBucketRequestPaymentWithContext(ctx aws.Context, input *s3.GetBucketRequestPaymentInput, opts ...request.Option) (*s3.GetBucketRequestPaymentOutput, error) {
	req, out := s.GetBucketRequestPaymentRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) GetBucketTaggingWithContext(ctx aws.Context, input *s3.GetBucketTaggingInput, opts ...request.Option) (*s3.GetBucketTaggingOutput, error) {
	req, out := s.GetBucketTaggingRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) GetBucketVersioningWithContext(ctx aws.Context, input *s3.GetBucketVersioningInput, opts ...request.Option) (*s3.GetBucketVersioningOutput, error) {
	req, out := s.GetBucketVersioningRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) GetBucketWebsiteWithContext(ctx aws.Context, input *s3.GetBucketWebsiteInput, opts ...request.Option) (*s3.GetBucketWebsiteOutput, error) {
	req, out := s.GetBucketWebsiteRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) GetObjectWithContext(ctx aws.Context, input *s3.GetObjectInput, opts ...request.Option) (*s3.GetObjectOutput, error) {
	req, out := s.GetObjectRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) GetObjectAclWithContext(ctx aws.Context, input *s3.GetObjectAclInput, opts ...request.Option) (*s3.GetObjectAclOutput, error) {
	req, out := s.GetObjectAclRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) GetObjectAttributesWithContext(ctx aws.Context, input *s3.GetObjectAttributesInput, opts ...request.Option) (*s3.GetObjectAttributesOutput, error) {
	req, out := s.GetObjectAttributesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) GetObjectLegalHoldWithContext(ctx aws.Context, input *s3.GetObjectLegalHoldInput, opts ...request.Option) (*s3.GetObjectLegalHoldOutput, error) {
	req, out := s.GetObjectLegalHoldRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) GetObjectLockConfigurationWithContext(ctx aws.Context, input *s3.GetObjectLockConfigurationInput, opts ...request.Option) (*s3.GetObjectLockConfigurationOutput, error) {
	req, out := s.GetObjectLockConfigurationRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) GetObjectRetentionWithContext(ctx aws.Context, input *s3.GetObjectRetentionInput, opts ...request.Option) (*s3.GetObjectRetentionOutput, error) {
	req, out := s.GetObjectRetentionRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) GetObjectTaggingWithContext(ctx aws.Context, input *s3.GetObjectTaggingInput, opts ...request.Option) (*s3.GetObjectTaggingOutput, error) {
	req, out := s.GetObjectTaggingRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) GetObjectTorrentWithContext(ctx aws.Context, input *s3.GetObjectTorrentInput, opts ...request.Option) (*s3.GetObjectTorrentOutput, error) {
	req, out := s.GetObjectTorrentRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) GetPublicAccessBlockWithContext(ctx aws.Context, input *s3.GetPublicAccessBlockInput, opts ...request.Option) (*s3.GetPublicAccessBlockOutput, error) {
	req, out := s.GetPublicAccessBlockRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) HeadBucketWithContext(ctx aws.Context, input *s3.HeadBucketInput, opts ...request.Option) (*s3.HeadBucketOutput, error) {
	req, out := s.HeadBucketRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) HeadObjectWithContext(ctx aws.Context, input *s3.HeadObjectInput, opts ...request.Option) (*s3.HeadObjectOutput, error) {
	req, out := s.HeadObjectRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) ListBucketAnalyticsConfigurationsWithContext(ctx aws.Context, input *s3.ListBucketAnalyticsConfigurationsInput, opts ...request.Option) (*s3.ListBucketAnalyticsConfigurationsOutput, error) {
	req, out := s.ListBucketAnalyticsConfigurationsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) ListBucketIntelligentTieringConfigurationsWithContext(ctx aws.Context, input *s3.ListBucketIntelligentTieringConfigurationsInput, opts ...request.Option) (*s3.ListBucketIntelligentTieringConfigurationsOutput, error) {
	req, out := s.ListBucketIntelligentTieringConfigurationsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) ListBucketInventoryConfigurationsWithContext(ctx aws.Context, input *s3.ListBucketInventoryConfigurationsInput, opts ...request.Option) (*s3.ListBucketInventoryConfigurationsOutput, error) {
	req, out := s.ListBucketInventoryConfigurationsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) ListBucketMetricsConfigurationsWithContext(ctx aws.Context, input *s3.ListBucketMetricsConfigurationsInput, opts ...request.Option) (*s3.ListBucketMetricsConfigurationsOutput, error) {
	req, out := s.ListBucketMetricsConfigurationsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) ListBucketsWithContext(ctx aws.Context, input *s3.ListBucketsInput, opts ...request.Option) (*s3.ListBucketsOutput, error) {
	req, out := s.ListBucketsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) ListMultipartUploadsWithContext(ctx aws.Context, input *s3.ListMultipartUploadsInput, opts ...request.Option) (*s3.ListMultipartUploadsOutput, error) {
	req, out := s.ListMultipartUploadsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) ListMultipartUploadsPagesWithContext(ctx aws.Context, input *s3.ListMultipartUploadsInput, fn func(*s3.ListMultipartUploadsOutput, bool) bool, opts ...request.Option) error {
	p := request.Pagination{
		NewRequest: func() (*request.Request, error) {
			var inCpy *s3.ListMultipartUploadsInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := s.ListMultipartUploadsRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}

	for p.Next() {
		if !fn(p.Page().(*s3.ListMultipartUploadsOutput), !p.HasNextPage()) {
			break
		}
	}

	return p.Err()
}

func (s *S3Link) ListObjectVersionsWithContext(ctx aws.Context, input *s3.ListObjectVersionsInput, opts ...request.Option) (*s3.ListObjectVersionsOutput, error) {
	req, out := s.ListObjectVersionsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) ListObjectVersionsPagesWithContext(ctx aws.Context, input *s3.ListObjectVersionsInput, fn func(*s3.ListObjectVersionsOutput, bool) bool, opts ...request.Option) error {
	p := request.Pagination{
		NewRequest: func() (*request.Request, error) {
			var inCpy *s3.ListObjectVersionsInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := s.ListObjectVersionsRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}

	for p.Next() {
		if !fn(p.Page().(*s3.ListObjectVersionsOutput), !p.HasNextPage()) {
			break
		}
	}

	return p.Err()
}

func (s *S3Link) ListObjectsWithContext(ctx aws.Context, input *s3.ListObjectsInput, opts ...request.Option) (*s3.ListObjectsOutput, error) {
	req, out := s.ListObjectsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) ListObjectsPagesWithContext(ctx aws.Context, input *s3.ListObjectsInput, fn func(*s3.ListObjectsOutput, bool) bool, opts ...request.Option) error {
	p := request.Pagination{
		NewRequest: func() (*request.Request, error) {
			var inCpy *s3.ListObjectsInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := s.ListObjectsRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}

	for p.Next() {
		if !fn(p.Page().(*s3.ListObjectsOutput), !p.HasNextPage()) {
			break
		}
	}

	return p.Err()
}

func (s *S3Link) ListObjectsV2WithContext(ctx aws.Context, input *s3.ListObjectsV2Input, opts ...request.Option) (*s3.ListObjectsV2Output, error) {
	req, out := s.ListObjectsV2Request(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) ListObjectsV2PagesWithContext(ctx aws.Context, input *s3.ListObjectsV2Input, fn func(*s3.ListObjectsV2Output, bool) bool, opts ...request.Option) error {
	p := request.Pagination{
		NewRequest: func() (*request.Request, error) {
			var inCpy *s3.ListObjectsV2Input
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := s.ListObjectsV2Request(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}

	for p.Next() {
		if !fn(p.Page().(*s3.ListObjectsV2Output), !p.HasNextPage()) {
			break
		}
	}

	return p.Err()
}

func (s *S3Link) ListPartsWithContext(ctx aws.Context, input *s3.ListPartsInput, opts ...request.Option) (*s3.ListPartsOutput, error) {
	req, out := s.ListPartsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) ListPartsPagesWithContext(ctx aws.Context, input *s3.ListPartsInput, fn func(*s3.ListPartsOutput, bool) bool, opts ...request.Option) error {
	p := request.Pagination{
		NewRequest: func() (*request.Request, error) {
			var inCpy *s3.ListPartsInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := s.ListPartsRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}

	for p.Next() {
		if !fn(p.Page().(*s3.ListPartsOutput), !p.HasNextPage()) {
			break
		}
	}

	return p.Err()
}

func (s *S3Link) PutBucketAccelerateConfigurationWithContext(ctx aws.Context, input *s3.PutBucketAccelerateConfigurationInput, opts ...request.Option) (*s3.PutBucketAccelerateConfigurationOutput, error) {
	req, out := s.PutBucketAccelerateConfigurationRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) PutBucketAclWithContext(ctx aws.Context, input *s3.PutBucketAclInput, opts ...request.Option) (*s3.PutBucketAclOutput, error) {
	req, out := s.PutBucketAclRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) PutBucketAnalyticsConfigurationWithContext(ctx aws.Context, input *s3.PutBucketAnalyticsConfigurationInput, opts ...request.Option) (*s3.PutBucketAnalyticsConfigurationOutput, error) {
	req, out := s.PutBucketAnalyticsConfigurationRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) PutBucketCorsWithContext(ctx aws.Context, input *s3.PutBucketCorsInput, opts ...request.Option) (*s3.PutBucketCorsOutput, error) {
	req, out := s.PutBucketCorsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) PutBucketEncryptionWithContext(ctx aws.Context, input *s3.PutBucketEncryptionInput, opts ...request.Option) (*s3.PutBucketEncryptionOutput, error) {
	req, out := s.PutBucketEncryptionRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) PutBucketIntelligentTieringConfigurationWithContext(ctx aws.Context, input *s3.PutBucketIntelligentTieringConfigurationInput, opts ...request.Option) (*s3.PutBucketIntelligentTieringConfigurationOutput, error) {
	req, out := s.PutBucketIntelligentTieringConfigurationRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) PutBucketInventoryConfigurationWithContext(ctx aws.Context, input *s3.PutBucketInventoryConfigurationInput, opts ...request.Option) (*s3.PutBucketInventoryConfigurationOutput, error) {
	req, out := s.PutBucketInventoryConfigurationRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) PutBucketLifecycleWithContext(ctx aws.Context, input *s3.PutBucketLifecycleInput, opts ...request.Option) (*s3.PutBucketLifecycleOutput, error) {
	req, out := s.PutBucketLifecycleRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) PutBucketLifecycleConfigurationWithContext(ctx aws.Context, input *s3.PutBucketLifecycleConfigurationInput, opts ...request.Option) (*s3.PutBucketLifecycleConfigurationOutput, error) {
	req, out := s.PutBucketLifecycleConfigurationRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) PutBucketLoggingWithContext(ctx aws.Context, input *s3.PutBucketLoggingInput, opts ...request.Option) (*s3.PutBucketLoggingOutput, error) {
	req, out := s.PutBucketLoggingRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) PutBucketMetricsConfigurationWithContext(ctx aws.Context, input *s3.PutBucketMetricsConfigurationInput, opts ...request.Option) (*s3.PutBucketMetricsConfigurationOutput, error) {
	req, out := s.PutBucketMetricsConfigurationRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) PutBucketNotificationWithContext(ctx aws.Context, input *s3.PutBucketNotificationInput, opts ...request.Option) (*s3.PutBucketNotificationOutput, error) {
	req, out := s.PutBucketNotificationRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) PutBucketNotificationConfigurationWithContext(ctx aws.Context, input *s3.PutBucketNotificationConfigurationInput, opts ...request.Option) (*s3.PutBucketNotificationConfigurationOutput, error) {
	req, out := s.PutBucketNotificationConfigurationRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) PutBucketOwnershipControlsWithContext(ctx aws.Context, input *s3.PutBucketOwnershipControlsInput, opts ...request.Option) (*s3.PutBucketOwnershipControlsOutput, error) {
	req, out := s.PutBucketOwnershipControlsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) PutBucketPolicyWithContext(ctx aws.Context, input *s3.PutBucketPolicyInput, opts ...request.Option) (*s3.PutBucketPolicyOutput, error) {
	req, out := s.PutBucketPolicyRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) PutBucketReplicationWithContext(ctx aws.Context, input *s3.PutBucketReplicationInput, opts ...request.Option) (*s3.PutBucketReplicationOutput, error) {
	req, out := s.PutBucketReplicationRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) PutBucketRequestPaymentWithContext(ctx aws.Context, input *s3.PutBucketRequestPaymentInput, opts ...request.Option) (*s3.PutBucketRequestPaymentOutput, error) {
	req, out := s.PutBucketRequestPaymentRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) PutBucketTaggingWithContext(ctx aws.Context, input *s3.PutBucketTaggingInput, opts ...request.Option) (*s3.PutBucketTaggingOutput, error) {
	req, out := s.PutBucketTaggingRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) PutBucketVersioningWithContext(ctx aws.Context, input *s3.PutBucketVersioningInput, opts ...request.Option) (*s3.PutBucketVersioningOutput, error) {
	req, out := s.PutBucketVersioningRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) PutBucketWebsiteWithContext(ctx aws.Context, input *s3.PutBucketWebsiteInput, opts ...request.Option) (*s3.PutBucketWebsiteOutput, error) {
	req, out := s.PutBucketWebsiteRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) PutObjectWithContext(ctx aws.Context, input *s3.PutObjectInput, opts ...request.Option) (*s3.PutObjectOutput, error) {
	req, out := s.PutObjectRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) PutObjectAclWithContext(ctx aws.Context, input *s3.PutObjectAclInput, opts ...request.Option) (*s3.PutObjectAclOutput, error) {
	req, out := s.PutObjectAclRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) PutObjectLegalHoldWithContext(ctx aws.Context, input *s3.PutObjectLegalHoldInput, opts ...request.Option) (*s3.PutObjectLegalHoldOutput, error) {
	req, out := s.PutObjectLegalHoldRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) PutObjectLockConfigurationWithContext(ctx aws.Context, input *s3.PutObjectLockConfigurationInput, opts ...request.Option) (*s3.PutObjectLockConfigurationOutput, error) {
	req, out := s.PutObjectLockConfigurationRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) PutObjectRetentionWithContext(ctx aws.Context, input *s3.PutObjectRetentionInput, opts ...request.Option) (*s3.PutObjectRetentionOutput, error) {
	req, out := s.PutObjectRetentionRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) PutObjectTaggingWithContext(ctx aws.Context, input *s3.PutObjectTaggingInput, opts ...request.Option) (*s3.PutObjectTaggingOutput, error) {
	req, out := s.PutObjectTaggingRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) PutPublicAccessBlockWithContext(ctx aws.Context, input *s3.PutPublicAccessBlockInput, opts ...request.Option) (*s3.PutPublicAccessBlockOutput, error) {
	req, out := s.PutPublicAccessBlockRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) RestoreObjectWithContext(ctx aws.Context, input *s3.RestoreObjectInput, opts ...request.Option) (*s3.RestoreObjectOutput, error) {
	req, out := s.RestoreObjectRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) SelectObjectContentWithContext(ctx aws.Context, input *s3.SelectObjectContentInput, opts ...request.Option) (*s3.SelectObjectContentOutput, error) {
	req, out := s.SelectObjectContentRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) UploadPartWithContext(ctx aws.Context, input *s3.UploadPartInput, opts ...request.Option) (*s3.UploadPartOutput, error) {
	req, out := s.UploadPartRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) UploadPartCopyWithContext(ctx aws.Context, input *s3.UploadPartCopyInput, opts ...request.Option) (*s3.UploadPartCopyOutput, error) {
	req, out := s.UploadPartCopyRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) WriteGetObjectResponseWithContext(ctx aws.Context, input *s3.WriteGetObjectResponseInput, opts ...request.Option) (*s3.WriteGetObjectResponseOutput, error) {
	req, out := s.WriteGetObjectResponseRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

func (s *S3Link) WaitUntilBucketExistsWithContext(ctx aws.Context, input *s3.HeadBucketInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilBucketExists",
		MaxAttempts: 20,
		Delay:       request.ConstantWaiterDelay(5 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:    request.SuccessWaiterState,
				Matcher:  request.StatusWaiterMatch,
				Expected: 200,
			},
			{
				State:    request.SuccessWaiterState,
				Matcher:  request.StatusWaiterMatch,
				Expected: 301,
			},
			{
				State:    request.SuccessWaiterState,
				Matcher:  request.StatusWaiterMatch,
				Expected: 403,
			},
			{
				State:    request.RetryWaiterState,
				Matcher:  request.StatusWaiterMatch,
				Expected: 404,
			},
		},
		Logger: s.logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *s3.HeadBucketInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := s.HeadBucketRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

func (s *S3Link) WaitUntilBucketNotExistsWithContext(ctx aws.Context, input *s3.HeadBucketInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilBucketNotExists",
		MaxAttempts: 20,
		Delay:       request.ConstantWaiterDelay(5 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:    request.SuccessWaiterState,
				Matcher:  request.StatusWaiterMatch,
				Expected: 404,
			},
		},
		Logger: s.logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *s3.HeadBucketInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := s.HeadBucketRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

func (s *S3Link) WaitUntilObjectExistsWithContext(ctx aws.Context, input *s3.HeadObjectInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilObjectExists",
		MaxAttempts: 20,
		Delay:       request.ConstantWaiterDelay(5 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:    request.SuccessWaiterState,
				Matcher:  request.StatusWaiterMatch,
				Expected: 200,
			},
			{
				State:    request.RetryWaiterState,
				Matcher:  request.StatusWaiterMatch,
				Expected: 404,
			},
		},
		Logger: s.logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *s3.HeadObjectInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := s.HeadObjectRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

func (s *S3Link) WaitUntilObjectNotExistsWithContext(ctx aws.Context, input *s3.HeadObjectInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilObjectNotExists",
		MaxAttempts: 20,
		Delay:       request.ConstantWaiterDelay(5 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:    request.SuccessWaiterState,
				Matcher:  request.StatusWaiterMatch,
				Expected: 404,
			},
		},
		Logger: s.logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *s3.HeadObjectInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := s.HeadObjectRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}
