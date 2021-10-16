package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Product struct {
	id          string
	title       string
	description string
	price       float64
}

func (prod *Product) store() {
	file, _ := os.Create(prod.id + ".txt")

	content := fmt.Sprintf(
		"ID: %v\nTitle: %v\nDescription: %v\nPrice: %v",
		prod.id,
		prod.title,
		prod.description,
		prod.price,
	)

	file.WriteString(content)
	file.Close()
}

func (prod *Product) printData() {
	fmt.Println(*prod)
}

func NewProduct(id string, title string, description string, price float64) *Product {
	product := Product{
		id,
		title,
		description,
		price,
	}

	return &product
}

func main() {
	createdProduct := getProduct()

	createdProduct.printData()
	createdProduct.store()
}

func getProduct() *Product {
	fmt.Println("Enter product data.")
	fmt.Println("-------------------")

	reader := bufio.NewReader(os.Stdin)

	idInput := readUserInput(reader, "Product ID: ")
	titleInput := readUserInput(reader, "Product Title: ")
	descriptionInput := readUserInput(reader, "Product Description: ")
	priceInput := readUserInput(reader, "Product Price: ")

	priceValue, _ := strconv.ParseFloat(priceInput, 64)

	product := NewProduct(idInput, titleInput, descriptionInput, priceValue)
	return product
}

func readUserInput(reader *bufio.Reader, promptText string) string {
	fmt.Print(promptText)
	userInput, _ := reader.ReadString('\n')
	userInput = strings.Replace(userInput, "\r\n", "", -2)

	return userInput
}
