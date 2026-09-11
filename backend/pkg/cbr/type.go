package cbr

import "encoding/xml"

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

type KRItem struct {
	Date string `xml:"DT" json:"date"`
	Rate string `xml:"Rate" json:"rate"`
}

type KeyRateResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		KeyRateXMLResponse struct {
			XMLName xml.Name `xml:"http://web.cbr.ru/ KeyRateXMLResponse"`
			Result  struct {
				KeyRate struct {
					Items []KRItem `xml:"KR"`
				} `xml:"KeyRate"`
			} `xml:"KeyRateXMLResult"`
		} `xml:"KeyRateXMLResponse"`
	} `xml:"Body"`
}

type RuoniaItem struct {
	Date string `xml:"D0" json:"date"`
	Rate string `xml:"ruo" json:"rate"`
}

type RuoniaResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		RuoniaXMLResponse struct {
			XMLName         xml.Name `xml:"http://web.cbr.ru/ RuoniaXMLResponse"`
			RuoniaXMLResult struct {
				Ruonia struct {
					Items []RuoniaItem `xml:"ro"`
				} `xml:"Ruonia"`
			} `xml:"RuoniaXMLResult"`
		} `xml:"RuoniaXMLResponse"`
	} `xml:"Body"`
}

type XmlPayloadType = string

const keyRate XmlPayloadType = "keyRate"
const ruoniaRate XmlPayloadType = "ruoniaRate"
