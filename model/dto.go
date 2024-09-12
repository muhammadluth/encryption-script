package model

type RequestRSAEncryption struct {
	PrivateKey        string `json:"private_key"`
	PublicKey         string `json:"public_key"`
	SecretKeyPassword string `json:"secret_key_password"`
	SecretKeyFilename string `json:"secret_key_filename"`
	OutputEncoding    string `json:"output_encoding"`
	Message           string `json:"message"`
}

type RequestHMACEncryptionAES struct {
	Type           string `json:"type"`
	SecretKey      string `json:"secret_key"`
	OutputEncoding string `json:"output_encoding"`
	Message        string `json:"message"`
}

type RequestHMACDecryptionAES struct {
	Type          string `json:"type"`
	SecretKey     string `json:"secret_key"`
	InputEncoding string `json:"input_encoding"`
	Message       string `json:"message"`
}

type RequestUploadSecretKeyFile struct {
	SecretKeyFilename string `json:"secret_key_file"`
}
