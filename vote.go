package goimgur

//Vote The base model for a vote.
type Vote struct {
	Ups   int `json:"ups"`   //Number of upvotes
	Downs int `json:"downs"` //Number of downvotes
}
