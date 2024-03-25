package model

type RequestHMACEncryption struct {
	SecretKey string `json:"secret_key"`
	Message   string `json:"message"`
}

type RequestRSAEncryption struct {
	PrivateKey     string `json:"private_key"`
	PublicKey      string `json:"public_key"`
	OutputEncoding string `json:"output_encoding"`
	Message        string `json:"message"`
}
