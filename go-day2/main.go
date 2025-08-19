package main

import (
	"fmt"
	"errors"
	"encoding/json"
)

type Student struct {
	Name string `json:"name"`
	Score int `json:score`
}


func main() {

	// Step 1. 變數與常數
	fmt.Println("Step 1. 變數與常數")
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
	fmt.Println("-----\n")


	// Step 2. 基本型別
	fmt.Println("Step 2. 基本型別")
	var i int16
	var f float64
	var s string
	var bo bool
	fmt.Println(i, f, s, bo)
	// 💡 Idioms
	// 建議直接用零值，而不是「初始化為空」。
	// Go 不存在 uninitialized 狀態
	fmt.Println("-----\n")


	// Step 3. 函式
	fmt.Println("Step 3. 函式")
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("錯誤：", err)
	} else {
		fmt.Println("結果：", result)
	}
	fmt.Println("-----\n")


	// Step 4. 陣列 vs 動態陣列
	fmt.Println("Step 4. 陣列 vs 動態陣列")

	arr1  := [3]int{1, 2, 3}   // 直接指定 , [1, 2, 3]
	fmt.Println(arr1)

	arr2  := [...]int{1} // 自動推導產生陣列長度 [1]
	fmt.Println(arr2)

	arr3  := [3]int{1,2} // 部分初始化（其餘補零值） , 會產生 [1,2,0]
	fmt.Println(arr3)

	arr4  := [5]int{0: 1, 2:3} // 指定索引初始化 [1,0,3,0,0]
	fmt.Println(arr4)

	arr5  := [...]int{0: 1, 2:5} // 猜測為 [1,0,5]
	fmt.Println(arr5)


	matrix := [2][3]int{} // 有 {} 就表示全帶入0 , 也可用 var matrix [2][3]int
	matrix[1][1] = 4
	fmt.Println(matrix) // [[0 0 0] [0 4 0]]

	nums := []int{1,2,3,4} // cap = 4
	nums = append(nums, 5) // [1,2,3,4,5] , 塞不下原本slice, 所以產生一個更大的陣列(2倍長度)，將舊資料搬移cap=8
	fmt.Println(nums)
	fmt.Println(len(nums)) // 長度為 5
	fmt.Println(cap(nums)) // cap = 8 , cap為容量，如果 newLen <= 2*oldCap , 新容量就會是就容量的兩倍, 超過1024 就會採用1.25倍

	// 💡 Idioms
	// 常用 append 擴充。
	// len 取長度，cap 取容量。
	// 切片是 reference type，傳遞時不會 copy。 但是如果影響更動到 cap, 就會產生新的array進行重新對應

	passnums := []int{1,2,3}
	modify(passnums)
	fmt.Println(passnums) // [100 2 3]
	modifyAppend(passnums)
	fmt.Println("in main = ", passnums) // [100 2 3]

	// 補充
	// Go runtime 裡 slice header 的樣子
	// type slice struct {
	//     ptr *array         // 指向底層 array
	//     len int            // 切片長度
	//     cap int            // 容量
	// }
	fmt.Println("-----\n")

	// Step 5. maps
	scores := map[string]int {
		"Jodie": 100,
		"Beard": 80,
	}
	scores["Archie"] = 96

	v, ok := scores["Beard"]
	if ok {
		fmt.Println("Beard 成績為：",v)
	} else {
		fmt.Println("無 Beard 的成績資料")
	}
	fmt.Println("-----\n")

	// Step 6. final 小練習
	// 👉 建立一個小程式：
	// 宣告一個 map[string][]int，存放學生名字與多筆成績。
	// 計算每位學生的平均分數。
	// 用 func 回傳 (avg float64, err error)，若學生沒有成績則回傳錯誤。
	// [OUT] Alice 平均分數: 88.33
	// [OUT] Bob 平均分數: 75.00
	// [OUT] Charlie 錯誤: 沒有成績


	// 練習版Ａ
	myClassA := map[string][]int {
		"Beard": {100, 98, 90},
		"Jodie": {100, 100, 99},
	}

	score, err := averageScore(myClassA, "Beard")
	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Printf("Beard 平均分數：%.2f\n", score)
	}

	score, err = averageScore(myClassA, "Archide")
	if err != nil {
		fmt.Printf("Archide 錯誤: 沒有成績\n")
	}else {
		fmt.Printf("Archide 平均分數：%.2f\n", score)
	}

	score, err = averageScore(myClassA, "Jodie")
	if err != nil {
		fmt.Printf("Jodie 錯誤: 沒有成績\n")
	}else {
		fmt.Printf("Jodie 平均分數：%.2f\n", score)
	}

	fmt.Println("-----\n")
	// 練習版Ｂ
	myClassB := map[string][]int {
		"Beard": {95,93,91},
		"Jodie": {100, 97, 99},
		"Archie": {},
	}

	for _, name := range []string{"Beard","Jodie","Archide","Archie"} {
		score, err := averageScoreVer2(myClassB, name)
		if err != nil {
			switch {
				case errors.Is(err, errors.New("查無成績")):
					fmt.Printf("%s 錯誤: 沒有成績\n", name)
				case errors.Is(err, errors.New("學生不存在")):
					fmt.Printf("%s 錯誤: 查無此學生\n", name)
				default:
					fmt.Printf("%s 錯誤: %v\n", name, err)
			}
			continue
		}

		fmt.Printf("%s 平均分數：%.2f\n", name, score)
	}

	fmt.Println("-----\n")
	// Step 6. slice <--> map <--> JSON

	// 6-1. slice → JSON
	students := []Student{
		{Name: "Alice", Score: 90},
		{Name: "Bob", Score: 85},
		{Name: "Charlie", Score: 95},
	}

	jsonBytes, err := json.MarshalIndent(students, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Println("Slice -> JSON:")
	fmt.Println(string(jsonBytes))

	// 6-2. JSON -> slice
	var decoded []Student
	if err := json.Unmarshal(jsonBytes, &decoded); err!= nil {
		panic(err)
	}
	fmt.Println("\nJSON -> Slice:")
	fmt.Println(decoded)

	// 6-3. slice -> map (以Name當key)
	scoreMap := make(map[string]int)
	for _, s := range students {
		scoreMap[s.Name] = s.Score
	}
	fmt.Println("\nSlice -> Map")
	fmt.Println(scoreMap)
}

func averageScore(classScore map[string][]int, name string )(avg float64, err error) {
	scoreList, ok := classScore[name]

	arraySum := 0
	arrayAvg := 0.0
	if ok {
		for i := 0; i < len(scoreList); i++ {
			arraySum += scoreList[i]
		}

		arrayAvg = float64(arraySum) / float64(len(scoreList))
		return arrayAvg, nil
	}else{
		return arrayAvg, fmt.Errorf("查無成績")
	}

}

func averageScoreVer2(classScore map[string][]int, name string)(avg float64, err error) {
	scoreList, ok := classScore[name]

	if !ok {
		return 0, errors.New("學生不存在");
	}

	if len(scoreList) == 0 {
		return 0, errors.New("查無成績")
	}

	sum := 0
	for _, score := range scoreList {
		sum += score
	}
	return float64(sum) / float64(len(scoreList)), nil
}


// Step 3. 函式
// 傳回多個值
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("除數不可為0")
	}
	return a / b, nil
}


// Step 4. 測試: slice 操作
func modify(s []int) {
    s[0] = 100
}
func modifyAppend(s []int) {
	s = append(s, 4) // 原本cap不足，所以進行建立一個新的array, pointer已經不同
	// 因為 slice 是數值複製，len 的改動（append 造成的長度變化）只影響函式內的這份 slice header。
	fmt.Println("in func modifyAppend = ", s) // [100,2,3,4]
}
