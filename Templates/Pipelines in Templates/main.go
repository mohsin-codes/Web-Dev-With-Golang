package main

import (
	"log"
	"math"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("1.gohtml"))
}

func Double(x int) int {
	return x + x
}

func Square(x int) float64 {
	return math.Pow(float64(x), 2)
}

func sqRoot(x float64) float64 {
	return math.Sqrt(x)
}

var fm = template.FuncMap{
	"fDouble": Double,
	"fSquare": Square,
	"fSqrt":   sqRoot,
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "1.gohtml", 3)
	if err != nil {
		log.Fatalln(err)
	}
}
