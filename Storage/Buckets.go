/*
 * Vueghost Tech, FZE © 2018- 2019. Unauthorized copying of this file,
 * via any medium is strictly prohibited proprietary and confidential,
 * can not be copied and/or distributed without the express permission of Vueghost
 */

package Storage

import (
	"bytes"
	"fmt"
	"mime/multipart"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	"VGCoreSystem/Utilities"
)

const (
	PUBLIC_READ       = "public-read"
	PUBLIC_READ_WRITE = "public-read-write"
	PRIVATE           = "private"
)

type Buckets struct {
	openSession *session.Session
	bucketName  string
	hasSession  bool
}

func (b *Buckets) createSession() (openSession *session.Session, success bool) {
	if b.hasSession {
		return b.openSession, true
	} else {
		b.hasSession = true
	}
	bucketCredentials := credentials.NewStaticCredentials("AKIAIGABQMVPQR6HCN3A", "wI6/ESMfQn4LWIs08TQeQt+Kzh5M6ESM/8OF30mt", "")
	s, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-2"),
		Credentials: bucketCredentials,
	})
	if err != nil {
		b.hasSession = false
		return nil, false
	}
	throwError(err)
	b.openSession = s
	return s, true
}
func (b *Buckets) Use(BucketName string) {
	b.bucketName = BucketName
}

// Objects Resolvers
func (b Buckets) getContentLength(source []byte) int64 {
	return int64(len(source))
}
func (b Buckets) getContentType(fileHeader *multipart.FileHeader) string {
	return fileHeader.Header.Get("Content-Type")
}

// Functions: Put objects in to buckets
type BucketPutArgs struct {
	FileName string
	Folder   string
	Cache    bool
	Policy   string
}
type MultipartFile struct {
	File       multipart.File
	FileHeader *multipart.FileHeader
	Error      error
}

func (b Buckets) PutFile(multipartFile MultipartFile, args BucketPutArgs) {
	s3Session, success := b.createSession()

	file := multipartFile.File
	fileHeader := multipartFile.FileHeader
	err := multipartFile.Error
	defer file.Close()
	throwError(err)

	fileSize := fileHeader.Size
	buffer := make([]byte, fileSize)
	_, err = file.Read(buffer)
	throwError(err)

	contentLength := int64(fileSize)
	contentType := b.getContentType(fileHeader)
	contentBody := bytes.NewReader(buffer)
	objectKey := fmt.Sprintf("%s/%s", args.Folder, args.FileName)

	var Tagging *string
	if args.Cache {
		Tagging = aws.String("cache=true")
	}

	var ACL *string
	if args.Policy != "" {
		ACL = aws.String(args.Policy)
	} else {
		ACL = aws.String("private")
	}

	if success {
		objectInputs := &s3.PutObjectInput{
			Bucket:               aws.String(b.bucketName),
			Key:                  aws.String(objectKey),
			ACL:                  ACL,
			Body:                 contentBody,
			ContentLength:        aws.Int64(contentLength),
			ContentType:          aws.String(contentType),
			ServerSideEncryption: aws.String("AES256"),
			Tagging:              Tagging,
		}
		_, err := s3.New(s3Session).PutObject(objectInputs)
		throwError(err)
	}
}

type BucketPutFilePhotoArgs struct {
	FileName string
	Folder   string
	Cache    bool
	Policy   string
}

func (b Buckets) PutFilePhoto(multipartFile MultipartFile, imagesSizes []*Utilities.CompressImageSize, args BucketPutFilePhotoArgs) {
	_, success := b.createSession()
	if success {
		com := Utilities.CompressImage{
			File: Utilities.MultipartFile{
				File:       multipartFile.File,
				FileHeader: multipartFile.FileHeader,
				Error:      multipartFile.Error,
			},
		}
		var Tagging *string
		if args.Cache {
			Tagging = aws.String("cache=true")
		}
		var ACL *string
		if args.Policy != "" {
			ACL = aws.String(args.Policy)
		} else {
			ACL = aws.String("private")
		}
		photos, succ := com.Compress(imagesSizes)
		contentType := b.getContentType(multipartFile.FileHeader)

		if succ {
			for _, value := range photos {
				if value.CompressOutput.Success {
					objectKey := fmt.Sprintf("%s/%s-%s.jpg", args.Folder, args.FileName, value.Tag)
					contentBody := bytes.NewReader(value.CompressOutput.ImageBits)
					contentLength := int64(len(value.CompressOutput.ImageBits))
					objectInputs := &s3.PutObjectInput{
						Bucket:               aws.String(b.bucketName),
						Key:                  aws.String(objectKey),
						ACL:                  ACL,
						Body:                 contentBody,
						ContentLength:        aws.Int64(contentLength),
						ContentType:          aws.String(contentType),
						ServerSideEncryption: aws.String("AES256"),
						Tagging:              Tagging,
					}
					_, err := s3.New(b.openSession).PutObject(objectInputs)
					throwError(err)
				}
			}
		}
	}
}

// Functions: Delete objects in to buckets
type BucketDeleteObjectArgs struct {
	ObjectPath string
}

func (b *Buckets) DeleteObject(args BucketDeleteObjectArgs) (*s3.DeleteObjectOutput, bool) {
	_, success := b.createSession()
	if success {
		output, err := b.deleteSingleObject(args.ObjectPath)
		throwError(err)
		if err != nil {
			return nil, false
		}
		return output, true
	} else {
		return nil, false
	}
}
func (b *Buckets) DeleteObjects(objectsArgs []BucketDeleteObjectArgs) (isSuccess bool) {
	_, success := b.createSession()
	if success {
		for _, object := range objectsArgs {
			_, err := b.deleteSingleObject(object.ObjectPath)
			throwError(err)
			if err != nil {
			}
		}
		return true
	} else {
		return false
	}
}
func (b Buckets) deleteSingleObject(objectKey string) (*s3.DeleteObjectOutput, error) {
	deleteObject := &s3.DeleteObjectInput{
		Bucket: aws.String(b.bucketName),
		Key:    aws.String(objectKey),
	}
	handler := s3.New(b.openSession)
	output, err := handler.DeleteObject(deleteObject)
	return output, err
}

// Functions: Mode objects
type BucketMoveObjectArgs struct {
	Object string
	From   string
	To     string
	Cache  bool
	Policy string
}

func (b *Buckets) MoveObject(args BucketMoveObjectArgs) {
	_, success := b.createSession()
	if success {
		b.moveSingleObject(args.Object, args.From, args.To, args.Cache, args.Policy)
	}
}
func (b *Buckets) MoveObjects(objectsArgs []BucketMoveObjectArgs) {
	_, success := b.createSession()
	if success {
		for _, object := range objectsArgs {
			b.moveSingleObject(object.Object, object.From, object.To, object.Cache, object.Policy)
		}
	}
}
func (b Buckets) moveSingleObject(object string, from string, to string, cache bool, policy string) {
	_, err := b.copySingleObject(object, from, to, cache, policy)
	if err == nil {
		objectPath := fmt.Sprintf("/%s/%s", from, object)
		_, err = b.deleteSingleObject(objectPath)
	}
}

// Functions: Copy bucket objects
type BucketCopyObjectArgs struct {
	Object string
	From   string
	To     string
	Cache  bool
	Policy string
}

func (b *Buckets) CopyObject(args BucketCopyObjectArgs) (*s3.CopyObjectOutput, bool) {
	_, success := b.createSession()
	if success {
		output, err := b.copySingleObject(args.Object, args.From, args.To, args.Cache, args.Policy)
		throwError(err)
		if err != nil {
			return nil, false
		}
		return output, true
	} else {
		return nil, false
	}
}
func (b *Buckets) CopyObjects(objectsArgs []BucketCopyObjectArgs) (isSuccess bool) {
	_, success := b.createSession()
	if success {
		for _, object := range objectsArgs {
			_, err := b.copySingleObject(object.Object, object.From, object.To, object.Cache, object.Policy)
			throwError(err)
			if err != nil {
			}
		}
		return true
	} else {
		return false
	}
}
func (b Buckets) copySingleObject(object string, from string, to string, cache bool, policy string) (*s3.CopyObjectOutput, error) {
	copySourceFrom := fmt.Sprintf("/%s/%s/%s", b.bucketName, from, object)
	copySourceTo := fmt.Sprintf("/%s/%s", to, object)
	copyObject := &s3.CopyObjectInput{
		Bucket:     aws.String(b.bucketName),
		CopySource: aws.String(copySourceFrom),
		Key:        aws.String(copySourceTo),
	}

	if cache == true {
		copyObject.SetTagging("cache=true")

	} else if cache == false {
		copyObject.SetTagging("cache=false")
	}

	if policy != "" {
		copyObject.ACL = aws.String(policy)
	}

	handler := s3.New(b.openSession)
	output, err := handler.CopyObject(copyObject)
	return output, err
}

func throwError(err error) {
	if err != nil {
		panic(fmt.Sprintf("Bucket Error: %s", err.Error()))
	}
}
