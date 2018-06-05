package ys7

//接口地址
const (
	//[设备]管理
	DEVICEADD     = "https://open.ys7.com/api/lapp/device/add"         //添加设备到账号下
	DEVICEDELETE  = "https://open.ys7.com/api/lapp/device/delete"      //删除账号下指定设备
	DEVICEUPDATE  = "https://open.ys7.com/api/lapp/device/name/update" //修改设备名称
	DEVICECAPTURE = "https://open.ys7.com/api/lapp/device/capture"     //抓拍设备的当前画面
	URLIPCADD     = "https://open.ys7.com/api/lapp/device/ipc/add"     //NVR设备关联IPC
	URLIPCDELETE  = "hhttps://open.ys7.com/api/lapp/device/ipc/delete" //NVR设备删除IPC
	//修改设备视频加密密码
	//生成设备扫描配网二维码二进制数据
	CAMERANAMEUPDATE = "https://open.ys7.com/api/lapp/camera/name/update" // 修改通道名称

	//[设备]查询
	DEVICELIST = "https://open.ys7.com/api/lapp/device/list" //获取用户下的设备列表
	DEVICEINFO = "https://open.ys7.com/api/lapp/device/info" //获取指定设备的信息
	CAMERALIST = "https://open.ys7.com/api/lapp/camera/list" //获取用户下的摄像头列表
	//设备互联互通根据UUID查询抓拍的图片
	//根据序列号获取设备的状态信息
	DEVICECAMERALIST = "https://open.ys7.com/api/lapp/device/camera/list" //根据序列号获取设备的通道信息
	//根据设备型号以及设备版本号查询设备是否支持萤石协议
	//根据时间获取录像信息

	//[设备]配置
	// --全部待实现

	//[设备]升级
	// --全部待实现
)

//AddDevice 添加设备
func (ys *Ys7) AddDevice(deviceSerial, validateCode string) (err error) {
	params := make(map[string]interface{})
	params["deviceSerial"] = deviceSerial
	params["validateCode"] = validateCode
	_, err = ys.authorizeRequset("POST", DEVICEADD, params, nil)
	if err != nil {
		return err
	}
	return nil
}

//DeleteDevice 删除设备
func (ys *Ys7) DeleteDevice(deviceSerial string) (err error) {
	params := make(map[string]interface{})
	params["deviceSerial"] = deviceSerial
	_, err = ys.authorizeRequset("POST", DEVICEDELETE, params, nil)
	if err != nil {
		return err
	}
	return nil
}

//UpdateDeviceName 修改通道名称
func (ys *Ys7) UpdateDeviceName(deviceSerial, deviceName string) (err error) {
	params := make(map[string]interface{})
	params["deviceSerial"] = deviceSerial
	params["deviceName"] = deviceName

	_, err = ys.authorizeRequset("POST", DEVICEUPDATE, params, nil)
	if err != nil {
		return
	}
	return nil
}

// GetDeviceCapture 设备抓拍图片
func (ys *Ys7) GetDeviceCapture(deviceSerial string, channelNo int) (picURL string, err error) {
	params := make(map[string]interface{})
	params["deviceSerial"] = deviceSerial
	params["channelNo"] = channelNo
	type st struct {
		PicURL string `json:"picUrl"`
	}
	var in st
	_, err = ys.authorizeRequset("POST", DEVICECAPTURE, params, &in)
	if err != nil {
		return
	}
	return in.PicURL, nil
}

//AddDeviceIpc NVR设备关联IPC
func (ys *Ys7) AddDeviceIpc(deviceSerial, ipcSerial string, channelNo int, validateCode string) (err error) {
	params := make(map[string]interface{})
	params["deviceSerial"] = deviceSerial
	params["ipcSerial"] = ipcSerial
	params["channelNo"] = channelNo
	params["validateCode"] = validateCode
	_, err = ys.authorizeRequset("POST", URLIPCADD, params, nil)
	if err != nil {
		return err
	}
	return nil
}

//DeleteDeviceIpc NVR设备删除IPC
func (ys *Ys7) DeleteDeviceIpc(deviceSerial, ipcSerial string, channelNo int) (err error) {
	params := make(map[string]interface{})
	params["deviceSerial"] = deviceSerial
	params["ipcSerial"] = ipcSerial
	params["channelNo"] = channelNo
	_, err = ys.authorizeRequset("POST", URLIPCDELETE, params, nil)
	if err != nil {
		return err
	}
	return nil
}

//GetAllDeviceList 获取所有设备列表
func (ys *Ys7) GetAllDeviceList() (devices []Device, err error) {
	var page Page
	devices, page, err = ys.GetDeviceList(0, 50)
	if err != nil {
		return
	}
	for pageNum := 1; pageNum < page.Total/50; pageNum++ {
		list, _, err := ys.GetDeviceList(pageNum, 50)
		if err != nil {
			return nil, err
		}
		devices = append(devices, list...)
	}
	return
}

//GetDeviceList 获取设备列表
func (ys *Ys7) GetDeviceList(pageStart, pageSize int) (devices []Device, page Page, err error) {
	params := make(map[string]interface{})
	params["pageStart"] = pageStart
	params["pageSize"] = pageSize
	_, err = ys.authorizeRequset("POST", DEVICELIST, params, &devices, &page) //获取用户下的设备列表
	if err != nil {
		return nil, page, err
	}
	return devices, page, nil
}

//InfoDevice 获取单个设备信息
func (ys *Ys7) InfoDevice(deviceSerial string) (deviceinfo DeviceInfo, err error) {
	params := make(map[string]interface{})
	params["deviceSerial"] = deviceSerial
	_, err = ys.authorizeRequset("POST", DEVICEINFO, params, &deviceinfo)
	if err != nil {
		return
	}
	return
}

//GetAllCameraList 获取所有摄像头列表
func (ys *Ys7) GetAllCameraList() (cameras []Camera, err error) {
	var page Page
	cameras, page, err = ys.GetCameraList(0, 50)
	if err != nil {
		return
	}
	for pageNum := 1; pageNum < page.Total/50; pageNum++ {
		list, _, err := ys.GetCameraList(pageNum, 50)
		if err != nil {
			return nil, err
		}
		cameras = append(cameras, list...)
	}
	return
}

//GetCameraList 获取摄像头列表
func (ys *Ys7) GetCameraList(pageStart, pageSize int) (cameras []Camera, page Page, err error) {
	params := make(map[string]interface{})
	params["pageStart"] = pageStart
	params["pageSize"] = pageSize
	_, err = ys.authorizeRequset("POST", CAMERALIST, params, &cameras, &page)
	if err != nil {
		return nil, page, err
	}
	return cameras, page, nil
}

//GetDeviceCameraList 获取指定设备的通道信息
func (ys *Ys7) GetDeviceCameraList(deviceSerial string) (cameras []Camera, err error) {
	params := make(map[string]interface{})
	params["deviceSerial"] = deviceSerial
	_, err = ys.authorizeRequset("POST", CAMERALIST, params, &cameras)
	if err != nil {
		return nil, err
	}
	return cameras, nil
}

//UpdateCameraName 修改通道名称
func (ys *Ys7) UpdateCameraName(deviceSerial, name string, channelNo int) (err error) {
	params := make(map[string]interface{})
	params["deviceSerial"] = deviceSerial
	params["name"] = name
	params["channelNo"] = channelNo

	_, err = ys.authorizeRequset("POST", CAMERANAMEUPDATE, params, nil)
	if err != nil {
		return
	}
	return nil
}
