package polymarketmonitor

import (
	"time"

	"github.com/akulsharma1/godiscord"
)

func SendWebhook(title string, slug string, enddate string, image string, webhook string, profilepic string, groupname string) {
	ts := time.Now()
	wbTimestamp := ts.Format("2006-01-02 3:04:05 PM")

	e := godiscord.NewEmbed("New Bet Found! (Polymarket New)", title, "https://polymarket.com/market/"+slug)
	e.AddField("End Date", enddate, true)
	e.SetThumbnail(image)

	e.SetFooter(groupname+" monitor • Created by splash#0003 • "+wbTimestamp, profilepic)
	e.SendToWebhook(webhook)
}