package main

import "github.com/Pratham16112202/LLD_go/prototype"

// func renderUI(factory absfact.GUIFactory) {
// 	button := factory.CreateButton()
// 	checkbox := factory.CreateCheckbox()
// 	button.Render()
// 	checkbox.Render()

// }

func ExampleEmailService() {
	welcomeTemplate := prototype.NewWelcomeEmail()
	resetTemplate := prototype.NewPasswordResetEmail()

	welcome := welcomeTemplate.Clone()
	welcome.SetContent("Hi Pratham, welcome to our platform")
	welcome.Send("prathamdhiman@rediffmail.com")

	resetEmail := resetTemplate.Clone()
	resetEmail.SetContent("Hi Bob , click on the below link to reset your password ")
	resetEmail.SetSubject("Reset password")
	resetEmail.Send("bob@exmpale.com")

}

func main() {
	ExampleEmailService()
}
