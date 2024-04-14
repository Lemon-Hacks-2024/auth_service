package handler

import (
	"bufio"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func (h *Handler) TestFunc(ctx *gin.Context) {
	results, err := ReadFile()
	if err != nil {
		log.Fatalln(err)
	}

	for _, result := range results {
		log.Println(result)

		//	Запись в базу
		err := h.services.StoreAddresses.AddStoreAddress(result)
		if err != nil {
			log.Println(err)
		}
	}
}

func ReadFile() ([]string, error) {
	var data []string

	file, err := os.Open("data.txt")
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	file.Close()

	return data, nil
}
