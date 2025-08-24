package main

import (
	"fmt"
	"math"
)

type Student struct {
	Name string
	Score int
}

func main() {
	tryStruct()
	tryMethod()
	tryInterface()
	printValue(100)
	printValue("Hello")

	hw()
}

func tryStruct() {
	// struct
	s1 := Student{Name: "Alice", Score: 90}
	s2 := Student{"Bob", 85} // 短寫, 這種寫法需注意順序

	fmt.Println(s1)
	fmt.Println(s2.Name, s2.Score)
}

type Rectangle struct {
	Width, Height float64
}

// 值接收器
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// 指標接收器
func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}

func tryMethod() {
	rect := Rectangle{10, 5}
	fmt.Println("面積:", rect.Area())

	rect.Scale(2)
	fmt.Println("放大2倍後面積:", rect.Area())
}


type Shape interface {
	// https://ithelp.ithome.com.tw/articles/10332708?sc=rss.iron
	// 介面型別會包含一系列函式與方法特徵(method signatures)，而其他型別只要定義相同方法，就等於**實作(implement)**此介面，並可以當成該介面型別來看待。
	Area() float64
}

type Circle struct {
	Radius float64
}

// Circle 實作了 Area, 等於實現了sharp 介面
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func tryInterface() {
	var s Shape
	s = Circle{10}
	fmt.Println("圓面積:", s.Area())
}

func printValue(v interface{}) {
	// 空介面 interface{} 可以容納任何型別
	switch val := v.(type) {
		case int:
			fmt.Println("整數：", val)
		case string:
			fmt.Println("字串：", val)
		default:
			fmt.Println("其他型別：",val)
	}
}


type HwShape interface {
	Area() float64
	Perimeter() float64
	ShapeName() string
}

type HwRectangle struct { width, height float64 }
func (r HwRectangle) Area() float64 { return r.width * r.height }
func (r HwRectangle) Perimeter() float64 { return (r.width + r.height) * 2 }
func (r HwRectangle) ShapeName() string { return "矩形" }

type HwCircle struct { radius float64 }
func (c HwCircle) Area() float64 { return c.radius * c.radius * math.Pi }
func (c HwCircle) Perimeter() float64 { return 2 * c.radius * math.Pi }
func (c HwCircle) ShapeName() string { return "圓形" }

func PrintShapeInfo(s HwShape) {
	fmt.Printf("%s: 面積=%.2f, 周長=%.2f\n", s.ShapeName(), s.Area(), s.Perimeter())
}



func hw() {
	// 👉 建立一個 HwShape 系統：
	// 。定義介面 HwShape，有 Area() 與 Perimeter()。
	// 。建立 HwRectangle 與 HwCircle，實作介面。
	// 。寫一個函式 PrintShapeInfo(s HwShape)，印出面積與周長。
	// 。在 main 建立不同形狀，放進 []HwShape slice，逐一列印。

	// 輸出範例：
	// 矩形: 面積=50.00, 周長=30.00
	// 圓形: 面積=314.16, 周長=62.83

	rec1 := HwRectangle{10, 20}
	rec2 := HwRectangle{3, 4}
	circle1 := HwCircle{2}
	circle2 := HwCircle{5}
	shapelist := []HwShape{rec1, rec2, circle1, circle2}

	// 這個 _ 在 Go 裡是個「空白辨識子 (blank identifier)」，用來丟棄不需要的值
	// range 會回傳: index, value , 如果有沒用到的var, go會出錯，所以沒用到就用 _ 處理
	for i, s := range shapelist {
		fmt.Println("----\nHwShape ", i)
		PrintShapeInfo(s)
	}

	// 。嘗試做一個 TotalArea(shapes []HwShape) float64，計算所有形狀的總面積。
	fmt.Printf("\n總面積:%.2f\n", TotalArea(shapelist))
}

func TotalArea(shapes []HwShape) float64 {
	total := 0.0
	for _, v := range shapes {
		total += v.Area()
	}
	return total
}

// note:

// 一般函式 (function)
// func 函式名(參數...) 回傳值 {...}
// 方法 (method)
// func (接收者 ReceiverType) 方法名(...) 回傳值 {...}

// Q: 為什麼有些要用 method，有些用 function？
// A: 跟某個型別緊密相關的行為 → method
// 例如「圓形的面積」、「矩形的周長」→ 這些都是該型別的行為，用 method 比較合理。
// 跟多個型別通用的行為，或單純工具邏輯 → function
// 例如 PrintShapeInfo(s HwShape)，它只需要介面就能完成，不需要和某一個 struct 綁定。
