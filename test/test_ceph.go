package main

import (
	"network-disk/store/ceph"
	"golang.org/x/tools/go/ssa/interp/testdata/src/fmt"
	"gopkg.in/amz.v1/s3"
	"os"
)

func main()  {
	bucket := ceph.GetCephBucket("userfile")

	d, _ := bucket.Get("")
	tmpFile, _ := os.Create("/tmp/test_file")
	tmpFile.Write(d)
	return

	// 创建一个新的bucket
	err := bucket.PutBucket(s3.PublicRead)
	fmt.Println("create bucket err:%v\n", err)

	res, err := bucket.List("", "", "", 100)
	fmt.Println("object keys :%+v\n", res)

	//// 新上传一个对象
	//err = bucket.Put("/testupload/a.txt", []byte("just for test"), "octet-stream", s3.PublicRead)
	//fmt.Printf("upload err: %+v\n", err)
	//
	//// 查询这个bucket下面指定条件的object keys
	//res, err = bucket.List("", "", "", 100)
	//fmt.Printf("object keys: %+v\n", res)
}