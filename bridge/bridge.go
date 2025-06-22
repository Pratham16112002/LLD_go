package bridge

import "fmt"

type VideoQuality interface {
	load(title string)
}

type HDQuality struct {
}

func (hd *HDQuality) load(title string) {
	fmt.Println("Loading title in HD", title)
}

type FourKQuality struct{}

func (f *FourKQuality) load(title string) {
	fmt.Println("Laoding title in 4k")
}

type Player interface {
	Play(title string)
}

type SmartTVPlayer struct {
	quality VideoQuality
}

func (s *SmartTVPlayer) Play(title string) {
	s.quality.load(title)
	fmt.Printf("Playing %s on smart tv at\n", title)
}

type MobilePlayer struct {
	quality VideoQuality
}

func NewMobilePlayer(quality VideoQuality) *MobilePlayer {
	return &MobilePlayer{
		quality: quality,
	}
}

func (s *MobilePlayer) Play(title string) {
	s.quality.load(title)
	fmt.Printf("Playing %s on mobile at\n", title)
}

type WebPlayer struct {
	quality VideoQuality
}

func NewWebPlayer(quality VideoQuality) *WebPlayer {
	return &WebPlayer{
		quality: quality,
	}
}

func (w *WebPlayer) Play(title string) {
	w.quality.load(title)
	fmt.Printf("Playing %s on web player\n", title)
}
