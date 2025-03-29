package services

func (s *service) CreateChannel() string {
	id := s.rooms.CreateChannel()

	return id
}
