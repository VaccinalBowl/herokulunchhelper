package main

import (
	"github.com/kataras/iris"
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
	"strings"
	"strconv"
)

var current int
var indices []int

type Restaurant struct {
	name     string
	foodType string
	distance string
	price    string
	address  string
	lat      float64
	lng      float64
}

var restaurants = []Restaurant{
	{"Mercedes Canteen", "Generic", "close", "cheap(<6€)", "Arnulfstraße 61, 80636 München", 48.146091, 11.536832},
	{"Richel's", "Mainly Schnitzel", "not so far", "average(7-10€)", "Richelstraße 10, 80634 München", 48.146103, 11.534684},
	{"Edeka", "Supermarket", "close", "cheap", "Erika-Mann-Straße 68, 80335 München", 48.143750, 11.535762},
	{"Pappasito", "Mexican", "close", "average", "Erika-Mann-Straße 60, 80636 München", 48.143745, 11.536954},
	{"Viet Ha", "Vietnamese", "not so far but not so close either", "average(7-10)€", "Landsberger Str. 104, 80339 München", 48.140104, 11.536692},
	{"La Forschetta", "Italian", "not so far but not so close either", "average(7-10)€", "Klaus-Mann-Platz 1, 80636 München", 48.143029, 11.545894},
	{"Foodtrucks", "various", "close", "average(7-10€)", "Rainer-Werner-Fassbinder-Platz 1, 80636 München", 48.143704, 11.537177},
	{"Dean&David", "Salad", "close", "average(7-10€)", "Rainer Werner Fassbinder Platz, 4, 80636 München", 48.143690, 11.537466},
	{"Yak & Yeti Himalayan Food House","Indian","not so far but not so close either","Don't know yet","Blücherstraße 1, 80634 München",48.146755, 11.53316},
	{"Sappralot","Dunno yet.","not so far but not so close either","Dunno yet.","Donnersbergerstraße 37, 80634 München",48.148780, 11.533900},
	{"Bollywood Indian Restaurant","Indian","kind of close","average(7-10)€","Donnersbergerstraße 44, 80634 München",48.147742, 11.534823},
	{"Il Castagno","Italian","Miles Away. Bring hiking shoes.","average(7-10)€","Grasserstraße 10, 80339 München",48.140898, 11.548970},
	{"Hans im Glück","(Pretentious & Overrated) Burger","Miles Away. Bring hiking shoes.","average(7-10)€","Nymphenburger Str. 69, 80335 München",48.149539, 11.546905},
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
func printStringSlice(s []string) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func initRandomSequence(min, max int) {
	indices = make([]int, 0)
	url := fmt.Sprintf("https://www.random.org/integers/?num=100&min=%d&max=%d&col=1&base=10&format=plain&rnd=new", min, max)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	} else {
		defer response.Body.Close()
		responseData, _ := ioutil.ReadAll(response.Body)
		numbersAsStrings := string(responseData)
		numbers := strings.Split(numbersAsStrings, "\n")
		//printStringSlice(numbers)
		for i := 0; i < len(numbers)-1; i++ {
			var number int
			number, _ = strconv.Atoi(numbers[i])
			indices = append(indices, number)
			//printSlice(indices)
		}
	}
}

func getIndex() int {
	//printSlice(indices)
	//fmt.Println("Current:%d",current)
	index := indices[current];
	current++
	if current == 100 {
		initRandomSequence(0, len(restaurants)-1)
		current = 0
	}
	return index
}

func main() {
	current = 0
	initRandomSequence(0, len(restaurants)-1)
	app := iris.New()
	app.RegisterView(iris.HTML("./views", ".html"))
	app.Get("/", func(ctx iris.Context) {

		index := getIndex()
		fmt.Println(index)
		var rest Restaurant = restaurants[index]
		ctx.ViewData("name", rest.name)
		ctx.ViewData("foodType", rest.foodType)
		ctx.ViewData("distance", rest.distance)
		ctx.ViewData("price", rest.price)
		ctx.ViewData("address", rest.address)
		ctx.ViewData("lat", rest.lat)
		ctx.ViewData("lng", rest.lng)
		ctx.View("hello.html")
	})
	app.Run(iris.Addr(":8080"))
}
