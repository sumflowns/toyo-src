package check_code

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sumflowns/toyo-src/book-ticket-common/basic/common"
	"github.com/sumflowns/toyo-src/book-ticket-common/basic/utils/conversation"
	"github.com/sumflowns/toyo-src/book-ticket-common/basic/utils/http_util"
	"github.com/sumflowns/toyo-src/book-ticket-common/ticket/static/api"
	"io/ioutil"
	"log"
	"math/rand"
	"mime/multipart"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type VCode struct {
	Image             string // 验证码
	Timestamp         string // 请求验证码时的时间戳
	CallbackParameter string // 请求验证码时验证时间戳的回调参数
}

type CodeResponse struct {
	Image          string
	Result_message string
	Result_code    string
}

type CheckResponse struct {
	Code    int64
	Data    []string
	Massage string
}
type ImgCheck struct {
	Msg string
	Result []int
}

//获取验证码
func getVCode(conversation *conversation.Conversation) (vCode VCode, err error) {
	log.Println("[getVCode] 获取验证码...")
	defer func() {
		if re := recover(); re != nil {
			if err == nil {
				err = errors.New(fmt.Sprintf("%v", re))
			}
		}
	}()
	vCode = VCode{}
	//生成随机数
	cbpara := "jQuery1910"
	for i := 0; i < 16; i++ {
		cbpara += strconv.Itoa(rand.Intn(9))
	}
	vCode.CallbackParameter = cbpara + "_" + strconv.FormatInt(time.Now().UnixNano()/10000000, 10)
	vCode.Timestamp = strconv.FormatInt(time.Now().UnixNano()/10000000, 10)
	req, _ := http.NewRequest(http.MethodGet, api.GET_CHECK_CODE+"?login_site=E&module=login&rand=sjrand&"+vCode.Timestamp+"=&callback="+vCode.CallbackParameter+"&_="+vCode.Timestamp, nil)
	req.Header.Set("Origin", "https://kyfw.12306.cn")
	req.Header.Set("Referer", "https://kyfw.12306.cn/otn/resources/login.html")
	http_util.SetReqHeader(req)
	resp, err := conversation.Client.Do(req)
	if err != nil {
		return
	}
	http_util.CookieChange(conversation, resp.Cookies())
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("error: %s", err)
		return
	}
	str := string(body)
	codeResponse := &CodeResponse{}
	err = http_util.TimeStrapStringJsonToBean(str, codeResponse)
	if err != nil {
		log.Printf("error: %s", err)
		return
	}
	if codeResponse.Result_code != "0" {
		err = errors.New(codeResponse.Result_message)
	}
	vCode.Image = codeResponse.Image
	return
}

//验证码识别
func checkCodeIdentify(vCode VCode) (result string, err error) {
	log.Println("[checkCodeIdentify]识别验证码...")
	defer func() {
		if re := recover(); re != nil {
			if err == nil {
				err = errors.New(fmt.Sprintf("[checkCodeIdentify] %v", re))
				return
			}
		}
	}()
	defer func() {
		if re := recover(); re != nil {
			if err != nil {
				err = errors.New(fmt.Sprintf("%v", re))
			}
		}
	}()
	data := url.Values{}
	data.Add("img", vCode.Image)
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("img", vCode.Image)
	err = writer.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	client := &http.Client {
	}
	req, err := http.NewRequest("POST", common.POST_CHECK_CODE_FROM_MY_SERVER, payload)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	resp, err := client.Do(req)

	//resp, err := http.PostForm(common.POST_CHECK_CODE_FROM_MY_SERVER, data)

	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	log.Println(string(body))
	if err != nil {
		log.Println(err)
		return
	}
	//checkResponse := CheckResponse{}
	imgCheck:=ImgCheck{}
	err = json.Unmarshal(body, &imgCheck)
	if err != nil {
		log.Printf("error: %s", err)
		return
	}
	log.Println(imgCheck)

	var codeList []string
	for k, onset := range imgCheck.Result{
		codeList[k]=string(onset)
	}
	return simulatedClick(codeList)
}

//模拟用户点击验证码
func simulatedClick(codeList []string) (result string, err error) {
	log.Println("[simulatedClick]模拟点击验证码...")
	if len(codeList) == 0 {
		err = errors.New("codeList is nil")
		return
	}
	offsetsX := "0"
	offsetsY := "0"
	for _, ofset := range codeList {
		switch ofset {
		case "1":
			offsetsY = "40"
			offsetsX = "40"
			break
		case "2":
			offsetsY = "40"
			offsetsX = "110"
			break
		case "3":
			offsetsY = "40"
			offsetsX = "184"
			break
		case "4":
			offsetsY = "40"
			offsetsX = "256"
			break
		case "5":
			offsetsY = "110"
			offsetsX = "40"
			break
		case "6":
			offsetsY = "110"
			offsetsX = "110"
			break
		case "7":
			offsetsY = "110"
			offsetsX = "184"
			break
		case "8":
			offsetsY = "110"
			offsetsX = "256"
			break
		default:
			offsetsY = "-1"
			offsetsX = "-1"
			break
		}
		if len(result) == 0 {
			result += offsetsX + "," + offsetsY
		} else {
			result += "," + offsetsX + "," + offsetsY
		}
	}
	log.Println("result:(" + result + ")")
	return

}

//向12306验证-验证码
func checkCodeTo12306(conversation *conversation.Conversation, code VCode, strIdentify string) (err error) {
	log.Println("[checkCodeTo12306]向12306验证-验证码...")
	defer func() {
		if re := recover(); re != nil {
			if err == nil {
				err = errors.New(fmt.Sprintf("[checkCodeTo12306] %v", re))
				return
			}
		}
	}()
	data := url.Values{}
	data.Set("callback", code.CallbackParameter)
	data.Set("answer", strIdentify)
	data.Set("rand", "sjrand")
	data.Set("login_site", "E")
	data.Set("_", strconv.FormatInt(time.Now().UnixNano()/10000000, 10))
	req, _ := http.NewRequest(http.MethodPost, "https://kyfw.12306.cn/passport/captcha/captcha-check", strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	//req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("Origin", "https://kyfw.12306.cn")
	req.Header.Set("Referer", "https://kyfw.12306.cn/otn/resources/login.html")
	req.Header.Set("Content-Length", strconv.Itoa(len(data.Encode())))
	http_util.AddReqCookie(conversation.C, req)
	http_util.SetReqHeader(req)
	resp, err := conversation.Client.Do(req)
	if err != nil {
		log.Printf("error: %s\n", err.Error())
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		return
	}
	http_util.CookieChange(conversation, resp.Cookies())
	log.Println(string(body) + "\n")
	return
}

func CheckCode(conversation *conversation.Conversation) (code string, err error) {
	defer func() {
		if re := recover(); re != nil {
			if err == nil {
				err = errors.New(fmt.Sprintf("[CheckCode] %v", re))
				return
			}
		}
	}()
	vCode, err := getVCode(conversation)
	if err != nil {
		fmt.Printf("获取验证码失败正在重试... %s\n", err.Error())
		return
	}
	result, err := checkCodeIdentify(vCode)
	if err != nil {
		fmt.Printf("验证-验证码失败，开始重试获取验证码...  %s\n", err.Error())
		return
	}
	err = checkCodeTo12306(conversation, vCode, result)
	if err != nil {
		fmt.Printf("将验证码发到12306验证，出错... %s\n", err.Error())
		return
	}
	return result, nil
}
