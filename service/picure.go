package service

import (
	"car/conf"
	"car/serializer"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type PictureService struct {
	PictureURL string `json:"picture_url" form:"picture_url"`
}

type accessTokenInfo struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

type WordResult struct {
	Number string `json:"number"`
}
type Data struct {
	WordsResult WordResult `json:"words_result"`
}

func (pictureService *PictureService) PictureOCR() serializer.Response {
	accessToken := GetAccessToken(conf.ClientId, conf.BDSecretKey)
	number, err := GetPlate(pictureService.PictureURL, accessToken)
	if err != nil {
		return serializer.Response{
			Status: 400,
			Data:   nil,
			Msg:    "车牌号识别出错",
			Error:  err,
		}
	}
	return serializer.Response{
		Status: 200,
		Data:   number,
		Msg:    "success",
	}
}

// GetAccessToken 获取accessToken
func GetAccessToken(appKey string, appSecret string) (accessToken string) {
	url := "https://aip.baidubce.com/oauth/2.0/token?grant_type=client_credentials&client_id=" + appKey + "&client_secret=" + appSecret

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return ""
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
		return ""
	}
	info := accessTokenInfo{}
	json.Unmarshal(data, &info)
	//log.Print("请求accessToken返回的数据:", string(data))
	return info.AccessToken
}

// GetPlate 识别车牌号
/*	读取上述图片url地址，获取图片的二进制流信息
	进行base64处理
	进行urlencode处理
	传入access_token和上一步的结果，调用车牌识别api，再根据这个token去调用车牌识别的接口
*/
func GetPlate(pictureUrl string, accessToken string) (plate string, err error) {
	rsp, err := http.Get(pictureUrl)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	image, _ := ioutil.ReadAll(rsp.Body)
	imageValue, err2 := url.Parse(base64.StdEncoding.EncodeToString(image)) //对图片进行base64处理
	if err2 != nil {
		log.Fatal(err)
		return "", err
	}
	toUrl := "https://aip.baidubce.com/rest/2.0/ocr/v1/license_plate?access_token=" + accessToken
	values := url.Values{}
	values.Add("image", imageValue.EscapedPath())
	values.Add("multi_detect", "false")
	rsp2, err := http.PostForm(toUrl, values)
	defer rsp2.Body.Close()
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	data, err := ioutil.ReadAll(rsp2.Body)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	//log.Println("请求车牌返回的数据:",string(data))
	m := Data{}
	err = json.Unmarshal(data, &m)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	//log.Println(m)
	return m.WordsResult.Number, nil
}
