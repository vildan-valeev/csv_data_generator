package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz"
const countRows = 500000
const outfile = "csv_file.csv"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func randomString(length int) string {
	return StringWithCharset(length, charset)
}

func randomEmail() string {
	emails := [3]string{"@mail.ru", "@gmail.com", "@yandex.ru"}
	rand.Seed(time.Now().UnixNano())
	randIdx := rand.Intn(len(emails))
	return emails[randIdx]
}

func randomDate() string {
	min := time.Date(1980, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Now().Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0).Format("2006-01-02 15:04:05")
}

func genRows() [][]string {
	var rows [][]string
	for i := 1; i < countRows; i++ {
		row := []string{randomString(10), randomString(10), randomDate(), randomString(5) + randomEmail()}
		rows = append(rows, row)
	}
	return rows

}

func saveCsv() {
	var rows [][]string
	rows = genRows()

	f, err := os.Create(outfile)
	defer f.Close()

	if err != nil {

		log.Fatalln("failed to open file", err)
	}

	w := csv.NewWriter(f)
	err = w.WriteAll(rows) // calls Flush internally

	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	start := time.Now()
	saveCsv()
	duration := time.Since(start)
	fmt.Println(duration)
}
