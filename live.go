package ys7

//接口地址
const (
	//[直播管理]直播接口
	LIVEVIDEOLIST      = "https://open.ys7.com/api/lapp/live/video/list"      //获取用户下的直播地址列表
	LIVEADDRESSLIMITED = "https://open.ys7.com/api/lapp/live/address/limited" //获取指定有效期的直播地址
	LIVEVIDEOOPEN      = "https://open.ys7.com/api/lapp/live/video/open"      //批量开通直播功能
	LIVEVIDEOCLOSE     = "https://open.ys7.com/api/lapp/live/video/close"     //批量关闭直播功能
	LIVEGET            = "https://open.ys7.com/api/lapp/live/address/get"     //批量获取设备的直播地址信息

)

//ListLiveVideo 获取直播列表
func (ys *Ys7) ListLiveVideo(pageStart, pageSize int) (live []Live, err error) {
	params := make(map[string]interface{})
	params["pageStart"] = pageStart
	params["pageSize"] = pageSize
	_, err = ys.authorizeRequset("POST", LIVEVIDEOLIST, params, &live)
	if err != nil {
		return
	}
	return
}

//GetLiveLimited 获取指定有效期直播流地址
func (ys *Ys7) GetLiveLimited(deviceSerial string, channelNo, expireTime int) (live Live, err error) {
	params := make(map[string]interface{})
	params["deviceSerial"] = deviceSerial
	params["channelNo"] = channelNo
	if expireTime != 0 {
		params["expireTime"] = expireTime
	}
	_, err = ys.authorizeRequset("POST", LIVEADDRESSLIMITED, params, &live)
	if err != nil {
		return
	}
	return
}

//OpenLive 批量开通直播功能
func (ys *Ys7) OpenLive(source string) (states []LiveState, err error) {
	params := make(map[string]interface{})
	params["source"] = source
	_, err = ys.authorizeRequset("POST", LIVEVIDEOOPEN, params, &states)
	if err != nil {
		return
	}
	return
}

//CloseLive 批量关闭直播功能
func (ys *Ys7) CloseLive(source string) (states []LiveState, err error) {
	params := make(map[string]interface{})
	params["source"] = source
	_, err = ys.authorizeRequset("POST", LIVEVIDEOCLOSE, params, &states)
	if err != nil {
		return
	}
	return
}

//LiveGet 批量获取设备的直播地址信息
func (ys *Ys7) LiveGet(source string) (states []Live, err error) {
	params := make(map[string]interface{})
	params["source"] = source
	_, err = ys.authorizeRequset("POST", LIVEGET, params, &states)
	if err != nil {
		return
	}
	return
}
