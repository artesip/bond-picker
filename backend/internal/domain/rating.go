package domain

import (
	"backend/pkg/cbr"
	"time"
)

type Rating struct {
	ID         UUID      `json:"id"`
	CompanyID  string    `json:"companyID"`
	Rating     string    `json:"ratingValue"`
	AgencyName string    `json:"agencyName"`
	ReleaseUrl string    `json:"releaseUrl" db:"url"`
	ObjectName string    `json:"objectName"`
	Date       time.Time `json:"releaseDate"`
}

func NewRating(rating cbr.Rating) (*Rating, error) {
	releaseDate, err := time.Parse("02.01.2006", rating.Date)

	if err != nil {
		return nil, err
	}

	return &Rating{
		CompanyID:  rating.Inn,
		Rating:     rating.Rating,
		AgencyName: rating.AgencyName,
		ReleaseUrl: rating.ReleaseUrl,
		ObjectName: rating.ObjectName,
		Date:       releaseDate,
	}, err
}
