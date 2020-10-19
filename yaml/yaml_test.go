package yaml_test

import (
	"fmt"

	lg "github.com/hiromaily/golibs/log"
	tu "github.com/hiromaily/golibs/testutil"
	. "github.com/hiromaily/golibs/yaml"

	//"gopkg.in/guregu/null.v3"
	"os"
	"testing"
	"time"

	yml "gopkg.in/yaml.v2"
)

//Test1
type T1 struct {
	TT `yaml:"tables"`
}

type TT struct {
	A string
	B struct {
		RenamedC int   `yaml:"c"`
		D        []int `yaml:",flow"`
	}
}

//Test2
type T2 struct {
	TT2 `yaml:"tables"`
}

type TT2 struct {
	Events EventA `yaml:"events"`
}

//TEST3
type T3 struct {
	TT3 `yaml:"tables"`
}

type TT3 struct {
	Events []EventA `yaml:"events"`
}

//TEST4
type T4 struct {
	TT4 `yaml:"tables"`
}

type TT4 struct {
	Events []EventB `yaml:"events"`
}

type EventA struct {
	Name                string   `yaml:"name"`
	OrganiserID         int      `yaml:"organiserId"`
	Start               string   `yaml:"start"`
	End                 string   `yaml:"end"`
	Header              string   `yaml:"header"`
	Description         string   `yaml:"description"`
	CreatedAt           string   `yaml:"createdAt"`
	UpdatedAt           string   `yaml:"updatedAt"`
	Genres              string   `yaml:"genres"`
	Website             string   `yaml:"website"`
	FacebookPage        string   `yaml:"facebookPage"`
	InstagramURL        string   `yaml:"instagramUrl"`
	SoundcloudURL       string   `yaml:"soundcloudUrl"`
	SpotifyURL          string   `yaml:"spotifyUrl"`
	TwitterURL          string   `yaml:"twitterUrl"`
	VimeoURL            string   `yaml:"vimeoUrl"`
	YoutubeURL          string   `yaml:"youtubeUrl"`
	FbEventID           string   `yaml:"fbEventId"`
	YtActive            int      `yaml:"ytActive"`
	LineUp              string   `yaml:"lineUp"`
	VenueID             int      `yaml:"venueId"`
	Logo                string   `yaml:"logo"`
	ScoreEpicness       int      `yaml:"scoreEpicness"`
	ScoreLineup         int      `yaml:"scoreLineup"`
	ScoreLocation       int      `yaml:"scoreLocation"`
	ScoreDecoration     int      `yaml:"scoreDecoration"`
	ScoreFood           int      `yaml:"scoreFood"`
	ScoreSound          int      `yaml:"scoreSound"`
	ScoreWaitingTime    int      `yaml:"scoreWaitingTime"`
	ScoreCleanliness    int      `yaml:"scoreCleanliness"`
	ScoreReachability   int      `yaml:"scoreReachability"`
	ScoreSustainability int      `yaml:"scoreSustainability"`
	ScoreFacilities     int      `yaml:"scoreFacilities"`
	ScorePrices         int      `yaml:"scorePrices"`
	ScoreAverage        int      `yaml:"scoreAverage"`
	TotalReviews        int      `yaml:"totalReviews"`
	TotalComments       int      `yaml:"totalComments"`
	Published           int      `yaml:"published"`
	Source              string   `yaml:"source"`
	Featured            int      `yaml:"featured"`
	Demo                int      `yaml:"demo"`
	TicketURL           string   `yaml:"ticketUrl"`
	TicketMin           string   `yaml:"ticketMin"`
	TicketMax           string   `yaml:"ticketMax"`
	Indoor              int      `yaml:"indoor"`
	Days                int      `yaml:"days"`
	Free                int      `yaml:"free"`
	Camping             int      `yaml:"camping"`
	Tags                []string `yaml:"tags"`
	IsActive            int      `yaml:"isActive"`
	Status              int      `yaml:"status"`
}

type EventB struct {
	Name                string     `yaml:"name"`
	OrganiserID         int        `yaml:"organiserId"`
	Start               *time.Time `yaml:"start"`
	End                 *time.Time `yaml:"end"`
	Header              *string    `yaml:"header"`
	Description         *string    `yaml:"description"`
	CreatedAt           *time.Time `yaml:"createdAt"`
	UpdatedAt           *time.Time `yaml:"updatedAt"`
	Genres              *string    `yaml:"genres"`
	Website             *string    `yaml:"website"`
	FacebookPage        *string    `yaml:"facebookPage"`
	InstagramURL        *string    `yaml:"instagramUrl"`
	SoundcloudURL       *string    `yaml:"soundcloudUrl"`
	SpotifyURL          *string    `yaml:"spotifyUrl"`
	TwitterURL          *string    `yaml:"twitterUrl"`
	VimeoURL            *string    `yaml:"vimeoUrl"`
	YoutubeURL          *string    `yaml:"youtubeUrl"`
	FbEventID           *string    `yaml:"fbEventId"`
	YtActive            int        `yaml:"ytActive"`
	LineUp              *string    `yaml:"lineUp"`
	VenueID             int        `yaml:"venueId"`
	Logo                *string    `yaml:"logo"`
	ScoreEpicness       int        `yaml:"scoreEpicness"`
	ScoreLineup         int        `yaml:"scoreLineup"`
	ScoreLocation       int        `yaml:"scoreLocation"`
	ScoreDecoration     int        `yaml:"scoreDecoration"`
	ScoreFood           int        `yaml:"scoreFood"`
	ScoreSound          int        `yaml:"scoreSound"`
	ScoreWaitingTime    int        `yaml:"scoreWaitingTime"`
	ScoreCleanliness    int        `yaml:"scoreCleanliness"`
	ScoreReachability   int        `yaml:"scoreReachability"`
	ScoreSustainability int        `yaml:"scoreSustainability"`
	ScoreFacilities     int        `yaml:"scoreFacilities"`
	ScorePrices         int        `yaml:"scorePrices"`
	ScoreAverage        int        `yaml:"scoreAverage"`
	TotalReviews        int        `yaml:"totalReviews"`
	TotalComments       int        `yaml:"totalComments"`
	Published           int        `yaml:"published"`
	Source              *string    `yaml:"source"`
	Featured            int        `yaml:"featured"`
	Demo                int        `yaml:"demo"`
	TicketURL           *string    `yaml:"ticketUrl"`
	TicketMin           *string    `yaml:"ticketMin"`
	TicketMax           *string    `yaml:"ticketMax"`
	Indoor              int        `yaml:"indoor"`
	Days                int        `yaml:"days"`
	Free                int        `yaml:"free"`
	Camping             int        `yaml:"camping"`
	Tags                []string   `yaml:"tags"`
	IsActive            int        `yaml:"isActive"`
	Status              int        `yaml:"status"`
}

// For table test
var yamlTests = []struct {
	yamlFile  string
	typeValue interface{}
	//type1 T1
}{
	{"./data/sample.yml", T1{}},
	{"./data/single.yml", T2{}},
	{"./data/multiple.yml", T3{}},
	{"./data/event.yml", T4{}},
	{"./data/event2.yml", EventB{}},
}

//-----------------------------------------------------------------------------
// Test Framework
//-----------------------------------------------------------------------------

func setup() {
	tu.InitializeTest("[YAML]")
}

func teardown() {
}

func TestMain(m *testing.M) {
	setup()

	code := m.Run()

	teardown()

	os.Exit(code)
}

//-----------------------------------------------------------------------------
// function
//-----------------------------------------------------------------------------
func checkMapData(t *testing.T, b []byte, yamlFile string) {
	m := make(map[interface{}]interface{})

	//3-1. unmarshal yaml
	err := yml.Unmarshal(b, &m)
	if err != nil {
		t.Errorf("yml.Unmarshal error on map: file: %s\n error: %s, ", yamlFile, err)
	}
	lg.Debugf("After unmarshal YAML on map: %v", m)

	//3-2. marshal from map data
	d, err := yml.Marshal(&m)
	if err != nil {
		t.Errorf("yml.Marshal error on map: file: %s\n error: %s, ", yamlFile, err)
	}
	lg.Debugf("After marshal YAML on map: %v", string(d))
}

//-----------------------------------------------------------------------------
// Test
//-----------------------------------------------------------------------------
func TestYAMLTable(t *testing.T) {
	var data []byte
	var err error

	for i, tt := range yamlTests {
		fmt.Printf("[%d] file:%s\n", i+1, tt.yamlFile)
		//data = nil

		//1. load YAML
		data, err = LoadYAMLFile(tt.yamlFile)
		if err != nil {
			t.Errorf("LoadYAMLFile error: file: %s\n error: %s, ", tt.yamlFile, err)
			continue
		}

		//2. using struct data
		//2-1. unmarshal yaml
		//TODO: assert interface to specific type
		if val, ok := tt.typeValue.(T1); ok {
			//TODO: this code is ugly.
			err := yml.Unmarshal(data, &val)
			//checkErrorForUnmarshalStruct(err, tt.yamlFile)
			if err != nil {
				t.Errorf("yml.Unmarshal error on struct: file: %s\n error: %s, ", tt.yamlFile, err)
				continue
			}
			lg.Debugf("After unmarshal YAML on struct: %v", val)

			//2-2. marchal from data
			d, err := yml.Marshal(&val)
			if err != nil {
				t.Errorf("yml.Marshal error on struct: file: %s\n error: %s, ", tt.yamlFile, err)
				continue
			}
			lg.Debugf("After marshal YAML on struct: %v", string(d))
		} else if val, ok := tt.typeValue.(T2); ok {
			err := yml.Unmarshal(data, &val)
			//checkErrorForUnmarshalStruct(err, tt.yamlFile)
			if err != nil {
				t.Errorf("yml.Unmarshal error on struct: file: %s\n error: %s, ", tt.yamlFile, err)
				continue
			}
			lg.Debugf("After unmarshal YAML on struct: %v", val)

			//2-2. marchal from data
			d, err := yml.Marshal(&val)
			if err != nil {
				t.Errorf("yml.Marshal error on struct: file: %s\n error: %s, ", tt.yamlFile, err)
				continue
			}
			lg.Debugf("After marshal YAML on struct: %v", string(d))
		} else if val, ok := tt.typeValue.(T3); ok {
			err := yml.Unmarshal(data, &val)
			//checkErrorForUnmarshalStruct(err, tt.yamlFile)
			if err != nil {
				t.Errorf("yml.Unmarshal error on struct: file: %s\n error: %s, ", tt.yamlFile, err)
				continue
			}
			lg.Debugf("After unmarshal YAML on struct: %v", val)

			//2-2. marchal from data
			d, err := yml.Marshal(&val)
			if err != nil {
				t.Errorf("yml.Marshal error on struct: file: %s\n error: %s, ", tt.yamlFile, err)
				continue
			}
			lg.Debugf("After marshal YAML on struct: %v", string(d))
		} else if val, ok := tt.typeValue.(T4); ok {
			err := yml.Unmarshal(data, &val)
			//checkErrorForUnmarshalStruct(err, tt.yamlFile)
			if err != nil {
				t.Errorf("yml.Unmarshal error on struct: file: %s\n error: %s, ", tt.yamlFile, err)
				continue
			}
			lg.Debugf("After unmarshal YAML on struct: %v", val)

			//2-2. marchal from data
			d, err := yml.Marshal(&val)
			if err != nil {
				t.Errorf("yml.Marshal error on struct: file: %s\n error: %s, ", tt.yamlFile, err)
				continue
			}
			lg.Debugf("After marshal YAML on struct: %v", string(d))
		} else if val, ok := tt.typeValue.(EventB); ok {
			err := yml.Unmarshal(data, &val)
			//checkErrorForUnmarshalStruct(err, tt.yamlFile)
			if err != nil {
				t.Errorf("yml.Unmarshal error on struct: file: %s\n error: %s, ", tt.yamlFile, err)
				continue
			}
			lg.Debugf("After unmarshal YAML on struct: %v", val)

			//2-2. marchal from data
			d, err := yml.Marshal(&val)
			if err != nil {
				t.Errorf("yml.Marshal error on struct: file: %s\n error: %s, ", tt.yamlFile, err)
				continue
			}
			lg.Debugf("After marshal YAML on struct: %v", string(d))
		} else {
			lg.Debug("Assert can not be.")
		}

		//3. as map interface
		checkMapData(t, data, tt.yamlFile)
	}
}

//-----------------------------------------------------------------------------
// Benchmark
//-----------------------------------------------------------------------------
func BenchmarkYaml(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		//
		//_ = CallSomething()
		//
	}
	b.StopTimer()
}
