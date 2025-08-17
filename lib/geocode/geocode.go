/**
 * go-mapbox Geocoding Module
 * Wraps the mapbox geocoding API for server side use
 * See https://www.mapbox.com/api-documentation/#geocoding for API information
 *
 * https://github.com/ryankurte/go-mapbox
 * Copyright 2017 Ryan Kurte
 */

package geocode

import (
	"fmt"

	"github.com/google/go-querystring/query"
	"github.com/icep1anet/go-mapbox/lib/base"
)

const (
	apiBaseName          = "search"
	apiName              = "geocode"
	apiVersion           = "v6"
)

// Type defines geocode location response types
type Type string

const (
	// Country level
	Country Type = "country"
	// Region level
	Region Type = "region"
	// Postcode level
	Postcode Type = "postcode"
	// District level
	District Type = "district"
	// Place level
	Place Type = "place"
	// Locality level
	Locality Type = "locality"
	// Neighborhood level
	Neighborhood Type = "neighborhood"
	// Address level
	Address Type = "address"
	// POI (Point of Interest) level
	POI Type = "poi"
)

// Geocode api wrapper instance
type Geocode struct {
	base *base.Base
}

// NewGeocode Create a new Geocode API wrapper
func NewGeocode(base *base.Base) *Geocode {
	return &Geocode{base}
}

// ForwardRequestOpts request options fo forward geocoding
type ForwardRequestOpts struct {
	Place 		 string           `url:"q"`
	Permanent 	 bool             `url:"permanent,omitempty"`
	Autocomplete bool             `url:"autocomplete,omitempty"`
	Language     string           `url:"language,omitempty"`
	Country      string           `url:"country,omitempty"`
	Proximity    []float64        `url:"proximity,omitempty"`
	Types        []Type           `url:"types,omitempty"`
	BBox         base.BoundingBox `url:"bbox,omitempty"`
	Limit        uint             `url:"limit,omitempty"`
}

// ForwardResponse is the response from a forward geocode lookup
type ForwardResponse struct {
	*base.FeatureCollection
	Query []string
}

// Forward geocode lookup
// Finds locations from a place name
func (g *Geocode) Forward(req *ForwardRequestOpts) (*ForwardResponse, error) {
	v, err := query.Values(req)
	if err != nil {
		return nil, err
	}

	resp := ForwardResponse{}

	// https://api.mapbox.com/search/geocode/v6/forward?q=<queryString>
	err = g.base.QueryBase(fmt.Sprintf("%s/%s/%s/forward", apiBaseName, apiName, apiVersion), &v, &resp)

	return &resp, err
}

// ReverseRequestOpts request options fo reverse geocoding
type ReverseRequestOpts struct {
	Longitude float64 `url:"longitude"`
	Latitude  float64 `url:"latitude"`
	Language  string  `url:"language,omitempty"`
	Country   string  `url:"country,omitempty"`
	Permanent bool    `url:"permanent,omitempty"`
	Types     []Type  `url:"types,omitempty"`
	Limit     uint    `url:"limit,omitempty"`
}

// ReverseResponse is the response to a reverse geocode request
type ReverseResponse struct {
	*base.FeatureCollection
	Query []float64
}

// Reverse geocode lookup
// Finds place names from a location
func (g *Geocode) Reverse(req *ReverseRequestOpts) (*ReverseResponse, error) {
	v, err := query.Values(req)
	if err != nil {
		return nil, err
	}

	resp := ReverseResponse{}

	// https://api.mapbox.com/search/geocode/v6/reverse?longitude={longitude}&latitude={latitude}
	err = g.base.QueryBase(fmt.Sprintf("%s/%s/%s/reverse", apiBaseName, apiName, apiVersion), &v, &resp)

	return &resp, err
}
