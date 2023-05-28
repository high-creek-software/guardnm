package main

import (
	_ "embed"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

//go:embed switchOn.png
var switchOnBytes []byte

//go:embed switchOff.png
var switchOffBytes []byte

type Switch struct {
	widget.DisableableWidget

	Checked bool

	OnChanged func(bool)
}

func (s *Switch) Tapped(event *fyne.PointEvent) {
	s.Checked = !s.Checked
	s.Refresh()
	s.OnChanged(s.Checked)
}

func (s *Switch) CreateRenderer() fyne.WidgetRenderer {
	img := canvas.NewImageFromResource(fyne.NewStaticResource("off", switchOffBytes))
	img.Resize(fyne.NewSize(40, 20))
	img.ScaleMode = canvas.ImageScaleSmooth
	img.FillMode = canvas.ImageFillOriginal
	return &switchRenderer{s: s, img: img}
}

func NewSwitch(changed func(bool)) *Switch {
	s := &Switch{
		DisableableWidget: widget.DisableableWidget{},
		OnChanged:         changed,
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
	imgSize := s.img.MinSize()
	pos := fyne.NewPos(theme.Padding(), theme.Padding())

	s.img.Move(pos)
	s.img.Resize(imgSize)
}

func (s *switchRenderer) MinSize() fyne.Size {
	return fyne.NewSize(40, 20)
}

func (s *switchRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{s.img}
}

func (s *switchRenderer) Refresh() {
	if s.s.Checked {
		s.img.Resource = fyne.NewStaticResource("on", switchOnBytes)
	} else {
		s.img.Resource = fyne.NewStaticResource("off", switchOffBytes)
	}
	s.img.Refresh()
}
