package event

import (
	"context"
	"rim-server/internal/app/rim/model"

	r "github.com/go-redis/redis/v8"
)

var redis *r.Client

type ImageProcessStatus struct {
	Image model.Image
}

func Start() {
	redis = r.NewClient(&r.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func setStatus(fileID string, status string) (err error) {
	res := redis.XAdd(context.Background(), &r.XAddArgs{
		Stream: fileID,
		Values: []interface{}{
			"s", status,
		},
	})
	return res.Err()
}

func (ips *ImageProcessStatus) CurrentStatus() string {
	res := redis.XRead(context.Background(), &r.XReadArgs{
		Streams: []string{ips.Image.FileID, "0"},
		Count:   0,
	})
	r, err := res.Result()
	if err != nil {
		panic(err)
	}
	msg := r[0].Messages
	return msg[len(msg)-1].Values["s"].(string)
}

//StartProcess func
func (ips *ImageProcessStatus) StartProcess() {
	setStatus(ips.Image.FileID, "started")
}

func (ips *ImageProcessStatus) ImageCreated() {
	setStatus(ips.Image.FileID, "created")
}

func (ips *ImageProcessStatus) CompleteProcess() {
	setStatus(ips.Image.FileID, "completed")
}

func (ips *ImageProcessStatus) WaitForImageProcessed() {
	for status := range ips.StatusNotifaction() {
		if status == "completed" {
			return
		}
	}
}

func (ips *ImageProcessStatus) StatusNotifaction() (c chan string) {
	c = make(chan string)
	go func() {
		c <- ips.CurrentStatus()
	}()
	go loop(ips.Image.FileID, "$", c)
	return c
}

func loop(fileID string, timeStamp string, c chan string) {
	res := redis.XRead(context.Background(), &r.XReadArgs{
		Streams: []string{fileID, timeStamp},
		// Count:   0,
		Block: 0,
	})
	r, err := res.Result()
	if err != nil {
		panic(err)
	}
	s := r[0].Messages[0].Values["s"].(string)
	c <- s
	if s == "completed" {
		close(c)
		return
	}
	loop(fileID, r[0].Messages[0].ID, c)
}
