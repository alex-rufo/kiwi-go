package kiwi

import (
	"github.com/google/go-querystring/query"
)

// Parameters represents all URI parameters that can be passed to the Kiwi API
type Parameters struct {
	FlyFrom                         string  `url:"flyFrom,omitempty"`
	To                              string  `url:"to,omitempty"`
	DateFrom                        string  `url:"dateFrom,omitempty"`
	DateTo                          string  `url:"dateTo,omitempty"`
	LongitudeFrom                   float32 `url:"longitudeFrom,omitempty"`
	LatitudeFrom                    float32 `url:"latitudeFrom,omitempty"`
	RadiusFrom                      int     `url:"radiusFrom,omitempty"`
	LongitudeTo                     float32 `url:"longitudeTo,omitempty"`
	LatitudeTo                      float32 `url:"latitudeTo,omitempty"`
	RadiusTo                        int     `url:"radiusTo,omitempty"`
	DaysInDestinationFrom           int     `url:"daysInDestinationFrom,omitempty"`
	DaysInDestinationTo             int     `url:"daysInDestinationTo,omitempty"`
	ReturnFrom                      string  `url:"returnFrom,omitempty"`
	ReturnTo                        string  `url:"returnTo,omitempty"`
	MaxFlyDuration                  int     `url:"maxFlyDuration,omitempty"`
	TypeFlight                      string  `url:"typeFlight,omitempty"`
	OneForCity                      bool    `url:"oneforcity,omitempty"`
	OnePerDate                      bool    `url:"one_per_date,omitempty"`
	Passengers                      int     `url:"passengers,omitempty"`
	Adults                          int     `url:"adults,omitempty"`
	Children                        int     `url:"children,omitempty"`
	Infants                         int     `url:"infants,omitempty"`
	FlyDays                         []int   `url:"flyDays,omitempty"`
	FlyDaysType                     string  `url:"flyDaysType,omitempty"`
	ReturnFlyDays                   []int   `url:"returnFlyDays,omitempty"`
	ReturnFlyDaysType               string  `url:"returnFlyDaysType,omitempty"`
	OnlyWorkingDays                 bool    `url:"onlyWorkingDays,omitempty"`
	OnlyWeekends                    bool    `url:"onlyWeekends,omitempty"`
	DirectFlights                   bool    `url:"directFlights,omitempty"`
	Partner                         string  `url:"partner,omitempty"`
	PartnerMarket                   string  `url:"partner_market,omitempty"`
	Currency                        string  `url:"curr,omitempty"`
	Locale                          string  `url:"locale,omitempty"`
	PriceFrom                       int     `url:"price_from,omitempty"`
	PriceTo                         int     `url:"price_to,omitempty"`
	DepartureTimeFrom               string  `url:"dtimefrom,omitempty"`
	DepartureTimeTo                 string  `url:"dtimeto,omitempty"`
	ArrivalTimeFrom                 string  `url:"atimefrom,omitempty"`
	ArrivalTimeTo                   string  `url:"atimeto,omitempty"`
	ReturnDepartureTimeFrom         string  `url:"returndtimefrom,omitempty"`
	ReturnDepartureTimeTo           string  `url:"returndtimeto,omitempty"`
	ReturnArrivalTimeFrom           string  `url:"returnatimefrom,omitempty"`
	ReturnArrivalTimeTo             string  `url:"returnatimeto,omitempty"`
	StopoverFrom                    string  `url:"stopoverfrom,omitempty"`
	StopoverTo                      string  `url:"stopoverto,omitempty"`
	MaxStopovers                    int     `url:"maxstopovers,omitempty"`
	ConnectionsOnDifferentAirport   int     `url:"connectionsOnDifferentAirport,omitempty"`
	ReturnFromDifferentAirport      int     `url:"returnFromDifferentAirport,omitempty"`
	ReturnToDifferentAirport        int     `url:"returnToDifferentAirport,omitempty"`
	InnerLimit                      int     `url:"innerLimit,omitempty"`
	SelectedAirlines                string  `url:"selectedAirlines,omitempty"`
	SelectedStopoverAirports        bool    `url:"selectedStopoverAirports,omitempty"`
	SelectedAirlinesExclude         bool    `url:"selectedAirlinesExclude,omitempty"`
	SelectedStopoverAirportsExclude string  `url:"selectedStopoverAirportsExclude,omitempty"`
	Offset                          int     `url:"offset,omitempty"`
	Limit                           int     `url:"limit,omitempty"`
	Sort                            string  `url:"sort,omitempty"`
	Asc                             int     `url:"asc,omitempty"`
	Version                         int     `url:"v,omitempty"`
	XML                             int     `url:"xml,omitempty"`
}

// Response has the information of different flights
type Response struct {
	SearchParameters map[string]string `json:"search_params"`
	Time             int               `json:"time"`
	Currency         string            `json:"currency"`
	CurrencyRate     float32           `json:"currency_rate"`
	Flights          []Flight          `json:"data"`
}

// Flight has the detailed information of a flight
type Flight struct {
	ID                  string            `json:"id"`
	MapIDFrom           string            `json:"mapIdfrom"`
	CityFrom            string            `json:"cityFrom"`
	CountryFrom         map[string]string `json:"countryFrom"`
	FlyFrom             string            `json:"flyFrom"`
	ReturnDuration      string            `json:"return_duration"`
	MapIDTo             string            `json:"mapIdto"`
	CityTo              string            `json:"cityTo"`
	CountryTo           map[string]string `json:"countryTo"`
	FlyTo               string            `json:"flyTo"`
	FlyDuration         string            `json:"fly_duration"`
	Conversion          map[string]int    `json:"conversion"`
	DeepLink            string            `json:"deep_link"`
	BookingToken        string            `json:"booking_token"`
	NightsInDestination int               `json:"nightsInDest"`
	BagLimit            map[string]int    `json:"baglimit"`
	BagsPrice           map[string]int    `json:"bags_price"`
	DepartureTimeUTC    int               `json:"dTimeUTC"`
	DepartureTime       int               `json:"dTime"`
	ArrivalTimeUTC      int               `json:"aTimeUTC"`
	ArrivalTime         int               `json:"aTime"`
	Distance            float32           `json:"distance"`
	Price               int               `json:"price"`
	TypeFlights         []string          `json:"type_flights"`
	Route               []Route           `json:"route"`
}

// Route has the information each of the flight during the whole trip
type Route struct {
	ID                  string  `json:"id"`
	Airline             string  `json:"airline"`
	FlightNumber        int     `json:"flight_no"`
	Price               int     `json:"price"`
	BagsRecheckRequired bool    `json:"bags_recheck_required"`
	DepartureTimeUTC    int     `json:"dTimeUTC"`
	DepartureTime       int     `json:"dTime"`
	ArrivalTimeUTC      int     `json:"aTimeUTC"`
	ArrivalTime         int     `json:"aTime"`
	MapIDFrom           string  `json:"mapIdfrom"`
	CityFrom            string  `json:"cityFrom"`
	FlyFrom             string  `json:"flyFrom"`
	LatitudeFrom        float32 `json:"latFrom"`
	LongitudeFrom       float32 `json:"lngFrom"`
	MapIDTo             string  `json:"mapIdto"`
	FlyTo               string  `json:"flyTo"`
	CityTo              string  `json:"cityTo"`
	LatitudeTo          float32 `json:"latTo"`
	LongitudeTo         float32 `json:"lngTo"`
}

// GetFlights calls the Kiwi API to retrieve all the flights that match the given parameters
func (k *Kiwi) GetFlights(p *Parameters) (*Response, error) {
	qs, err := query.Values(p)
	if err != nil {
		return nil, err
	}

	var r Response
	err = k.call("/flights?"+qs.Encode(), &r)

	return &r, err
}
