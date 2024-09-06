package upload

import (
	"bytes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/empnefsi/mop-service/internal/config"
	"io"
	"mime/multipart"
	"path/filepath"
	"strings"
)

func File(file *multipart.FileHeader, filename, path string) (*string, error) {
	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	buffer := new(bytes.Buffer)
	_, err = buffer.ReadFrom(src)
	if err != nil && err != io.EOF {
		return nil, err
	}

	fileExtension := strings.ToLower(filepath.Ext(file.Filename))
	filename = filename + fileExtension
	key := filepath.Join(path, filename)
	spaceName := config.GetSpaceName()

	_, err = config.GetS3().PutObject(&s3.PutObjectInput{
		Bucket: aws.String(spaceName),
		Key:    aws.String(key),
		Body:   bytes.NewReader(buffer.Bytes()),
		ACL:    aws.String("public-read"),
	})
	if err != nil {
		return nil, err
	}

	uploadedFilePath := fmt.Sprintf("%s/%s", spaceName, key)
	return &uploadedFilePath, nil
}
