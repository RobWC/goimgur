package goimgur

//Basic This is the basic response for requests that do not return data. If the POST request has a Basic model it will return the id.
type Basic struct {
	Data    string `json:"data"`    //Is null, boolean, or integer value. If it's a post then this will contain an object with the all generated values, such as an ID.
	Success bool   `json:"success"` //Was the request successful
	Status  int    `json:"status"`  //HTTP Status Code
}
