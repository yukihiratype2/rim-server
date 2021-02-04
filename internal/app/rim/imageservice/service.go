package imageservice

import (
	"github.com/minio/minio-go/v7"
)

type ImageService struct {
	s3 *minio.Client
}

func NewImageService(s3 *minio.Client) *ImageService {
	return &ImageService{s3}
}

func (is *ImageService) Start() {
	// for n := range is.s3.ListenBucketNotification(context.Background(), "test-img", "", "", []string{"s3:ObjectCreated:*"}) {
	// 	if n.Err != nil {
	// 		panic(n.Err)
	// 	}
	// 	var event notification.Event
	// 	event = n.Records[0]
	// 	key := event.S3.Object.Key
	// }
}

func (is *ImageService) cropImage() {
}
