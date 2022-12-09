package helper

import (
	_util "backend/utils"
	"context"
	"fmt"
	"log"
	"mime/multipart"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func CloudinaryUpload(ctx context.Context, source multipart.File, userId string) (string, error) {
	cloudinaryCloud := _util.GetConfig("CLOUDINARY_CLOUD")
	cloudinaryKey := _util.GetConfig("CLOUDINARY_KEY")
	cloudinarySecret := _util.GetConfig("CLOUDINARY_SECRET")

	cld, _ := cloudinary.NewFromParams(cloudinaryCloud, cloudinaryKey, cloudinarySecret)

	// Upload image and set the PublicID to userId.
	resp, err := cld.Upload.Upload(
		ctx,
		source,
		uploader.UploadParams{
			PublicID: fmt.Sprintf("user-%s", userId),
			Format:   "jpg",
			Folder:   "office-booking-profile-photo-user",
		},
	)

	url := resp.SecureURL

	return url, err
}

func CloudinaryUploadOfficeImgs(files []*multipart.FileHeader, officeName string) ([]string, error) {
	ctx := context.Background()
	cloudinaryCloud := _util.GetConfig("CLOUDINARY_CLOUD")
	cloudinaryKey := _util.GetConfig("CLOUDINARY_KEY")
	cloudinarySecret := _util.GetConfig("CLOUDINARY_SECRET")

	cld, _ := cloudinary.NewFromParams(cloudinaryCloud, cloudinaryKey, cloudinarySecret)
	
	var imageURLs []string
	var err error

	for i := len(files) - 1; i >= 0; i-- {
		src, err := files[i].Open()
		
		if err != nil {
			log.Println(err)
			return imageURLs, err
		}

		fileName := fmt.Sprintf("office-%s-gambar-ke-%d", officeName, i)

		// upload image and set the PublicID to fileName.
		resp, err := cld.Upload.Upload(
			ctx,
			src,
			uploader.UploadParams{
				PublicID: fileName,
				Format:   "jpg",
				Folder:   "better-space/office-images-test",
			},
		)

		if err != nil {
			log.Println(err)
			return imageURLs, err
		}

		url := resp.SecureURL

		imageURLs = append(imageURLs, url)

		defer src.Close()
	}
	
	return imageURLs, err
}