package minioobject

import (
	"bytes"
	"context"
	"fmt"
	"log"

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

func (repo *banksObjectRepository) GetBanksLogo(objectName string) ([]byte, string, error) {

	var err error
	var object *minio.Object

	object, err = repo.minio.GetObject(context.Background(), "bank-icon", objectName, minio.GetObjectOptions{})

	if err != nil {
		log.Panicln(err)
		return nil, "", fmt.Errorf("Service Unavailable")
	}

	var objectInfo minio.ObjectInfo

	objectInfo, err = object.Stat()

	if err != nil {
		log.Panicln(err)
		return nil, "", fmt.Errorf("Service Unavailable")
	}

	var objectBuffer *bytes.Buffer = new(bytes.Buffer)

	_, err = objectBuffer.ReadFrom(object)

	if err != nil {
		log.Panicln(err)
		return nil, "", fmt.Errorf("Service Unavailable")
	}

	return objectBuffer.Bytes(), objectInfo.ContentType, nil

}
