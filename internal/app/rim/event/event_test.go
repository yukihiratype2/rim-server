package event_test

import (
	"rim-server/internal/app/rim/event"
	"rim-server/internal/app/rim/model"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestStart(t *testing.T) {
	event.Start()
}

func TestReceiveStream(t *testing.T) {
	event.Start()
	imageStatus := event.ImageProcessStatus{Image: model.Image{FileID: uuid.New().String() + ".jpg"}}
	imageStatus.ImageCreated()
	go func() {
		time.Sleep(2 * time.Second)
		imageStatus.StartProcess()
	}()
	for status := range imageStatus.StatusNotifaction() {
		t.Log(status)
		break
	}
}

func TestCompleteEvent(t *testing.T) {
	event.Start()
	imageStatus := event.ImageProcessStatus{Image: model.Image{FileID: uuid.New().String() + ".jpg"}}
	imageStatus.ImageCreated()
	imageStatus.CompleteProcess()
	time.Sleep(time.Second)
	// for range imageStatus.StatusNotifaction() {
	// t.Log(status)
	// if status != "completed" {
	// 	t.Error()
	// 	return
	// }
	// 	return
	// }
}

func TestReadLastStatus(t *testing.T) {
	event.Start()
	imageStatus := event.ImageProcessStatus{Image: model.Image{FileID: uuid.New().String() + ".jpg"}}
	imageStatus.ImageCreated()
	// imageStatus.CompleteProcess()
	// time.Sleep(time.Second)
	status := imageStatus.CurrentStatus()
	if status != "created" {
		t.Log(status)
		t.Error()
	}
}
