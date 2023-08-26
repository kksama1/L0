package moodel

type Item struct {
	ChrtId      int    `json:"chrt_id" validate:"required, unique"`
	TrackNumber string `json:"track_number" validate:"required"`
	Price       int    `json:"price" validate:"required"`
	Rid         string `json:"rid" validate:"required, unique"`
	Name        string `json:"name" validate:"required"`
	Sale        int    `json:"sale"`
	Size        string `json:"size" validate:"required"`
	TotalPrice  int    `json:"total_price" validate:"required"`
	NmId        int    `json:"nm_id" validate:"required, unique"`
	Brand       string `json:"brand"`
	Status      int    `json:"status" validate:"required"`
}
