package imgur

//Conversation The base model for a conversation.
type Conversation struct {
	ID                 string    `json:"id"`                   //The ID for the Conversation
	LastMessagePreview string    `json:"last_message_preview"` //Preview of the last message
	DateTime           int       `json:"datetime"`             //Time inserted into the gallery, epoch time
	WithAccountId      int       `json:"with_account_id"`      //Account ID of the other user in conversation
	WithAccount        string    `json:"with_account"`         //Account username of the other user in conversation
	MessageCount       int       `json:"message_count"`        //Total number of messages in the conversation
	Messages           []Message `json:"messeges"`             //OPTIONAL: (only available when requesting a specific conversation) Reverse sorted such that most recent message is at the end of the array.
	Done               bool      `json:"done"`                 //OPTIONAL: (only available when requesting a specific conversation) Flag to indicate whether you've reached the beginning of the thread.
	Page               int       `json:"page"`                 //OPTIONAL: (only available when requesting a specific conversation) Flag to indicate whether you've reached the beginning of the thread.
}
