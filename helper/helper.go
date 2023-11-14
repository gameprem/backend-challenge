package helper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
)

// Ex.1
func findMaxPathSum(input [][]int) int {
	rows := len(input)

	// สร้าง slice 2 มิติสำหรับเก็บผลรวมที่ได้ที่แต่ละจุด
	sums := make([][]int, rows)
	for i := range sums {
		sums[i] = make([]int, len(input[i]))
	}

	// กำหนดค่าเริ่มต้นในบริเวณที่มีเพียงตัวเดียว (ด้านฐาน)
	sums[rows-1] = input[rows-1]

	// เริ่มต้นจากต่ำไปสูง
	for i := rows - 2; i >= 0; i-- {
		for j := 0; j < len(input[i]); j++ {
			// หาผลรวมที่มากที่สุดของเส้นทางที่เป็นไปได้
			sums[i][j] = input[i][j] + max(sums[i+1][j], sums[i+1][j+1])
		}
	}

	// ผลรวมที่มากที่สุดจะอยู่ที่ sums[0][0]
	return sums[0][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getDataEx1(url string) [][]int {

	// สร้าง HTTP client
	client := &http.Client{}

	// สร้าง request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil
	}

	// สามารถเพิ่มการตั้งค่า TLS ได้ (ถ้าจำเป็น)

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
func Example1() {
	input := [][]int{{59}, {73, 41}, {52, 40, 53}, {26, 53, 6, 34}}
	output := findMaxPathSum(input)
	fmt.Println("ข้อ 1. Test Case(1) ผลรวมที่มากที่สุด:", output)

	input = getDataEx1("https://raw.githubusercontent.com/7-solutions/backend-challenge/main/files/hard.json")
	output = findMaxPathSum(input)
	fmt.Println("ข้อ 1. Test Case(2) ผลรวมที่มากที่สุด:", output)

}

// Ex.2

// ฟังก์ชันสำหรับแปลงข้อความที่เข้ารหัสเป็นตัวเลขชุด
func decodeString(encoded string) string {
	// สร้าง slice เพื่อเก็บค่าตัวเลข
	numbers := make([]int, len(encoded)+1)

	// ตั้งค่าตัวเลขเริ่มต้น
	numbers[0] = 2

	// แปลงแต่ละตัวอักษรใน encoded เป็นตัวเลข
	for i, char := range encoded {
		if char == 'L' {
			numbers[i+1] = numbers[i] + 1
		} else if char == 'R' {
			numbers[i+1] = numbers[i] - 1
		} else { // char == '='
			numbers[i+1] = numbers[i]
		}
	}

	// เรียงลำดับตัวเลข
	sort.Ints(numbers)

	// สร้างข้อความจากตัวเลขที่เรียงลำดับ
	decoded := ""
	for i := 1; i < len(numbers); i++ {
		decoded += fmt.Sprintf("%d", numbers[i]-numbers[i-1])
	}

	return decoded
}

func Example2() {
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

// Ex.3
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

func getDataEx3(url string) string {

	// สร้าง HTTP client
	client := &http.Client{}

	// สร้าง request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return ""
	}

	// สามารถเพิ่มการตั้งค่า TLS ได้ (ถ้าจำเป็น)

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

func contains(arr []string, value string) bool {
	for _, item := range arr {
		if item == value {
			return true
		}
	}
	return false
}

func convertToDesiredFormat(input map[string]int) map[string]map[string]int {
	result := make(map[string]map[string]int)

	result["beef"] = make(map[string]int)
	arr := []string{"t-bone", "fatback", "pastrami", "pork", "meatloaf", "jowl", "enim", "bresaola"}
	for key, value := range input {
		if contains(arr, key) {
			result["beef"][key] = value
		}
	}

	return result
}

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

func prettyprint(b []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "  ")
	return out.Bytes(), err
}
