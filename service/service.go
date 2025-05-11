package service

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"sync"

	"github.com/go-playground/validator/v10"
)

var (
	validate            = validator.New()
	_defaultRows        = 5
	_defaultCols        = 5
	_defaultMinDistance = 6
)

type CinemaService interface {
	Configure(w http.ResponseWriter, r *http.Request)
	GetAvailableSeats(w http.ResponseWriter, r *http.Request)
	ReserveSeats(w http.ResponseWriter, r *http.Request)
	CancelSeats(w http.ResponseWriter, r *http.Request)
	ListSeats(w http.ResponseWriter, r *http.Request)
}

type cinemaServiceImpl struct {
	rows        int
	cols        int
	minDistance int
	seats       [][]int

	mu sync.Mutex
}

func NewCinemaService() CinemaService {
	cs := &cinemaServiceImpl{
		rows:        _defaultRows,
		cols:        _defaultCols,
		minDistance: _defaultMinDistance,
		seats:       make([][]int, _defaultRows),
	}
	for i := range cs.seats {
		cs.seats[i] = make([]int, _defaultCols)
	}

	return cs
}

func (cs *cinemaServiceImpl) Configure(w http.ResponseWriter, r *http.Request) {
	var req ConfigureRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if err := validate.Struct(req); err != nil {
		http.Error(w, fmt.Sprintf("Validation error %v", err), http.StatusBadRequest)
		return
	}

	cs.mu.Lock()
	defer cs.mu.Unlock()

	cs.rows = req.Rows
	cs.cols = req.Cols
	cs.minDistance = req.MinDistance
	cs.seats = make([][]int, req.Rows)
	for i := range cs.seats {
		cs.seats[i] = make([]int, req.Cols)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Configured"))
}

func (cs *cinemaServiceImpl) GetAvailableSeats(w http.ResponseWriter, r *http.Request) {
	var req GetAvailableSeatsRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	cs.mu.Lock()
	defer cs.mu.Unlock()

	groupSize := req.GroupSize
	availableSeats := [][]Seat{}

	for i := 0; i < cs.rows; i++ {
		for j := 0; j < cs.cols; j++ {
			if cs.seats[i][j] == 0 && cs.isValidGroup(i, j, groupSize) {
				group := []Seat{}
				for k := 0; k < groupSize; k++ {
					group = append(group, Seat{Row: i, Col: j + k})
				}
				availableSeats = append(availableSeats, group)
			}
		}
	}

	json.NewEncoder(w).Encode(availableSeats)
}

func (cs *cinemaServiceImpl) ReserveSeats(w http.ResponseWriter, r *http.Request) {
	var req ReserveRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	fmt.Printf("req %+v", req)
	if err := validate.Struct(req); err != nil {
		http.Error(w, fmt.Sprintf("Validation error %v", err), http.StatusBadRequest)
		return
	}

	cs.mu.Lock()
	defer cs.mu.Unlock()

	for _, seat := range req.Seats {
		if seat.Row < 0 || seat.Row >= cs.rows || seat.Col < 0 || seat.Col >= cs.cols {
			http.Error(w, "Invalid seat", http.StatusBadRequest)
			return
		}
		if cs.seats[seat.Row][seat.Col] == 1 {
			http.Error(w, "Seat already booked", http.StatusBadRequest)
			return
		}
		cs.seats[seat.Row][seat.Col] = 1
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Seats booked"))
}

func (cs *cinemaServiceImpl) CancelSeats(w http.ResponseWriter, r *http.Request) {
	var req CancelRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	if err := validate.Struct(req); err != nil {
		http.Error(w, fmt.Sprintf("Validation error %v", err), http.StatusBadRequest)
		return
	}

	cs.mu.Lock()
	defer cs.mu.Unlock()

	for _, seat := range req.Seats {
		if seat.Row < 0 || seat.Row >= cs.rows || seat.Col < 0 || seat.Col >= cs.cols {
			http.Error(w, "Invalid seat", http.StatusBadRequest)
			return
		}
		if cs.seats[seat.Row][seat.Col] == 0 {
			http.Error(w, "Seat not booked", http.StatusBadRequest)
			return
		}
		cs.seats[seat.Row][seat.Col] = 0
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Seats cancelled"))
}

func (cs *cinemaServiceImpl) ListSeats(w http.ResponseWriter, r *http.Request) {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	seats := make([][]int, cs.rows)
	for i := range seats {
		seats[i] = make([]int, cs.cols)
		copy(seats[i], cs.seats[i])
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(seats)
}

func (cs *cinemaServiceImpl) isValidGroup(row, col, groupSize int) bool {
	if col+groupSize > cs.cols {
		return false
	}

	for k := 0; k < groupSize; k++ {
		if cs.seats[row][col+k] == 1 {
			return false
		}
	}

	for i := 0; i < cs.rows; i++ {
		for j := 0; j < cs.cols; j++ {
			if cs.seats[i][j] == 1 {
				for k := 0; k < groupSize; k++ {
					if calculateDistance(Seat{Row: i, Col: j}, Seat{Row: row, Col: col + k}) < cs.minDistance {
						return false
					}
				}
			}
		}
	}

	return true
}

func calculateDistance(seat1, seat2 Seat) int {
	return int(math.Abs(float64(seat1.Row-seat2.Row)) + math.Abs(float64(seat1.Col-seat2.Col)))
}
