package main

import "fmt"

// func main() {
// 	/*
// 		fmt.Println("Hello World!", " Concat", " DONG!")
// 		fmt.Print("Print\n")
// 		fmt.Print("doang")
// 	*/

// 	//cukup tau lah ya

// 	var (
// 		name string = "Whiz"
// 		star int    = 2
// 		// var nama = true
// 		// nama := 3
// 		// angka1, angka2, angka3 int = 1, 2, 3
// 		_, _ = "variabel", "ga dipake"

// 		a       uint8   = 8    //0-255
// 		float   float64 = 1.03 //0-banyak
// 		boolean bool    = true
// 	)

// 	fmt.Println(name, "is a", star, "starred hotel")

// 	name = "Harris"
// 	star = 3

// 	fmt.Println("And", name, "is a", star, "starred hotel")
// 	//best practice for json
// 	fmt.Printf("%s is a decent hotel", name)
// 	fmt.Printf("\n %v is int", a)
// 	fmt.Printf("\n %.3f is float", float)
// 	fmt.Printf("\n %v is bool", boolean)
// }

// func main() {

// 	// constants & iota
// 	var message = `Selamat datang di
// 	salam kenal
// 	letsgooooo.`
// 	const full_name string = "Puji Choirul Huda"
// 	const (
// 		//buat responsecode
// 		v1 = 404 + iota
// 		v2 = 404 + iota
// 		v3 = iota
// 	)

// 	fmt.Println(message)
// 	fmt.Println(v1)
// 	fmt.Println(v2)
// 	fmt.Println(v2)
// 	fmt.Println(v3, "\n")

// 	//arithmetic & operators

// 	var a = 10
// 	var b = 20
// 	var jumlah = a + b - b*2/2
// 	var perbandingan = jumlah > 10
// 	var logical = jumlah > 10 && jumlah != 10 || jumlah == 10

// 	fmt.Println(perbandingan, jumlah, logical)

// }

// func main() {

// 	//conditional
// 	var currentYear = 2023

// 	if age := currentYear - 1989; age <= 18 {
// 		fmt.Println("Umurnya ini dibawah 18")
// 	} else {
// 		fmt.Println("Umurnya ini diatas 18")
// 	}

// 	var score = 11

// 	// switch score {
// 	// case 8:
// 	// 	fmt.Println("ok")
// 	// case 10:
// 	// 	fmt.Println("beyond")
// 	// case 7:
// 	// 	fmt.Println("good")
// 	// default:
// 	// 	fmt.Println("mayan")
// 	// }

// 	switch {
// 	case score == 8:
// 		fmt.Println("ok")
// 		fallthrough
// 	case score > 10:
// 		fmt.Println(" AND beyond")
// 	case score == 7:
// 		fmt.Println("good")
// 	default:
// 		fmt.Println("mayan")
// 	}

// 	fmt.Println("\n===================================================================================================\n")

// 	//loops

// 	for i := 0; i < 5; i++ {
// 		fmt.Println(i)
// 	}

// 	var i = 0
// 	for i < 3 {
// 		fmt.Println(i)
// 		i++
// 	}

// 	fmt.Println("\n")

// 	i = 0
// 	for {
// 		fmt.Println(i)
// 		i++
// 		if i > 5 {
// 			break
// 		}
// 	}

// 	fmt.Println("\nnested\n")

// 	for i := 0; i < 2; i++ {
// 		for i := 0; i < 5; i++ {
// 			fmt.Print(i)
// 		}
// 		fmt.Println(i)
// 	}

// }

// func main() {

// 	var repeat = 15
// 	// var i = 0

// 	for i := 0; i < repeat; i++ {
// 		if i%3 == 0 {
// 			fmt.Print("Fizz")
// 		}
// 		if i%5 == 0 {
// 			fmt.Print("Buzz")
// 		}
// 		fmt.Print(i)
// 	}

// }

// func main() {

// 	//arrays

// 	var numbers [2]int
// 	numbers = [2]int{1, 2}

// 	var strings = [3]string{"Apel", "Pisang", "Ceri"}
// 	strings[1] = "banana"

// 	fmt.Printf("%v\n", numbers)
// 	fmt.Printf("%v\n", strings)

// 	fmt.Printf("\n")

// 	for i, satuan := range strings {
// 		fmt.Printf("Index: %v, value: %v\n", i, satuan)
// 	}

// 	multidimensional := [2][3]int{{1, 2}, {3, 4, 5}}

// 	for _, arr := range multidimensional {
// 		for i, satuan := range arr {
// 			fmt.Printf("Index: %v, value: %v\n", i, satuan)
// 		}
// 		fmt.Printf("\n")
// 	}

// }

// func main() {

// 	//slices

// 	var variabel = []interface{}{"apel", 1}

// 	_ = variabel

// 	fmt.Printf("%#v", variabel)

// 	// fruits := make([]string, 0)
// 	// fruits2 := [2]string{"a", "b"}
// 	// fruits = append(fruits, fruits2)

// 	// fruits := make([]string, 0)
// 	// fruits2 := []string{"a", "b"}
// 	// fruits = append(fruits, fruits2...)

// 	// fmt.Print(fruits)

// 	fruits := []string{"CCC", "DDDD"}
// 	fruits2 := []string{"a", "b"}
// 	fruitcombine := copy(fruits, fruits2)

// 	fmt.Printf("%#v", fruitcombine)

// }

// func main() {

// 	//map interface

// 	var person map[string]string
// 	person = map[string]string{}

// 	person["name"] = "poje"
// 	person["age"] = "26"

// 	var persons map[string]interface{}
// 	persons = map[string]interface{}{}

// 	persons["name"] = "poje"
// 	persons["age"] = "26"

// 	fmt.Printf("%#v", persons)
// 	fmt.Printf("Name: ", persons["name"])
// 	fmt.Printf("age: ", persons["age"])
// }

// func main() {

// 	//map

// 	var person = map[string]string{
// 		"name": "poje",
// 		"age":  "26",
// 	}

// 	value, exist := person["age"]

// 	if exist {
// 		fmt.Println(value, "1")
// 	} else {
// 		fmt.Println("no")
// 	}

// 	delete(person, "age")

// 	value, exist = person["age"]

// 	if exist {
// 		fmt.Println(value)
// 	} else {
// 		fmt.Println("no")
// 	}

// 	for key, value := range person {
// 		fmt.Println(key, ":", value)
// 	}

// }

// FUNCTION DAY 2====================================================================================

// func main() {
// 	list := print("a", "b", "c", "d")
// 	// ints := []int{}{1, 2, 3, 4}

// 	// fmt.Println(ints...)
// 	// fmt.Println(sum(ints...))
// 	fmt.Println(list)
// 	fmt.Println(printInt("a", "b", true, "d"))
// 	fmt.Println(greet("yoman", 23))

// }

// func greet(name string, age int8) (returnnya string, pesan string) {
// 	returnnya = fmt.Sprintf("nama saya %v umur saya %v", name, age) // dipake buat return
// 	pesan = fmt.Sprintf("\nyaudahlah duduk dulu sini")
// 	return
// }

// func print(names ...string) map[string]string {
// 	var result = map[string]string{}
// 	for _, name := range names {
// 		result[name] = name
// 	}
// 	// fmt.Println
// 	return result
// }

// func printInt(namesInt ...interface{}) []interface{} {
// 	var result = []interface{}{}
// 	for _, name := range namesInt {
// 		result = append(result, name)
// 	}
// 	// fmt.Println
// 	return result
// }

// func sum(nums ...int8) int8 {
// 	var total int8
// 	for _, num := range nums {
// 		total += num
// 	}
// 	// fmt.Println
// 	return total
// }

// func profile(name string, fav ...string) {
// 	for _, food := range fav {
// 		fmt.Println(food)
// 	}

// }

//CLOSURE
// func main() {
// 	var evenNumbers = func(numbers ...int) []int {
// 		var result []int
// 		for _, v := range numbers {
// 			if v%2 == 0 {
// 				result = append(result, v)
// 			}
// 		}
// 		return result
// 	}

// 	var evenNumbers2 = func(numbers ...int) []int {
// 		var result []int
// 		for _, v := range numbers {
// 			if v%2 == 0 {
// 				result = append(result, v)
// 			}
// 		}
// 		return result
// 	}(1, 2, 3, 4, 5, 6)

// 	fmt.Printf("%v", evenNumbers(1, 2, 3, 4))
// 	fmt.Printf("%v", evenNumbers2)

// 	var result = profile("")

// 	fmt.Println(result("oke laa"))

// }

// func profile(students string) func(string) string {

// 	return func(s string) string {
// 		return fmt.Sprintln(s)
// 	}
// }

//STRUCT
// type User struct {
// 	Name  string
// 	Age   string
// 	Roles []Roles
// }

// type Roles struct {
// 	NameRole string
// 	Kode     string
// }

// func main() {
// 	var user User
// 	user.Name = "Dominic"
// 	user.Roles = []Roles{
// 		{NameRole: "Admin", Kode: "007"},
// 		{NameRole: "AO", Kode: "600"},
// 	}

// 	// user = struct {
// 	// 	Name string
// 	// 	Age int
// 	// }{
// 	// 	Name = "Dom"
// 	// }

// 	user.ChangeName1()
// 	fmt.Println(user.Name)

// 	user.ChangeName2()
// 	fmt.Println(user.Name)
// }

// func (p User) ChangeName1() {
// 	p.Name = "Mailo"
// }

// func (p *User) ChangeName2() {
// 	p.Name = "Mailo"
// }

//INTERFACE
type Kotak struct {
	panjang int
	lebar   int
}

type Balok struct {
	panjang int
	lebar   int
	tinggi  int
}

type hitung interface {
	luas() int
	keliling() int
	volume() int
}

func (k Kotak) luas() int {
	return k.panjang * k.lebar
}

func (b Balok) value() int {
	return b.panjang * b.lebar * b.tinggi
}

func main() {
	var persegi = Kotak{}
	persegi.lebar = 10
	persegi.panjang = 10

	fmt.Println(persegi.luas())
}
