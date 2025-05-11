package service

type ConfigureRequest struct {
	Rows        int `json:"rows" validate:"gte=1"`
	Cols        int `json:"cols" validate:"gte=1"`
	MinDistance int `json:"min_distance" validate:"gte=1"`
}

type Seat struct {
	Row int `json:"row" validate:"gte=0"`
	Col int `json:"col" validate:"gte=0"`
}

type GetAvailableSeatsRequest struct {
	GroupSize int `json:"group_size" validate:"required,min=1"`
}

type ReserveRequest struct {
	Seats []Seat `json:"seats" validate:"required,dive"`
}

type CancelRequest struct {
	Seats []Seat `json:"seats" validate:"required,dive"`
}
