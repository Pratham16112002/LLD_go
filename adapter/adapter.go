package adapter

import "fmt"

type Client struct {
}

func (c *Client) InsertIntoLightiningPortIntoComputer(com Computer) {
	com.InsertIntoLightiningPort()
}

type Computer interface {
	InsertIntoLightiningPort()
}

type Mac struct{}

func (m *Mac) InsertIntoLightiningPort() {
	fmt.Println("lightining connector is plugged into mac machine.")
}

type Windows struct {
}

func (w *Windows) InsertIntoUSBPort() {
	fmt.Println("usb port is plugged into windows machine.")
}

type WindowsAdapter struct {
	WindowsMachine *Windows
}

func (w *WindowsAdapter) InsertIntoLightiningPort() {
	fmt.Println("adapter convert lighting signal to USB")
	w.WindowsMachine.InsertIntoUSBPort()
}
