package goimgur

//FollowingData ID and name or the custom galleries in which the current user is following the requested tag.
type FollowingData struct {
	ID   string `json:"id"`   //TheIDof the following data
	Name string `json:"name"` //The name of the following data
}
