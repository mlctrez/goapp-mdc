module github.com/mlctrez/goapp-mdc/scripts

go 1.17

replace (
	github.com/mlctrez/goapp-mdc latest => /home/mattman/golang/goapp-mdc
)

require (
	github.com/aws/aws-sdk-go-v2 v1.10.0
	github.com/aws/aws-sdk-go-v2/config v1.9.0
	github.com/aws/aws-sdk-go-v2/service/s3 v1.17.0
	github.com/mlctrez/goapp-mdc latest
)

require (
	github.com/aws/aws-sdk-go-v2/credentials v1.5.0 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.7.0 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.2.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.4.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.4.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.8.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.5.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/sts v1.8.0 // indirect
	github.com/aws/smithy-go v1.8.1 // indirect
)
