package main

import (
	"fmt"

	"github.com/koala-proptech/koala-aws-library/s3manager"
)

var Config = s3manager.S3Configuration{
	
	"AWS_Key_ID_XXXX",
	"Secret_key_XXX",
	"BucketXXX",
	"Basepath-xxx", // env(staging,production)
	"public-read",
	"ap-southeast-1",
}
var A s3manager.IS3Manager = s3manager.NewS3Manager(Config)

func main() {

	path, err := A.Upload("/test/", "agustusan.jpg")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(path)

	// s3 url
	// https://koalaprop.s3-ap-southeast-1.amazonaws.com/staging/test/agustusan.jpg
	// a:=fmt.Sprintf(`/%s/%s`,Config.BasePath,"test/agustusan.jp")
	// fmt.Println(a)
	// err := A.Delete(fmt.Sprintf(`/%s/%s`, Config.BasePath, "test/agustusan.jpg"))
	// if err != nil {
	// 	fmt.Println(err)
	// }

	fmt.Println("success")
}
