package minioobject

import (
	"bytes"
	"context"
	"fmt"
	"log"

	"auctionkuy.wildangbudhi.com/domain"
	"auctionkuy.wildangbudhi.com/domain/v1/account"
	"github.com/minio/minio-go/v7"
)

type userObjectRepository struct {
	minio      *minio.Client
	bucketName string
}

func NewUserObjectRepository(minio *minio.Client) account.UserObjectRepository {
	return &userObjectRepository{
		minio:      minio,
		bucketName: "user-profile",
	}
}

func (repo *userObjectRepository) GetUserAvatar(userID *domain.UUID) ([]byte, string, error, domain.RepositoryErrorType) {

	var err error
	var object *minio.Object

	object, err = repo.minio.GetObject(context.Background(), repo.bucketName, userID.GetValue(), minio.GetObjectOptions{})

	if err != nil {
		log.Println(err)
		return nil, "", fmt.Errorf("Service Unavailable"), domain.RepositoryError
	}

	var objectInfo minio.ObjectInfo

	objectInfo, err = object.Stat()

	if err != nil {

		var errResponse minio.ErrorResponse = minio.ToErrorResponse(err)

		if errResponse.Code == "NoSuchKey" {
			return nil, "", fmt.Errorf("Data not found"), domain.RepositoryDataNotFound
		}

		log.Println(errResponse.Code)
		return nil, "", fmt.Errorf("Service Unavailable"), domain.RepositoryError

	}

	var objectBuffer *bytes.Buffer = new(bytes.Buffer)

	_, err = objectBuffer.ReadFrom(object)

	if err != nil {
		log.Println(err)
		return nil, "", fmt.Errorf("Service Unavailable"), domain.RepositoryError
	}

	return objectBuffer.Bytes(), objectInfo.ContentType, nil, 0

}

func (repo *userObjectRepository) PutUserAvatar(userID *domain.UUID, data []byte, contentType string) (error, domain.RepositoryErrorType) {

	var err error
	var fileReader *bytes.Reader = bytes.NewReader(data)

	if fileReader.Size() <= 0 {
		log.Println("User Avatar File Invalid with Bytes Length of 0")
		return fmt.Errorf("Service Unavailable"), domain.RepositoryError
	}

	_, err = repo.minio.PutObject(context.Background(), repo.bucketName, userID.GetValue(), fileReader, fileReader.Size(), minio.PutObjectOptions{ContentType: contentType})

	if err != nil {
		log.Println(err)
		return fmt.Errorf("Service Unavailable"), domain.RepositoryError
	}

	return nil, 0

}

func (repo *userObjectRepository) RemoveUserAvatar(userID *domain.UUID) (error, domain.RepositoryErrorType) {

	var err error

	err = repo.minio.RemoveObject(context.Background(), repo.bucketName, userID.GetValue(), minio.RemoveObjectOptions{})

	if err != nil {
		log.Println(err)
		return fmt.Errorf("Service Unavailable"), domain.RepositoryError
	}

	return nil, 0

}
