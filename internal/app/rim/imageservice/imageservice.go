package imageservice

import (
	"context"
	"fmt"
	"image"
	"io"
	"rim-server/internal/app/rim/event"
	"rim-server/internal/app/rim/model"
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
	imageStatus := event.ImageProcessStatus{Image: model.Image{FileID: objetKey}}
	imageStatus.StartProcess()
	object, err := s3.Client.GetObject(context.Background(), "test-img", objetKey, minio.GetObjectOptions{})
	img, err := imaging.Decode(object)
	if err != nil {
		panic(err)
	}
	r, w := io.Pipe()
	go cropImage(img, w)
	err = saveImage(objetKey, r)
	if err != nil {
		panic(err)
	}
	res, err := s3.Client.StatObject(context.Background(), "test-img-thumbnail", objetKey, minio.StatObjectOptions{})
	fmt.Printf("%+v\n", res.Size)
	imageStatus.CompleteProcess()
}

func cropImage(img image.Image, destWriter *io.PipeWriter) {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	size := min(img.Bounds().Max.X, img.Bounds().Max.Y)
	croped := imaging.CropCenter(img, size, size)
	result := imaging.Resize(croped, 400, 400, imaging.Lanczos)
	imaging.Encode(destWriter, result, imaging.JPEG)
	destWriter.Close()
}

func saveImage(name string, r io.Reader) (err error) {
	_, err = s3.Client.PutObject(context.Background(), "test-img-thumbnail", name, r, -1, minio.PutObjectOptions{})
	return err
}
