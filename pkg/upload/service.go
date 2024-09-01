package upload

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"

	b64 "encoding/base64"

	"github.com/google/uuid"
	awss3 "github.com/ongyoo/roomkub-api/pkg/awsS3"
	"github.com/ongyoo/roomkub-api/pkg/crypto"
	"go.mongodb.org/mongo-driver/bson/primitive"

	// Crypto
	// JWT
	jwt "github.com/ongyoo/roomkub-api/pkg/middleware"
)

type Service interface {
	Upload(ctx context.Context, key string, fileHeader multipart.FileHeader) (*UploadResponse, error)
	GetObject(ctx context.Context, key, fileName string) ([]byte, string, error)
}

type service struct {
	awsS3 awss3.Storage
}

func NewService(awsS3 awss3.Storage) *service {
	return &service{awsS3}
}

func (s service) Upload(ctx context.Context, key string, fileHeader multipart.FileHeader) (*UploadResponse, error) {
	businessID := primitive.NewObjectID()
	userClaims, _, err := jwt.GetContextUserClaims(ctx)
	if err != nil {
		return nil, errors.New(err.Error() + " [jwt_error] กรุณาติดต่อผู้ดูแลระบบ")
	}

	if userClaims.Payload.BusinessID != "" {
		objID, err := primitive.ObjectIDFromHex(userClaims.Payload.BusinessID)
		if err != nil {
			return nil, errors.New(err.Error() + "กรุณาติดต่อผู้ดูแลระบบ")
		}
		businessID = objID
	}
	file, err := fileHeader.Open()
	defer file.Close()
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(nil)
	_, err = io.Copy(buf, file)
	if err != nil {
		return nil, err
	}

	documentId := uuid.New().String()
	contentType := fileHeader.Header.Get("content-type")
	filename := fileHeader.Filename
	meta := map[string]string{
		"content-type": contentType,
		"filename":     filename,
		"documentId":   documentId,
	}

	s3Key := fmt.Sprintf("key/%s/%s", businessID.Hex(), documentId)
	encryptedBody, err := crypto.EncryptAes256(buf.String())
	if err != nil {
		return nil, err
	}

	location, err := s.awsS3.PutObject(ctx, s3Key, strings.NewReader(encryptedBody), meta)
	if err != nil {
		return nil, err
	}
	id := *location
	hostURL := os.Getenv("HOST_URL")
	publicUrl := hostURL + "/api/v1/upload/public/"
	privateUrl := hostURL + "/api/v1/upload/private/" + id

	encodeID, err := crypto.EncryptAes256(*location)
	if err == nil {
		sEnc := b64.StdEncoding.EncodeToString([]byte(encodeID))
		publicUrl = publicUrl + sEnc
	}

	res := UploadResponse{
		PublicID:   encodeID,
		PrivateID:  id,
		PublicUrl:  publicUrl,
		PrivateUrl: privateUrl,
	}

	return &res, nil
}

func (s service) GetObject(ctx context.Context, key, fileName string) ([]byte, string, error) {
	output, err := s.awsS3.GetObject(ctx, key, fileName)
	if err != nil {
		return nil, "", err
	}

	buf := bytes.NewBuffer(nil)
	_, err = buf.ReadFrom(output)
	if err != nil {
		return nil, "", err
	}
	decryptedContent, err := crypto.DecryptAes256(buf.String())
	if err != nil {
		return nil, "", err
	}

	bufArr := []byte(decryptedContent)
	contentType := http.DetectContentType(bufArr)
	return bufArr, contentType, nil
}
