package kiwi

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestItReturnsAnErrorIfTheRequestFails(t *testing.T) {
	k := New()
	k.endpoint = "invalid url"

	_, err := k.GetFlights(&Parameters{FlyFrom: "BCN"})
	if err == nil {
		t.Errorf("An error should be thrown if the request is malformed")
	}
}

func TestItReturnsAnErrorIfTheStatusCodeIsNot200(t *testing.T) {
	k := New()
	close := mockAPI(k, http.StatusNotFound)
	defer close()

	_, err := k.GetFlights(&Parameters{FlyFrom: "BCN"})
	if err == nil {
		t.Errorf("An error should be thrown if the response of the API is not 200")
	}
}

func TestItGetsTheFlights(t *testing.T) {
	k := New()
	close := mockAPI(k, http.StatusOK)
	defer close()

	response, err := k.GetFlights(&Parameters{FlyFrom: "BCN"})
	if err != nil {
		t.Errorf("An error occurred when getting the flights: %s", err)
	}

	if len(response.Flights) == 0 {
		t.Errorf("No flights where returned")
	}
}

func BenchmarkGetFlights(b *testing.B) {
	k := New()
	close := mockAPI(k, http.StatusOK)
	defer close()

	params := Parameters{FlyFrom: "BCN"}
	for i := 0; i < b.N; i++ {
		k.GetFlights(&params)
	}
}

func mockAPI(k *Kiwi, statusCode int) func() {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(statusCode)
		w.Write([]byte(getValidResponse()))
	}))
	k.endpoint = ts.URL

	return ts.Close
}

func getValidResponse() string {
	return `{
		"data": [
		  {
			"mapIdfrom": "edinburgh",
			"return_duration": "10h 55m",
			"flyTo": "WMI",
			"conversion": {
			  "PLN": 1154,
			  "EUR": 271
			},
			"deep_link": "https://www.skypicker.com/deep?type=abc-m14.5&from=EDI&to=WMI&departure=19-12-2015&return=23-12-2015&flightsId=240466858|236557749|245258276|234220297&affilid=skyscanner_pl&price=271.5362&passengers=1&lang=pl&currency=pln&booking_token=NwyrOTUtH6wv4BdWZTvOrbnFTa5SOifCjnpGoAl4cKO2kB2HhT7afLhR/TDxb+eFu4YtIczyK0WbM+XCt51oGwDK7vohxDpNs2GE3fTG8FoBhxY59V7abVbbkEWcRIOXaJcmIZB1tq37mciGQWi4Sc29PEWEgFYVwxJYBIOJMNLYkB8FOhhUNMDmnaVNHDOgySMJd4it4r5fhopJczkKnJIKoMFEelHE0x5203/3dHkfA7iBk4/c1OYTA61XBIAlsCuTH4D0j0IHASTFV4VXT36kZS9oxvJuya4nyDoff2017Up/UHkaXQbNBAoD5W0A",
			"mapIdto": "warsaw",
			"nightsInDest": 3,
			"id": "240466858|236557749|245258276|234220297",
			"fly_duration": "24h 10m",
			"countryTo": {
			  "code": "PL",
			  "name": "Rzeczpospolita Polska"
			},
			"baglimit": {
			  "hand_width": 40,
			  "hand_length": 55
			},
			"aTimeUTC": 1450643100,
			"distance": 1624.4,
			"price": 1154,
			"type_flights": [
			  "lcc"
			],
			"bags_price": {
			  "1": 100
			},
			"cityTo": "Warszawa",
			"flyFrom": "EDI",
			"dTimeUTC": 1450556100,
			"p2": 232,
			"countryFrom": {
			  "code": "GB",
			  "name": "Zjednoczone Królestwo Wielkiej Brytanii"
			},
			"p1": 232,
			"dTime": 1450556100,
			"booking_token": "NwyrOTUtH6wv4BdWZTvOrbnFTa5SOifCjnpGoAl4cKO2kB2HhT7afLhR/TDxb+eFu4YtIczyK0WbM+XCt51oGwDK7vohxDpNs2GE3fTG8FoBhxY59V7abVbbkEWcRIOXaJcmIZB1tq37mciGQWi4Sc29PEWEgFYVwxJYBIOJMNLYkB8FOhhUNMDmnaVNHDOgySMJd4it4r5fhopJczkKnJIKoMFEelHE0x5203/3dHkfA7iBk4/c1OYTA61XBIAlsCuTH4D0j0IHASTFV4VXT36kZS9oxvJuya4nyDoff2017Up/UHkaXQbNBAoD5W0A",
			"cityFrom": "Edynburg",
			"aTime": 1450646700,
			"route": [
			  {
				"bags_recheck_required": false,
				"aTimeUTC": 1450560300,
				"mapIdfrom": "edinburgh",
				"mapIdto": "dublin",
				"flight_no": 819,
				"dTime": 1450556100,
				"latTo": 53.3331,
				"flyTo": "DUB",
				"return": 0,
				"source": null,
				"id": "240466858",
				"airline": "FR",
				"lngTo": -6.24889,
				"cityTo": "Dublin",
				"cityFrom": "Edynburg",
				"lngFrom": -3.1936633587,
				"aTime": 1450560300,
				"flyFrom": "EDI",
				"price": 21,
				"latFrom": 55.9485946759,
				"dTimeUTC": 1450556100
			  }
			]
		  }
		]
	  }`
}
