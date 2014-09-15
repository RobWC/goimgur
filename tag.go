package goimgur

//Tag The base model for a tag
type Tag struct {
	Name       string          `json:"name"`        //Name of the tag.
	Followers  int             `json:"followers"`   //Number of followers for the tag.
	TotalItems int             `json:"total_items"` //Total number of gallery items for the tag.
	Following  []FollowingData //Array of objects with ID/name values for the custom galleries in which the current user is following the requested tag. Null if not logged in.
	Items      []Gallery       `json:"items"` //
}

//TagVote The base model for a tag vote.
type TagVote struct {
	Ups    int    `json:"ups"`    //Number of upvotes.
	Downs  int    `json:"downs"`  //Number of downvotes.
	Name   string `json:"name"`   //Name of the tag.
	Author string `json:"author"` //Author of the tag.
}
