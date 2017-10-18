package mock

import (
	"fmt"
	"io/ioutil"
)

const (
	MockAccountID          = "224y234ywerhwerhw345y2"
	MockApiKey             = "NOTAREALAPIKEY"
	MockAuthID             = "26723jk72l3gh2l34gj72l"
	MockUserID             = "a6dsga789dg69adg6a9g22"
	MockWorkID             = "dgsgg7979dsg896gs98s9d"
	MockProviderID         = "2g46k24g62jk346g2462k4"
	MockSegmentID1         = "f186a334ad7109bbe08880"
	MockSegmentID2         = "erg9erg8er6gerge90r8er"
	MockSegmentCollection  = "92630989bbea40f7b830f2"
	MockCampaignID         = "4a5984e1138646bb9d692c"
	MockVariationID        = "572d930921c447348ed424"
	MockSegmentMLID        = "all::active_prospects"
	MockTopicID            = "mock"
	MockFacebookAuthID     = "94a151a3c14aa2ee9c5574876b20c468"
	MockFacebookAccountID  = "10150970483556602"
	MockFacebookAudienceID = "6081729845018"
)

func readJsonFile(name string) string {
	jsonData, err := ioutil.ReadFile(fmt.Sprintf("mock/json/%s.json", name))
	if err != nil {
		panic(err)
	}

	return string(jsonData)
}
