package domain

type Company struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CompanyWithRating struct {
	Company
	Ratings []Rating `json:"ratings"`
}
