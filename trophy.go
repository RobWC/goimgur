package goimgur

//Trophy The data model for a trophy
type Trophy struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	NameClean   string `json:"name_clean"`
	Description string `json:"description"`
	Data        string `json:"data"`
	DataLink    string `json:"data_link"`
	Datatime    int    `json:"datetime"`
}
