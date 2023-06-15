package polymarketmonitor

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"polymarket-monitor/resources"
	"reflect"
	"time"

	"github.com/grafov/bcast"
)

type PolyMarketNewMonitorTask struct {
	WebsiteBaseURL string
	WebsiteName    string
	OldData        []resources.PolyMarketNewListingAPIResp
	Delay          time.Duration
	FirstRun bool
	BroadCastGroup *bcast.Group
}

var Headers = http.Header {
	"accept": {"text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9"},
	"accept-language": {"en-US,en;q=0.9"},
	"sec-ch-ua": {"sec-ch-ua"},
	"sec-ch-ua-mobile": {"sec-ch-ua-mobile"},
	"sec-ch-ua-platform": {"\"Windows\""},
	"sec-fetch-dest": {"document"},
	"sec-fetch-mode": {"navigate"},
	"sec-fetch-site": {"none"},
	"sec-fetch-user": {"?1"},
	"upgrade-insecure-requests": {"1"},
	"user-agent": {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36"},
}


func (m *PolyMarketNewMonitorTask) RunMonitor() {
	// monitor runs in an infinite loop
	var pageJson []byte
	for {
		// request to their new homepage api
		req, err := http.NewRequest(http.MethodGet, m.WebsiteBaseURL+"/api/homepage/new", nil)
		req.Header = Headers

		if err != nil {
			log.Println(err)
			time.Sleep(m.Delay * time.Millisecond)
			continue
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Println(err)
			time.Sleep(m.Delay * time.Millisecond)
			continue
		}

		// reads response data
		pageJson, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println("err", err)
		}
		resp.Body.Close() // dont use defer in loops
		switch resp.StatusCode {

		// if the request was successful, move on
		case 200:
			// just iterates through the infinite for loop again if theres no response
			if len(pageJson) <= 0 {
				time.Sleep(m.Delay * time.Millisecond)
				log.Println("len too small")
				continue
			}

			// unmarshals the json to a struct
			data := []resources.PolyMarketNewListingAPIResp{}
			json.Unmarshal(pageJson, &data)
			// if this is the first time the monitor is running, do this
			if (m.FirstRun) {
				log.Println("starting polymarket monitor")
				m.FirstRun = false
				m.OldData = data
				
			} else { // otherwise, do this
				log.Println("running polymarket monitor")
				SendList := []resources.PolyMarketNewListingAPIResp{}
				// looks through the polymarket api data
				for _, i := range data {
					previouslyExists := false
					// if any listing in the new data doesn't exist in the old data, then add it to a slice which will be sent to the webhooks
					for _, j := range m.OldData {
						if reflect.DeepEqual(i, j) {
							previouslyExists = true
						}
					}
					if (!previouslyExists) {
						SendList = append(SendList, i)
					}
				}


				m.OldData = data

				// just adds the store name for future reference
				for _, v := range SendList {
					v.Store = resources.POLYMARKET
				}

				// broadcasts the slice to the webhook tasks
				if len(SendList) > 0 {
					m.BroadCastGroup.Send(SendList)
				}
				
			}
		}
		time.Sleep(m.Delay * time.Millisecond)
	}
}
