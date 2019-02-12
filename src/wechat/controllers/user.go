package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

type WxOpenIdData struct {
	Openid string `json:"openid"`
}

/**
* 微信获取openId host
 */
const WX_OPENID_HOST = "https://api.weixin.qq.com/sns/jscode2session"

//获取用户的openId
func (this *UserController) GetOpenId() {
	appId := this.GetString("appid")
	secret := this.GetString("secret")
	js_code := this.GetString("js_code")
	url := WX_OPENID_HOST + "?appid=" + appId + "&secret=" + secret + "&js_code=" + js_code + "&grant_type=authorization_code"

	resp, err := http.Get(url)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"success": -1, "data": nil}
		this.ServeJSON()
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"success": -1, "data": nil}
		this.ServeJSON()
	}
	var wxOpenidData WxOpenIdData
	if err := json.Unmarshal(body, &wxOpenidData); err != nil {

		this.Data["json"] = map[string]interface{}{"success": -1, "data": nil}
		this.ServeJSON()
	}
	this.Data["json"] = map[string]interface{}{"success": 1, "data": wxOpenidData}
	this.ServeJSON()
}
