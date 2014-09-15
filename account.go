package goimgur

//Account This model is used to represent the basic account information.
type Account struct {
	ID            int    `json:"id"`             //The account id for the username requested.
	URL           string `json:"url"`            //The account username, will be the same as requested in the URL
	Bio           string `json:"bio"`            //A basic description the user has filled out
	Reputation    int    `json:"reputation"`     //The reputation for the account, in it's numerical format.
	Created       int    `json:"created"`        //The epoch time of account creation
	ProExperation int    `json:"pro_experation"` //False if not a pro user, their expiration date if they are.
}
