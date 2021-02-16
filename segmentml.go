package lytics

import (
	"fmt"
	"time"
)

const (
	segmentMLEndpoint     = "segmentml/:id"
	segmentMLListEndpoint = "segmentml"
	segmentMLDepEndpoint  = "segmentml/:id/_dependencies"
)

type SegmentML struct {
	Name     string    `json:"name"`
	State    string    `json:"state"`
	Reason   string    `json:"reason"`
	Created  time.Time `json:"created"`
	AuthorID string    `json:"author_id"`
	Conf     struct {
		Source struct {
			ID                 string      `json:"id"`
			Updated            time.Time   `json:"updated"`
			Created            time.Time   `json:"created"`
			Aid                int         `json:"aid"`
			AccountID          string      `json:"account_id"`
			ShortID            string      `json:"short_id"`
			Name               string      `json:"name"`
			Kind               string      `json:"kind"`
			IsPublic           bool        `json:"is_public"`
			PublicName         string      `json:"public_name"`
			SlugName           string      `json:"slug_name"`
			Description        string      `json:"description"`
			Table              string      `json:"table"`
			AuthorID           string      `json:"author_id"`
			Invalid            bool        `json:"invalid"`
			InvalidReason      string      `json:"invalid_reason"`
			Deleted            bool        `json:"deleted"`
			DatemathCalc       bool        `json:"datemath_calc"`
			ForwardDatemath    bool        `json:"forward_datemath"`
			SaveHist           bool        `json:"save_hist"`
			FieldUpdates       bool        `json:"field_updates"`
			FieldChangesFields interface{} `json:"field_changes_fields"`
			ScheduleExit       bool        `json:"schedule_exit"`
			SegmentQl          string      `json:"segment_ql"`
			Tags               []string    `json:"tags"`
			Ast                struct {
				Ident string `json:"ident"`
			} `json:"ast"`
			Fields interface{} `json:"fields"`
		} `json:"source"`
		Target struct {
			ID                 string      `json:"id"`
			Updated            time.Time   `json:"updated"`
			Created            time.Time   `json:"created"`
			Aid                int         `json:"aid"`
			AccountID          string      `json:"account_id"`
			ShortID            string      `json:"short_id"`
			Name               string      `json:"name"`
			Kind               string      `json:"kind"`
			IsPublic           bool        `json:"is_public"`
			SlugName           string      `json:"slug_name"`
			Description        string      `json:"description"`
			Table              string      `json:"table"`
			AuthorID           string      `json:"author_id"`
			Invalid            bool        `json:"invalid"`
			InvalidReason      string      `json:"invalid_reason"`
			Deleted            bool        `json:"deleted"`
			DatemathCalc       bool        `json:"datemath_calc"`
			ForwardDatemath    bool        `json:"forward_datemath"`
			SaveHist           bool        `json:"save_hist"`
			FieldUpdates       bool        `json:"field_updates"`
			FieldChangesFields interface{} `json:"field_changes_fields"`
			ScheduleExit       bool        `json:"schedule_exit"`
			SegmentQl          string      `json:"segment_ql"`
			Tags               []string    `json:"tags"`
			Ast                struct {
				Op   string `json:"op"`
				Args []struct {
					Op   string `json:"op"`
					Args []struct {
						Ident string `json:"ident,omitempty"`
						Val   string `json:"val,omitempty"`
					} `json:"args"`
				} `json:"args"`
			} `json:"ast"`
			Fields []string `json:"fields"`
		} `json:"target"`
		Additional       interface{} `json:"additional"`
		Collections      interface{} `json:"collections"`
		CustomSegmentIds interface{} `json:"custom_segment_ids"`
		Collect          int         `json:"collect"`
		UseScores        bool        `json:"use_scores"`
		UseContent       bool        `json:"use_content"`
		BuildOnly        bool        `json:"build_only"`
		AutoTune         bool        `json:"auto_tune"`
		ModelName        string      `json:"model_name"`
		ModelType        string      `json:"model_type"`
		ReRun            bool        `json:"re_run"`
	} `json:"conf"`
	Features []Feature `json:"features"`
	Summary  struct {
		Conf struct {
			FalsePositive int `json:"FalsePositive"`
			TruePositive  int `json:"TruePositive"`
			FalseNegative int `json:"FalseNegative"`
			TrueNegative  int `json:"TrueNegative"`
		} `json:"conf"`
		Mse                float64        `json:"mse"`
		Rsq                float64        `json:"rsq"`
		Success            map[string]int `json:"success"`
		Fail               map[string]int `json:"fail"`
		Auc                float64        `json:"auc"`
		Threshold          float64        `json:"threshold"`
		Accuracy           int            `json:"accuracy"`
		Reach              int            `json:"reach"`
		AudienceSimilarity float64        `json:"audience_similarity"`
		ModelHealth        string         `json:"model_health"`
		Msgs               []struct {
			Text     string   `json:"text"`
			Tags     []string `json:"tags"`
			Severity string   `json:"severity"`
		} `json:"msgs"`
	} `json:"summary"`
}

/*
type Feature struct {
	Kind        string  `json:"kind, omitempty"`
	Type        string  `json:"type, omitempty"`
	Name        string  `json:"name, omitempty"`
	Importance  float64 `json:"importance, omitempty"`
	Correlation float64 `json:"correlation, omitempty"`
	Impact      struct {
		Lift      float64 `json:"lift, omitempty"`
		Threshold float64 `json:"threshold, omitempty"`
	} `json:"impact, omitempty"`
}

type SegmentML struct {
	Name string `json:"name"`

	Features []Feature `json:"features,omitempty"`

	Summary struct {
		Mse float64 `json:"mse"`
		Rsq float64 `json:"rsq"`
		AUC float64 `json:"auc"`

		Conf struct {
			FalsePositive int `json:"FalsePositive,omitempty"`
			TruePositive  int `json:"TruePositive,omitempty"`
			FalseNegative int `json:"FalseNegative,omitempty"`
			TrueNegative  int `json:"TrueNegative,omitempty"`
		} `json:"Conf,omitempty"`

		Fail    map[string]int `json:"fail"`
		Success map[string]int `json:"success"`
	} `json:"summary,omitempty"`

	Conf struct {
		Source      Segment  `json:"source"`
		Target      Segment  `json:"target"`
		Additional  []string `json:"additional"`
		Collections []string `json:"collections"`
		Collect     int      `json:"collect"`
		UseScores   bool     `json:"use_scores"`
		UseContent  bool     `json:"use_content"`
		WriteToGcs  bool     `json:"write_to_gcs"`
	} `json:"conf,omitempty"`
}

type Dependencies struct {
	Fields map[string][][2]float64 "fields"
}

type Prediction struct {
	Val       string `json:"val"`
	FailCt    int    `json:"fail_ct"`
	SuccessCt int    `json:"success_ct"`
}
*/
type Prediction struct {
	Val       string `json:"val"`
	FailCt    int    `json:"fail_ct"`
	SuccessCt int    `json:"success_ct"`
}

type Dependencies struct {
	Fields map[string][][2]float64 "fields"
}
type Feature struct {
	Kind            string  `json:"kind"`
	Type            string  `json:"type"`
	Name            string  `json:"name"`
	Importance      float64 `json:"importance"`
	FieldPrevalence struct {
		Source float64 `json:"source"`
		Target float64 `json:"target"`
	} `json:"field_prevalence"`
	Correlation float64 `json:"correlation"`
	Impact      struct {
		Lift      float64 `json:"lift"`
		Threshold float64 `json:"threshold"`
	} `json:"impact"`
}

// GetSegmentMLModel returns the details for a single segmentML Model based on id
// https://www.getlytics.com/developers/rest-api#segment-m-l
func (l *Client) GetSegmentMLModel(id string) (SegmentML, error) {
	res := ApiResp{}
	data := SegmentML{}

	// make the request
	err := l.Get(parseLyticsURL(segmentMLEndpoint, map[string]string{"id": id}), nil, nil, &res, &data)
	if err != nil {
		return SegmentML{}, err
	}
	return data, nil
}

// GetSegmentMLModels returns all models for the account
// https://www.getlytics.com/developers/rest-api#segment-m-l
func (l *Client) GetSegmentMLModels() ([]SegmentML, error) {
	res := ApiResp{}
	data := []SegmentML{}

	// make the request
	err := l.Get(segmentMLListEndpoint, nil, nil, &res, &data)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (l *Client) GetSegmentMLDependencies(id string) (Dependencies, error) {
	res := ApiResp{}
	data := Dependencies{}

	err := l.Get(parseLyticsURL(segmentMLDepEndpoint, map[string]string{"id": id}), nil, nil, &res, &data)
	if err != nil {
		return Dependencies{}, err
	}

	return data, nil
}

func (f Feature) Headers() []interface{} {
	return []interface{}{
		"Field_Name", "Field_Type", "Field_Kind", "Importance", "Correlation", "Impact-Lift", "Impact-Threshold",
	}
}
func (f Feature) Row() []interface{} {
	return []interface{}{
		f.Name, f.Type, f.Kind, f.Importance, f.Correlation, f.Impact.Lift, f.Impact.Threshold,
	}
}

func (s *SegmentML) Headers() []interface{} {
	return []interface{}{
		"Name", "Source", "Target", "MSE", "RSQ", "AUC", "False_Negative", "False_Positive", "True_Negative", "True_Positive",
	}
}

func (s *SegmentML) Row() []interface{} {
	sum := s.Summary
	c := s.Conf
	return []interface{}{
		s.Name, c.Source.Name, c.Source.Name, sum.Mse, sum.Rsq, sum.Auc, sum.Conf.FalseNegative, sum.Conf.FalsePositive, sum.Conf.TrueNegative, sum.Conf.TruePositive,
	}
}

func (p Prediction) Headers() []interface{} {
	return []interface{}{
		"Prediction", "Fail_Count", "Success_Count",
	}
}

func (p Prediction) Row() []interface{} {
	return []interface{}{
		p.Val, p.FailCt, p.SuccessCt,
	}
}

func (s *SegmentML) GetPredictions() []Prediction {
	var predictions []Prediction

	failPredMap := make(map[string]int)
	for failX, failCt := range s.Summary.Fail {
		failPredMap[failX] = failCt
	}

	successPredMap := make(map[string]int)
	for successX, successCt := range s.Summary.Success {
		successPredMap[successX] = successCt
	}

	for i := 0; i < 101; i++ {
		xVal := fmt.Sprintf("%.2f", float64(i)/100)
		sl := Prediction{}

		sl.Val = xVal
		sl.FailCt = failPredMap[xVal]
		sl.SuccessCt = successPredMap[xVal]

		predictions = append(predictions, sl)

	}
	return predictions
}
