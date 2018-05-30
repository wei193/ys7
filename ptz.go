package ys7

//接口地址
const (
	// [设备]云台
	URLPTZSTAR      = "https://open.ys7.com/api/lapp/device/ptz/start"
	URLPTZSTOP      = "https://open.ys7.com/api/lapp/device/ptz/stop"
	URLPTZMIRROR    = "https://open.ys7.com/api/lapp/device/ptz/mirror"
	URLPRESETADD    = "https://open.ys7.com/api/lapp/device/preset/add"
	URLPPRESETMOVE  = "https://open.ys7.com/api/lapp/device/preset/move"
	URLPPRESETCLEAR = "https://open.ys7.com/api/lapp/device/preset/clear"
)

//StartPtz 开始云台控制
func (ys *Ys7) StartPtz(deviceSerial string, channelNo, direction, speed int) (err error) {
	params := make(map[string]interface{})
	params["deviceSerial"] = deviceSerial
	params["channelNo"] = channelNo
	params["direction"] = direction
	params["speed"] = speed

	_, err = ys.authorizeRequset("POST", URLPTZSTAR, params, nil)
	if err != nil {
		return
	}
	return nil
}

//StopPtz 停止云台转动
func (ys *Ys7) StopPtz(deviceSerial string, channelNo, direction int) (err error) {
	params := make(map[string]interface{})
	params["deviceSerial"] = deviceSerial
	params["channelNo"] = channelNo
	params["direction"] = direction

	_, err = ys.authorizeRequset("POST", URLPTZSTOP, params, nil)
	if err != nil {
		return
	}
	return nil
}

//MirrorPtz 镜像翻转
func (ys *Ys7) MirrorPtz(deviceSerial string, channelNo, command int) (err error) {
	params := make(map[string]interface{})
	params["deviceSerial"] = deviceSerial
	params["channelNo"] = channelNo
	params["command"] = command

	_, err = ys.authorizeRequset("POST", URLPTZMIRROR, params, nil)
	if err != nil {
		return
	}
	return nil
}

//AddPreset 添加预置点
func (ys *Ys7) AddPreset(deviceSerial string, channelNo int) (index int, err error) {
	params := make(map[string]interface{})
	params["deviceSerial"] = deviceSerial
	params["channelNo"] = channelNo

	type st struct {
		Index int `json:"index"`
	}
	var data st
	_, err = ys.authorizeRequset("POST", URLPRESETADD, params, &data)
	if err != nil {
		return
	}
	return data.Index, nil
}

//MovePreset 调用预置点
func (ys *Ys7) MovePreset(deviceSerial string, channelNo, index int) (err error) {
	params := make(map[string]interface{})
	params["deviceSerial"] = deviceSerial
	params["channelNo"] = channelNo
	params["index"] = index

	_, err = ys.authorizeRequset("POST", URLPPRESETMOVE, params, nil)
	if err != nil {
		return
	}
	return nil
}

//ClearPreset 清除预置点
func (ys *Ys7) ClearPreset(deviceSerial string, channelNo, index int) (err error) {
	params := make(map[string]interface{})
	params["deviceSerial"] = deviceSerial
	params["channelNo"] = channelNo
	params["index"] = index

	_, err = ys.authorizeRequset("POST", URLPPRESETCLEAR, params, nil)
	if err != nil {
		return
	}
	return nil
}
