package factory

import (
	"testing"

	"github.com/stretchr/testify/suite"
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
	}
}
