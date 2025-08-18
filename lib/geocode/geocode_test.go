/**
 * go-mapbox Geocoding Module Tests
 * Wraps the mapbox geocoding API for server side use
 * See https://www.mapbox.com/api-documentation/#geocoding for API information
 *
 * https://github.com/ryankurte/go-mapbox
 * Copyright 2017 Ryan Kurte
 */

package geocode

import (
	"os"
	"testing"

	"github.com/icep1anet/go-mapbox/lib/base"
)

func TestGeocoder(t *testing.T) {

	b, err := base.NewBase(os.Getenv("MAPBOX_TOKEN"))
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	geocode := NewGeocode(b)

	t.Run("Can geocode", func(t *testing.T) {
		var reqOpt ForwardRequestOpts
		reqOpt.Limit = 1

		reqOpt.Place = "2 lincoln memorial circle nw"

		res, err := geocode.Forward(&reqOpt)
		if err != nil {
			t.Error(err)
		}

		if res.Type != "FeatureCollection" {
			t.Errorf("Invalid response type: %s", res.Type)
		}

	})

	t.Run("Can reverse geocode", func(t *testing.T) {
		var reqOpt ReverseRequestOpts
		reqOpt.Limit = 1

		reqOpt.Latitude = 34.074122
		reqOpt.Longitude = 72.438939

		res, err := geocode.Reverse(&reqOpt)
		if err != nil {
			t.Error(err)
		}

		if res.Type != "FeatureCollection" {
			t.Errorf("Invalid response type: %s", res.Type)
		}

	})

}
