package main

import "fmt"

type Computer interface {
	InsertIntoLightningPort()
}

type Mac struct {
}

type Windows struct {
}

func (m *Mac) InsertIntoLightningPort() {
	fmt.Printf("lightning mac...")
}

func (w *Windows) InsertIntoUSBPort() {
	fmt.Println("lightning windows with USB...")
}

type WindowsAdapter struct {
	windowsMachine *Windows
}

func (w *WindowsAdapter) InsertIntoLightningPort() {
	fmt.Println("Adapter convert Lightning signal to USB.")
	w.windowsMachine.InsertIntoUSBPort()
}

// Client : 调用方
type Client struct {
}

func (c *Client) InsertLightningConnectorIntoComputer(com Computer) {
	fmt.Println("Client inserts Lightning connector into computer.")
	com.InsertIntoLightningPort()
}
