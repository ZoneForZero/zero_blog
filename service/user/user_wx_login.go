package service

import (
	"bytes"
	"crypto/sha1"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	MODEL "zero_blog/model"
)

type WXLoginResp struct {
	OpenId     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionId    string `json:"unionid"`
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

// 这个函数以 code 作为输入, 返回调用微信接口得到的对象指针和异常情况
func WXLogin(code string, userName string) (*WXLoginResp, error) {
	url := "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
	// 合成url, 这里的appId和secret是在微信公众平台上获取的
	url = fmt.Sprintf(url, os.Getenv("WxAppId"), os.Getenv("WxAppSecret"), code)
	// 创建http get请求
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 解析http请求中body 数据到我们定义的结构体中
	wxResp := WXLoginResp{}
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&wxResp); err != nil {
		return nil, err
	}
	// 判断微信接口返回的是否是一个异常情况
	if wxResp.ErrCode != 0 {
		return nil, errors.New(fmt.Sprintf("ErrCode:%d  ErrMsg:%s", wxResp.ErrCode, wxResp.ErrMsg))
	}
	// 账号没错，创建用户
	user := MODEL.User{
		NickName: userName,
		OpenId:   wxResp.OpenId,
	}
	if createUserErr := user.UpsertUserByOpenId(); createUserErr != nil {
		return nil, createUserErr
	}

	// fmt.Printf("\n 1: %v 2: %v 3: %v 4: %v 5:%v\n", wxResp.OpenId, wxResp.SessionKey, wxResp.UnionId, wxResp.ErrCode, wxResp.ErrMsg)
	return &wxResp, nil
}

// 校验微信返回的用户数据
func ValidateUserInfo(rawData, sessionKey, signature string) bool {
	signature2 := GetSha1(rawData + sessionKey)
	return signature == signature2
}

// SHA-1 加密
func GetSha1(str string) string {
	data := []byte(str)
	has := sha1.Sum(data)
	res := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return res
}

func (WxRes *WXLoginResp) String() string {
	b, err := json.Marshal(*WxRes)
	if err != nil {
		return fmt.Sprintf("%+v", *WxRes)
	}
	var out bytes.Buffer
	err = json.Indent(&out, b, "", "    ")
	if err != nil {
		return fmt.Sprintf("%+v", *WxRes)
	}
	return out.String()
}
