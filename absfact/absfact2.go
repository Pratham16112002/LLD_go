package absfact

import "fmt"

type Button interface {
	Render()
}

type Checkbox interface {
	Render()
}

type WindowsButton struct{}

func (wb *WindowsButton) Render() {
	fmt.Println("Rendering windows button")
}

type WindowsCheckBox struct{}

func (wb *WindowsCheckBox) Render() {
	fmt.Println("Rendering windows Checkbox")
}

type MacOSButton struct{}

func (mb *MacOSButton) Render() {
	fmt.Println("Render Macos button")
}

type MacOSCheckbox struct{}

func (mc *MacOSCheckbox) Render() {
	fmt.Println("Rendering Macos Checkbox")
}

type GUIFactory interface {
	CreateButton() Button
	CreateCheckbox() Checkbox
}

type WindowsFactory struct{}

func (wf *WindowsFactory) CreateButton() Button {
	return &WindowsButton{}
}

func (wf *WindowsFactory) CreateCheckbox() Checkbox {
	return &WindowsCheckBox{}
}

type MacOsFactory struct{}

func (mo *MacOsFactory) CreateButton() Button {
	return &MacOSButton{}
}

func (mo *MacOsFactory) CreateCheckbox() Checkbox {
	return &MacOSCheckbox{}
}
