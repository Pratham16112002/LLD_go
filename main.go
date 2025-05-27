package main

import (
	"sync"

	"github.com/Pratham16112202/LLD_go/adapter"
	"github.com/Pratham16112202/LLD_go/singleton"
)

// func renderUI(factory absfact.GUIFactory) {
// 	button := factory.CreateButton()
// 	checkbox := factory.CreateCheckbox()
// 	button.Render()
// 	checkbox.Render()

// }

// func ExampleEmailService() {
// 	// welcomeTemplate := prototype.NewWelcomeEmail()
// 	// resetTemplate := prototype.NewPasswordResetEmail()

// 	// welcome := welcomeTemplate.Clone()
// 	// welcome.SetContent("Hi Pratham, welcome to our platform")
// 	// welcome.Send("prathamdhiman@rediffmail.com")

// 	// resetEmail := resetTemplate.Clone()
// 	// resetEmail.SetContent("Hi Bob , click on the below link to reset your password ")
// 	// resetEmail.SetSubject("Reset password")
// 	// resetEmail.Send("bob@exmpale.com")

// }

func TestingSingleton() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		log := singleton.GetInstance()
		log.Println("Log message from go routine1")
	}()

	go func() {
		defer wg.Done()
		log := singleton.GetInstance()
		log.Println("Log message from go routine2")
	}()

	// wait for both the go routines to get completed
	wg.Wait()

}

func TestingAdapterPattern() {
	client := &adapter.Client{}
	mac := &adapter.Mac{}
	window := &adapter.Windows{}
	client.InsertIntoLightiningPortIntoComputer(mac)
	// to connect windows client we need an Usb to Lighting adapter
	windowAdapter := &adapter.WindowsAdapter{
		window,
	}
	client.InsertIntoLightiningPortIntoComputer(windowAdapter)

}

func main() {
	TestingAdapterPattern()
}
