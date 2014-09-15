package goimgur

//Message The base model for a message.
type Message struct {
	ID             int    `json:"id"`              //TheIDfor the message
	From           string `json:"from"`            //Account username of person sending the message
	AccountID      int    `json:"account_id"`      //The AccountIDof the person receiving the message
	SenderID       int    `json:"sender_id"`       //The AccountIDof the person who sent the message
	Body           string `json:"body"`            //Text of the message
	ConversationID int    `json:"conversation_id"` //ID for the overall conversation
	DateTime       int    `json:"datetime"`        //Time message was sent, epoch time
}
