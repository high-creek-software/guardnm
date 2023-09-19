package connections

import (
	"log"
	"os/exec"
	"strings"
)

const (
	typeWireguard = "wireguard"
)

type Manager struct {
}

func NewManager() *Manager {
	return &Manager{}
}

func (m *Manager) ListWireguardConnections() []*Connection {
	cmd := exec.Command("nmcli", "-t", "-f", "name,type", "con", "show")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil
	}

	connections := make(map[string]*Connection)

	lines := strings.Split(string(out), "\n")
	for _, l := range lines {
		components := strings.Split(l, ":")
		if len(components) == 2 && components[1] == typeWireguard {
			connections[components[0]] = &Connection{Name: components[0], Status: Inactive}
		}
	}

	cmd = exec.Command("nmcli", "-t", "-f", "name,type", "con", "show", "--active")
	out, err = cmd.CombinedOutput()
	if err != nil {
		values := make([]*Connection, len(connections))
		count := 0

		for _, connection := range connections {
			values[count] = connection
			count++
		}

		return values
	}

	lines = strings.Split(string(out), "\n")
	for _, l := range lines {
		components := strings.Split(l, ":")
		if len(components) == 2 && components[1] == typeWireguard {
			name := components[0]
			if con, ok := connections[name]; ok {
				con.Status = Active
				connections[name] = con
			}
		}
	}

	values := make([]*Connection, len(connections))
	count := 0

	for _, connection := range connections {
		values[count] = connection
		count++
	}

	return values
}

func (m *Manager) ToggleConnection(c *Connection, turnOn bool) {
	action := "down"
	if turnOn {
		action = "up"
	}

	cmd := exec.Command("nmcli", "con", action, c.Name)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Println("error toggling connection", err)
	}
	log.Println(string(out))
}
