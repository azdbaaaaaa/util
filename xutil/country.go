package xutil

type Country string

const (
	CountryNG    = "NG"
	CountryKE    = "KE"
	CountryOther = "OTHER"
)

type CountryDetail struct {
	FullCountry string
	Country     string
	Currency    string
}

var CountryMap = map[Country]CountryDetail{
	CountryNG: {
		Country:     CountryNG,
		FullCountry: "Nigeria",
		Currency:    "Naira",
	},
	CountryKE: {
		Country:     CountryKE,
		FullCountry: "Kenya",
		Currency:    "KSH",
	},
	CountryOther: {
		Country:     CountryOther,
		FullCountry: "Other",
		Currency:    "USD",
	},
}

func (c Country) Get() CountryDetail {
	if v, ok := CountryMap[c]; ok {
		return v
	}
	return CountryMap[CountryOther]
}
