package settings

import (
	"encoding/json"
	"io/ioutil"
	"time"

	"github.com/aws/aws-sdk-go/aws/endpoints"
)

type ELBSettings struct {
	Delay      time.Duration
	Period     time.Duration
	QueryRange time.Duration
	AWSRegion  string
	TagName    string
	TagValue   string
	Metrics    []ELBMetric
}

type ELBMetric struct {
	Name              string
	Statistic         string
	ExtendedStatistic string
}

func NewSettings(filepath string) (*ELBSettings, error) {
	raw, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var settings ELBSettings
	json.Unmarshal(raw, &settings)
	return &settings, nil
}

func (this *ELBSettings) UnmarshalJSON(data []byte) error {
	type Alias ELBSettings

	read := &Alias{
		Delay:      60 * time.Second,
		Period:     60 * time.Second,
		QueryRange: 60 * time.Second,
		AWSRegion:  endpoints.UsEast1RegionID,
		TagName:    "KubernetesCluster",
		TagValue:   "MyCluster",
		Metrics:    nil,
	}

	if err := json.Unmarshal(data, &read); err != nil {
		return err
	}

	*this = ELBSettings(*read)

	return nil
}