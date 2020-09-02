package storage

import (
  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/credentials"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/s3/s3manager"
  "mime/multipart"
  "os"
  "strconv"
  "time"
)

func UploadFileToS3(file multipart.File, fileKey string) bool {

  sess, errWhileCreatingSession := session.NewSession(&aws.Config{
    Region: aws.String("us-east-2"),
    Credentials: credentials.NewStaticCredentials(
      os.Getenv("ACCOUNTABILITY_AWS_ID"),
      os.Getenv("ACCOUNTABILITY_AWS_SECRET"),
      os.Getenv("ACCOUNTABILITY_AWS_TOKEN")),
  })
  if errWhileCreatingSession != nil {
    return false
  }
  svc := s3manager.NewUploader(sess)
  input := &s3manager.UploadInput{
    Bucket: aws.String("accountability-user-data"),
    Key:    aws.String("/task-updates-media/" + fileKey),
    Body:   file,
  }
  _, errWhileUploading := svc.Upload(input)
  if errWhileUploading != nil {
    return false
  }
  return true
}


func GetUniqueS3Key(fileName string) string {
  t := time.Now()
  return strconv.FormatInt(t.UnixNano(), 10) + "_" + fileName
}
