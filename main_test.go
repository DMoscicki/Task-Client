package main

import (
	"DesktopApp/connection"
	"DesktopApp/handlers"
	"DesktopApp/model"
	"fyne.io/fyne/v2"
	"reflect"
	"testing"
)

var w fyne.Window

func checking(a, b string) bool {
	var result bool
	if reflect.TypeOf(a) == reflect.TypeOf(b) {
		result = true
	} else {
		result = false
	}
	return result
}

func CheckingUser(data interface{}) bool {
	var result bool
	var dataType []handlers.User
	if reflect.TypeOf(data) != reflect.TypeOf(dataType) {
		result = false
		return result
	} else {
		result = true
		return result
	}
}

func CheckingTask(data interface{}) bool {
	var result bool
	var dataType []handlers.Task
	if reflect.TypeOf(data) != reflect.TypeOf(dataType) {
		result = false
		return result
	} else {
		result = true
		return result
	}
}

func TestLoginForm(t *testing.T) {
	email := "test@example.org"
	password := "qazwsxedc94"
	_, user, pass := model.MessageForm(email, password, w)
	if checking(user, pass) == false {
		t.Error(checking(user, pass))
		//t.Logf("%s %s", user, pass)
	} else {
		t.Log("Тест пройден")
	}
}

func TestConnection(t *testing.T) {
	db := connection.OpenConnection()
	err := db.Ping()
	if err != nil {
		t.Errorf("%s", err)
	}
	defer db.Close()
}

func TestChart(t *testing.T) {
	data, err := handlers.GetTask()
	if err != nil {
		t.Errorf("%s", err)
	}
	if len(data) == 0 && CheckingTask(data) == false {
		t.Error("Пусто")
	}
}

func TestAxis(t *testing.T) {
	data, err := handlers.GetAxis()
	if err != nil {
		t.Errorf("%s", err)
	}
	if len(data) == 0 && CheckingTask(data) == false {
		t.Error("пусто")
	}
}

func TestMonthly(t *testing.T) {
	data, err := handlers.GetTaskMonthly()
	if err != nil {
		t.Errorf("%s", err)
	}
	if len(data) == 0 && CheckingTask(data) {
		t.Error("Пусто")
	}
}

func TestPerson(t *testing.T) {
	fio := "YOLAND"
	data, err := handlers.GetTaskByUserLite(fio)
	if err != nil {
		t.Errorf("%s", err)
	}
	if len(data) == 0 && CheckingTask(data) == false {
		t.Errorf("%s", data)
	}
}

func TestList(t *testing.T) {
	fio := "YOLAND"
	data, err := handlers.GetUsers(fio)
	//var dataType []handlers.User
	if err != nil {
		t.Errorf("%s", err)
	}
	if len(data) == 0 || CheckingUser(data) == false {
		t.Error("Пусто")
	}
}

func TestUsers(t *testing.T) {
	_, err := handlers.GetAllUsers()
	if err != nil {
		t.Errorf("%s", err)
	} else {
		t.Log("Ok")
	}
}
