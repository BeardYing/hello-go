package main
import (
	"fmt"
)

func main() {
	var a, b float64
	var op string

	fmt.Print("請輸入第一個數字：")
	fmt.Scan(&a)

	fmt.Print("輸入運算子 (+ - * /)")
	fmt.Scan(&op)

	fmt.Print("請輸入第二個數字：")
	fmt.Scan(&b)

	switch op {
		case "+":
			fmt.Printf("結果：%.2f\n", a+b)
		case "-":
			fmt.Printf("結果：%.2f\n", a-b)
		case "*":
			fmt.Printf("結果：%.2f\n", a*b)
		case "/":
			if b != 0 {
				fmt.Printf("結果：%.2f\n", a/b)
			}else{
				fmt.Println("除數不可為0")
			}
		default:
			fmt.Println("不支援的運算符號")
	}
}