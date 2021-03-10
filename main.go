package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/koala-proptech/koala-aws-library/s3manager"
	"github.com/koala-proptech/koala-aws-library/sqsmanager"
)

var S3Config = s3manager.S3Configuration{
	
	"AWS_Key_ID_XXXX",
	"Secret_key_XXX",
	"BucketXXX",
	"Basepath-xxx", // env(staging,production)
	"public-read",
	"ap-southeast-1",
}

var SqsConfig=sqsmanager.SqsConfiguration{
	"xxx",
	"yyy",
	"ap-southeast-1",
}
var S3 s3manager.IS3Manager = s3manager.NewS3Manager(S3Config)
var SQS sqsmanager.ISqsManager=sqsmanager.NewSqsManager(SqsConfig)

func GenereateAttributeMessage(Param string, NotificationType string, PhoneNumber int64) map[string]*sqs.MessageAttributeValue {
	fmt.Println(string(PhoneNumber))
	MessageAttributes := map[string]*sqs.MessageAttributeValue{
		"notification_type": &sqs.MessageAttributeValue{
			DataType:    aws.String("String"),
			StringValue: aws.String(NotificationType),
		},
		"data": &sqs.MessageAttributeValue{
			DataType:    aws.String("String"),
			StringValue: aws.String(Param),
		},
		"phone_number": &sqs.MessageAttributeValue{
			DataType:    aws.String("String"),
			StringValue: aws.String(fmt.Sprintf("%d", PhoneNumber)),
		},
	}
	return MessageAttributes

}
func main() {

	path, err := S3.Upload("/test/", "agustusan.jpg")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(path)

	// s3 url
	// https://koalaprop.s3-ap-southeast-1.amazonaws.com/staging/test/agustusan.jpg
	// a:=fmt.Sprintf(`/%s/%s`,Config.BasePath,"test/agustusan.jp")
	// fmt.Println(a)
	// err := S3.Delete(fmt.Sprintf(`/%s/%s`, Config.BasePath, "test/agustusan.jpg"))
	// if err != nil {
	// 	fmt.Println(err)
	// }

	
	Message:=GenereateAttributeMessage("test","whatsapp",6287777000056)
	err=SQS.SendMessage(Message,"XXXXX-XXXX-QW","https://sqs.ap-southeast-1.amazonaws.com/737690422155/staging_notifcation_services.fifo")

	fmt.Println("success")
}
