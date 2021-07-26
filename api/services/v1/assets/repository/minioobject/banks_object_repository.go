package minioobject

import (
	"bytes"
	"context"
	"fmt"
	"log"

	"auctionkuy.wildangbudhi.com/domain"
	"auctionkuy.wildangbudhi.com/domain/v1/assets"
	"github.com/minio/minio-go/v7"
)

type banksObjectRepository struct {
	minio *minio.Client
}

func NewAppSettingsRepository(minio *minio.Client) assets.BanksObjectRepository {
	return &banksObjectRepository{
		minio: minio,
	}
}

func (repo *banksObjectRepository) GetBanksLogo(objectName string) ([]byte, string, error, domain.RepositoryErrorType) {

	var err error
	var object *minio.Object

	object, err = repo.minio.GetObject(context.Background(), "bank-icon", objectName, minio.GetObjectOptions{})

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
		} else {
			log.Println(errResponse.Code)
			return nil, "", fmt.Errorf("Service Unavailable"), domain.RepositoryError
		}

	}

	var objectBuffer *bytes.Buffer = new(bytes.Buffer)

	_, err = objectBuffer.ReadFrom(object)

	if err != nil {
		log.Println(err)
		return nil, "", fmt.Errorf("Service Unavailable"), domain.RepositoryError
	}

	return objectBuffer.Bytes(), objectInfo.ContentType, nil, 0

}
