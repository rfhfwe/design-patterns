package main

type Mediator interface {
	canArrive(Train) bool
	notifyAboutDeparture()
}

type StationManager struct {
	isPlatformFree bool
	trainQueue     []Train
}

func newStationManager() *StationManager {
	return &StationManager{
		isPlatformFree: true,
	}
}

func (s *StationManager) canArrive(t Train) bool {
	if s.isPlatformFree {
		s.isPlatformFree = false
		return true
	}
	s.trainQueue = append(s.trainQueue, t)
	return false
}

func (s *StationManager) notifyAboutDeparture() {
	if !s.isPlatformFree {
		s.isPlatformFree = true
	}
	if len(s.trainQueue) > 0 {
		firstTrainQueue := s.trainQueue[0]
		s.trainQueue = s.trainQueue[1:]
		firstTrainQueue.permitArrival()
	}
}
