package goimgur

//Gallery An interface for gallery model types
type Gallery interface {
}

//GalleryImages A struct containg a slice of GalleryImage contained within Data
type GalleryImages struct {
	Data []GalleryImage `json:"data"`
}

//GalleryImage The data model formatted for gallery images
type GalleryImage struct {
	ID          string `json:"id"`          //TheIDfor the image
	Title       string `json:"title"`       //The title of the image.
	Description string `json:"description"` //Description of the image.
	DateTime    int    `json:"datetime"`    //Time inserted into the gallery, epoch time
	Type        string `json:"type"`        //Image MIME type.
	Animated    bool   `json:"animated"`    //Is the image animated
	Width       int    `json:"width"`       //The width of the image in pixels
	Height      int    `json:"height"`      //The height of the image in pixels
	Size        int    `json:"size"`        //The size of the image in bytes
	Views       int    `json:"views"`       //The number of image views
	Bandwidth   int    `json:"bandwidth"`   //Bandwidth consumed by the image in bytes
	Deletehash  string `json:"deletehash"`  //OPTIONAL, the deletehash, if you're logged in as the image owner
	Link        string `json:"link"`        //The direct link to the the image
	Vote        string `json:"vote"`        //The current user's vote on the album. null if not signed in or if the user hasn't voted on it.
	Favorite    bool   `json:"favorite"`    //Indicates if the current user favorited the image. Defaults to false if not signed in.
	Nsfw        bool   `json:"nsfw"`        //Indicates if the image has been marked as nsfw or not. Defaults to null if information is not available.
	Section     string `json:"section"`     //If the image has been categorized by our backend then this will contain the section the image belongs in. (funny, cats, adviceanimals, wtf, etc)
	AccountURL  string `json:"account_url"` //The username of the account that uploaded it, or null.
	AccountID   int    `json:"account_id"`  //The AccountID of the account that uploaded it, or null.
	Ups         int    `json:"ups"`         //Upvotes for the image
	Downs       int    `json:"downs"`       //Number of downvotes for the image
	Score       int    `json:"score"`       //Imgur popularity score
	IsAlbum     bool   `json:"is_album"`    //If it's an album or not
}

//GalleryProfile The totals for a users gallery information
type GalleryProfile struct {
	TotalGalleryComments    int      `json:"total_gallery_comments"`
	TotalGalleryFavorites   int      `json:"total_gallery_favorites"`
	TotalGallerySubmissions int      `json:"total_galley_submissions"`
	Trophies                []Trophy `json:"trophies"`
}
