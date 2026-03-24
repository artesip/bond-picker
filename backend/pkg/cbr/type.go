package cbr

type Rating struct {
	Inn        string `json:"inn"`
	Rating     string `json:"ratingValue"`
	AgencyName string `json:"kraName"`
	ReleaseUrl string `json:"releaseUrl"`
	ObjectName string `json:"objectName"`
	Date       string `json:"releaseDate"`
}

type PaginationAnswer struct {
	PageCount int      `json:"pageCount"`
	Items     []Rating `json:"itemList"`
}

type Errors struct {
	Message    string     `json:"message"`
	CustomData CustomData `json:"customData"`
}

type CustomData struct {
	Csrf string `json:"csrf"`
}

type RatingResponse struct {
	Status string           `json:"status"`
	Data   PaginationAnswer `json:"data"`
	Errors []Errors         `json:"errors"`
}
