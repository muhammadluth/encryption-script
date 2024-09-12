package usecase

import (
	"encryption-script/model"
	"encryption-script/src"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

type UploadSecretKeyUsecase struct {
}

func NewUploadSecretKeyUsecase() src.IUploadSecretKeyUsecase {
	return &UploadSecretKeyUsecase{}
}

func (u *UploadSecretKeyUsecase) DoUploadSecretKey(traceId string, secretKeyFile *multipart.FileHeader, message model.Message) model.Response {
	if secretKeyFile == nil {
		return model.FResponseDefault(http.StatusBadRequest, "Invalid Request")
	}

	sourceFile, err := secretKeyFile.Open()
	if err != nil {
		return model.FResponseDefault(http.StatusInternalServerError, "Server Error")
	}
	defer sourceFile.Close()

	destFile, err := os.Create(secretKeyFile.Filename)
	if err != nil {
		return model.FResponseDefault(http.StatusInternalServerError, "Server Error")
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return model.FResponseDefault(http.StatusInternalServerError, "Server Error")
	}

	data := model.RequestUploadSecretKeyFile{
		SecretKeyFilename: secretKeyFile.Filename,
	}
	return model.FResponseData(http.StatusOK, "Successfully Upload Secret Key", data)
}
