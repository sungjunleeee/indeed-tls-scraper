package main

import (
	"os"
	"strings"

	"github.com/labstack/echo"
	"github.com/sungjunleeee/learngo/scraper"
)

const filename string = "jobs.csv"

func main() {
	e := echo.New()

	// Routes
	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
func handleHome(c echo.Context) error {
	return c.File("home.html")
}

func handleScrape(c echo.Context) error {
	defer os.Remove(filename)
	job := strings.ToLower(scraper.CleanString(c.FormValue("job")))
	scraper.Scrape(job)
	return c.Attachment(filename, job+"_"+filename)
}
