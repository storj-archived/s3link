# s3link
An experiment in using the [Storj Network](https://storj.io) via the [AWS SDK for Go](https://github.com/aws/aws-sdk-go/)

# background

In order to increase user adoption, Storj runs an S3-compliant web service called Gateway-MT.
Gateway-MT has increased user adoption, but compromises on a core feature of the Storj Network in that it requires centralized HTTP endpoints.
Analysis of various Amazon SDKs indicates it would be difficult to code generate a "one size fits all" S3-API-compatible bindings across languages.
Indeed [and perhaps because they span service types], most Amazon SDKs can be described as having "HTTP as the architecture".

Thus an initial attempt to integrate with the AWS SDK for Go by implementing [an interface](https://github.com/aws/aws-sdk-go/blob/main/service/s3/s3iface/interface.go) was abandoned.  This more language agnostic solution accepts the overhead of serializing code to HTTP streams, which are then processed without actually sending those streams.  Application logic from Storj's [Gateway-ST](https://github.com/storj/gateway-st/) project handles adapting the request back to native [libuplink](https://github.com/storj/uplink) Storj Network requests.

# license

Due to re-use of the AWS SDK for Go, we'll preserve the Apache2 license.