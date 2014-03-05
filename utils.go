package gobotOpencv

import (
	"github.com/hybridgroup/go-opencv/opencv"
)

func loadHaarClassifierCascade(haar string) *opencv.HaarCascade {
	return opencv.LoadHaarClassifierCascade(haar)
}

func DetectFaces(haar string, image *opencv.IplImage) []*opencv.Rect {
	cascade := loadHaarClassifierCascade(haar)
	rect := cascade.DetectObjects(image)
	return rect
}

func DrawRectangles(image *opencv.IplImage, rect []*opencv.Rect) *opencv.IplImage {
	for _, value := range rect {
		opencv.Rectangle(image,
			opencv.Point{value.X() + value.Width(), value.Y()},
			opencv.Point{value.X(), value.Y() + value.Height()},
			opencv.ScalarAll(255.0), 1, 1, 0)
	}
	return image
}
