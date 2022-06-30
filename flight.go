package kiwi

import (
	"github.com/google/go-querystring/query"
)

// Parameters represents all URI parameters that can be passed to the Kiwi API
type Parameters struct {
	FlyFrom                         string `url:"fly_from,omitempty"`
	FlyTo                           string `url:"fly_to,omitempty"`
	Partner                         string `url:"partner,omitempty"`
	PartnerMarket                   string `url:"partner_market,omitempty"`
	FlyDays                         []int  `url:"fly_days,omitempty"`
	FlyDaysType                     string `url:"fly_days_type,omitempty"`
	ReturnFlyDaysType               string `url:"ret_fly_days_type,omitempty"`
	DepartureFrom                   string `url:"depart_after,omitempty"`
	DepartureTo                     string `url:"depart_before,omitempty"`
	ArrivalFrom                     string `url:"arrive_after,omitempty"`
	ArrivalTo                       string `url:"arrive_before,omitempty"`
	ReturnDepartureFrom             string `url:"rt_depart_after,omitempty"`
	ReturnDepartureTo               string `url:"rt_depart_before,omitempty"`
	ReturnArrivalFrom               string `url:"rt_arrive_after,omitempty"`
	ReturnArrivalTo                 string `url:"rt_arrive_before,omitempty"`
	NightsInDestinationFrom         int    `url:"nights_in_dst_from,omitempty"`
	NightsInDestinationTo           int    `url:"nights_in_dst_to,omitempty"`
	DepartureTimeFrom               string `url:"dtime_from,omitempty"`
	ArrivalTimeFrom                 string `url:"atime_from,omitempty"`
	DepartureTimeTo                 string `url:"dtime_to,omitempty"`
	ArrivalTimeTo                   string `url:"atime_to,omitempty"`
	ReturnDepartureTimeFrom         string `url:"ret_dtime_from,omitempty"`
	ReturnArrivalTimeFrom           string `url:"ret_atime_from,omitempty"`
	ReturnDepartureTimeTo           string `url:"ret_dtime_to,omitempty"`
	ReturnArrivalTimeTo             string `url:"ret_atime_to,omitempty"`
	MaxFlyDuration                  int    `url:"max_fly_duration,omitempty"`
	VehicleType                     string `url:"vehicle_type,omitempty"`
	SelectedAirlines                string `url:"select_airlines,omitempty"`
	SelectedAirlinesExclude         bool   `url:"select_airlines_exclude,omitempty"`
	SelectedStopoverAirports        string `url:"select_stopover_airports,omitempty"`
	SelectedStopoverAirportsExclude bool   `url:"select_stopover_airports_exclude,omitempty"`
	StopoverFrom                    string `url:"stopover_from,omitempty"`
	StopoverTo                      string `url:"stopover_to,omitempty"`
	MaxStopovers                    int    `url:"max_stopovers,omitempty"`
	MaxSectorStopovers              int    `url:"max_sector_stopovers,omitempty"`
	ConnectionsOnDifferentAirport   int    `url:"conn_on_diff_airport,omitempty"`
	SelectedCabins                  string `url:"selected_cabins,omitempty"`
	MixWithCabins                   string `url:"selected_cabins,omitempty"`
	Adults                          int    `url:"adults,omitempty"`
	Children                        int    `url:"children,omitempty"`
	Infants                         int    `url:"infants,omitempty"`
	AdultHoldBag                    string `url:"adult_hold_bag,omitempty"`
	AdultHandBag                    string `url:"adult_hand_bag,omitempty"`
	ChildHoldBag                    string `url:"child_hold_bag,omitempty"`
	ChildHandBag                    string `url:"child_hand_bag,omitempty"`
	PriceFrom                       int    `url:"price_from,omitempty"`
	PriceTo                         int    `url:"price_to,omitempty"`
	ReturnFromDifferentAirport      bool   `url:"ret_from_diff_airport,omitempty"`
	ReturnToDifferentAirport        bool   `url:"ret_to_diff_airport,omitempty"`
	OneForCity                      bool   `url:"one_for_city,omitempty"`
	OnePerDate                      bool   `url:"one_per_date,omitempty"`
	Currency                        string `url:"curr,omitempty"`
	Locale                          string `url:"locale,omitempty"`
	Sort                            string `url:"sort,omitempty"`
	Limit                           int    `url:"limit,omitempty"`
	Asc                             bool   `url:"asc,omitempty"`
}

// Response has the information of different flights
type Response struct {
	SearchId string   `json:"search_id"`
	Currency string   `json:"currency"`
	FxRate   float64  `json:"fx_rate"`
	Flights  []Flight `json:"data"`
}

// Flight has the detailed information of a flight
type Flight struct {
	ID                  string             `json:"id"`
	FlyFrom             string             `json:"flyFrom"`
	FlyTo               string             `json:"flyTo"`
	CityFrom            string             `json:"cityFrom"`
	CityCodeFrom        string             `json:"cityCodeFrom"`
	CityCodeTo          string             `json:"cityCodeTo"`
	CityTo              string             `json:"cityTo"`
	CountryFrom         map[string]string  `json:"countryFrom"`
	CountryTo           map[string]string  `json:"countryTo"`
	MapIDFrom           string             `json:"mapIdfrom"`
	MapIDTo             string             `json:"mapIdto"`
	DepartureTimeUTC    int                `json:"dTimeUTC"`
	DepartureTime       int                `json:"dTime"`
	ArrivalTimeUTC      int                `json:"aTimeUTC"`
	ArrivalTime         int                `json:"aTime"`
	BagsPrice           map[string]float64 `json:"bags_price"`
	BagLimit            map[string]int     `json:"baglimit"`
	Price               float64            `json:"price"`
	Conversion          map[string]int     `json:"conversion"` //not as in the doc
	FlyDuration         string             `json:"fly_duration"`
	ReturnDuration      string             `json:"return_duration"`
	Duration            map[string]int     `json:"duration"`
	NightsInDestination int                `json:"nightsInDest"`
	Distance            float32            `json:"distance"`
	VirtualInterlining  bool               `json:"virtual_interlining"`
	ThrowAwayTicketing  bool               `json:"throw_away_ticketing"`
	HiddenCityTicketing bool               `json:"hidden_city_ticketing"`
	PnrCount            int                `json:"pnr_count"`
	HasAirportChange    bool               `json:"has_airport_changes"`
	TechnicalStops      int                `json:"technical_stops"`
	DeepLink            string             `json:"deep_link"`
	BookingToken        string             `json:"booking_token"`
	Route               []Route            `json:"route"`
}

// Route has the information each of the flight during the whole trip
type Route struct {
	ID                     string  `json:"id"`
	FlyFrom                string  `json:"flyFrom"`
	FlyTo                  string  `json:"flyTo"`
	CityFrom               string  `json:"cityFrom"`
	CityTo                 string  `json:"cityTo"`
	CityCodeFrom           string  `json:"cityCodeFrom"`
	CityCodeTo             string  `json:"cityCodeTo"`
	MapIDFrom              string  `json:"mapIdfrom"`
	MapIDTo                string  `json:"mapIdto"`
	LatitudeFrom           float32 `json:"latFrom"`
	LongitudeFrom          float32 `json:"lngFrom"`
	LatitudeTo             float32 `json:"latTo"`
	LongitudeTo            float32 `json:"lngTo"`
	StationFrom            string  `json:"stationFrom"`
	StationTo              string  `json:"stationTo"`
	DepartureTimeUTC       int     `json:"dTimeUTC"`
	DepartureTime          int     `json:"dTime"`
	ArrivalTimeUTC         int     `json:"aTimeUTC"`
	ArrivalTime            int     `json:"aTime"`
	VehicleType            string  `json:"vehicle_type"`
	Airline                string  `json:"airline"`
	FlightNumber           int     `json:"flight_no"`
	OperatingCarrier       string  `json:"operating_carrier"`
	OperatingFlightNo      string  `json:"operating_flight_no"`
	Equipment              string  `json:"equipment"`
	FareBasis              string  `json:"fare_basis"`
	FareCategory           string  `json:"fare_category"`
	FareFamily             string  `json:"fare_family"`
	FareClasses            string  `json:"fare_classes"`
	Return                 int     `json:"return"`
	BagsRecheckRequired    bool    `json:"bags_recheck_required"`
	Guarantee              bool    `json:"guarantee"`
	ViConnection           bool    `json:"vi_connection"`
	FollowingAirportChange bool    `json:"following_airport_change"`
	FollowingTechnicalStop bool    `json:"following_technical_stop"`
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
