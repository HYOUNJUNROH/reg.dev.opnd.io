package s3

import (
	"fmt"
	"os"

	"git.dev.opnd.io/gc/backend-admin/pkg/config"
	awss3 "github.com/aws/aws-sdk-go/service/s3"
	"github.com/qor/oss/s3"
)

var Client *s3.Client

func Init() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
			os.Exit(2)
		}
	}()

	Client = s3.New(&s3.Config{
		AccessID:         config.Config.S3.AccessKeyID,
		AccessKey:        config.Config.S3.SecretAccessKey,
		Region:           config.Config.S3.Region,
		Bucket:           config.Config.S3.Bucket,
		Endpoint:         config.Config.S3.Endpoint,
		S3Endpoint:       config.Config.S3.Endpoint,
		ACL:              awss3.BucketCannedACLPublicRead,
		S3ForcePathStyle: config.Config.S3.ForcePathStyle,
	})
}
