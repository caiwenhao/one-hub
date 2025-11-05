package storage

import (
    "one-api/common/storage/drives"

    "github.com/spf13/viper"
    "one-api/common/config"
)

type Storage struct {
	drives map[string]StorageDrive
}

func InitStorage() {
	InitImgurStorage()
	InitSMStorage()
	InitALIOSSStorage()
	InitS3Storage()
}

func InitALIOSSStorage() {
	endpoint := viper.GetString("storage.alioss.endpoint")
	if endpoint == "" {
		return
	}
	accessKeyId := viper.GetString("storage.alioss.accessKeyId")
	if accessKeyId == "" {
		return
	}
	accessKeySecret := viper.GetString("storage.alioss.accessKeySecret")
	if accessKeySecret == "" {
		return
	}
	bucketName := viper.GetString("storage.alioss.bucketName")
	if bucketName == "" {

		return
	}

	aliUpload := drives.NewAliOSSUpload(endpoint, accessKeyId, accessKeySecret, bucketName)
	AddStorageDrive(aliUpload)
}

func InitSMStorage() {
	smSecret := viper.GetString("storage.smms.secret")
	if smSecret == "" {
		return
	}

	smUpload := drives.NewSMUpload(smSecret)
	AddStorageDrive(smUpload)
}

func InitImgurStorage() {
	imgurClientId := viper.GetString("storage.imgur.client_id")
	if imgurClientId == "" {
		return
	}

	imgurUpload := drives.NewImgurUpload(imgurClientId)
	AddStorageDrive(imgurUpload)
}

func InitS3Storage() {
    // 优先使用面板配置（config），否则回退到 viper（文件配置）
    endpoint := config.S3Endpoint
    if endpoint == "" {
        endpoint = viper.GetString("storage.s3.endpoint")
    }
    if endpoint == "" {
        return
    }

    accessKeyId := config.S3AccessKeyId
    if accessKeyId == "" {
        accessKeyId = viper.GetString("storage.s3.accessKeyId")
    }
    if accessKeyId == "" {
        return
    }

    accessKeySecret := config.S3AccessKeySecret
    if accessKeySecret == "" {
        accessKeySecret = viper.GetString("storage.s3.accessKeySecret")
    }
    if accessKeySecret == "" {
        return
    }

    bucketName := config.S3BucketName
    if bucketName == "" {
        bucketName = viper.GetString("storage.s3.bucketName")
    }
    if bucketName == "" {
        return
    }

    cdnurl := config.S3CDNURL
    if cdnurl == "" {
        cdnurl = viper.GetString("storage.s3.cdnurl")
    }
    if cdnurl == "" {
        cdnurl = endpoint
    }

    expirationDays := config.S3ExpirationDays
    if expirationDays == 0 {
        expirationDays = viper.GetInt("storage.s3.expirationDays")
    }

    s3Upload := drives.NewS3Upload(endpoint, accessKeyId, accessKeySecret, bucketName, cdnurl, expirationDays)
    AddStorageDrive(s3Upload)
}
