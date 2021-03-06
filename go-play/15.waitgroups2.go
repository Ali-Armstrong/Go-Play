package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"
)
type City struct {
	Name       string
	Population int
	Households int
	Temp       int
	Forecast   string
	Elevation  int
	Lat        float64
	Long       float64
}
// Main function Synchronous
// func main() {
// 	defer timeTrack(time.Now(), "Fetching city info")

// 	fmt.Println("Starting synchronous calls...")

// 	city := City{Name: "Richmond"}

// 	getDemoInfo(&city)
// 	getCurrentWeather(&city)
// 	getMapInfo(&city)

// 	data, _ := json.Marshal(city)
// 	fmt.Printf("%s\n", data)
// }

func main() {
	defer timeTrack(time.Now(), "Fetching city info")
	fmt.Println("Starting concurrent calls...")

	var waitgroup sync.WaitGroup
	waitgroup.Add(3)
	
	city := City{Name: "Richmond"}

	go func() {
		getDemoInfo(&city)
		waitgroup.Done()
	}()

	go func() {
		getCurrentWeather(&city)
		waitgroup.Done()
	}()

	go func() {
		getMapInfo(&city)
		waitgroup.Done()
	}()

	waitgroup.Wait()
	
	data, _ := json.Marshal(city)
	fmt.Printf("%s\n", data)
}

func getDemoInfo(c *City) {
	// Calling Sleep method (lag)
	time.Sleep(2100 * time.Millisecond)
	c.Population = 230436
	c.Households = 90301
}

func getCurrentWeather(c *City) {
	// Calling Sleep method (lag)
	time.Sleep(1200 * time.Millisecond)
	
	c.Temp = 32
	c.Forecast = "Snow, turning to rain"
}

func getMapInfo(c *City) {
	// Calling Sleep method (lag)
	time.Sleep(3300 * time.Millisecond)
	c.Elevation = 457
	c.Lat = 37.533333333
	c.Long = -77.466666666
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}