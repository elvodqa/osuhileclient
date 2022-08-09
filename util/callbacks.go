package helper

import (
	"fmt"
	"time"

	"github.com/pwiecz/go-fltk"
)

func CreditsCallback() {
	hakkinda_win := fltk.NewWindow(700, 400)
	hakkinda_win.SetColor(fltk.ColorFromRgb(88, 85, 83))
	hakkinda_win.SetLabel("Hakkinda")
	box := fltk.NewBox(fltk.FLAT_BOX, 0, 0, 700, 300, "")
	hakkinda_img, err := fltk.NewPngImageLoad("resources/credits.png")
	if err != nil {
		fmt.Printf("An error occured: %s\n", err)
	} else {
		box.SetColor(fltk.ColorFromRgb(88, 85, 83))
		box.SetImage(hakkinda_img)
	}
	barisuraz_text := fltk.NewBox(fltk.NO_BOX, 115, 275, 0, 0, "barisuraz")
	barisuraz_text.SetAlign(fltk.ALIGN_CENTER)
	barisuraz_text.SetLabelSize(40)
	barisuraz_text.SetLabelColor(fltk.ColorFromRgb(255, 255, 255))

	barisuraz_credits_text := fltk.NewBox(fltk.NO_BOX, 115, 310, 0, 0, "website")
	barisuraz_credits_text.SetAlign(fltk.ALIGN_CENTER)
	barisuraz_credits_text.SetLabelSize(20)
	barisuraz_credits_text.SetLabelColor(fltk.ColorFromRgb(66, 135, 245))

	humanovan_text := fltk.NewBox(fltk.NO_BOX, 583, 275, 0, 0, "humanovan")
	humanovan_text.SetAlign(fltk.ALIGN_CENTER)
	humanovan_text.SetLabelSize(40)
	humanovan_text.SetLabelColor(fltk.ColorFromRgb(255, 255, 255))

	humanovan_credits_text := fltk.NewBox(fltk.NO_BOX, 583, 310, 0, 0, "client/api")
	humanovan_credits_text.SetAlign(fltk.ALIGN_CENTER)
	humanovan_credits_text.SetLabelSize(20)
	humanovan_credits_text.SetLabelColor(fltk.ColorFromRgb(66, 135, 245))

	discord_button := fltk.NewButton(300, 300, 100, 24, "Discord")
	discord_button.SetColor(fltk.ColorFromRgb(88, 85, 83))
	discord_button.SetAlign(fltk.ALIGN_CENTER)
	discord_button.SetLabelSize(20)
	discord_button.SetLabelColor(fltk.ColorFromRgb(66, 135, 245))
	discord_button.SetCallback(func() {
		OpenBrowser("https://discord.gg/VmmrcMb2")
	})

	website_button := fltk.NewButton(300, 340, 100, 24, "Website")
	website_button.SetSelectionColor(fltk.ColorFromRgb(88, 85, 83))
	website_button.SetColor(fltk.ColorFromRgb(88, 85, 83))
	website_button.SetAlign(fltk.ALIGN_CENTER)
	website_button.SetLabelSize(20)
	website_button.SetLabelColor(fltk.ColorFromRgb(66, 135, 245))
	website_button.SetCallback(func() {
		OpenBrowser("https://www.youtube.com/watch?v=dQw4w9WgXcQ")
	})
	//hakkinda_win.Add(barisuraz_text)
	hakkinda_win.End()
	hakkinda_win.Show()
}

func LoginCallback() {
	login_win := fltk.NewWindow(400, 600)
	bar_value := 0.0

	load_timer := time.NewTimer(time.Millisecond * 10)

	fbi_box := fltk.NewBox(fltk.FLAT_BOX, 50, 50, 300, 300, "")
	fbi_logo, err := fltk.NewPngImageLoad("resources/fbi_logo.png")
	if err != nil {
		fmt.Printf("An error occured: %s\n", err)
	} else {
		fbi_logo.Scale(300, 300, true, true)
		fbi_box.SetImage(fbi_logo)
	}
	fbi_box.Activate()

	cia_box := fltk.NewBox(fltk.FLAT_BOX, 50, 50, 300, 300, "")
	cia_logo, err := fltk.NewPngImageLoad("resources/cia_logo.png")

	if err != nil {
		fmt.Printf("An error occured: %s\n", err)
	} else {
		cia_logo.Scale(300, 300, true, true)
		cia_box.SetImage(cia_logo)
	}
	cia_box.Activate()
	cia_box.Hide()

	mossad_box := fltk.NewBox(fltk.FLAT_BOX, 50, 50, 300, 300, "")
	mossad_logo, err := fltk.NewPngImageLoad("resources/mossad_logo.png")
	if err != nil {
		fmt.Printf("An error occured: %s\n", err)
	} else {
		mossad_logo.Scale(300, 300, true, true)
		mossad_box.SetImage(mossad_logo)
	}
	mossad_box.Activate()
	mossad_box.Hide()

	load_info := fltk.NewBox(fltk.NO_BOX, 200, 380, 0, 0, "Establishing connection to FBI database...")
	load_info.Activate()
	load_info.SetAlign(fltk.ALIGN_CENTER)
	load_info.SetLabelColor(fltk.ColorFromRgb(56, 38, 49))

	load_bar := fltk.NewProgress(10, 410, 380, 20)
	load_bar.Activate()

	login_win.SetEventHandler(func(e fltk.Event) bool {

		go func() {
			<-load_timer.C
			bar_value += 1
			load_timer.Reset(time.Millisecond * 100)
			load_bar.SetValue(float64(bar_value))
		}()

		if load_bar.Value() > 25 && load_bar.Value() < 60 {
			load_info.SetLabel("Bypassing CIA firewall...")
			fbi_box.Hide()
			cia_box.Show()
		}
		if load_bar.Value() > 60 && load_bar.Value() < 101 {
			load_info.SetLabel("Logging into Mossad...")
			cia_box.Hide()
			mossad_box.Show()
		}
		return false
	})

	login_win.Show()
	login_win.End()

}
