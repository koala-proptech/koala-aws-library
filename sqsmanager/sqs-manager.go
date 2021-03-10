package sqsmanager

import(
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/google/uuid"
)

type ISqsManager interface{
	SendMessage(data map[string]*sqs.MessageAttributeValue, MessageGrouptId string, QueueUrl string)(error)
}

type SqsManager struct{
	Config SqsConfiguration
}

func NewSqsManager(config SqsConfiguration)ISqsManager{
	return &SqsManager{Config: config,}
}

func(r *SqsManager)SendMessage(data map[string]*sqs.MessageAttributeValue, MessageGrouptId string, QueueUrl string)(error){
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(r.Config.Region),
		Credentials: credentials.NewStaticCredentials(r.Config.KeyId, r.Config.SecretKey, ""),
	})
	// _=sess
	if err!=nil{
		// fmt.Println(err)
		return err
	}
	_ = err
	
	svc := sqs.New(sess)
	output, err := svc.SendMessage(&sqs.SendMessageInput{
		DelaySeconds:      aws.Int64(0),
		MessageAttributes: data,
		MessageBody:       aws.String(uuid.New().String()),
		QueueUrl:          aws.String(QueueUrl),
		MessageGroupId:    aws.String(MessageGrouptId),
	})
	if err!=nil{
		// fmt.Println(err)
		return err
	}
	// fmt.Println(output)
	_ = output
	return nil
}