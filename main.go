package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"github.com/veo/jscdecrypt/Canvas"
)

func main() {
	app := app.New()
	app.Settings().SetTheme(theme.LightTheme())
	w := app.NewWindow("jscdecrypt https://github.com/veo/jscdecrypt")
	Canvas.Key.SetPlaceHolder("key")
	w.SetContent(widget.NewVBox(
		Canvas.Key,
		Canvas.JscfileCanvas(),
		Canvas.OutfileCanvas(),
		widget.NewHBox(layout.NewSpacer(), widget.NewLabel("is zip ?:"), Canvas.Iszip, layout.NewSpacer()),
		widget.NewButton("decrypt", func() {
			pr := dialog.NewProgressInfinite("", "decrypting", w)
			pr.Show()
			//Canvas.Decrypt(Canvas.Jscpath.Text)
			go func() {
				Canvas.Decrypt(Canvas.Jscpath.Text)
				pr.Hide()
			}()
		}),
		Canvas.Cmdout,
		Canvas.Fileslist,
		layout.NewSpacer(),
		widget.NewButton("Quit", func() {
			app.Quit()
		})))
	w.Resize(fyne.NewSize(570, 512))
	w.ShowAndRun()
}
