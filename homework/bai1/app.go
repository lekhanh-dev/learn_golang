package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type Product struct {
	Name     string
	Category string
	Price    int
}

func randomInt(min int, max int) int {
	return rand.Intn(max-min+1) + min
}

func main() {
	category := [4]string{"fashion", "electronics", "sport", "food"}
	products := [20]Product{}
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < len(products); i++ {
		products[i] = Product{
			fmt.Sprintf("%s %d", "Product", i),
			category[rand.Intn(len(category))],
			randomInt(100, 200),
		}
	}

	// Tim kiem theo ten san pham
	input := bufio.NewReader(os.Stdin)
	fmt.Print("Nhap ten san pham: ")
	productName, _ := input.ReadString('\n')
	productName = formatText(productName)

	for i := 0; i < len(products); i++ {
		if products[i].Name == productName {
			fmt.Println(products[i])
		}
	}

	// Tim kiem theo ten the loai
	fmt.Print("Nhap ten the loai: ")
	categoryName, _ := input.ReadString('\n')
	categoryName = formatText(categoryName)

	for i := 0; i < len(products); i++ {
		if products[i].Category == categoryName {
			fmt.Println(products[i])
		}
	}

	// Tim kiem theo khoang gia
	fmt.Print("Nhap gia thap nhat: ")
	priceMin, _ := input.ReadString('\n')
	priceMin = formatText(priceMin)
	x, err1 := strconv.ParseInt(priceMin, 10, 64)

	fmt.Print("Nhap gia cao nhat: ")
	priceMax, _ := input.ReadString('\n')
	priceMax = formatText(priceMax)
	y, err2 := strconv.ParseInt(priceMax, 10, 64)
	if err2 != nil {
		fmt.Println(err2, err1)
	}

	for i := 0; i < len(products); i++ {
		if int64(products[i].Price) >= x && int64(products[i].Price) <= y {
			fmt.Println(products[i])
		}
	}
}

func formatText(text string) string {
	if runtime.GOOS == "windows" {
		text = strings.TrimRight(text, "\r\n")
	} else {
		text = strings.TrimRight(text, "\n")
	}
	return text
}
