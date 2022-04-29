# s3link
An experiment in using the [Storj Network](https://storj.io) via the [AWS SDK for Go](https://github.com/aws/aws-sdk-go/)

# background

In order to increase user adoption, Storj runs an S3-compliant web service called Gateway-MT.
Gateway-MT has increased user adoption, but compromises on a core feature of the Storj Network in that it requires centralized HTTP endpoints.
Analysis of various Amazon SDKs indicates it would be difficult to code generate "one size fits all" S3-API-compatible bindings across languages.
Indeed [and perhaps because they span service types], most Amazon SDKs can be described as having "HTTP as the architecture".

Thus an initial attempt to integrate with the AWS SDK for Go by implementing [an interface](https://github.com/aws/aws-sdk-go/blob/main/service/s3/s3iface/interface.go) was abandoned.  This more language agnostic solution accepts the overhead of serializing code to HTTP streams, which are then processed without actually sending those streams.  Application logic from Storj's [Gateway-ST](https://github.com/storj/gateway-st/) project handles adapting the request back to native [libuplink](https://github.com/storj/uplink) Storj Network requests.

# usage

The Gateway-ST / Minio libraries used by this code implement request signing.  Thus, the credentials and host names must match between the AWS SDK for Go session variable and the S3Link logic.  EG: 


		awsSession, _ := session.NewSession(&aws.Config{
			Credentials:      credentials.NewStaticCredentials("placeholder", "placeholder", ""),
			Endpoint:         aws.String("http://s3.amazonaws.com"),
			Region:           aws.String("us-east-1"),
			S3ForcePathStyle: aws.Bool(true),
		})

		client := s3.New(awsSession, aws.NewConfig().WithHTTPClient(s3link.NewHTTPClient()))
		downloader := s3manager.NewDownloader(awsSession, func(d *s3manager.Downloader) { d.S3 = client }),
		uploader := s3manager.NewUploader(awsSession, func(u *s3manager.Uploader) { u.S3 = client }),