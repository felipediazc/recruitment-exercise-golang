package factory

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type factoryUnitTestSuite struct {
	suite.Suite
	adapter *Factory
}

func (s *factoryUnitTestSuite) SetupSuite() {

	s.adapter = &Factory{}
}

func TestFactoryUnitTestSuite(t *testing.T) {
	suite.Run(t, &factoryUnitTestSuite{})
}

func (s *factoryUnitTestSuite) TestSamble() {
	//code here
	factory := New()
	vehicleList := factory.generateVehicleLots(5)
	i := 0
	//testing generateVehicleLots
	for _, vehicle := range vehicleList {
		s.Assert().Equal(i, vehicle.Id)
		i++
	}
	//testing AssembleVehicle
	vehicle := vehicleList[1]
	idleSpot := <-factory.AssemblingSpots
	idleSpot.SetVehicle(&vehicle)
	vehicleCar, err := idleSpot.AssembleVehicle()
	if err == nil {
		s.Assert().Equal(1, vehicleCar.Id)
		//testing logs
		vehicle.EngineStarted = true
		vehicle.TestingLog = factory.testCar(&vehicle)
		vehicle.AssembleLog = idleSpot.GetAssembledLogs()
		s.Assert().NotEqual("", vehicleCar.AssembleLog)
		s.Assert().NotEqual("", vehicleCar.TestingLog)
	}
}
