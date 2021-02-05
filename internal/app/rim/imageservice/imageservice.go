package imageservice

import (
	"context"
	"image"
	"io"
	"rim-server/internal/app/rim/s3"

	"github.com/disintegration/imaging"
	"github.com/minio/minio-go/v7"
)

func Start() {
	for n := range s3.Client.ListenBucketNotification(context.Background(), "test-img", "", "", []string{"s3:ObjectCreated:*"}) {
		if n.Err != nil {
			panic(n.Err)
		}
		fileName := n.Records[0].S3.Object.Key
		fetchImage(fileName)
	}
}

func fetchImage(objetKey string) {
	object, err := s3.Client.GetObject(context.Background(), "test-img", objetKey, minio.GetObjectOptions{})
	img, err := imaging.Decode(object)
	if err != nil {
		panic(err)
	}
	r, w := io.Pipe()
	go cropImage(img, w)
	saveImage(objetKey, r)
}

func cropImage(img image.Image, destWriter *io.PipeWriter) {
	result := imaging.Resize(img, 128, 128, imaging.Lanczos)
	imaging.Encode(destWriter, result, imaging.JPEG)
	destWriter.Close()
}

func saveImage(name string, r io.Reader) {
	s3.Client.PutObject(context.Background(), "test-img-thumbnail", name, r, -1, minio.PutObjectOptions{})
}
