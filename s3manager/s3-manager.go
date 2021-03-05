package s3manager

import (
	"bytes"
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type IS3Manager interface {
	Upload(s3path string, filename string) (path string, err error)
	Delete(pathUrl string) error
}

type S3Manager struct {
	Config S3Configuration
}

func NewS3Manager(config S3Configuration) IS3Manager {
	return &S3Manager{config}
}

func (s3m *S3Manager) Delete(pathUrl string) error {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(s3m.Config.Region),
		Credentials: credentials.NewStaticCredentials(s3m.Config.KeyId, s3m.Config.SecretKey, ""),
	})
	svc := s3.New(sess)
	request := &s3.DeleteObjectInput{
		Bucket: aws.String(s3m.Config.BucketName),
		Key:    aws.String(pathUrl),
	}
	_, err = svc.DeleteObject(request)
	if err != nil {
		return err
	}
	return nil

}

func (s3 *S3Manager) Upload(s3path string, filename string) (path string, err error) {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(s3.Config.Region),
		Credentials: credentials.NewStaticCredentials(s3.Config.KeyId, s3.Config.SecretKey, ""),
	})
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	s3Fullpath, err := s3.uploadFileToS3(sess, s3path, filename)
	if err != nil {
		
		return "", err
	}
	err = os.Remove(filename)
	if err != nil {
		return "", err
	}
	
	return s3Fullpath, nil
}

func (r *S3Manager) uploadFileToS3(s *session.Session, s3path string, fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()
	fileInfo, _ := file.Stat()
	var size = fileInfo.Size()
	buffer := make([]byte, size)
	file.Read(buffer)
	e, s3err := s3.New(s).PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(r.Config.BucketName),
		Key:                  aws.String(fmt.Sprintf("%s%s%s", r.Config.BasePath, s3path, fileName)),
		ACL:                  aws.String(r.Config.ACL),
		Body:                 bytes.NewReader(buffer),
		ContentLength:        aws.Int64(size),
		ContentType:          aws.String(http.DetectContentType(buffer)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
	})
	if s3err != nil {
		return "", s3err
	}
	_ = e
	var s3FullPath string = fmt.Sprintf("%s%s%s", r.Config.BasePath, s3path, fileName)
	return s3FullPath, nil
}
