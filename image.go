package goimgur

//Image The base model for an image
type Image struct {
	ID          string `json:"id"`          //TheIDfor the image
	Title       string `json:"title"`       //The title of the image
	Description string `json:"description"` //Description of the image
	DateTime    int    `json:"datetime"`    //Time inserted into the gallery, epoch time
	Type        string `json:"type"`        //Image MIME type
	Animated    bool   `json:"animated"`    //Is the image animated
	Width       int    `json:"width"`       //The width of the image in pixels
	Height      int    `json:"height"`      //The height of the image in pixels
	Size        int    `json:"size"`        //The size of the image in bytes
	Views       int    `json:"views"`       //The number of image views
	Bandwidth   int    `json:"bandwidth"`   //Bandwidth consumed by the image in bytes
	Deletehash  string `json:"deletehash"`  //OPTIONAL, the deletehash, if you're logged in as the image owner
	Section     string `json:"section"`     //If the image has been categorized by our backend then this will contain the section the image belongs in. (funny, cats, adviceanimals, wtf, etc)
	Link        string `json:"link"`        //The direct link to the the image
	Favorite    bool   `json:"favorite"`    //Indicates if the current user favorited the image. Defaults to false if not signed in.
	Nsfw        bool   `json:"ns+ fw"`      //Indicates if the image has been marked as nsfw or not. Defaults to null if information is not available.
	Vote        string `json:"vote"`        //The current user's vote on the album. null if not signed in, if the user hasn't voted on it, or if not submitted to the gallery.
	AccountURL  string `json:"account_url"` //he username of the account that uploaded it, or null.
}
