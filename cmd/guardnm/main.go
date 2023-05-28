package main

import (
	_ "embed"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"github.com/high-creek-software/guardnm/connections"
	"log"
	"time"
)

//go:embed icon.png
var iconBytes []byte

func main() {
	g := newGuardnm()
	g.start()
}

type guardnm struct {
	app        fyne.App
	mainWindow fyne.Window
	desk       desktop.App
	menu       *fyne.Menu

	manager *connections.Manager

	connectionItems map[string]*connectionDisplay
	box             *fyne.Container

	//connections []*connections.Connection
}

func newGuardnm() *guardnm {

	iconRes := fyne.NewStaticResource("icon", iconBytes)

	g := &guardnm{connectionItems: make(map[string]*connectionDisplay)}
	g.app = app.NewWithID("io.highcreeksoftware.guardnm")
	g.app.SetIcon(iconRes)
	g.mainWindow = g.app.NewWindow("Guard NM")
	g.mainWindow.Resize(fyne.NewSize(400, 250))
	g.manager = connections.NewManager()

	if desk, ok := g.app.(desktop.App); ok {
		g.desk = desk
		menu := fyne.NewMenu("GuardNM", fyne.NewMenuItem("Show", func() {
			g.mainWindow.Show()
		}))

		g.menu = menu
		desk.SetSystemTrayIcon(iconRes)
		desk.SetSystemTrayMenu(menu)
	}

	g.box = container.NewVBox()

	return g
}

func (g *guardnm) start() {
	g.mainWindow.SetContent(container.NewPadded(container.NewScroll(container.NewBorder(nil, nil, nil, nil, g.box))))
	g.startTicker()
	g.loadConnections()
	g.mainWindow.ShowAndRun()
}

func (g *guardnm) toggleConnection(c *connections.Connection, turnOn bool) {
	g.manager.ToggleConnection(c, turnOn)
}

func (g *guardnm) startTicker() {
	tckr := time.NewTicker(5 * time.Second)
	go func() {
		for {
			select {
			case <-tckr.C:
				log.Println("Ticker ticked again...")
				g.loadConnections()
			}
		}
	}()
}

func (g *guardnm) loadConnections() {
	cons := g.manager.ListWireguardConnections()
	for _, c := range cons {
		if cd, ok := g.connectionItems[c.Name]; ok {
			cd.updateConnection(c)
		} else {
			cd := newConnectionDisplay(c)
			cd.OnToggle = g.toggleConnection
			g.connectionItems[c.Name] = cd
			g.box.Add(cd)
		}
	}
}

/*func (g *guardnm) count() int {
	return len(g.connections)
}

func (g *guardnm) create() fyne.CanvasObject {
	cd := newConnectionDisplay(nil)
	cd.OnToggle = g.toggleConnection
	return cd
}*/
