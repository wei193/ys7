package ys7

//接口地址
const (
	//[流量管理]流量接口
	TRAFFICUSERTOTAL    = "https://open.ys7.com/api/lapp/traffic/user/total"    //该接口用于查询账号下流量消耗汇总
	TRAFFICUSERDETAIL   = "https://open.ys7.com/api/lapp/traffic/user/detail"   //查询账户下流量消耗详情
	TRAFFICDAYDETAIL    = "https://open.ys7.com/api/lapp/traffic/day/detail"    //查询账户下某天流量消耗详情
	TRAFFICDEVICEDETAIL = "https://open.ys7.com/api/lapp/traffic/device/detail" //查询指定设备在某一时间段消耗流量数据
)

//GetTrafficUserTotal 该接口用于查询账号下流量消耗汇总
func (ys *Ys7) GetTrafficUserTotal() (tra *TrafficTotal, err error) {
	params := make(map[string]interface{})
	tra = &TrafficTotal{}
	_, err = ys.authorizeRequset("POST", TRAFFICUSERTOTAL, params, &tra)
	if err != nil {
		return
	}
	return
}

//GetTrafficUserDetail 查询账户下流量消耗详情
func (ys *Ys7) GetTrafficUserDetail(startTime, endTime int64, pageStart, pageSize int) (traffic []TrafficUserDetail, err error) {
	params := make(map[string]interface{})
	params["startTime"] = startTime
	params["endTime"] = endTime
	params["pageStart"] = pageStart
	params["pageSize"] = pageSize
	_, err = ys.authorizeRequset("POST", TRAFFICUSERDETAIL, params, &traffic)
	if err != nil {
		return
	}
	return
}

//GetTrafficDayDetail 查询账户下某天流量消耗详情
func (ys *Ys7) GetTrafficDayDetail(flowTime int64, pageStart, pageSize int) (traffic []TrafficUserDetail, err error) {
	params := make(map[string]interface{})
	params["flowTime"] = flowTime
	params["pageStart"] = pageStart
	params["pageSize"] = pageSize
	_, err = ys.authorizeRequset("POST", TRAFFICDAYDETAIL, params, &traffic)
	if err != nil {
		return
	}
	return

}

//GetTrafficDeviceDetail 查询指定设备在某一时间段消耗流量数据
func (ys *Ys7) GetTrafficDeviceDetail(deviceSerial string, startTime, endTime int64, pageStart, pageSize int) (traffic []TrafficUserDetail, err error) {
	params := make(map[string]interface{})
	params["deviceSerial"] = deviceSerial
	params["startTime"] = startTime
	params["endTime"] = endTime
	params["pageStart"] = pageStart
	params["pageSize"] = pageSize
	_, err = ys.authorizeRequset("POST", TRAFFICDEVICEDETAIL, params, &traffic)
	if err != nil {
		return
	}
	return
}
