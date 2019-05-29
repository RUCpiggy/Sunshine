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
	UserID   int
	Name     string
	Email    string
	Password string
	Tel      string
	Address  string
	Balance  int
}
