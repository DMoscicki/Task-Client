package model

import (
	"DesktopApp/charts"
	"DesktopApp/handlers"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	gomail "gopkg.in/mail.v2"
	"log"
	"net/url"
	"time"
)

func parseURL(urlStr string) *url.URL {
	link, err := url.Parse(urlStr)
	if err != nil {
		fyne.LogError("Не могу открыть данную ссылку", err)
	}
	return link
}

type StartPage struct {
	Title, Intro string
	View         func(w fyne.Window) fyne.CanvasObject
	SupportWeb   bool
}

var (
	StartPages = map[string]StartPage{
		"welcome": {"Приветствие", "Добро пожаловать на тестовую страницу приложения " +
			"написанного на Golang", WelcomingPage, true},
		"outputBitrix":  {"Задачи из Битрикс", "", TaskWelcome, true},
		"pie":           {"Диаграмма", "", WeekTask, true},
		"table":         {"Таблица задач", "За текущую неделю", TableTask, false},
		"BarChart":      {"График по дням", "Текущая неделя", AxisPage, true},
		"outputMonthly": {"Задачи за месяц", "Текущий месяц", MonthlyTask, true},
		"userTask":      {"Задачи пользователя", "Текущий месяц", TableUser, true},
		"newUserList":   {"Список пользователей", "", TableListUser, false},
	}
	StartPagesIndex = map[string][]string{
		"":             {"welcome", "outputBitrix", "newUserList"},
		"outputBitrix": {"table", "pie", "BarChart", "outputMonthly", "userTask"},
	}
)

func WelcomingPage(_ fyne.Window) fyne.CanvasObject {
	logo := canvas.NewImageFromFile("assets/gopher.svg")
	logo.FillMode = canvas.ImageFillContain
	logo.SetMinSize(fyne.NewSize(256, 256))
	contain := container.NewCenter(container.NewVBox(
		widget.NewLabelWithStyle("Приложение по учёту задач персонала", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		logo,
		container.NewHBox(
			widget.NewLabel(""),
			widget.NewLabel(""),
			widget.NewHyperlink("GitHub", parseURL("https://github.com/DMoscicki")),
			widget.NewLabel("-"),
			widget.NewHyperlink("documentation", parseURL("https://developer.fyne.io/")),
			widget.NewLabel(" "),
		),
		widget.NewLabel(""),
	))
	return contain
}

/// ---------------------------SQLITE FUNC-----------------------------

//func HashUsers(fio string) (map[string][]string, int, int) {
//	data, err := handlers.GetUsers(fio)
//	if err != nil {
//		panic(err)
//	}
//	pairs := make(map[string][]string)
//	for _, value := range data {
//		pairs["Fio"] = append(pairs["Fio"], value.FIO)
//		pairs["Email"] = append(pairs["Email"], value.Email)
//		pairs["Position"] = append(pairs["Position"], value.Position)
//		pairs["Phone"] = append(pairs["Phone"], value.Phone)
//	}
//
//	var num_columns = len(pairs)
//	var num_rows = len(data)
//
//	return pairs, num_columns, num_rows
//}

//func TableUsersList(_ fyne.Window) fyne.CanvasObject {
//	app := fyne.CurrentApp()
//	win := app.NewWindow("Форма добавления юзера")
//	fio := widget.NewEntry()
//	fio.SetPlaceHolder("Пример: Петров")
//	data, num_columns, num_rows := HashUsers(fio.Text)
//	table := widget.NewTable(
//		func() (int, int) { return num_rows, num_columns },
//		func() fyne.CanvasObject { return widget.NewLabel("") },
//		func(id widget.TableCellID, obj fyne.CanvasObject) {
//			label := obj.(*widget.Label)
//			switch id.Col {
//			case 0:
//				label.SetText(fmt.Sprintf("%d", id.Row+1))
//			case 1:
//				label.SetText(fmt.Sprintf("%s", data["Fio"][id.Row]))
//			case 2:
//				label.SetText(fmt.Sprintf("%s", data["Email"][id.Row]))
//			case 3:
//				label.SetText(fmt.Sprintf("%s", data["Position"][id.Row]))
//			case 4:
//				label.SetText(fmt.Sprintf("%s", data["Phone"][id.Row]))
//			default:
//				label.SetText(fmt.Sprintf("%s", "Пусто"))
//			}
//		},
//	)
//	table.SetColumnWidth(0, 34)
//	table.SetColumnWidth(1, 230)
//	table.SetColumnWidth(2, 300)
//	table.SetColumnWidth(3, 200)
//	table.SetColumnWidth(4, 300)
//	button := widget.NewButtonWithIcon("Поиск", theme.ConfirmIcon(), func() {
//		data, num_columns, num_rows = HashUsers(fio.Text)
//		table.Refresh()
//	})
//	addUser := widget.NewButton("Добавить пользователя", func() {
//		username := widget.NewEntry()
//		username.SetPlaceHolder("John Smith")
//		mail := widget.NewEntry()
//		mail.SetPlaceHolder("test@example.net")
//		phone := widget.NewEntry()
//		phone.SetPlaceHolder("+7(812)XXX-XX-XX")
//		position := widget.NewEntry()
//		position.SetPlaceHolder("Engineer")
//		form := &widget.Form{
//			Items: []*widget.FormItem{
//				{Text: "Имя", Widget: username, HintText: "Введите имя"},
//				{Text: "Почта", Widget: mail, HintText: "Введите почту"},
//				{Text: "Телефон", Widget: phone, HintText: "Введите телефон"},
//				{Text: "Должность", Widget: position, HintText: "Введите позицию"},
//			},
//			OnCancel: func() {
//				win.Close()
//			},
//			OnSubmit: func() {
//				num, err := handlers.InsertUser(username.Text, mail.Text, phone.Text, position.Text)
//				if err != nil {
//					log.Fatal(err)
//				} else {
//					fmt.Println(num)
//				}
//				dialog.ShowInformation("Добавлен", "Успешно добавлен "+username.Text, win)
//			},
//		}
//		//form.Append("Имя", username)
//		//form.Append("Почта", mail)
//		//form.Append("Телефон", phone)
//		//form.Append("Позиция", position)
//		win.Resize(fyne.NewSize(500, 500))
//		win.SetContent(form)
//		win.Show()
//	})
//	rows := container.NewGridWithRows(2, fio, button, addUser)
//	allBox := container.NewBorder(rows, nil, nil, nil, table)
//	return allBox
//}

//--------------------------Более изящный вариант с возможностью редактирования БД----------------------------

func TaskWelcome(_ fyne.Window) fyne.CanvasObject {
	content := container.NewVBox(widget.NewLabelWithStyle("Прошу выбрать из выпадающего списка тип визуализации данных",
		fyne.TextAlignLeading, fyne.TextStyle{Italic: true}))

	return container.NewCenter(content)
}

func MessageForm(email, password string, w fyne.Window) (fyne.CanvasObject, string, string) {

	receiver := widget.NewEntry()
	receiver.SetPlaceHolder("Введите почту получателя")
	subject := widget.NewSelectEntry([]string{"Ошибка", "Доработка", "Предложение по улучшению"})
	subject.SetPlaceHolder("Введите тему")
	messageBody := widget.NewMultiLineEntry()
	messageBody.SetPlaceHolder("Введите сообщение....")
	var saveDir fyne.URIReadCloser
	fileAdd := widget.NewButtonWithIcon("Файл", theme.FileImageIcon(), func() {
		dialog.ShowFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, w)
				return
			}
			if reader == nil {
				fmt.Println("Действие отменено")
				return
			}
			saveDir = reader
			fmt.Println(saveDir.URI().Path())
		}, w)
	})
	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Получатель", Widget: receiver, HintText: "test@example.org"},
			{Text: "Тема", Widget: subject, HintText: "Введите тему письма"},
		},
		OnCancel: func() {
			w.Close()
			//return
		},
		OnSubmit: func() {
			message := gomail.NewMessage()
			message.SetHeader("From", email)                // ПОЧТА ОТПРАВИТЕЛЯ
			message.SetHeader("To", receiver.Text)          // ПОЧТА ПОЛУЧАТЕЛЯ
			message.SetHeader("Subject", subject.Text)      // ТЕМА ПИСЬМА
			message.SetBody("text/plain", messageBody.Text) // ТЕКСТ ПИСЬМА
			message.Attach(fmt.Sprintf("%s", saveDir.URI().Path()))
			// ЗДЕСЬ ПРОПИСАТЬ ВАШ SMTP HOST, PORT
			dialing_info := gomail.NewDialer("SMTP HOST ", 587, email, password)
			err := dialing_info.DialAndSend(message)
			if err != nil {
				dialog.ShowError(err, w)
			} else {
				fyne.CurrentApp().SendNotification(&fyne.Notification{Title: subject.Text,
					Content: "Ваше письмо отправлено"})
				fmt.Println("Письмо отправлено")
			}
		},
	}
	form.Append("Файл", fileAdd)
	form.Append("Текст", messageBody)
	form.Resize(fyne.NewSize(500, 300))
	return form, email, password
}

func MonthlyTask(w fyne.Window) fyne.CanvasObject {
	app := fyne.CurrentApp()
	req, err := charts.PieChartMonthly()
	if err != nil {
		dialog.ShowError(err, w)
	}
	image := canvas.NewImageFromFile(req)
	button := widget.NewButtonWithIcon("Обновить", theme.ConfirmIcon(), func() {
		req, err = charts.PieChartMonthly()
		w.Canvas().Refresh(w.Content())
		if err != nil {
			dialog.ShowError(err, w)
		} else {
			image = canvas.NewImageFromFile(req)
			image.FillMode = canvas.ImageFillOriginal
			image.Resize(fyne.NewSize(400, 500))
			image.Refresh()
		}
	})
	exit_button := widget.NewButton("Выход", func() {
		app.Quit()
	})
	image.FillMode = canvas.ImageFillOriginal
	image.Resize(fyne.NewSize(400, 500))
	image.Refresh()
	hbBox := container.NewGridWithRows(2, exit_button, button)
	allBox := container.NewBorder(hbBox, nil, nil, nil, container.NewCenter(image))
	return allBox
}

func DataUser(fio string) (c map[string][]string, a int, b int, err error) {
	s, err := handlers.GetTaskByUserLite(fio)
	pairs := make(map[string][]string)
	if err != nil {
		log.Fatal(err)
	}
	for _, val := range s {
		layouts := "2006-01-02"
		layoutCorr := "02-01-2006"
		t, err := time.Parse(layouts, val.Responsible)
		if err != nil {
			log.Fatal(err)
		}
		pairs["Title"] = append(pairs["Title"], val.Title)
		pairs["Status"] = append(pairs["Status"], val.Status)
		pairs["Created"] = append(pairs["Created"], val.Created_date)
		pairs["Respons"] = append(pairs["Respons"], t.Format(layoutCorr))
	}
	var num_columns = len(pairs)
	var num_rows = len(s)
	return pairs, num_rows, num_columns, err
}

func TableUser(_ fyne.Window) fyne.CanvasObject {
	fio := widget.NewEntry()
	fio.SetPlaceHolder("Пример: Петров")
	pairs, num_rows, num_columns, _ := DataUser(fio.Text)
	table := widget.NewTable(
		func() (int, int) { return num_rows, num_columns + 1 },
		func() fyne.CanvasObject { return widget.NewLabel("") },
		func(id widget.TableCellID, obj fyne.CanvasObject) {
			label := obj.(*widget.Label)
			switch id.Col {
			case 0:
				label.SetText(fmt.Sprintf("%d", id.Row+1))
			case 1:
				label.SetText(pairs["Status"][id.Row])
			case 2:
				label.SetText(pairs["Created"][id.Row])
			case 3:
				label.SetText(pairs["Respons"][id.Row])
			case 4:
				label.SetText(pairs["Title"][id.Row])
			default:
			}
		},
	)
	table.SetColumnWidth(0, 34)
	table.SetColumnWidth(1, 110)
	table.SetColumnWidth(2, 200)
	table.SetColumnWidth(3, 200)
	table.SetColumnWidth(4, 100)
	button := widget.NewButtonWithIcon("Поиск", theme.ConfirmIcon(), func() {
		pairs, num_rows, num_columns, _ = DataUser(fio.Text)
		table.Refresh()
	})
	rows := container.NewGridWithRows(2, fio, button)
	allBox := container.NewBorder(rows, nil, nil, nil, table)
	return allBox
}

func WeekTask(w fyne.Window) fyne.CanvasObject {
	app := fyne.CurrentApp()
	req, err := charts.PieChart()
	if err != nil {
		dialog.ShowError(err, w)
	}
	card := canvas.NewImageFromFile(req)
	button := widget.NewButtonWithIcon("Обновить", theme.ConfirmIcon(), func() {
		req, err = charts.PieChart()
		w.Canvas().Refresh(w.Content())
		if err != nil {
			dialog.ShowError(err, w)
		} else {
			card = canvas.NewImageFromFile(req)
			card.FillMode = canvas.ImageFillOriginal
			card.Resize(fyne.NewSize(300, 400))
			card.Refresh()
		}
	})
	exit_button := widget.NewButton("Выход", func() {
		app.Quit()
	})
	card.FillMode = canvas.ImageFillOriginal
	card.Resize(fyne.NewSize(400, 500))
	card.Refresh()
	hbBox := container.NewGridWithRows(2, exit_button, button)
	vBox := container.NewCenter(card)
	allBox := container.NewBorder(hbBox, nil, nil, nil, vBox)
	return allBox
}

func TableTask(_ fyne.Window) fyne.CanvasObject {
	s, _ := handlers.GetTask()
	pairs := make(map[string][]string)
	for _, val := range s {
		pairs["ID"] = append(pairs["ID"], val.ID)
		pairs["Title"] = append(pairs["Title"], val.Title)
		pairs["Status"] = append(pairs["Status"], val.Status)
		pairs["Created"] = append(pairs["Created"], val.Responsible)
		pairs["Respons"] = append(pairs["Respons"], val.Created_date)
	}
	var num_columns = len(pairs)
	var num_rows = len(s)
	table := widget.NewTable(
		func() (int, int) { return num_rows, num_columns + 1 },
		func() fyne.CanvasObject { return widget.NewLabel("") },
		func(id widget.TableCellID, obj fyne.CanvasObject) {
			label := obj.(*widget.Label)
			switch id.Col {
			case 0:
				label.SetText(fmt.Sprintf("%d", id.Row+1))
			case 1:
				label.SetText(pairs["Status"][id.Row])
			case 2:
				label.SetText(pairs["Created"][id.Row])
			case 3:
				label.SetText(pairs["Respons"][id.Row])
			case 4:
				label.SetText(pairs["Title"][id.Row])
			default:
			}
		},
	)
	table.SetColumnWidth(0, 34)
	table.SetColumnWidth(1, 110)
	table.SetColumnWidth(2, 250)
	table.SetColumnWidth(3, 290)
	table.SetColumnWidth(4, 400)
	return table
}

func AxisPage(w fyne.Window) fyne.CanvasObject {
	img, err := charts.AxisChart()
	if err != nil {
		dialog.ShowError(err, w)
	}
	image := canvas.NewImageFromFile(img)
	image.FillMode = canvas.ImageFillOriginal
	content := container.NewCenter(image)

	return content
}

func TableListUser(w fyne.Window) fyne.CanvasObject {
	name := widget.NewEntry()
	name.SetPlaceHolder("Имя")
	data, err := handlers.GetAllUsers()
	num_rows := len(data)
	if err != nil {
		log.Fatal(err)
	}

	var list *widget.List

	list = widget.NewList(
		func() int { return num_rows },
		func() fyne.CanvasObject {
			nameLabel := widget.NewLabel("Имя")
			emaiLable := widget.NewLabel("Почта")
			phoneLable := widget.NewLabel("Телефон")
			positionLable := widget.NewLabel("Должность")
			userContainer := container.NewGridWithColumns(4, nameLabel, emaiLable, phoneLable, positionLable)
			edit := widget.NewButtonWithIcon("", theme.DocumentCreateIcon(), nil)
			deleteRow := widget.NewButtonWithIcon("", theme.DeleteIcon(), nil)
			buttonContainer := container.NewHBox(edit, deleteRow)
			return container.NewBorder(nil, nil, nil, buttonContainer, userContainer)
		},
		func(i widget.ListItemID, obj fyne.CanvasObject) {
			item := obj.(*fyne.Container)

			userContainer := item.Objects[0].(*fyne.Container)
			buttonContainer := item.Objects[1].(*fyne.Container)

			nameLabel := userContainer.Objects[0].(*widget.Label)
			emaiLabel := userContainer.Objects[1].(*widget.Label)
			phoneLable := userContainer.Objects[2].(*widget.Label)
			positionLable := userContainer.Objects[3].(*widget.Label)

			edit := buttonContainer.Objects[0].(*widget.Button)
			deleteRow := buttonContainer.Objects[1].(*widget.Button)

			userID := data[i].ID

			nameLabel.SetText(data[i].FIO)
			emaiLabel.SetText(data[i].Email)
			phoneLable.SetText(data[i].Phone)
			positionLable.SetText(data[i].Position)

			edit.OnTapped = func() {

				entryNameLabel := widget.NewEntry()
				entryEmaiLable := widget.NewEntry()
				entryPhoneLable := widget.NewEntry()
				entryPosition := widget.NewEntry()

				nameForm := widget.NewFormItem("Имя", entryNameLabel)
				emailForm := widget.NewFormItem("Почта", entryEmaiLable)
				phoneForm := widget.NewFormItem("Телефон", entryPhoneLable)
				posForm := widget.NewFormItem("Должность", entryPosition)

				formItems := []*widget.FormItem{nameForm, emailForm, phoneForm, posForm}

				firstDialog := dialog.NewForm("Изменение контакта", "Сохранить", "Отмена", formItems, func(b bool) {
					if b {
						user := handlers.User{
							ID:       userID,
							FIO:      entryNameLabel.Text,
							Email:    entryEmaiLable.Text,
							Phone:    entryPhoneLable.Text,
							Position: entryPosition.Text,
						}
						if err = handlers.UpdateUserInfo(user); err != nil {
							log.Fatal(err)
						}
						data, err = handlers.GetAllUsers()
						if err != nil {
							panic(err)
						}
					}
					list.Refresh()
				}, w)

				entryNameLabel.SetText(nameLabel.Text)
				entryEmaiLable.SetText(emaiLabel.Text)
				entryPosition.SetText(positionLable.Text)
				entryPhoneLable.SetText(phoneLable.Text)

				firstDialog.Resize(fyne.NewSize(500, 300))

				firstDialog.Show()

			}

			deleteRow.OnTapped = func() {

				secondDialog := dialog.NewConfirm("Удалить пользователя", "Вы хотите удалить пользователя?", func(b bool) {
					if b {
						if err = handlers.DeleteUser(userID); err != nil {
							log.Fatal(err)
						}
						data, err = handlers.GetAllUsers()
						if err != nil {
							log.Fatal(err)
						}
					}
					list.Refresh()
				},
					w,
				)
				secondDialog.Resize(fyne.NewSize(300, 200))
				secondDialog.Show()

			}
		})
	list.OnSelected = func(id widget.ListItemID) {
		list.UnselectAll()
	}

	add := widget.NewButton("Добавить", func() {

		entryNameLabel := widget.NewEntry()
		entryEmaiLable := widget.NewEntry()
		entryPhoneLable := widget.NewEntry()
		entryPosition := widget.NewEntry()

		nameForm := widget.NewFormItem("Имя", entryNameLabel)
		emailForm := widget.NewFormItem("Почта", entryEmaiLable)
		phoneForm := widget.NewFormItem("Телефон", entryPhoneLable)
		posForm := widget.NewFormItem("Должность", entryPosition)

		formItems := []*widget.FormItem{nameForm, emailForm, phoneForm, posForm}

		dialogAdd := dialog.NewForm("Добавить пользователя", "Добавить", "Отмена", formItems, func(b bool) {
			if b {
				user := handlers.User{
					FIO:      entryNameLabel.Text,
					Email:    entryEmaiLable.Text,
					Phone:    entryPhoneLable.Text,
					Position: entryPosition.Text,
				}

				if err = handlers.InsertUser(user); err != nil {
					panic(err)
				}
				data, err = handlers.GetAllUsers()
				if err != nil {
					panic(err)
				}
				list.Refresh()
			}

		}, w)

		dialogAdd.Resize(fyne.NewSize(500, 300))
		dialogAdd.Show()
	})

	searchButton := widget.NewButton("Поиск", func() {
		data, err = handlers.GetUsers(name.Text)
		if err != nil {
			panic(err)
		}
		num_rows = len(data)
		list.Refresh()
	})
	searchFields := container.NewGridWithRows(2, name, searchButton)
	content := container.NewBorder(searchFields, container.New(layout.NewVBoxLayout(), add), nil, nil, list)

	return content
}

//------------------------------------------------------------------------------------------------------------------
