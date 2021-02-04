package imageservice

import (
	"context"
	"image"
	"io"

	"github.com/disintegration/imaging"
	"github.com/minio/minio-go/v7"
)

type ImageService struct {
	s3 *minio.Client
}

func NewImageService(s3 *minio.Client) *ImageService {
	return &ImageService{s3}
}

func (is *ImageService) Start() {
	for n := range is.s3.ListenBucketNotification(context.Background(), "test-img", "", "", []string{"s3:ObjectCreated:*"}) {
		if n.Err != nil {
			panic(n.Err)
		}
		fileName := n.Records[0].S3.Object.Key
		is.fetchImage(fileName)
	}
}

func (is *ImageService) fetchImage(objetKey string) {
	object, err := is.s3.GetObject(context.Background(), "test-img", objetKey, minio.GetObjectOptions{})
	img, err := imaging.Decode(object)
	if err != nil {
		panic(err)
	}
	r, w := io.Pipe()
	go cropImage(img, w)
	is.saveImage(objetKey, r)
}

func cropImage(img image.Image, destWriter *io.PipeWriter) {
	result := imaging.Resize(img, 128, 128, imaging.Lanczos)
	imaging.Encode(destWriter, result, imaging.JPEG)
	destWriter.Close()
}

func (is *ImageService) saveImage(name string, r io.Reader) {
	is.s3.PutObject(context.Background(), "test-img-thumbnail", name, r, -1, minio.PutObjectOptions{})
}
