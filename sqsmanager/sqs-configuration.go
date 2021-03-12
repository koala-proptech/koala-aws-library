package sqsmanager

type SqsConfiguration struct {
	KeyId      string `json:"key_id"`
	SecretKey  string `json:"secret_key"`
	Region     string `json:"region"`
}