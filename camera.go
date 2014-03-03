package gobotOpencv

import (
	"github.com/hybridgroup/go-opencv/opencv"
	"github.com/hybridgroup/gobot"
	"time"
)

type Camera struct {
	gobot.Driver
	Adaptor *Opencv
	camera  *opencv.Capture
}

type CameraInterface interface {
}

func NewCamera(adaptor *Opencv) *Camera {
	d := new(Camera)
	d.Events = make(map[string]chan interface{})
	d.Adaptor = adaptor
	d.Commands = []string{}
	return d
}

func (me *Camera) Start() bool {
	me.Events["Frame"] = make(chan interface{}, 0)
	me.camera = opencv.NewCameraCapture(0)
	go func() {
		for {
			if me.camera.GrabFrame() {
				image := me.camera.RetrieveFrame(1)
				if image != nil {
					gobot.Publish(me.Events["Frame"], image)
				}
			}
			time.Sleep(10 * time.Millisecond)
		}
	}()
	return true
}
