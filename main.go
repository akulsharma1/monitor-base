package main

import (
	"encoding/json"
	"log"
	"polymarket-monitor/polymarketmonitor"
	"polymarket-monitor/resources"
	"sync"

	_ "net/http/pprof"

	"io/ioutil"

	"github.com/grafov/bcast"
)



type MonitorUser struct {
	GroupName string `json:"group_name"`
	Webhook      string `json:"webhook_url"`
	ProfilePic   string `json:"profile_picture"`
	Stores []string `json:"stores"`
}

type Users struct {
	MonitorUsers []MonitorUser `json:"users"`
}

func initializePolymarketMonitor(polymarketGroup *bcast.Group, wg *sync.WaitGroup) {
	// polymarket monitor task struct thing
	PolyMarketMonitorRunTask := &polymarketmonitor.PolyMarketNewMonitorTask {
		WebsiteBaseURL: "https://polymarket.com/",
		WebsiteName: resources.POLYMARKET,
		OldData: []resources.PolyMarketNewListingAPIResp{},
		Delay: 1200,
		FirstRun: true,
		BroadCastGroup: polymarketGroup,
	}

	// starts polymarket run monitor (it uses a struct task)
	go PolyMarketMonitorRunTask.RunMonitor()
	wg.Add(1)
}


func main() {
	var wg sync.WaitGroup

	jsonData, err := ioutil.ReadFile("users.json")
	if err != nil {
		log.Fatal("Error reading users.json,", err)
	}

	var response Users
	json.Unmarshal(jsonData, &response)

	/*
	Monitor tasks:
	- The idea is that there are monitor users and each monitor user wants to monitor different sites.
	- So no matter what, each individual monitor typeruns constantly, and whenever something new comes it will automatically send it to the users that want stuff
	*/
	go func() {

		// creates polymarket broadcast group
		broadcastGroup := bcast.NewGroup()
		go broadcastGroup.Broadcast(0)
	
		// initializes polymarket monitor
		go initializePolymarketMonitor(broadcastGroup, &wg)



		MonitorUsers := response.MonitorUsers
	
	
		// goes through the monitor users and sends polymarket data - note that in the future this will check what sites it wants & only send based on that
		for i := range MonitorUsers {
			go func() {
				// joins polymarket broadcast group & sends data whenever new data is found (once again will only do this if in the future it wants polymarket)
				member := broadcastGroup.Join()
				val := member.Recv()

				switch val.(resources.Message).Store {
				case resources.POLYMARKET:
					data := val.(resources.Message).PolymarketMessage
					for _, v := range data {
						polymarketmonitor.SendWebhook(v.Title, v.Slug, v.EndDate, v.Image, MonitorUsers[i].Webhook, MonitorUsers[i].ProfilePic, MonitorUsers[i].GroupName)
					}
				}
			}()
			wg.Add(1)
		}
	}()
	wg.Add(1)
	
	wg.Wait()

}