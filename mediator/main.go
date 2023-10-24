package main

var (
	stationManager *StationManager
	passengerTrain *PassengerTrain
	freightTrain   *FreightTrain
)

func Init() {
	stationManager = newStationManager()
	passengerTrain = &PassengerTrain{
		mediator: stationManager,
	}
	freightTrain = &FreightTrain{
		mediator: stationManager,
	}
}

func main() {
	Init()
	passengerTrain.arrive()
	freightTrain.arrive()
	passengerTrain.depart()

}
