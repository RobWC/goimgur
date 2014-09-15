package goimgur

//BlockedUsers users that are blocked
type BlockedUsers struct {
	BlockedID  int `json:"blocked_id"`  //Blocked user ID
	BlockedURL int `json:"blocked_url"` //Blocked UrL
}

//AccountSettings The account settings, only accessible if you're logged in as the user.
type AccountSettings struct {
	Email                string         `json:"email"`                  //The users email address
	HighQuality          bool           `json:"high_quality"`           //The users ability to upload higher quality images, there will be less compression
	PublicImages         bool           `json:"public_images"`          //Automatically allow all images to be publicly accessible
	AlbumPrivacy         string         `json:"album_privacy"`          //Set the album privacy to this privacy setting on creation
	ProExpiration        int            `json:"pro_experation"`         //False if not a pro user, their expiration date if they are.
	AcceptedGalleryTerms bool           `json:"accepted_gallery_terms"` //True if the user has accepted the terms of uploading to the Imgur gallery.
	ActiveEmails         []string       `json:"active_emails"`          //The email addresses that have been activated to allow uploading
	MessagingEnabled     bool           `json:"messaging_enabled"`      //If the user is accepting incoming messages or not
	BlockedUsers         []BlockedUsers `json:"blocked_users"`          //An array of users that have been blocked from messaging, the object is blocked_id and blocked_url.
}
