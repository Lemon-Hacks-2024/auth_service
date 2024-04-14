package handler

import (
	"bufio"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"strings"
)

func (h *Handler) TestFunc(ctx *gin.Context) {
	results, err := ReadFile()
	if err != nil {
		log.Fatalln(err)
	}

	for _, result := range results {
		//	Разбирать строку
		splitted := strings.Split(result, "\t")
		city := splitted[0] + " " + splitted[1] + ", " + splitted[2] + " " + splitted[3]
		log.Println(city)

		err := h.services.City.AddCity(city)
		if err != nil {
			log.Println(err)
		}
	}
}

func ReadFile() ([]string, error) {
	var data []string

	file, err := os.Open("city.txt")
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
