package charts

import (
	"DesktopApp/handlers"
	"github.com/wcharczuk/go-chart/v2"
	"log"
	"os"
	"strconv"
	"time"
)

type ResData struct {
	a int
	b int
	c int
	d int
	f int
}

type WeekCount struct {
	monday    int
	tuesday   int
	wednesday int
	thursday  int
	friday    int
	saturday  int
	sunday    int
}

///-------------------------------SQLITE------------------------------

func PieChartMonthly() (fs string, err error) {
	res := make(map[string]int)
	var data ResData
	s, erro := handlers.GetTaskMonthly()
	for _, el := range s {
		switch el.Status {
		case "Выполняется":
			data.a++
		case "Завершена":
			data.b++
		case "Просрочена":
			data.c++
		case "Новая":
			data.d++
		case "Проверяется":
			data.f++
		}
	}
	res["Выполняется"] = data.a
	res["Завершена"] = data.b
	res["Просрочена"] = data.c
	res["Новая"] = data.d
	res["Проверяется"] = data.f
	items := make([]chart.Value, 0)
	for name, val := range res {
		b := float64(val)
		c := strconv.Itoa(val)
		items = append(items, chart.Value{
			Value: b,
			Label: name + " " + c,
		})
	}
	pie := chart.PieChart{
		Width:  600,
		Height: 700,
		Values: items,
		DPI:    80,
	}
	var name = "tasks/pie_Monthly.png"
	file, ok := os.Create(name)
	if ok != nil {
		log.Fatal(ok)
	}
	defer file.Close()
	err = pie.Render(chart.PNG, file)
	if err != nil {
		log.Fatal(err)
	}
	return name, erro
}

func PieChart() (fs string, err error) {
	res := make(map[string]int)
	var data ResData
	s, erro := handlers.GetTask()
	for _, el := range s {
		switch el.Status {
		case "Выполняется":
			data.a++
		case "Завершена":
			data.b++
		case "Просрочена":
			data.c++
		case "Новая":
			data.d++
		case "Проверяется":
			data.f++
		}
	}
	res["Выполняется"] = data.a
	res["Завершена"] = data.b
	res["Просрочена"] = data.c
	res["Новая"] = data.d
	res["Проверяется"] = data.f
	items := make([]chart.Value, 0)
	for name, val := range res {
		b := float64(val)
		c := strconv.Itoa(val)
		items = append(items, chart.Value{
			Value: b,
			Label: name + " " + c,
		})
	}
	pie := chart.PieChart{
		Width:  600,
		Height: 700,
		Values: items,
		DPI:    80,
	}
	var name = "tasks/pie_output.png"
	file, ok := os.Create(name)
	if ok != nil {
		log.Fatal(ok)
	}
	defer file.Close()
	err = pie.Render(chart.PNG, file)
	if err != nil {
		log.Fatal(err)
	}
	return name, erro
}

func AxisChart() (fs string, err error) {
	s, err := handlers.GetAxis()
	if err != nil {
		log.Fatal(err)
	}
	var data WeekCount
	res := make(map[string]int)
	for _, val := range s {
		layout := "2006-01-02"
		t, erri := time.Parse(layout, val.Responsible)
		if erri != nil {
			log.Fatal(erri)
		}
		switch t.Weekday().String() {
		case "Monday":
			data.monday++
		case "Tuesday":
			data.tuesday++
		case "Wednesday":
			data.wednesday++
		case "Thursday":
			data.thursday++
		case "Friday":
			data.friday++
		case "Saturday":
			data.saturday++
		case "Sunday":
			data.sunday++
		}
	}
	res["Понедельник"] = data.monday
	res["Вторник"] = data.tuesday
	res["Среда"] = data.wednesday
	res["Четверг"] = data.thursday
	res["Пятница"] = data.friday
	res["Суббота"] = data.saturday
	res["Воскресенье"] = data.sunday
	items := make([]chart.Value, 0)
	var max int
	for name, val := range res {
		value := float64(val)
		if val > max {
			max = val
		}
		items = append(items, chart.Value{
			Value: value,
			Label: name,
		})
	}
	graphic := chart.BarChart{
		Title: "График по дням",
		Background: chart.Style{
			Padding: chart.Box{
				Top: max,
			},
		},
		Bars: items,
	}
	name_file := "tasks/axis.png"
	//newName := "tasks/axis_min.svg"
	f, _ := os.Create(name_file)
	defer f.Close()
	graphic.Render(chart.PNG, f)
	//cmd := exec.Command("svgo", name_file, "-o", newName)
	//output, err := cmd.CombinedOutput()
	//if err != nil {
	//	log.Fatal(err)
	//	fmt.Println("Проверьте, установлен ли у Вас svgo?")
	//}
	//fmt.Println(string(output))
	//fix_file, err := os.Open(newName)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer fix_file.Close()
	//scanner := bufio.NewScanner(fix_file)
	//for scanner.Scan() {
	//	str := strings.Replace(scanner.Text(), ">\\n<path", "><path", -1)
	//	n, ok := os.Create("tasks/axis_min.svg")
	//	if ok != nil {
	//		log.Printf("Ошибка1 %v", ok)
	//	}
	//	defer n.Close()
	//	_, err2 := n.WriteString(str)
	//	if err2 != nil {
	//		log.Printf("Mistake %v", err2)
	//	} else {
	//		fmt.Println("Сделано")
	//	}
	//}
	return name_file, err
}
