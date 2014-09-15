package goimgur

//Notification The base model for a notification.
type Notification struct {
	ID        int        `json:"id"`         //TheIDfor the notification
	AccountID int        `json:"account_id"` //The AccountIDfor the notification
	Viewed    bool       `json:"viewed"`     //Has the user viewed the image yet?
	Content   ImgurModel `json:"content"`    //This can be any other model, currently only using comments and conversation metadata.
}
