package lytics

import (
	"net/url"
)

const (
	facebookLookalikeStatusEndpoint = "provider/facebook/lookalike/status"
	facebookLookalikeCreateEndpoint = "provider/facebook/lookalike/create"
)

type FBLookalikeAudience struct {
	Name             string                  `json:"name"`               // Custom Audience name (required)
	OriginAudienceId string                  `json:"origin_audience_id"` // ID of Custom Audience. Origin audiences must have at least 100 members. (required)
	LookalikeSpec    FBLookalikeAudienceSpec `json:"lookalike_spec"`     // Lookalike Spec Object (required)
}

type FBLookalikeAudienceSpec struct {
	Type                    string  `json:"type,omitempty"`                      // (similarity | reach) (type or ratio required)
	Ratio                   float64 `json:"ratio,omitempty"`                     // 0.01-0.20 incremented by 0.01. Top x% of original audience in a selected country (type or ratio required)
	StartingRatio           float64 `json:"starting_ratio,omitempty"`            // Start percentage for lookalike. For example, starting_ratio 0.01 and ratio 0.02 creates a lookalike from 1% to 2% of a lookalike segment.
	AllowInternationalSeeds bool    `json:"allow_international_seeds,omitempty"` // At least 100 seed audience members from a country.
	Country                 string  `json:"country,omitempty"`                   // Find lookalike audience members in this country

	// note: there are more detailed targeting params related to geo/location that I have not added as they are not necessary for
	// the current use cases. just mentioning in case they come up. full docs here https://developers.facebook.com/docs/marketing-api/lookalike-audience-targeting/
}

type FBLookalikeAudienceStatus struct {
	Id             string `json:"id"`
	DeliveryStatus struct {
		Code        int    `json:"code"`
		Description string `json:"description"`
	} `json:"delivery_status"`
}

// GetLookalikeStatus gets the status of a previously created lookalike audience
func (l *Client) GetLookalikeStatus(authId, audienceId string) (FBLookalikeAudienceStatus, error) {
	res := ApiResp{}
	data := FBLookalikeAudienceStatus{}

	params := url.Values{}
	params.Add("auth_id", authId)
	params.Add("audience_id", audienceId)

	// make the request
	err := l.Get(parseLyticsURL(facebookLookalikeStatusEndpoint, map[string]string{}), params, nil, &res, &data)
	if err != nil {
		return data, err
	}

	return data, nil
}

// CreateLookalike creates a lookalike audience in facebook for a given account, auth and audience spec
func (l *Client) CreateLookalike(authId, adsAccountId string, lookalike FBLookalikeAudience) (string, error) {
	var (
		err error
	)

	res := ApiResp{}
	var data struct {
		Id string `json:"id"`
	}

	params := url.Values{}
	params.Add("auth_id", authId)
	params.Add("ads_account_id", adsAccountId)

	err = l.Post(facebookLookalikeCreateEndpoint, params, lookalike, &res, &data)
	if err != nil {
		return data.Id, err
	}

	return data.Id, nil
}
