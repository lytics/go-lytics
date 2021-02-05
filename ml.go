package lytics

import (
	"time"
)

const (
	mlEndpoint     = "ml/:id"
	mlListEndpoint = "ml"
	// MLDepEndpoint  = "segmentml/:id/_dependencies"
)

// type MLConfigs struct {
// 	mlConfigs []ML
// }

// ML Struct
type ML struct {
	Aid    int `json:"aid"`
	Config struct {
		AutoTune         bool        `json:"auto_tune"`
		BuildOnly        bool        `json:"build_only"`
		Collect          int         `json:"collect"`
		CustomSegmentIds interface{} `json:"custom_segment_ids"`
		Internal         bool        `json:"internal"`
		ModelType        string      `json:"model_type"`
		ReRun            bool        `json:"re_run"`
		TuneModel        bool        `json:"tune_model"`
		UseContent       bool        `json:"use_content"`
		UseEntStore      bool        `json:"use_ent_store"`
		UseScores        bool        `json:"use_scores"`
	} `json:"config"`
	Created   time.Time `json:"created"`
	ID        string    `json:"id"`
	IsActive  bool      `json:"is_active"`
	IsHealthy bool      `json:"is_healthy"`
	Label     string    `json:"label"`
	Name      string    `json:"name"`
	Source    string    `json:"source"`
	State     string    `json:"state"`
	Target    string    `json:"target"`
	Type      string    `json:"type"`
	Updated   time.Time `json:"updated"`
	WorkIds   []string  `json:"work_ids"`
}

// GetMLModel returns the details for a single ML Model based on id
// https://www.getlytics.com/developers/rest-api#segment-m-l
func (l *Client) GetMLModel(id string) (ML, error) {
	res := ApiResp{}
	data := ML{}

	// make the request
	err := l.Get(parseLyticsURL(mlEndpoint, map[string]string{"id": id}), nil, nil, &res, &data)
	if err != nil {
		return ML{}, err
	}
	return data, nil
}

// GetMLModels returns all ML models for the account
// https://www.getlytics.com/developers/rest-api#segment-m-l
func (l *Client) GetMLModels() ([]ML, error) {
	res := ApiResp{}
	data := []ML{}

	// make the request
	err := l.Get(mlListEndpoint, nil, nil, &res, &data)
	if err != nil {
		return data, err
	}

	return data, nil
}

// Headers returns headers for tablewriter
func (m ML) Headers() []interface{} {
	return []interface{}{
		"Name", "ID", "Source", "Target", "Type", "Active", "Healthy",
	}
}

// Row returns a row for tablewriter
func (m ML) Row() []interface{} {
	// c := m.Config
	return []interface{}{
		m.Name, m.ID, m.Source, m.Target, m.Type, m.IsActive, m.IsHealthy,
	}
}
