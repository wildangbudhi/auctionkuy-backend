package usecase

import (
	"fmt"
	"log"

	"auctionkuy.wildangbudhi.com/domain"
	"github.com/minio/minio-go/v7"
)

func (usecase *assetsUsecase) BankObject(objectName string) (*minio.Object, *minio.ObjectInfo, error, domain.HTTPStatusCode) {

	var err error
	var object *minio.Object

	if objectName == "" {
		return nil, nil, fmt.Errorf("Object name invalid"), 400
	}

	objectName += ".png"

	log.Println(objectName)

	object, err = usecase.banksObjectRepository.GetBanksLogo(objectName)

	if err != nil {
		return nil, nil, err, 500
	}

	var stats minio.ObjectInfo

	stats, err = object.Stat()

	if err != nil {
		log.Println(err)
		return nil, nil, fmt.Errorf("Failed to process object"), 500
	}

	return object, &stats, nil, 200

}
