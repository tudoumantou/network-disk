package store

import (
	"gopkg.in/amz.v1/s3"
	"gopkg.in/amz.v1/aws"
)


var cephConn *s3.S3

func GetCephConnection() *s3.S3{
	if(cephConn!=nil){
		return cephConn
	}
	// 1.初始化ceph的一些信息
	auth := aws.Auth{
		AccessKey:"",
		SecretKey:"",
	}
	// 2.创建s3类型的连接
	curRegion := aws.Region{
		Name:                 "default",
		EC2Endpoint:          "http://127.0.0.1:9080",
		S3Endpoint:           "http://127.0.0.1:9080",
		S3BucketEndpoint:     "",
		S3LocationConstraint: false,
		S3LowercaseBucket:    false,
		Sign:                 aws.SignV2,
	}
	return s3.New(auth, curRegion)
}
