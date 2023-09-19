package main

import (
	_ "embed"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

const (
	defaultImgWidth     = 40
	defaultImgHeight    = 40
	defaultImagePadding = 5
)

//go:embed switchOn.png
var switchOnBytes []byte

//go:embed switchOff.png
var switchOffBytes []byte

var onResource = fyne.NewStaticResource("on", switchOnBytes)
var offResource = fyne.NewStaticResource("off", switchOffBytes)

type SwitchOpt = func(s *Switch) error

type Switch struct {
	widget.DisableableWidget

	Checked bool

	OnChanged func(bool)

	imgWidth, imgHeight float32
	padding             float32
}

func (s *Switch) Tapped(event *fyne.PointEvent) {
	s.Checked = !s.Checked
	s.Refresh()
	s.OnChanged(s.Checked)
}

func (s *Switch) CreateRenderer() fyne.WidgetRenderer {
	img := canvas.NewImageFromResource(fyne.NewStaticResource("off", switchOffBytes))
	img.Resize(fyne.NewSize(s.imgWidth, s.imgHeight))
	img.ScaleMode = canvas.ImageScaleSmooth
	img.FillMode = canvas.ImageFillOriginal
	return &switchRenderer{s: s, img: img}
}

func NewSwitch(changed func(bool), opts ...SwitchOpt) *Switch {
	s := &Switch{
		DisableableWidget: widget.DisableableWidget{},
		OnChanged:         changed,
		imgWidth:          defaultImgWidth,
		imgHeight:         defaultImgHeight,
		padding:           defaultImagePadding,
	}

	for _, opt := range opts {
		opt(s)
	}

	s.ExtendBaseWidget(s)
	s.Refresh()
	return s
}

type switchRenderer struct {
	s *Switch

	img *canvas.Image
}

func (s *switchRenderer) Destroy() {

}

func (s *switchRenderer) Layout(size fyne.Size) {
	imgSize := fyne.NewSize(s.s.imgWidth, s.s.imgHeight)
	pos := fyne.NewPos(size.Width/2-imgSize.Width/2, size.Height/2-imgSize.Height/2)

	s.img.Move(pos)
	s.img.Resize(imgSize)
}

func (s *switchRenderer) MinSize() fyne.Size {
	return fyne.NewSize(s.s.imgWidth+s.s.padding, s.s.imgHeight+s.s.padding)
}

func (s *switchRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{s.img}
}

func (s *switchRenderer) Refresh() {
	if s.s.Checked {
		s.img.Resource = onResource
	} else {
		s.img.Resource = offResource
	}
	s.img.Refresh()
}

func SetImageWidth(width float32) SwitchOpt {
	return func(s *Switch) error {
		s.imgWidth = width
		return nil
	}
}

func SetImageHeight(height float32) SwitchOpt {
	return func(s *Switch) error {
		s.imgHeight = height
		return nil
	}
}

func SetImagePadding(padding float32) SwitchOpt {
	return func(s *Switch) error {
		s.padding = padding
		return nil
	}
}
