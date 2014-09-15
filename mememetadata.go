package goimgur

//MemeMetadata This model is used to represent the basic meme metadata.
type MemeMetadata struct {
	MemeName   string `json:"meme_name"`   //The name of the meme used.
	TopText    string `json:"top_text"`    //The top text of the meme.
	BottomText string `json:"bottom_text"` //The bottom text of the meme.
	BgImage    string `json:"bg_image"`    //The imageIDof the background image of the meme.
}
