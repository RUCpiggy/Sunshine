package sqlite

import (
	"database/sql"
	"structure"
)

func AddArt(db *sql.DB, args ...interface{}) {
	sql := "Insert INTO artworks(artworkID,artist,imageFileName,title,description,yearOfWork,genre,width,height,price,view,ownerID,orderID,timeReleased) value(?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	Insert(db, sql, args...)
}

func UpdateArtInfo(db *sql.DB, artworkID string, info string, value string) {
	sql := ""
	switch info {
	case "artist":
		sql = "UPDATE artworks SET artist=? WHERE artworkID=?"
	case "inageFileName":
		sql = "UPDATE artworks SET inageFileName=? WHERE artworkID=?"
	case "title":
		sql = "UPDATE artworks SET title=? WHERE artworkID=?"
	case "description":
		sql = "UPDATE artworks SET description=? WHERE artworkID=?"
	case "yearOfWork":
		sql = "UPDATE artworks SET yearOfWork=? WHERE artworkID=?"
	case "genre":
		sql = "UPDATE artworks SET genre=? WHERE artworkID=?"
	case "width":
		sql = "UPDATE artworks SET width=? WHERE artworkID=?"
	case "height":
		sql = "UPDATE artworks SET height=? WHERE artworkID=?"
	case "price":
		sql = "UPDATE artworks SET price=? WHERE artworkID=?"
	case "view":
		sql = "UPDATE artworks SET view=? WHERE artworkID=?"
	case "ownerID":
		sql = "UPDATE artworks SET ownerID=? WHERE artworkID=?"
	case "orderID":
		sql = "UPDATE artworks SET orderID=? WHERE artworkID=?"
	case "timeReleased":
		sql = "UPDATE artworks SET timeReleased=? WHERE artworkID=?"
	}
	Update(db, sql, value, artworkID)
}

func DeleteArt(db *sql.DB, artworkID string) {
	sql := "DELETE FROM artworks WHERE artworkID=?"
	Delete(db, sql, artworkID)
}

func SearchArt(db *sql.DB, judge string) ([]structure.Artwork, bool) {

	sql := "SELECT * FROM artworks WHERE " + judge
	rows, err := db.Query(sql)
	checkErr(err)

	//map_columns := make(map[string]string)
	var artworks []structure.Artwork

	for rows.Next() {
		var artwork structure.Artwork

		err = rows.Scan(&artwork.ArtworkID, &artwork.Artist, &artwork.ImageFileName, &artwork.Title, &artwork.Description, &artwork.YearOfWork,
			&artwork.Genre, &artwork.Width, &artwork.Height, &artwork.Price, &artwork.View, &artwork.OwnerID, &artwork.OrderID, &artwork.TimeReleased)

		artworks = append(artworks, artwork)
		//println(artwork.ArtworkID)
	}
	if len(artworks) != 0 {
		return artworks, true
	} else {
		return artworks, false
	}

}
