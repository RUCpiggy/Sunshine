package structure

import "database/sql"

type Artwork struct {
	ArtworkID     int    `json:"artworkID"`
	Artist        string `json:"artist"`
	ImageFileName string `json:"imageFileName"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	YearOfWork    string `json:"yearOfWork"`
	Genre         string `json:"genre"`
	Width         string `json:"width"`
	Height        string `json:"height"`
	Price         string `json:"price"`
	View          int    `json:"view"`
	OwnerID       int    `json:"ownerID"`
	OrderID       sql.NullInt64
	TimeReleased  string
	Tocken        string `json:"tocken"`
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
