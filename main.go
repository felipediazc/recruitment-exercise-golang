package main

import (
	"fmt"
	"github.com/felipediazc/recluitment-exercise-golang/factory"
	"github.com/felipediazc/recluitment-exercise-golang/vehicle"
	"time"
)

const carsAmount = 100

func printVehicleData(blockNumber int, carNumber int, ch chan *vehicle.Car) {
	vehicle := <-ch
	fmt.Println("BLOCK: ", blockNumber, "CAR ", carNumber, "CAR_ID", vehicle.Id)
	fmt.Println("AssembleLog", vehicle.AssembleLog)
	fmt.Println("TestingLog", vehicle.TestingLog)
}

func main() {

	//Hint: change appropriately for making factory give each vehicle once assembled, even though the others have not been assembled yet,
	//each vehicle delivered to main should display testinglogs and assemblelogs with the respective vehicle id

	ch := make(chan *vehicle.Car, 5)

	for i := 0; i < carsAmount; i++ {
		factory := factory.New()
		go factory.StartAssemblingProcess(1, ch)
		go factory.StartAssemblingProcess(1, ch)
		go factory.StartAssemblingProcess(1, ch)
		go factory.StartAssemblingProcess(1, ch)
		go factory.StartAssemblingProcess(1, ch)

		time.Sleep(7 * time.Second)
		printVehicleData(i, 1, ch)
		printVehicleData(i, 2, ch)
		printVehicleData(i, 3, ch)
		printVehicleData(i, 4, ch)
		printVehicleData(i, 5, ch)
	}

}
