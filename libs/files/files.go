package files

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/spf13/viper"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func addFileToS3(fileDir string) (string, error) {
	s3REGION := viper.GetString("s3_region")
	s3BUCKET := viper.GetString("s3_bucket")
	s3URL := viper.GetString("s3_url")

	//new s3 session
	s, err := session.NewSession(&aws.Config{Region: aws.String(s3REGION)})

	// Open the file for use
	file, err := os.OpenFile(fileDir, os.O_RDONLY, 0444)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Get file size and read the file content into a buffer
	fileInfo, _ := file.Stat()
	var size = fileInfo.Size()
	buffer := make([]byte, size)
	file.Read(buffer)

	// Config settings: this is where you choose the bucket, filename, content-type etc.
	// of the file you're uploading.
	output, err := s3.New(s).PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(s3BUCKET),
		Key:                  aws.String(fileDir),
		ACL:                  aws.String("public-read"),
		Body:                 bytes.NewReader(buffer),
		ContentLength:        aws.Int64(size),
		ContentType:          aws.String(http.DetectContentType(buffer)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
	})
	fmt.Println(output)
	return s3URL, nil

}

//UploadFile uploads a file to the server under uploads dir and calls the addFileToS3 function
func UploadFile(key string, req *http.Request) string {

	file, handler, err := req.FormFile(key)
	if err != nil {
		fmt.Println(err)
	}

	f, err := os.OpenFile("uploads/"+handler.Filename, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Println(err)

	}

	data, _ := ioutil.ReadAll(file)
	f.Write(data)

	location, errr := addFileToS3(f.Name())

	if err != nil {
		fmt.Println(errr)
	}

	fileLocation := fmt.Sprintf("%s/%s", location, f.Name())

	return fileLocation

}
