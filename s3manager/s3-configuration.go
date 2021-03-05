package s3manager

type S3Configuration struct {
	KeyId      string `json:"key_id"`
	SecretKey  string `json:"secret_key"`
	BucketName string `json:"bucket_name"`
	BasePath   string `json:"base_path"`
	ACL        string `json:"acl"`
	Region     string `json:"region"`
}
