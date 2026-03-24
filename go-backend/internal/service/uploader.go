package service

import (
	"context"
	"mime/multipart"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type Uploader interface {
	Upload(ctx context.Context, file multipart.File) (url string, publicId string, err error);
}


type CloudinaryUploader struct {
	Cloudinary *cloudinary.Cloudinary
}

func (u *CloudinaryUploader) Upload(ctx context.Context, file multipart.File) (string, string, error) {
	result, err := u.Cloudinary.Upload.Upload(ctx, file, uploader.UploadParams{})

	if err != nil {
		return "", "", ErrUploadingImage
	}

	return result.SecureURL, result.PublicID, nil
}

func NewCloudinaryUploader(cdl *cloudinary.Cloudinary) *CloudinaryUploader {
	return &CloudinaryUploader{
		Cloudinary: cdl,
	}
}

