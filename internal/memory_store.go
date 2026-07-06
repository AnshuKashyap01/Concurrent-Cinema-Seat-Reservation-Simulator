package booking


type MemoryStore struct {
	bookings map[string]Booking
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		bookings: map[string]Booking{}, //Yeh bookings field ko empty map se initialize kar raha hai.
	}
}

// Methods of struct MemoryStore
func (s *MemoryStore) Book(b Booking) error {
	//seat is taken
	//if not then populate

	if _, exists := s.bookings[b.SeatID]; exists {
		return ErrSeatAlreadyBooked
	}

	s.bookings[b.SeatID] = b
	
	return nil

}

func (s *MemoryStore) ListBookings(movieID string) []Booking{
	var result []Booking

	for _, b := range s.bookings {
		if b.MovieID == movieID{
			result = append(result, b)
		}
	}

	return result
}
