package goimgur

//Album This model represents the data for albums
type Album struct {
	ID          string   `json:"id"`           //The ID for the image
	Title       string   `json:"title"`        //The title of the image.
	Description string   `json:"description"`  //Description of the image.
	Datetime    int      `json:"datetime"`     //Time inserted into the gallery, epoch time
	Cover       string   `json:"cover"`        //The ID of the album cover image
	CoverWidth  int      `json:"cover_width"`  //The width, in pixels, of the album cover image
	CoverHeight int      `json:"cover_heigth"` //The height, in pixels, of the album cover image
	AccountURL  string   `json:"account_url"`  //The username of the account that uploaded it, or null.
	Privacy     string   `json:"privacy"`      //The privacy level of the album, you can only view public if not logged in as album owner
	Layout      string   `json:"layout"`       //The view layout of the album.
	Views       int      `json:"views"`        //The number of image views
	Link        string   `json:"link"`         //The direct link to the the image
	Favorite    bool     `json:"favorite"`     //Indicates if the current user favorited the image. Defaults to false if not signed in.
	Nsfw        bool     `json:"nsfw"`         //Indicates if the image has been marked as nsfw or not. Defaults to null if information is not available.
	Section     string   `json:"section"`      //If the image has been categorized by our backend then this will contain the section the image belongs in. (funny, cats, adviceanimals, wtf, etc)
	Order       int      `json:"order"`        //Order number of the album on the user's album page (defaults to 0 if their albums haven't been reordered)
	Deletehash  string   `json:"deletehash"`   //OPTIONAL, the deletehash, if you're logged in as the image owner
	ImagesCount int      `json:"images_count"` //The total number of images in the album
	Images      []Images `json:"images"`       //An slice of all the images in the album (only available when requesting the direct album)
}
