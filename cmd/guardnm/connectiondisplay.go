package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/high-creek-software/guardnm/connections"
)

type connectionDisplay struct {
	widget.BaseWidget

	connection *connections.Connection
	OnToggle   func(c *connections.Connection, turnOn bool)

	swtch *Switch
	name  *widget.Label
}

func (c *connectionDisplay) CreateRenderer() fyne.WidgetRenderer {

	c.swtch = NewSwitch(func(on bool) {
		c.OnToggle(c.connection, on)
	}, SetImageWidth(50), SetImageHeight(25), SetImagePadding(10))
	if c.connection != nil {
		c.swtch.Checked = c.connection.Status == connections.Active
	}

	nameTxt := ""
	if c.connection != nil {
		nameTxt = c.connection.Name
	}
	c.name = widget.NewLabel(nameTxt)

	//return &connectionDisplayRenderer{cd: c, name: c.name, swtch: c.swtch}

	return widget.NewSimpleRenderer(container.NewBorder(nil, nil, nil, c.swtch, c.name))
	//return widget.NewSimpleRenderer(container.NewGridWithColumns(2, c.name, c.swtch))
}

func (c *connectionDisplay) updateConnection(con *connections.Connection) {
	c.connection = con
	//c.Refresh()
	c.name.SetText(con.Name)
	c.swtch.Checked = c.connection.Status == connections.Active

}

func newConnectionDisplay(connection *connections.Connection) *connectionDisplay {
	cd := &connectionDisplay{connection: connection}
	cd.ExtendBaseWidget(cd)
	cd.Refresh()

	return cd
}

type connectionDisplayRenderer struct {
	cd *connectionDisplay

	name  *widget.Label
	swtch *Switch
}

func (c *connectionDisplayRenderer) Destroy() {

}

func (c *connectionDisplayRenderer) Layout(size fyne.Size) {
	pos := fyne.NewPos(theme.Padding(), theme.Padding())
	c.name.Move(pos)

	checkSize := c.swtch.MinSize()
	yOffset := fyne.Max(theme.Padding(), size.Height/2-checkSize.Height/2)
	checkPos := fyne.NewPos(size.Width-theme.Padding()-checkSize.Width, yOffset)
	c.swtch.Move(checkPos)
	//c.swtch.Resize(checkSize)
}

func (c *connectionDisplayRenderer) MinSize() fyne.Size {
	nameSize := c.name.MinSize()
	checkSize := c.swtch.MinSize()

	return fyne.NewSize(nameSize.Width+checkSize.Width, fyne.Max(nameSize.Height, checkSize.Height))
}

func (c *connectionDisplayRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{c.name, c.swtch}
}

func (c *connectionDisplayRenderer) Refresh() {
	if c.cd.connection == nil {
		return
	}

	c.name.SetText(c.cd.connection.Name)
	c.swtch.Checked = c.cd.connection.Status == connections.Active
	c.swtch.Refresh()
}
