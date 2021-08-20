package main

import (
	"fmt"
	"github.com/felipediazc/recluitment-exercise-golang/factory"
	"time"
)

const carsAmount = 100
const carsAmountPerSecond = 5

//var mutex = &sync.Mutex{}

func buildSingleVehicle(factory *factory.Factory) {
	vehicle := factory.StartAssemblingProcess(1)
	fmt.Println(vehicle.AssembleLog)
	fmt.Println(vehicle.TestingLog)
}

func buildVehicleBlockPerSecond(factory *factory.Factory) {
	for i := 0; i < carsAmount; i++ {
		go buildSingleVehicle(factory)
	}
	time.Sleep(1 * time.Second)
}

func main() {
	factory := factory.New()

	//Hint: change appropriately for making factory give each vehicle once assembled, even though the others have not been assembled yet,
	//each vehicle delivered to main should display testinglogs and assemblelogs with the respective vehicle id

	for i := 0; i < carsAmount; i++ {
		buildVehicleBlockPerSecond(factory)
	}
}
