package s3link

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func (s *S3Link) AbortMultipartUpload(input *s3.AbortMultipartUploadInput) (*s3.AbortMultipartUploadOutput, error) {
	return s.AbortMultipartUploadWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) CompleteMultipartUpload(input *s3.CompleteMultipartUploadInput) (*s3.CompleteMultipartUploadOutput, error) {
	return s.CompleteMultipartUploadWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) CopyObject(input *s3.CopyObjectInput) (*s3.CopyObjectOutput, error) {
	return s.CopyObjectWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) CreateBucket(input *s3.CreateBucketInput) (*s3.CreateBucketOutput, error) {
	return s.CreateBucketWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) CreateMultipartUpload(input *s3.CreateMultipartUploadInput) (*s3.CreateMultipartUploadOutput, error) {
	return s.CreateMultipartUploadWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) DeleteBucket(input *s3.DeleteBucketInput) (*s3.DeleteBucketOutput, error) {
	return s.DeleteBucketWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) DeleteBucketAnalyticsConfiguration(input *s3.DeleteBucketAnalyticsConfigurationInput) (*s3.DeleteBucketAnalyticsConfigurationOutput, error) {
	return s.DeleteBucketAnalyticsConfigurationWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) DeleteBucketCors(input *s3.DeleteBucketCorsInput) (*s3.DeleteBucketCorsOutput, error) {
	return s.DeleteBucketCorsWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) DeleteBucketEncryption(input *s3.DeleteBucketEncryptionInput) (*s3.DeleteBucketEncryptionOutput, error) {
	return s.DeleteBucketEncryptionWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) DeleteBucketIntelligentTieringConfiguration(input *s3.DeleteBucketIntelligentTieringConfigurationInput) (*s3.DeleteBucketIntelligentTieringConfigurationOutput, error) {
	return s.DeleteBucketIntelligentTieringConfigurationWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) DeleteBucketInventoryConfiguration(input *s3.DeleteBucketInventoryConfigurationInput) (*s3.DeleteBucketInventoryConfigurationOutput, error) {
	return s.DeleteBucketInventoryConfigurationWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) DeleteBucketLifecycle(input *s3.DeleteBucketLifecycleInput) (*s3.DeleteBucketLifecycleOutput, error) {
	return s.DeleteBucketLifecycleWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) DeleteBucketMetricsConfiguration(input *s3.DeleteBucketMetricsConfigurationInput) (*s3.DeleteBucketMetricsConfigurationOutput, error) {
	return s.DeleteBucketMetricsConfigurationWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) DeleteBucketOwnershipControls(input *s3.DeleteBucketOwnershipControlsInput) (*s3.DeleteBucketOwnershipControlsOutput, error) {
	return s.DeleteBucketOwnershipControlsWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) DeleteBucketPolicy(input *s3.DeleteBucketPolicyInput) (*s3.DeleteBucketPolicyOutput, error) {
	return s.DeleteBucketPolicyWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) DeleteBucketReplication(input *s3.DeleteBucketReplicationInput) (*s3.DeleteBucketReplicationOutput, error) {
	return s.DeleteBucketReplicationWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) DeleteBucketTagging(input *s3.DeleteBucketTaggingInput) (*s3.DeleteBucketTaggingOutput, error) {
	return s.DeleteBucketTaggingWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) DeleteBucketWebsite(input *s3.DeleteBucketWebsiteInput) (*s3.DeleteBucketWebsiteOutput, error) {
	return s.DeleteBucketWebsiteWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) DeleteObject(input *s3.DeleteObjectInput) (*s3.DeleteObjectOutput, error) {
	return s.DeleteObjectWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) DeleteObjectTagging(input *s3.DeleteObjectTaggingInput) (*s3.DeleteObjectTaggingOutput, error) {
	return s.DeleteObjectTaggingWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) DeleteObjects(input *s3.DeleteObjectsInput) (*s3.DeleteObjectsOutput, error) {
	return s.DeleteObjectsWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) DeletePublicAccessBlock(input *s3.DeletePublicAccessBlockInput) (*s3.DeletePublicAccessBlockOutput, error) {
	return s.DeletePublicAccessBlockWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) GetBucketAccelerateConfiguration(input *s3.GetBucketAccelerateConfigurationInput) (*s3.GetBucketAccelerateConfigurationOutput, error) {
	return s.GetBucketAccelerateConfigurationWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) GetBucketAcl(input *s3.GetBucketAclInput) (*s3.GetBucketAclOutput, error) {
	return s.GetBucketAclWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) GetBucketAnalyticsConfiguration(input *s3.GetBucketAnalyticsConfigurationInput) (*s3.GetBucketAnalyticsConfigurationOutput, error) {
	return s.GetBucketAnalyticsConfigurationWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) GetBucketCors(input *s3.GetBucketCorsInput) (*s3.GetBucketCorsOutput, error) {
	return s.GetBucketCorsWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) GetBucketEncryption(input *s3.GetBucketEncryptionInput) (*s3.GetBucketEncryptionOutput, error) {
	return s.GetBucketEncryptionWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) GetBucketIntelligentTieringConfiguration(input *s3.GetBucketIntelligentTieringConfigurationInput) (*s3.GetBucketIntelligentTieringConfigurationOutput, error) {
	return s.GetBucketIntelligentTieringConfigurationWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) GetBucketInventoryConfiguration(input *s3.GetBucketInventoryConfigurationInput) (*s3.GetBucketInventoryConfigurationOutput, error) {
	return s.GetBucketInventoryConfigurationWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) GetBucketLifecycle(input *s3.GetBucketLifecycleInput) (*s3.GetBucketLifecycleOutput, error) {
	return s.GetBucketLifecycleWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) GetBucketLifecycleConfiguration(input *s3.GetBucketLifecycleConfigurationInput) (*s3.GetBucketLifecycleConfigurationOutput, error) {
	return s.GetBucketLifecycleConfigurationWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) GetBucketLocation(input *s3.GetBucketLocationInput) (*s3.GetBucketLocationOutput, error) {
	return s.GetBucketLocationWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) GetBucketLogging(input *s3.GetBucketLoggingInput) (*s3.GetBucketLoggingOutput, error) {
	return s.GetBucketLoggingWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) GetBucketMetricsConfiguration(input *s3.GetBucketMetricsConfigurationInput) (*s3.GetBucketMetricsConfigurationOutput, error) {
	return s.GetBucketMetricsConfigurationWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) GetBucketNotification(input *s3.GetBucketNotificationConfigurationRequest) (*s3.NotificationConfigurationDeprecated, error) {
	return s.GetBucketNotificationWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) GetBucketNotificationConfiguration(input *s3.GetBucketNotificationConfigurationRequest) (*s3.NotificationConfiguration, error) {
	return s.GetBucketNotificationConfigurationWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) GetBucketOwnershipControls(input *s3.GetBucketOwnershipControlsInput) (*s3.GetBucketOwnershipControlsOutput, error) {
	return s.GetBucketOwnershipControlsWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) GetBucketPolicy(input *s3.GetBucketPolicyInput) (*s3.GetBucketPolicyOutput, error) {
	return s.GetBucketPolicyWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) GetBucketPolicyStatus(input *s3.GetBucketPolicyStatusInput) (*s3.GetBucketPolicyStatusOutput, error) {
	return s.GetBucketPolicyStatusWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) GetBucketReplication(input *s3.GetBucketReplicationInput) (*s3.GetBucketReplicationOutput, error) {
	return s.GetBucketReplicationWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) GetBucketRequestPayment(input *s3.GetBucketRequestPaymentInput) (*s3.GetBucketRequestPaymentOutput, error) {
	return s.GetBucketRequestPaymentWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) GetBucketTagging(input *s3.GetBucketTaggingInput) (*s3.GetBucketTaggingOutput, error) {
	return s.GetBucketTaggingWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) GetBucketVersioning(input *s3.GetBucketVersioningInput) (*s3.GetBucketVersioningOutput, error) {
	return s.GetBucketVersioningWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) GetBucketWebsite(input *s3.GetBucketWebsiteInput) (*s3.GetBucketWebsiteOutput, error) {
	return s.GetBucketWebsiteWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) GetObject(input *s3.GetObjectInput) (*s3.GetObjectOutput, error) {
	return s.GetObjectWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) GetObjectAcl(input *s3.GetObjectAclInput) (*s3.GetObjectAclOutput, error) {
	return s.GetObjectAclWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) GetObjectAttributes(input *s3.GetObjectAttributesInput) (*s3.GetObjectAttributesOutput, error) {
	return s.GetObjectAttributesWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) GetObjectLegalHold(input *s3.GetObjectLegalHoldInput) (*s3.GetObjectLegalHoldOutput, error) {
	return s.GetObjectLegalHoldWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) GetObjectLockConfiguration(input *s3.GetObjectLockConfigurationInput) (*s3.GetObjectLockConfigurationOutput, error) {
	return s.GetObjectLockConfigurationWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) GetObjectRetention(input *s3.GetObjectRetentionInput) (*s3.GetObjectRetentionOutput, error) {
	return s.GetObjectRetentionWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) GetObjectTagging(input *s3.GetObjectTaggingInput) (*s3.GetObjectTaggingOutput, error) {
	return s.GetObjectTaggingWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) GetObjectTorrent(input *s3.GetObjectTorrentInput) (*s3.GetObjectTorrentOutput, error) {
	return s.GetObjectTorrentWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) GetPublicAccessBlock(input *s3.GetPublicAccessBlockInput) (*s3.GetPublicAccessBlockOutput, error) {
	return s.GetPublicAccessBlockWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) HeadBucket(input *s3.HeadBucketInput) (*s3.HeadBucketOutput, error) {
	return s.HeadBucketWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) HeadObject(input *s3.HeadObjectInput) (*s3.HeadObjectOutput, error) {
	return s.HeadObjectWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) ListBucketAnalyticsConfigurations(input *s3.ListBucketAnalyticsConfigurationsInput) (*s3.ListBucketAnalyticsConfigurationsOutput, error) {
	return s.ListBucketAnalyticsConfigurationsWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) ListBucketIntelligentTieringConfigurations(input *s3.ListBucketIntelligentTieringConfigurationsInput) (*s3.ListBucketIntelligentTieringConfigurationsOutput, error) {
	return s.ListBucketIntelligentTieringConfigurationsWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) ListBucketInventoryConfigurations(input *s3.ListBucketInventoryConfigurationsInput) (*s3.ListBucketInventoryConfigurationsOutput, error) {
	return s.ListBucketInventoryConfigurationsWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) ListBucketMetricsConfigurations(input *s3.ListBucketMetricsConfigurationsInput) (*s3.ListBucketMetricsConfigurationsOutput, error) {
	return s.ListBucketMetricsConfigurationsWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) ListBuckets(input *s3.ListBucketsInput) (*s3.ListBucketsOutput, error) {
	return s.ListBucketsWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) ListMultipartUploads(input *s3.ListMultipartUploadsInput) (*s3.ListMultipartUploadsOutput, error) {
	return s.ListMultipartUploadsWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) ListMultipartUploadsPages(input *s3.ListMultipartUploadsInput, fn func(*s3.ListMultipartUploadsOutput, bool) bool) error {
	return s.ListMultipartUploadsPagesWithContext(aws.BackgroundContext(), input, fn)
}

func (s *S3Link) ListObjectVersions(input *s3.ListObjectVersionsInput) (*s3.ListObjectVersionsOutput, error) {
	return s.ListObjectVersionsWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) ListObjectVersionsPages(input *s3.ListObjectVersionsInput, fn func(*s3.ListObjectVersionsOutput, bool) bool) error {
	return s.ListObjectVersionsPagesWithContext(aws.BackgroundContext(), input, fn)
}

func (s *S3Link) ListObjects(input *s3.ListObjectsInput) (*s3.ListObjectsOutput, error) {
	return s.ListObjectsWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) ListObjectsPages(input *s3.ListObjectsInput, fn func(*s3.ListObjectsOutput, bool) bool) error {
	return s.ListObjectsPagesWithContext(aws.BackgroundContext(), input, fn)
}

func (s *S3Link) ListObjectsV2(input *s3.ListObjectsV2Input) (*s3.ListObjectsV2Output, error) {
	return s.ListObjectsV2WithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) ListObjectsV2Pages(input *s3.ListObjectsV2Input, fn func(*s3.ListObjectsV2Output, bool) bool) error {
	return s.ListObjectsV2PagesWithContext(aws.BackgroundContext(), input, fn)
}

func (s *S3Link) ListParts(input *s3.ListPartsInput) (*s3.ListPartsOutput, error) {
	return s.ListPartsWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) ListPartsPages(input *s3.ListPartsInput, fn func(*s3.ListPartsOutput, bool) bool) error {
	return s.ListPartsPagesWithContext(aws.BackgroundContext(), input, fn)
}

func (s *S3Link) PutBucketAccelerateConfiguration(input *s3.PutBucketAccelerateConfigurationInput) (*s3.PutBucketAccelerateConfigurationOutput, error) {
	return s.PutBucketAccelerateConfigurationWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) PutBucketAcl(input *s3.PutBucketAclInput) (*s3.PutBucketAclOutput, error) {
	return s.PutBucketAclWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) PutBucketAnalyticsConfiguration(input *s3.PutBucketAnalyticsConfigurationInput) (*s3.PutBucketAnalyticsConfigurationOutput, error) {
	return s.PutBucketAnalyticsConfigurationWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) PutBucketCors(input *s3.PutBucketCorsInput) (*s3.PutBucketCorsOutput, error) {
	return s.PutBucketCorsWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) PutBucketEncryption(input *s3.PutBucketEncryptionInput) (*s3.PutBucketEncryptionOutput, error) {
	return s.PutBucketEncryptionWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) PutBucketIntelligentTieringConfiguration(input *s3.PutBucketIntelligentTieringConfigurationInput) (*s3.PutBucketIntelligentTieringConfigurationOutput, error) {
	return s.PutBucketIntelligentTieringConfigurationWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) PutBucketInventoryConfiguration(input *s3.PutBucketInventoryConfigurationInput) (*s3.PutBucketInventoryConfigurationOutput, error) {
	return s.PutBucketInventoryConfigurationWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) PutBucketLifecycle(input *s3.PutBucketLifecycleInput) (*s3.PutBucketLifecycleOutput, error) {
	return s.PutBucketLifecycleWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) PutBucketLifecycleConfiguration(input *s3.PutBucketLifecycleConfigurationInput) (*s3.PutBucketLifecycleConfigurationOutput, error) {
	return s.PutBucketLifecycleConfigurationWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) PutBucketLogging(input *s3.PutBucketLoggingInput) (*s3.PutBucketLoggingOutput, error) {
	return s.PutBucketLoggingWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) PutBucketMetricsConfiguration(input *s3.PutBucketMetricsConfigurationInput) (*s3.PutBucketMetricsConfigurationOutput, error) {
	return s.PutBucketMetricsConfigurationWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) PutBucketNotification(input *s3.PutBucketNotificationInput) (*s3.PutBucketNotificationOutput, error) {
	return s.PutBucketNotificationWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) PutBucketNotificationConfiguration(input *s3.PutBucketNotificationConfigurationInput) (*s3.PutBucketNotificationConfigurationOutput, error) {
	return s.PutBucketNotificationConfigurationWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) PutBucketOwnershipControls(input *s3.PutBucketOwnershipControlsInput) (*s3.PutBucketOwnershipControlsOutput, error) {
	return s.PutBucketOwnershipControlsWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) PutBucketPolicy(input *s3.PutBucketPolicyInput) (*s3.PutBucketPolicyOutput, error) {
	return s.PutBucketPolicyWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) PutBucketReplication(input *s3.PutBucketReplicationInput) (*s3.PutBucketReplicationOutput, error) {
	return s.PutBucketReplicationWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) PutBucketRequestPayment(input *s3.PutBucketRequestPaymentInput) (*s3.PutBucketRequestPaymentOutput, error) {
	return s.PutBucketRequestPaymentWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) PutBucketTagging(input *s3.PutBucketTaggingInput) (*s3.PutBucketTaggingOutput, error) {
	return s.PutBucketTaggingWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) PutBucketVersioning(input *s3.PutBucketVersioningInput) (*s3.PutBucketVersioningOutput, error) {
	return s.PutBucketVersioningWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) PutBucketWebsite(input *s3.PutBucketWebsiteInput) (*s3.PutBucketWebsiteOutput, error) {
	return s.PutBucketWebsiteWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) PutObject(input *s3.PutObjectInput) (*s3.PutObjectOutput, error) {
	return s.PutObjectWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) PutObjectAcl(input *s3.PutObjectAclInput) (*s3.PutObjectAclOutput, error) {
	return s.PutObjectAclWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) PutObjectLegalHold(input *s3.PutObjectLegalHoldInput) (*s3.PutObjectLegalHoldOutput, error) {
	return s.PutObjectLegalHoldWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) PutObjectLockConfiguration(input *s3.PutObjectLockConfigurationInput) (*s3.PutObjectLockConfigurationOutput, error) {
	return s.PutObjectLockConfigurationWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) PutObjectRetention(input *s3.PutObjectRetentionInput) (*s3.PutObjectRetentionOutput, error) {
	return s.PutObjectRetentionWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) PutObjectTagging(input *s3.PutObjectTaggingInput) (*s3.PutObjectTaggingOutput, error) {
	return s.PutObjectTaggingWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) PutPublicAccessBlock(input *s3.PutPublicAccessBlockInput) (*s3.PutPublicAccessBlockOutput, error) {
	return s.PutPublicAccessBlockWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) RestoreObject(input *s3.RestoreObjectInput) (*s3.RestoreObjectOutput, error) {
	return s.RestoreObjectWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) SelectObjectContent(input *s3.SelectObjectContentInput) (*s3.SelectObjectContentOutput, error) {
	return s.SelectObjectContentWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) UploadPart(input *s3.UploadPartInput) (*s3.UploadPartOutput, error) {
	return s.UploadPartWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) UploadPartCopy(input *s3.UploadPartCopyInput) (*s3.UploadPartCopyOutput, error) {
	return s.UploadPartCopyWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) WriteGetObjectResponse(input *s3.WriteGetObjectResponseInput) (*s3.WriteGetObjectResponseOutput, error) {
	return s.WriteGetObjectResponseWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) WaitUntilBucketExists(input *s3.HeadBucketInput) error {
	return s.WaitUntilBucketExistsWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) WaitUntilBucketNotExists(input *s3.HeadBucketInput) error {
	return s.WaitUntilBucketNotExistsWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) WaitUntilObjectExists(input *s3.HeadObjectInput) error {
	return s.WaitUntilObjectExistsWithContext(aws.BackgroundContext(), input)
}

func (s *S3Link) WaitUntilObjectNotExists(input *s3.HeadObjectInput) error {
	return s.WaitUntilObjectNotExistsWithContext(aws.BackgroundContext(), input)
}
