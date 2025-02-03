package test3

import (
	"io"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func Run() {
	app := fiber.New()

	app.Get("/beef/summary", func(c *fiber.Ctx) error {
		beef := newBeefSum("https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text")
		summary := beef.GetSummary()
		return c.Status(200).JSON(Response{Beef: summary})
	})

	app.Listen(":3000")
}

type Response struct {
	Beef map[string]uint `json:"beef"`
}

type Beef struct {
	url string
}

func newBeefSum(url string) *Beef {
	return &Beef{url: url}
}

func (b *Beef) getRawData() string {
	response, err := http.Get(b.url)
	if err != nil {
		log.Fatalf("Error making GET request: %v", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	return string(body)
}

func (b *Beef) GetSummary() map[string]uint {
	rawData := b.getRawData()
	temp := ""
	summary := make(map[string]uint)
	for _, char := range rawData {
		if char == ',' || char == '.' || char == ' ' || char == '\t' || char == '\n' {
			if len(temp) > 0 {
				summary[temp] += 1
				temp = ""
			}
		} else {
			temp += string(char)
		}
	}
	return summary
}
