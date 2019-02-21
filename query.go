package ys7

//接口地址
const (
	//DEVICELIST			= "https://open.ys7.com/api/lapp/device/list"
	//DEVICEINFO			= "https://open.ys7.com/api/lapp/device/info"
)

// ListDevice 获取设备列表
func (ys *Ys7) ListDevice(pageStart, pageSize int) (deviceList []DeviceList, err error) {
	params := make(map[string]interface{})
	params["pageStart"] = pageStart
	params["pageSize"] = pageSize
	_, err = ys.authorizeRequset("POST", DEVICEADD, params, &deviceList)
	if err != nil {
		return
	}
	return
}

// DeviceInfo 获取单个设备信息
func (ys *Ys7) DeviceInfo(deviceSerial int) (dev *DeviceInfo, err error) {
	params := make(map[string]interface{})
	params["deviceSerial"] = deviceSerial
	dev = &DeviceInfo{}
	_, err = ys.authorizeRequset("POST", DEVICEINFO, params, &dev)
	if err != nil {
		return
	}
	return
}

//
