package main

import (
	"DesktopApp/menu"
	"DesktopApp/model"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"log"
)

const preferenceCurrentPage = "currentPage"

var topWindow fyne.Window

func main() {
	myApp := app.NewWithID("Desktop")
	w := myApp.NewWindow("Приложение для учёта задач персонала")
	w.SetMaster()
	topWindow = w
	icon, _ := fyne.LoadResourceFromPath("Icon.png")
	w.SetIcon(icon)
	content := container.NewMax()
	title := widget.NewLabel("")
	intro := widget.NewLabel("")
	setPage := func(t model.StartPage) {
		title.SetText(t.Title)
		intro.SetText(t.Intro)
		content.Objects = []fyne.CanvasObject{t.View(w)}
		content.Refresh()
	}
	showPage := container.NewBorder(container.NewVBox(title, widget.NewSeparator(), intro), nil, nil,
		nil, content)
	data := MakeTree(setPage, false, w)
	split := container.NewHSplit(data, showPage)
	split.Offset = 0.2
	main_menu := fyne.NewMainMenu(menu.MyMenu(myApp, w))
	w.Resize(fyne.NewSize(1300, 1000))
	w.SetMainMenu(main_menu)
	w.SetContent(split)
	w.ShowAndRun()
}

func unsupportedPage(t model.StartPage) bool {
	return !t.SupportWeb && fyne.CurrentDevice().IsBrowser()
}

func MakeTree(setPage func(page model.StartPage), loadPrevious bool, w fyne.Window) fyne.CanvasObject {
	ap := fyne.CurrentApp()
	tree := &widget.Tree{
		ChildUIDs: func(uid string) []string {
			return model.StartPagesIndex[uid]
		},
		IsBranch: func(uid string) bool {
			children, ok := model.StartPagesIndex[uid]

			return ok && len(children) > 0
		},
		CreateNode: func(branch bool) fyne.CanvasObject {
			return widget.NewLabel("Тут должен быть виджет")
		},
		UpdateNode: func(uid string, branch bool, obj fyne.CanvasObject) {
			t, ok := model.StartPages[uid]
			if !ok {
				fyne.LogError("Пропавшая панель: "+uid, nil)
				return
			}
			obj.(*widget.Label).SetText(t.Title)
			if unsupportedPage(t) {
				obj.(*widget.Label).TextStyle = fyne.TextStyle{Italic: true}
			} else {
				obj.(*widget.Label).TextStyle = fyne.TextStyle{}
			}
		},
		OnSelected: func(uid string) {
			if t, ok := model.StartPages[uid]; ok {
				if unsupportedPage(t) {
					return
				}
				ap.Preferences().SetString(preferenceCurrentPage, uid)
				setPage(t)
			}
		},
	}
	if loadPrevious {
		currentPref := ap.Preferences().StringWithFallback(preferenceCurrentPage, "welcome")
		tree.Select(currentPref)
	}
	ex_button := container.NewGridWithRows(1, widget.NewButton("Выход", func() {
		ap.Quit()
	}))
	mailButton := widget.NewButtonWithIcon("Обратная связь", theme.MailForwardIcon(), func() {
		username := widget.NewEntry()
		password := widget.NewPasswordEntry()
		//remember := false
		items := []*widget.FormItem{
			widget.NewFormItem("Логин", username),
			widget.NewFormItem("Пароль", password),
		}
		subject := widget.NewEntry()
		subject.SetPlaceHolder("Тема")

		dialog.ShowForm("Введите данные", "Войти", "Отменить", items, func(b bool) {
			if !b {
				return
			} else {
				win := ap.NewWindow("Почтовый клиент")
				data, _, _ := model.MessageForm(username.Text, password.Text, win)
				win.SetContent(data)
				win.Resize(fyne.NewSize(700, 600))
				win.FullScreen()
				win.Show()
			}
			log.Println("Зашли под логином и паролем", username.Text, password.Text)
		}, w)
	})

	res := container.NewBorder(mailButton, ex_button, nil, nil, tree)
	return res
}
