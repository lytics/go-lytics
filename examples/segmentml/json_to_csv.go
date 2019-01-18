package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	lytics "github.com/lytics/go-lytics"
)

const (
	FeatureCSV    = "FeatureCSV"
	OverviewCSV   = "OverviewCSV"
	PredictionCSV = "PredictionsCSV"
)

var aid string

func main() {
	apikey := os.Getenv("LIOKEY")
	aid = os.Getenv("LIOAID")

	client := lytics.NewLytics(apikey, nil)
	if apikey == "" || aid == "" {
		fmt.Println("LIOKEY or LIOAID env not set")
		os.Exit(1)
	}

	segML1, err := client.GetSegmentMLModel("prospects::prospects_product_trial")
	if err != nil {
		fmt.Printf("error on SegML Get %v \n err %v", segML1, err)
		os.Exit(1)
	}
	segML2, err := client.GetSegmentMLModel("prospects_product_trial::customers")
	if err != nil {
		fmt.Printf("error on SegML Get %v \n err %v", segML2, err)
		os.Exit(1)
	}

	toCSV(FeatureCSV, []lytics.SegmentML{segML1})
	toCSV(FeatureCSV, []lytics.SegmentML{segML2})
	toCSV(OverviewCSV, []lytics.SegmentML{segML1, segML2})
	toCSV(PredictionCSV, []lytics.SegmentML{segML1})
	toCSV(PredictionCSV, []lytics.SegmentML{segML2})

}

func toCSV(format string, models []lytics.SegmentML) {
	var filename string
	var header []string
	switch format {
	case FeatureCSV:
		filename = fmt.Sprintf("%s_%s_Features.csv", aid, models[0].Name)
		header = []string{"Field_Name", "Field_Type", "Field_Kind", "Importance", "Correlation", "Impact-Lift", "Impact-Threshold"}
	case OverviewCSV:
		filename = fmt.Sprintf("%s_SegmentML_Models_Overview.csv", aid)
		header = []string{"Name", "Source", "Target", "MSE", "RSQ", "AUC", "False_Negative", "Fales_Positive", "True_Negative", "True_Positive"}
	case PredictionCSV:
		filename = fmt.Sprintf("%s_%s_Predictions.csv", aid, models[0].Name)
		header = []string{"Prediction", "Source_Count", "Target_Count"}
	}

	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var records [][]string
	records = append(records, header)

	switch format {
	case FeatureCSV:
		for _, feature := range models[0].Features {
			records = append(records, featuresToSlice(feature))
		}
	case OverviewCSV:
		for _, model := range models {
			records = append(records, overviewToSlice(model))
		}
	case PredictionCSV:
		records = predictionsTo2DSlice(models[0], header)
	}

	w := csv.NewWriter(file)
	err = w.WriteAll(records)
	if err != nil {
		panic(err)
	}

}

func overviewToSlice(model lytics.SegmentML) []string {
	sl := make([]string, 10)

	for i := 0; i < 10; i++ {
		var input string
		switch i {
		case 0:
			input = model.Name
		case 1:
			input = model.Conf.Source.Name
		case 2:
			input = model.Conf.Target.Name
		case 3:
			input = FloatToString(model.Summary.Mse)
		case 4:
			input = FloatToString(model.Summary.Rsq)
		case 5:
			input = FloatToString(model.Summary.AUC)
		case 6:
			input = strconv.Itoa(model.Summary.Conf.FalseNegative)
		case 7:
			input = strconv.Itoa(model.Summary.Conf.FalsePositive)
		case 8:
			input = strconv.Itoa(model.Summary.Conf.TrueNegative)
		case 9:
			input = strconv.Itoa(model.Summary.Conf.TruePositive)
		}

		sl[i] = input
	}
	return sl
}

func predictionsTo2DSlice(model lytics.SegmentML, header []string) [][]string {
	sl2D := [][]string{header}

	failPredMap := make(map[string]int)
	for failX, failCt := range model.Summary.Fail {
		failPredMap[failX] = failCt
	}

	successPredMap := make(map[string]int)
	for successX, successCt := range model.Summary.Success {
		successPredMap[successX] = successCt
	}

	for i := 0; i < 101; i++ {
		xVal := fmt.Sprintf("%.2f", float64(i)/100)
		sl := make([]string, 3)
		for j := 0; j < 3; j++ {
			var input string
			switch j {
			case 0:
				input = xVal
			case 1:
				input = strconv.Itoa(failPredMap[xVal])
			case 2:
				input = strconv.Itoa(successPredMap[xVal])
			}
			sl[j] = input
		}
		sl2D = append(sl2D, sl)
	}
	return sl2D
}

func featuresToSlice(feature lytics.Feature) []string {
	sl := make([]string, 7)

	for i := 0; i < 7; i++ {
		var input string
		switch i {
		case 0:
			input = feature.Name
		case 1:
			input = feature.Type
		case 2:
			input = feature.Kind
		case 3:
			input = FloatToString(feature.Importance)
		case 4:
			input = FloatToString(feature.Correlation)
		case 5:
			input = FloatToString(feature.Impact.Lift)
		case 6:
			input = FloatToString(feature.Impact.Threshold)
		}
		sl[i] = input
	}
	return sl
}

func FloatToString(num float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(num, 'f', 6, 64)
}
