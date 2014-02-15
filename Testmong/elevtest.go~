package main
import "driver"
import "fmt"
import "time"

const FLOORS = 4

func main () {
	driver.Init()
	fmt.Printf("Started!\n")
	go readButtons()
	go turnAround()
	go watchObs()
	time.Sleep(1*time.Second)
	elevdriver.MotorUp()
	for {
		select {
		case <-time.After(1*time.Second):
		}
	}
}

func readButtons () {
	var current [FLOORS][3] bool
	for {
		floor, dir := driver.GetButton()
		current[floor-1][dir] = !current[floor-1][dir]
		if current[floor-1][dir] {
			driver.SetLight(floor, dir)
		} else {
			driver.ClearLight(floor, dir)
		}
	}
}

func turnAround () {
	for {
		floor := driver.GetFloor()
		go driver.SetFloor(floor)
		switch floor {
		case 1:
			go driver.MotorUp()
		case 4:
			go driver.MotorDown()
		}
	}
}

func watchObs () {
	for {
		
		if driver.GetObs() {
			go driver.SetDoor()
		} else {
			go driver.ClearDoor()
		}
	}
}
