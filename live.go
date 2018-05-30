package ys7

//接口地址
const (
	//[直播管理]直播接口
	// URLVIDEOLIST  = "https://open.ys7.com/api/lapp/live/video/list"
	LIVEADDRESSLIMITED = "https://open.ys7.com/api/lapp/live/address/limited" //获取指定有效期的直播地址
	// URLVIDEOOPEN  = "https://open.ys7.com/api/lapp/live/video/open"
	// URLVIDEOCLOSE = "https://open.ys7.com/api/lapp/live/video/close"
	// URLLIVEGET    = "https://open.ys7.com/api/lapp/live/address/get"
)

//GetLimited 获取直播流地址
func (ys *Ys7) GetLimited(deviceSerial string, channelNo, expireTime int) (err error) {
	params := make(map[string]interface{})
	params["deviceSerial"] = deviceSerial
	params["channelNo"] = channelNo
	params["expireTime"] = expireTime
	_, err = ys.authorizeRequset("POST", LIVEADDRESSLIMITED, params, nil)
	if err != nil {
		return err
	}
	// fmt.Println(4, string(buf))
	return
}
