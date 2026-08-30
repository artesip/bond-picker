package domain

import (
	"backend/pkg/cbr"
	"time"
)

type Rating struct {
	ID         UUID      `json:"id" db:"id"`
	CompanyID  string    `json:"companyID" db:"company_id"`
	Rating     string    `json:"ratingValue" db:"rating"`
	AgencyName string    `json:"agencyName" db:"agency_name"`
	ReleaseUrl string    `json:"releaseUrl" db:"url"`
	ObjectName string    `json:"objectName" db:"object_name"`
	Date       time.Time `json:"releaseDate" db:"date"`
	IsRevoked  bool      `json:"isRevoked" db:"is_revoked"`
}

func NewRating(rating cbr.Rating) (*Rating, error) {
	releaseDate, err := time.Parse("02.01.2006", rating.Date)

	if err != nil {
		return nil, err
	}

	isRevoked := false
	resultRating := rating.Rating
	if rating.Rating == cbr.NoRating {
		isRevoked = true
		resultRating = ""
	}

	return &Rating{
		CompanyID:  rating.Inn,
		Rating:     resultRating,
		AgencyName: rating.AgencyName,
		ReleaseUrl: rating.ReleaseUrl,
		ObjectName: rating.ObjectName,
		Date:       releaseDate,
		IsRevoked:  isRevoked,
	}, err
}
