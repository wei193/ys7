package ys7

import (
	"fmt"
	"testing"
)

//TestLiveGet TestLiveGet
func TestLiveGetLimited(t *testing.T) {
	// tlog("!2")
	live, err := ys.GetLiveLimited(DeviceSerial, ChannelNo, 3600)
	if err != nil {
		tlog(err)
	} else {
		tlog(live)
	}

}

//TestLiveGet TestLiveGet
func TestLiveGet(t *testing.T) {
	// tlog("!2")
	live, err := ys.LiveGet(fmt.Sprintf("%s:%d", DeviceSerial, ChannelNo))
	if err != nil {
		tlog(err)
	} else {
		tlog(live)
	}
}

func TestLiveOpen(t *testing.T) {
	status, err := ys.OpenLive(fmt.Sprintf("%s:%d", DeviceSerial, ChannelNo))
	if err != nil {
		tlog(err)
	} else {
		tlog(status)
	}
}

func TestLiveClose(t *testing.T) {
	status, err := ys.CloseLive(fmt.Sprintf("%s:%d", DeviceSerial, ChannelNo))
	if err != nil {
		tlog(err)
	} else {
		tlog(status)
	}
}
