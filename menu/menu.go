package menu

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/cmd/fyne_settings/settings"
	"net/url"
	"os"
)

type Settings struct {
	fyneSettings app.SettingsSchema

	preview *canvas.Image
	colors  []fyne.CanvasObject

	userTheme fyne.Theme
}

func MyMenu(myApp fyne.App, win fyne.Window) (*fyne.Menu, *fyne.Menu) {
	item1 := fyne.NewMenuItem("Новый файл", func() {
		//Пример создания файла Создания файла
		os.Create("created.txt")
	})
	settingsItems := fyne.NewMenuItem("Настройки", func() {
		win = myApp.NewWindow("Настройки")
		win.SetContent(settings.NewSettings().LoadAppearanceScreen(win))
		win.Resize(fyne.NewSize(480, 480))
		win.Show()
	})
	menu1 := fyne.NewMenu("Файл", item1, settingsItems)
	documentation := fyne.NewMenuItem("Документация", func() {
		link, _ := url.Parse("https://developer.fyne.io")
		_ = myApp.OpenURL(link)
	})
	instruction := fyne.NewMenuItem("Инструкция", func() {
		link, _ := url.Parse("https://github.com/DMoscicki/Task-client/blob/main/README.md")
		_ = myApp.OpenURL(link)
	})
	helpMenu := fyne.NewMenu("Помощь", documentation, instruction)

	return menu1, helpMenu
}
