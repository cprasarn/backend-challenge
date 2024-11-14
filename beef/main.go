package main

import (
	"io"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type BeefData struct {
	Beef map[string]int `json:"beef"`
}

func GetBeefData() (*BeefData, error) {
	response, err := http.Get("https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text")
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	result := BeefData{}
	result.Beef = make(map[string]int)
	data := strings.ReplaceAll(string(body), ".", "")
	data = strings.ReplaceAll(data, ",", "")
	data = strings.ReplaceAll(data, "\n", " ")
	words := strings.Split(strings.ToLower(data), " ")
	for _, w := range words {
		if w != "" {
			result.Beef[w]++
		}
	}

	return &result, nil
}

func GetBeefSummary(c *fiber.Ctx) error {
	result, err := GetBeefData()
	if err != nil {
		return err
	}
	return c.JSON(result)
}

func main() {
	app := fiber.New()
	app.Get("/beef/summary", GetBeefSummary)
	app.Listen(":8080")
}
