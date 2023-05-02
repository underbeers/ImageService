package helper

import (
	config "ImageService/configs"
	"context"
	"fmt"
	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/admin"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"io"
	"os"
)

func ImageUploadHelper(input interface{}) (string, string, error) {

	ctx := context.Background()

	//create cloudinary instance
	cld, err := cloudinary.NewFromParams(config.EnvCloudName(), config.EnvCloudAPIKey(), config.EnvCloudAPISecret())
	if err != nil {
		return "", "", err
	}

	//upload file

	uploadParam, err := cld.Upload.Upload(ctx, input, uploader.UploadParams{
		Folder: config.EnvCloudUploadFolder()})
	if err != nil {
		return "", "", err
	}

	w := os.Stdout

	//publicIdImage := uploadParam.SecureURL
	publicIdImage := uploadParam.PublicID
	fmt.Printf("%s\n", publicIdImage)

	resp, err := cld.Admin.Asset(ctx, admin.AssetParams{PublicID: publicIdImage})
	if err != nil {
		return "", "", err
	}

	io.WriteString(w, "Get and use details of the image\nDetailed response:\n")
	fmt.Printf(" %s\n", resp)

	originURL := ""

	if resp.Width > resp.Height {
		updateResp, err := cld.Upload.Explicit(ctx, uploader.ExplicitParams{
			PublicID: publicIdImage,
			Type:     "upload",
			Eager:    "c_pad,w_1920"})
		if err != nil {
			return "", "", err
		} else {
			// Log the new tag to the console.
			io.WriteString(w, "New tag: ")
			fmt.Printf("%s\n", updateResp.URL)
			originURL = updateResp.URL
		}
	} else {
		updateResp, err := cld.Upload.Explicit(ctx, uploader.ExplicitParams{
			PublicID: publicIdImage,
			Type:     "upload",
			Eager:    "c_pad,h_1920"})
		if err != nil {
			return "", "", err
		} else {
			// Log the new tag to the console.
			io.WriteString(w, "New tag 2: ")
			fmt.Printf("%s\n", updateResp.Tags)
			originURL = updateResp.URL
		}
	}

	// Instantiate an object for the asset with public ID "my_image"
	qsImg, err := cld.Image(publicIdImage)
	if err != nil {
		return "", "", err
	}

	// Add the transformation
	qsImg.Transformation = "c_pad,h_200"

	// Generate and log the delivery URL
	thumbnailURL, err := qsImg.String()
	if err != nil {
		return "", "", err
	} else {

		io.WriteString(w, "Transform the image\nTransfrmation URL: ")
		fmt.Printf("%s\n", thumbnailURL)

	}

	return originURL, thumbnailURL, nil
}

func ImageUploadHelperUser(input interface{}) (string, error) {

	ctx := context.Background()

	//create cloudinary instance
	cld, err := cloudinary.NewFromParams(config.EnvCloudName(), config.EnvCloudAPIKey(), config.EnvCloudAPISecret())
	if err != nil {
		return "", err
	}

	//upload file

	uploadParam, err := cld.Upload.Upload(ctx, input, uploader.UploadParams{
		Folder:         config.EnvCloudUploadFolder(),
		Transformation: "c_fill,g_faces,h_200,w_200"})
	if err != nil {
		return "", err
	}

	return uploadParam.URL, nil

}
