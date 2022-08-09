package main

import (
	"fmt"
	"math/rand"
	util "osuhileclient/util"

	"github.com/pwiecz/go-fltk"
)

func main() {
	fltk.InitStyles()
	fltk.SetScheme("base")
	win := fltk.NewWindow(640, 480)

	win.SetLabel("ohi!client v0.0.1")
	menu_bar := fltk.NewMenuBar(0, 0, 640, 30, "Menu")
	menu_bar.Activate()
	menu_bar.Add("Hesap/Giris Yap", util.LoginCallback)
	menu_bar.Add("Yardim/Hakkinda", util.CreditsCallback)

	durum := fltk.NewBox(fltk.NO_BOX, 50, 50, 0, 0, "Durum: osu! calismiyor")
	durum.SetLabelSize(20)
	durum.SetAlign(fltk.ALIGN_RIGHT)
	durum.SetLabelColor(fltk.ColorFromRgb(255, 0, 0))
	durum.Show()

	durum_neden := fltk.NewBox(fltk.NO_BOX, 50, 80, 0, 0, "Giris yapilmamis")
	durum_neden.SetLabelSize(15)
	durum_neden.SetAlign(fltk.ALIGN_RIGHT)
	durum_neden.SetLabelColor(fltk.ColorFromRgb(255, 0, 0))
	durum_neden.Show()

	giris_yap_button := fltk.NewButton(65, 100, 130, 50, "Giris Yap")
	giris_yap_button.SetColor(fltk.ColorFromRgb(200, 30, 10))
	giris_yap_button.SetLabelColor(fltk.ColorFromRgb(255, 255, 255))
	giris_yap_button.Show()

	ohi_logo, err := fltk.NewPngImageLoad("resources/osuhilexyz_logo.png")
	if err != nil {
		fmt.Printf("An error occured: %s\n", err)
	} else {
		ohi_logo.Scale(150, 150, true, false)
		box := fltk.NewBox(fltk.FLAT_BOX, 480, 50, 150, 150, "")
		box.SetImage(ohi_logo)
	}

	pencere_ismini_degistir := fltk.NewButton(200, 100, 130, 50, "Pencere ismini \ndegistir")
	pencere_ismini_degistir.SetColor(fltk.ColorFromRgb(88, 85, 83))
	pencere_ismini_degistir.SetLabelColor(fltk.ColorFromRgb(255, 255, 255))
	pencere_ismini_degistir.SetCallback(func() {
		win.SetLabel(fmt.Sprintf("%v", rand.Int()))
	})
	pencere_ismini_degistir.Show()

	yenile := fltk.NewButton(335, 100, 130, 50, "Yenile")
	yenile.SetColor(fltk.ColorFromRgb(88, 85, 83))
	yenile.SetLabelColor(fltk.ColorFromRgb(255, 255, 255))
	yenile.Show()

	gecikme_drag := fltk.NewSpinner(300, 300, 100, 20, "Gecikme (+-)")
	gecikme_drag.Activate()

	win.End()
	win.Show()
	fltk.Run()
}
