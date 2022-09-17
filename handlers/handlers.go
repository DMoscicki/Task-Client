package handlers

import (
	"DesktopApp/connection"
	"fmt"
	"log"
)

type Task struct {
	ID           string `json:"id"`
	Title        string `json:"title"`
	Status       string `json:"status"`
	Created_date string `json:"created_date"`
	Responsible  string `json:"responsible_full_name"`
}

type User struct {
	ID       int    `json:"id"`
	FIO      string `json:"full_name"`
	Email    string `json:"email"`
	Position string `json:"work_position"`
	Phone    string `json:"phone"`
}

///-------------------------------------ХЭНДЛЕРЫ ИЗ ПОСТГРИ-----------------------------------

//func GetTask() ([]Task, error) {
//	db := connection.OpenConnection()
//	rows, err := db.Query("SELECT id, title, status, created_date, responsible_full_name FROM bitrix_data_test " +
//		"WHERE (created_date >= date_trunc('week', CURRENT_TIMESTAMP - interval '1 week') and created_date < " +
//		"date_trunc('week', CURRENT_TIMESTAMP)) group by id, title, status, created_date, responsible_full_name;")
//	if err != nil {
//		log.Fatal(err)
//	}
//	var task Task
//	var tasks []Task
//	for rows.Next() {
//		rows.Scan(&task.ID, &task.Title, &task.Status, &task.Responsible, &task.Created_date)
//		tasks = append(tasks, task)
//	}
//	defer rows.Close()
//	defer db.Close()
//	return tasks, err
//}
//
//func GetAxis() ([]Task, error) {
//	db := connection.OpenConnection()
//	rows, err := db.Query("SELECT id, title, status, created_date, responsible_full_name FROM bitrix_data_test " +
//		"WHERE (created_date >= date_trunc('week', CURRENT_TIMESTAMP - interval '1 week') and created_date < " +
//		"date_trunc('week', CURRENT_TIMESTAMP)) group by id, title, status, created_date, responsible_full_name;")
//	if err != nil {
//		log.Fatal(err)
//	}
//	var task Task
//	var tasks []Task
//	for rows.Next() {
//		rows.Scan(&task.ID, &task.Title, &task.Status, &task.Responsible, &task.Created_date)
//		tasks = append(tasks, task)
//	}
//
//	defer rows.Close()
//	defer db.Close()
//
//	return tasks, err
//}
//
//func GetTaskMonthly() ([]Task, error) {
//	db := connection.OpenConnection()
//	rows, err := db.Query("SELECT id, title, status, created_date, responsible_full_name FROM bitrix_data_test " +
//		"WHERE created_date BETWEEN (date_trunc('MONTH', current_date::date))" +
//		" AND (date_trunc('MONTH', created_date::date) + INTERVAL '1 MONTH - 1 day')::DATE " +
//		"group by id, title, status, created_date, responsible_full_name;")
//	if err != nil {
//		log.Fatal(err)
//	}
//	var task Task
//	var tasks []Task
//	for rows.Next() {
//		rows.Scan(&task.ID, &task.Title, &task.Status, &task.Responsible, &task.Created_date)
//		tasks = append(tasks, task)
//	}
//	defer rows.Close()
//	defer db.Close()
//
//	return tasks, err
//}
//
//func GetTaskbyPerson(fio string) ([]Task, error) {
//	db := connection.OpenConnection()
//	rows, err := db.Query("SELECT id, title, status, created_date, responsible_full_name FROM bitrix_data_test "+
//		"WHERE created_date BETWEEN (date_trunc('MONTH', current_date::date))"+
//		" AND (date_trunc('MONTH', created_date::date) + INTERVAL '1 MONTH - 1 day')::DATE and responsible_full_name ILIKE '%'||$1||'%' "+
//		"group by id, title, status, created_date, responsible_full_name;", fio)
//	if err != nil {
//		log.Fatal(err)
//	}
//	var task Task
//	var tasks []Task
//
//	for rows.Next() {
//		rows.Scan(&task.ID, &task.Title, &task.Status, &task.Responsible, &task.Created_date)
//		tasks = append(tasks, task)
//	}
//
//	defer rows.Close()
//	defer db.Close()
//
//	return tasks, nil
//}
//
//func GetUserList(fio string) ([]User, error) {
//	db := connection.OpenConnection()
//	rows, err := db.Query("SELECT email, full_name, work_position, dep_name FROM users_merge where "+
//		"full_name ILIKE '%'||$1||'%';", fio)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	var user User
//	var users []User
//
//	for rows.Next() {
//		rows.Scan(&user.Email, &user.FIO, &user.Position, &user.Department)
//		users = append(users, user)
//	}
//
//	defer rows.Close()
//	defer db.Close()
//
//	return users, err
//
//}

///-------------------------------------------------------------------------------------------

///-------------------------------------ХЭНДЛЕРЫ SQLite-----------------------------------

//func InsertUser(fio, email, phone, position string) (int, error) {
//	db := connection.OpenConnection()
//	res, err := db.Exec("INSERT INTO user (fio, email, phone, position) VALUES ($1, $2, $3, $4)", fio, email, phone, position)
//	if err != nil {
//		log.Fatal(err)
//	}
//	var id int64
//	if id, err = res.LastInsertId(); err != nil {
//		return 0, err
//	}
//
//	defer db.Close()
//	fmt.Println("Успешно")
//	fmt.Println(res)
//
//	return int(id), err
//}

func GetUsers(fio string) ([]User, error) {
	db := connection.OpenConnection()
	rows, err := db.Query("SELECT id, fio, email, phone, position FROM user where fio LIKE ''||$1||'%' group by id, fio, email, phone, position;", fio)
	if err != nil {
		panic(err)
	}

	var user User
	var users []User

	for rows.Next() {
		rows.Scan(&user.ID, &user.FIO, &user.Email, &user.Phone, &user.Position)
		users = append(users, user)
	}
	defer rows.Close()
	defer db.Close()

	return users, err
}

func GetTaskByUserLite(fio string) ([]Task, error) {
	db := connection.OpenConnection()
	rows, err := db.Query("SELECT title, status, created_date, responsible_full_name FROM tasks WHERE created_date BETWEEN datetime('now','localtime','start of month', '-1 month') AND date('now','localtime', 'start of month', '+1 month', '-1 day') "+
		"and responsible_full_name LIKE ''||$1||'%' group by title, status, created_date, responsible_full_name;", fio)
	if err != nil {
		panic(err)
	}

	var task Task
	var tasks []Task

	for rows.Next() {
		rows.Scan(&task.Title, &task.Status, &task.Responsible, &task.Created_date)

		tasks = append(tasks, task)
	}

	defer rows.Close()
	defer db.Close()

	return tasks, err
}

func GetTaskMonthly() ([]Task, error) {
	db := connection.OpenConnection()
	rows, err := db.Query("SELECT title, status, created_date, responsible_full_name FROM tasks WHERE created_date " +
		"BETWEEN datetime('now', 'start of month') AND datetime('now', 'localtime') group by title, status, created_date, responsible_full_name;")
	if err != nil {
		log.Fatal(err)
	}
	var task Task
	var tasks []Task
	for rows.Next() {
		rows.Scan(&task.Title, &task.Status, &task.Responsible, &task.Created_date)
		tasks = append(tasks, task)
	}
	defer rows.Close()
	defer db.Close()

	return tasks, err
}

func GetTask() ([]Task, error) {
	db := connection.OpenConnection()
	rows, err := db.Query("SELECT title, status, created_date, responsible_full_name FROM tasks WHERE created_date BETWEEN datetime('now', '-6 days') " +
		"AND datetime('now', 'localtime') group by title, status, created_date, responsible_full_name;")
	if err != nil {
		log.Fatal(err)
	}
	var task Task
	var tasks []Task
	for rows.Next() {
		rows.Scan(&task.Title, &task.Status, &task.Responsible, &task.Created_date)
		tasks = append(tasks, task)
	}
	defer rows.Close()
	defer db.Close()
	return tasks, err
}

func GetAxis() ([]Task, error) {
	db := connection.OpenConnection()
	rows, err := db.Query("SELECT title, status, created_date, responsible_full_name FROM tasks WHERE created_date BETWEEN datetime('now', 'start of month') " +
		"AND datetime('now','localtime') group by title, status, created_date, responsible_full_name;")
	if err != nil {
		log.Fatal(err)
	}
	var task Task
	var tasks []Task
	for rows.Next() {
		rows.Scan(&task.Title, &task.Status, &task.Responsible, &task.Created_date)

		tasks = append(tasks, task)

	}

	defer rows.Close()
	defer db.Close()

	return tasks, err
}

func UpdateUserInfo(user User) error {
	db := connection.OpenConnection()
	res, err := db.Exec("UPDATE user SET email = $1, fio = $2, phone = $3, position = $4 where id = $5", user.Email, user.FIO, user.Phone, user.Position, user.ID)
	if err != nil {
		panic(err)
	}

	outputResult, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Println(outputResult)

	defer db.Close()

	return err
}

func GetAllUsers() ([]User, error) {
	db := connection.OpenConnection()
	res, err := db.Query("SELECT * FROM user;")
	if err != nil {
		panic(err)
	}

	var user User
	var users []User

	for res.Next() {

		res.Scan(&user.ID, &user.FIO, &user.Email, &user.Phone, &user.Position)
		users = append(users, user)
	}

	defer db.Close()
	defer res.Close()

	return users, err
}

func DeleteUser(id int) error {
	db := connection.OpenConnection()
	res, err := db.Exec("DELETE FROM user where id = $1", id)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	fmt.Println(res)

	return err
}

func InsertUser(user User) error {
	db := connection.OpenConnection()

	res, err := db.Exec("INSERT INTO user (fio, email, phone, position) VALUES ($1, $2, $3, $4)", user.FIO, user.Email, user.Phone, user.Position)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	output, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Println(output)

	return err

}
