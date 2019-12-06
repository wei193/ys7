package ys7

import "encoding/json"

//接口地址
const (
	//[子账号]子账号接口
	ACCOUNTCREATE         = "https://open.ys7.com/api/lapp/ram/account/create"         //创建子账户
	ACCOUNTGET            = "https://open.ys7.com/api/lapp/ram/account/get"            //获取指定子账户信息
	ACCOUNTLIST           = "https://open.ys7.com/api/lapp/ram/account/list"           //分页获取应用下的子账户信息列表
	ACCOUNTUPDATEPASSWORD = "https://open.ys7.com/api/lapp/ram/account/updatePassword" //修改子账户密码
	POLICYSET             = "https://open.ys7.com/api/lapp/ram/policy/set"             //设置子账户的授权策略
	STATEMENTADD          = "https://open.ys7.com/api/lapp/ram/statement/add"          //增加子账户授权策略中的授权语句
	STATEMENTDELETE       = "https://open.ys7.com/api/lapp/ram/statement/delete"       //删除子账户授权策略中某个设备的所有授权语句
	RAMTOKENGET           = "https://open.ys7.com/api/lapp/ram/token/get"              //获取子账户AccessToken
	ACCOUNTDELETE         = "https://open.ys7.com/api/lapp/ram/account/delete"         //删除子账户
)

//CreateAccount 创建子账号
func (ys *Ys7) CreateAccount(accountName, password string) (acc *Account, err error) {
	params := make(map[string]interface{})
	params["accountName"] = accountName
	params["password"] = getPasswd(ys.AppKey, password)
	acc = &Account{}
	_, err = ys.authorizeRequset("POST", ACCOUNTCREATE, params, &acc)
	if err != nil {
		return nil, err
	}
	return acc, nil
}

//CreateAccountAndToken 创建子账号并返回子账号对象
func (ys *Ys7) CreateAccountAndToken(accountName, password string) (y *Ys7, err error) {
	params := make(map[string]interface{})
	params["accountName"] = accountName
	params["password"] = getPasswd(ys.AppKey, password)
	acc := &Account{}
	_, err = ys.authorizeRequset("POST", ACCOUNTCREATE, params, &acc)
	if err != nil {
		return nil, err
	}
	ac, err := ys.RAMGetAccessToken(acc.AccountID)
	if err != nil {
		return y, err
	}
	y = &Ys7{
		AppKey:      ys.AppKey,
		AccessToken: ac.AccessToken,
		ExpireTime:  ac.ExpireTime,
		IsRAM:       RAMACC,
		Password:    password,
		AccountID:   acc.AccountID,
	}
	return y, nil
}

// RAMAccountGet  获取单个子账户信息
func (ys *Ys7) RAMAccountGet(accountID, accountName string) (acc *RAMAccount, err error) {
	params := make(map[string]interface{})
	if accountID != "" {
		params["accountId"] = accountID
	} else {
		params["accountName"] = accountName
	}
	acc = &RAMAccount{}
	_, err = ys.authorizeRequset("POST", ACCOUNTGET, params, &acc)
	if err != nil {
		return nil, err
	}
	return
}

// RAMAccountList 获取子账户信息列表
func (ys *Ys7) RAMAccountList(pageStart, pageSize int) (acc []RAMAccount, page Page, err error) {
	params := make(map[string]interface{})
	params["pageStart"] = pageStart
	params["pageSize"] = pageSize
	_, err = ys.authorizeRequset("POST", ACCOUNTLIST, params, &acc, &page)
	if err != nil {
		return nil, page, err
	}
	return acc, page, nil
}

//RAMUpdatePassword 修改当前子账户密码
func (ys *Ys7) RAMUpdatePassword(accountID, old, new string) (err error) {
	params := make(map[string]interface{})
	params["accountId"] = accountID
	params["oldPassword"] = getPasswd(ys.AppKey, old)
	params["newPassword"] = getPasswd(ys.AppKey, new)
	_, err = ys.authorizeRequset("POST", ACCOUNTUPDATEPASSWORD, params, nil)
	if err != nil {
		return err
	}
	return nil
}

//RAMSetSetPolicy 设置子账号权限
func (ys *Ys7) RAMSetSetPolicy(accountID string, policy Policy) (err error) {
	policyByte, err := json.Marshal(policy)
	if err != nil {
		return err
	}
	params := make(map[string]interface{})
	params["accountId"] = accountID
	params["policy"] = string(policyByte)
	_, err = ys.authorizeRequset("POST", POLICYSET, params, nil)
	if err != nil {
		return err
	}
	return nil
}

// RAMAddStatement 增加子账户权限
func (ys *Ys7) RAMAddStatement(accountID string, statement Statement) (err error) {
	StatementByte, err := json.Marshal(statement)
	if err != nil {
		return err
	}
	params := make(map[string]interface{})
	params["accountId"] = accountID
	params["statement"] = string(StatementByte)
	_, err = ys.authorizeRequset("POST", STATEMENTADD, params, nil)
	if err != nil {
		return err
	}
	return nil
}

//RAMDeleteStatement 删除子账户权限
func (ys *Ys7) RAMDeleteStatement(accountID, deviceSerial string) (err error) {
	params := make(map[string]interface{})
	params["accountId"] = accountID
	params["deviceSerial"] = deviceSerial
	_, err = ys.authorizeRequset("POST", STATEMENTDELETE, params, nil)
	if err != nil {
		return err
	}
	return nil
}

//RAMGetAccessToken 获取B模式子账户accessToken
func (ys *Ys7) RAMGetAccessToken(accountID string) (ac *AccessToken, err error) {
	params := make(map[string]interface{})
	params["accountId"] = accountID
	ac = &AccessToken{}
	_, err = ys.authorizeRequset("POST", RAMTOKENGET, params, &ac)
	if err != nil {
		return nil, err
	}
	return ac, nil
}

// DeleteAccount  删除子账户
func (ys *Ys7) DeleteAccount(accountID string) (err error) {
	params := make(map[string]interface{})
	params["accountId"] = accountID
	_, err = ys.authorizeRequset("POST", ACCOUNTDELETE, params, nil)
	if err != nil {
		return err
	}
	return nil
}
