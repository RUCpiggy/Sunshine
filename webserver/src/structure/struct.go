package structure

import "database/sql"

type Artwork struct {
	ArtworkID     int
	Artist        string
	ImageFileName string
	Title         string
	Description   string
	YearOfWork    int
	Genre         string
	Width         float32
	Height        float32
	Price         int
	View          int
	OwnerID       int
	OrderID       sql.NullInt64
	TimeReleased  string
}
type Cart struct {
	CartID    int
	UserID    int
	ArtworkID int
}
type Order struct {
	OrderID     int
	OwnerID     int
	Sum         int
	TimeCreated string
}
type User struct {
	UserID   int    `json:"userID"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Tel      string `json:"tel"`
	Address  string `json:"address"`
	Balance  int    `json:"balance"`
}
