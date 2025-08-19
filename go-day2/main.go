package main

import "fmt"


func main() {

	// Step 1. 變數與常數
	// 定義明確型別
	var a int = 10
	// 型別自行推斷
	var b = 20
	// 短變數宣告（常見）
	c := 30
	// 常數
	const Pi = 3.14159
	fmt.Println(a, b, c, Pi)

	//💡 Idioms
	// 在函式內幾乎都用 :=。
	// 常數用 const，不可用 :=。



	// Step 2. 基本型別
	var i int16
	var f float64
	var s string
	var bo bool
	fmt.Println(i, f, s, bo)

	// 💡 Idioms
	// 建議直接用零值，而不是「初始化為空」。
	// Go 不存在 uninitialized 狀態


	// Step 3. 函式
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("錯誤：", err)
	} else {
		fmt.Println("結果：", result)
	}
}


// Step 3. 函式
// 傳回多個值
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("除數不可為0")
	}
	return a / b, nil
}


