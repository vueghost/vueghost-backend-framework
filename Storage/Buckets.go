/*
 * Vueghost Tech, FZE © 2018-2019. Unauthorized copying of this file,
 * via any medium is strictly prohibited. Proprietary and confidential,
 * cannot be copied and/or distributed without the express permission of Vueghost.
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

func (b *Buckets) createSession() (*session.Session, bool) {
	if b.hasSession {
		return b.openSession, true
	}

	
	s, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-2"),
		Credentials: bucketCredentials,
	})
	if err != nil {
		return nil, false
	}

	b.openSession = s
	b.hasSession = true
	return s, true
}

func (b *Buckets) Use(bucketName string) {
	b.bucketName = bucketName
}

// Object Helpers
func (b *Buckets) getContentLength(source []byte) int64 {
	return int64(len(source))
}

func (b *Buckets) getContentType(fileHeader *multipart.FileHeader) string {
	return fileHeader.Header.Get("Content-Type")
}

// Multipart Uploads
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

func (b *Buckets) PutFile(multipartFile MultipartFile, args BucketPutArgs) {
	s3Session, success := b.createSession()
	if !success {
		return
	}

	defer multipartFile.File.Close()
	throwError(multipartFile.Error)

	fileSize := multipartFile.FileHeader.Size
	buffer := make([]byte, fileSize)
	_, err := multipartFile.File.Read(buffer)
	throwError(err)

	contentType := b.getContentType(multipartFile.FileHeader)
	objectKey := fmt.Sprintf("%s/%s", args.Folder, args.FileName)

	tagging := aws.String("")
	if args.Cache {
		tagging = aws.String("cache=true")
	}

	acl := aws.String(PRIVATE)
	if args.Policy != "" {
		acl = aws.String(args.Policy)
	}

	objectInput := &s3.PutObjectInput{
		Bucket:               aws.String(b.bucketName),
		Key:                  aws.String(objectKey),
		ACL:                  acl,
		Body:                 bytes.NewReader(buffer),
		ContentLength:        aws.Int64(fileSize),
		ContentType:          aws.String(contentType),
		ServerSideEncryption: aws.String("AES256"),
		Tagging:              tagging,
	}

	_, err = s3.New(s3Session).PutObject(objectInput)
	throwError(err)
}

type BucketPutFilePhotoArgs struct {
	FileName string
	Folder   string
	Cache    bool
	Policy   string
}

func (b *Buckets) PutFilePhoto(multipartFile MultipartFile, imagesSizes []*Utilities.CompressImageSize, args BucketPutFilePhotoArgs) {
	_, success := b.createSession()
	if !success {
		return
	}

	compressor := Utilities.CompressImage{
		File: Utilities.MultipartFile{
			File:       multipartFile.File,
			FileHeader: multipartFile.FileHeader,
			Error:      multipartFile.Error,
		},
	}

	photos, succ := compressor.Compress(imagesSizes)
	if !succ {
		return
	}

	contentType := b.getContentType(multipartFile.FileHeader)

	tagging := aws.String("")
	if args.Cache {
		tagging = aws.String("cache=true")
	}

	acl := aws.String(PRIVATE)
	if args.Policy != "" {
		acl = aws.String(args.Policy)
	}

	for _, photo := range photos {
		if !photo.CompressOutput.Success {
			continue
		}

		objectKey := fmt.Sprintf("%s/%s-%s.jpg", args.Folder, args.FileName, photo.Tag)
		objectInput := &s3.PutObjectInput{
			Bucket:               aws.String(b.bucketName),
			Key:                  aws.String(objectKey),
			ACL:                  acl,
			Body:                 bytes.NewReader(photo.CompressOutput.ImageBits),
			ContentLength:        aws.Int64(int64(len(photo.CompressOutput.ImageBits))),
			ContentType:          aws.String(contentType),
			ServerSideEncryption: aws.String("AES256"),
			Tagging:              tagging,
		}

		_, err := s3.New(b.openSession).PutObject(objectInput)
		throwError(err)
	}
}

// Deletion
type BucketDeleteObjectArgs struct {
	ObjectPath string
}

func (b *Buckets) DeleteObject(args BucketDeleteObjectArgs) (*s3.DeleteObjectOutput, bool) {
	_, success := b.createSession()
	if !success {
		return nil, false
	}

	output, err := b.deleteSingleObject(args.ObjectPath)
	if err != nil {
		return nil, false
	}

	return output, true
}

func (b *Buckets) DeleteObjects(objectsArgs []BucketDeleteObjectArgs) bool {
	_, success := b.createSession()
	if !success {
		return false
	}

	for _, object := range objectsArgs {
		_, err := b.deleteSingleObject(object.ObjectPath)
		throwError(err)
	}

	return true
}

func (b *Buckets) deleteSingleObject(objectKey string) (*s3.DeleteObjectOutput, error) {
	input := &s3.DeleteObjectInput{
		Bucket: aws.String(b.bucketName),
		Key:    aws.String(objectKey),
	}

	return s3.New(b.openSession).DeleteObject(input)
}

// Moving
type BucketMoveObjectArgs struct {
	Object string
	From   string
	To     string
	Cache  bool
	Policy string
}

func (b *Buckets) MoveObject(args BucketMoveObjectArgs) {
	_, success := b.createSession()
	if !success {
		return
	}

	b.moveSingleObject(args)
}

func (b *Buckets) MoveObjects(objectsArgs []BucketMoveObjectArgs) {
	_, success := b.createSession()
	if !success {
		return
	}

	for _, object := range objectsArgs {
		b.moveSingleObject(object)
	}
}

func (b *Buckets) moveSingleObject(args BucketMoveObjectArgs) {
	_, err := b.copySingleObject(args.Object, args.From, args.To, args.Cache, args.Policy)
	if err == nil {
		sourceKey := fmt.Sprintf("%s/%s", args.From, args.Object)
		_, err = b.deleteSingleObject(sourceKey)
	}
}

// Copying
type BucketCopyObjectArgs struct {
	Object string
	From   string
	To     string
	Cache  bool
	Policy string
}

func (b *Buckets) CopyObject(args BucketCopyObjectArgs) (*s3.CopyObjectOutput, bool) {
	_, success := b.createSession()
	if !success {
		return nil, false
	}

	output, err := b.copySingleObject(args.Object, args.From, args.To, args.Cache, args.Policy)
	if err != nil {
		return nil, false
	}

	return output, true
}

func (b *Buckets) CopyObjects(objectsArgs []BucketCopyObjectArgs) bool {
	_, success := b.createSession()
	if !success {
		return false
	}

	for _, object := range objectsArgs {
		_, err := b.copySingleObject(object.Object, object.From, object.To, object.Cache, object.Policy)
		throwError(err)
	}

	return true
}

func (b *Buckets) copySingleObject(object, from, to string, cache bool, policy string) (*s3.CopyObjectOutput, error) {
	copySource := fmt.Sprintf("%s/%s/%s", b.bucketName, from, object)
	destinationKey := fmt.Sprintf("%s/%s", to, object)

	copyInput := &s3.CopyObjectInput{
		Bucket:     aws.String(b.bucketName),
		CopySource: aws.String(copySource),
		Key:        aws.String(destinationKey),
	}

	if cache {
		copyInput.SetTagging("cache=true")
	} else {
		copyInput.SetTagging("cache=false")
	}

	if policy != "" {
		copyInput.ACL = aws.String(policy)
	}

	return s3.New(b.openSession).CopyObject(copyInput)
}

func throwError(err error) {
	if err != nil {
		panic(fmt.Sprintf("Bucket Error: %s", err))
	}
}
