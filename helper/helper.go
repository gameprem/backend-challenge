package helper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	// "math"
	"net/http"
	"strings"
)

// <<<<< Ex.1 >>>>>>

func findMaxPathSum(input [][]int) int {
	rows := len(input) // [[59], [73, 41], [52, 40, 53], [26, 53, 6, 34]] => 4

	// สร้าง slice 2 มิติสำหรับเก็บผลรวมที่ได้ที่แต่ละจุด
	sums := make([][]int, rows) // len = 4
	// [[] [] [] []]

	for i := range sums {
		sums[i] = make([]int, len(input[i]))
	}
	// [[0] [0 0] [0 0 0] [0 0 0 0]]

	// กำหนดค่าเริ่มต้นในบริเวณที่มีเพียงตัวเดียว (ด้านฐาน)
	sums[rows-1] = input[rows-1]
	// [[0] [0 0] [0 0 0] [26 53 6 34]]

	// เริ่มต้นจากต่ำไปสูง
	for i := rows - 2; i >= 0; i-- {
		countArray := len(input[i]) // จำนวน array
		for j := 0; j < countArray; j++ {
			// หาผลรวมที่มากที่สุดของเส้นทางที่เป็นไปได้
			greater := input[i][j] // ค่าที่สูงสุดใน array
			// fmt.Println("L", input[i+1][j])                    // ค่าคู่ทางซ้าย
			// fmt.Println("R", input[i+1][j+1])                  // ค่าคู่ทางขวา
			sumsPathTotal := max(sums[i+1][j], sums[i+1][j+1]) // ต่าที่สูงสุดของผลรวมที่ทำการคำนวณก่อนหน้า
			sums[i][j] = greater + sumsPathTotal
			// fmt.Println("sums[i][j]", sums[i][j])
		}
	}

	// ผลรวมที่มากที่สุดจะอยู่ที่ sums[0][0]
	return sums[0][0]
}

// ฟังก์ชันสำหรับหาค่าที่มากที่สุด
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// ฟังก์ชันสำหรับรับข้อมูลจาก http
func getDataEx1(url string) [][]int {

	// สร้าง HTTP client
	client := &http.Client{}

	// สร้าง request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil
	}

	// ทำการ execute request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return nil
	}
	defer resp.Body.Close()

	// อ่านข้อมูลจาก response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil
	}

	// แปลงข้อมูลเป็น array int
	var input [][]int
	err = json.Unmarshal([]byte(string(body)), &input)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return nil
	}
	return input
}

// ฟังก์ชันสำหรับรัน
func Example1() {
	input := [][]int{{59}, {73, 41}, {52, 40, 53}, {26, 53, 6, 34}}
	output := findMaxPathSum(input)
	fmt.Println("ข้อ 1. Test Case(1) ผลรวมที่มากที่สุด:", output)

	input = getDataEx1("https://raw.githubusercontent.com/7-solutions/backend-challenge/main/files/hard.json")
	output = findMaxPathSum(input)
	fmt.Println("ข้อ 1. Test Case(2) ผลรวมที่มากที่สุด:", output)
}

// ======== Ex.1 ========

// <<<<< Ex.2 >>>>>>

// ฟังก์ชันสำหรับแปลงข้อความที่เข้ารหัสเป็นตัวเลขชุด
func decodeString(encoded string) string {
	// สร้าง slice เพื่อเก็บค่าตัวเลข

	numbers := []float64{0, 0, 0, 0, 0, 0}

	// // ตั้งค่าตัวเลขเริ่มต้น

	// // แปลงแต่ละตัวอักษรใน encoded เป็นตัวเลข
	for i, char := range encoded {
		if char == 'L' {
			numbers[i+1] = numbers[i] - 1
		} else if char == 'R' {
			numbers[i+1] = numbers[i] + 1
		} else if char == '=' { // char == '='
			numbers[i+1] = numbers[i]
		}
	}
	// fmt.Println("numbers ", numbers)

	// หาค่าตัวเลขที่น้อยที่สุด
	// minValue := math.Inf(1)
	// for _, item := range numbers {
	// 	if item < minValue {
	// 		minValue = item
	// 	}
	// }

	// // ถ้ามีค่าติดลบให้เพิ่มค่าให้ทุกตัวใน slice เพื่อทำให้เป็นบวกทั้งหมด
	// if minValue < 0 {
	// 	for i := range numbers {
	// 		numbers[i] += math.Abs(minValue)
	// 	}
	// }

	// fmt.Println("numbers after adjustment:", numbers)

	// สร้าง string จากตัวเลขที่ได้
	result := ""
	for _, num := range numbers {
		result += fmt.Sprintf("%d", int(num))
	}

	return result
}

// ฟังก์ชันสำหรับรัน
func Example2() {
	// สัญลักษณ์ “L” หมายความว่า ตัวเลขด้านซ้าย มีค่ามากกว่า ตัวเลขด้านขวา
	// สัญลักษณ์ “R” หมายความว่า ตัวเลขด้านขวา มีค่ามากกว่า ตัวเลขด้านซ้าย
	// สัญลักษณ์ “=“ หมายความว่า ตัวเลขด้านซ้าย มีค่าเท่ากับ ตัวเลขด้านขวา

	input1 := "LLRR="
	output1 := decodeString(input1)
	fmt.Println("ข้อ 2. Test Case(1) : input =", input1, "output =", output1)

	input1 = "==RLL"
	output1 = decodeString(input1)
	fmt.Println("ข้อ 2. Test Case(2) : input =", input1, "output =", output1)

	input1 = "=LLRR"
	output1 = decodeString(input1)
	fmt.Println("ข้อ 2. Test Case(3) : input =", input1, "output =", output1)
}

// ======== Ex.2 ========

// <<<<< Ex.3 >>>>>>

func countMeats(text string) map[string]int {
	meats := make(map[string]int)

	// แปลงทุกตัวอักษรให้เป็นตัวพิมพ์เล็กเพื่อไม่สนใจตัวหนา ตัวเอียง
	text = strings.ToLower(text)

	// แยกคำตาม space
	words := strings.Fields(text)

	for _, word := range words {
		// ลบ , . ออกจากคำ
		word = strings.Trim(word, ",.")
		meats[word]++
	}

	return meats
}

// ฟังก์ชันสำหรับรับข้อมูลจาก http
func getDataEx3(url string) string {

	// สร้าง HTTP client
	client := &http.Client{}

	// สร้าง request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return ""
	}

	// ทำการ execute request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return ""
	}
	defer resp.Body.Close()

	// อ่านข้อมูลจาก response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return ""
	}

	return string(body)
}

// ฟังก์ชันสำหรับเช็คข้อมูลใน array
func contains(arr []string, value string) bool {
	for _, item := range arr {
		if item == value {
			return true
		}
	}
	return false
}

// ฟังก์ชันสำหรับเพิ่มข้อมูล
func convertToDesiredFormat(input map[string]int) map[string]map[string]int {
	result := make(map[string]map[string]int)

	result["beef"] = make(map[string]int)
	// arr := []string{"t-bone", "fatback", "pastrami", "pork", "meatloaf", "jowl", "enim", "bresaola"}
	for key, value := range input {
		// if contains(arr, key) {
		result["beef"][key] = value
		// }
	}

	return result
}

// ฟังก์ชันสำหรับจัดโค้ดตามรูปแบบ json
func prettyprint(b []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "  ")
	return out.Bytes(), err
}

// ฟังก์ชันสำหรับรัน
func Example3() {
	text := getDataEx3("https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text")

	meatsCount := countMeats(text)

	fmt.Println("ข้อ 3. Test Case(1) รายชื่อเนื้อทั้งหมด และจำนวนของเนื้อแต่ละชนิด:")
	result := convertToDesiredFormat(meatsCount)
	jsonStr, err := json.Marshal(result)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	} else {
		b, _ := prettyprint(jsonStr)
		fmt.Printf("%s", b)
	}
}

// ======== Ex.3 ========
