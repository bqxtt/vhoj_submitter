package hdu

import (
	"github.com/bqxtt/vhoj_common/pkg/common/constants"
	"github.com/bqxtt/vhoj_submitter/pkg/loginer"
)

type HDULoginer struct {
	loginer.DefaultLoginerImpl
}

func (H *HDULoginer) GetOJInfo() *constants.RemoteOJInfo {
	panic("implement me")
}

func (H *HDULoginer) Login() error {
	return nil
	//url := constants.HDUInfo.Host + constants.HDUInfo.LoginUrl
	//method := "POST"
	//
	//payload := &bytes.Buffer{}
	//writer := multipart.NewWriter(payload)
	//_ = writer.WriteField("username", "bqxtt233")
	//_ = writer.WriteField("userpass", "tcg19981108")
	//err := writer.Close()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	////client := &http.Client{
	////	CheckRedirect: func(req *http.Request, via []*http.Request) error {
	////		return http.ErrUseLastResponse /* 不进入重定向 */
	////	},
	////}
	//req, err := http.NewRequest(method, url, payload)
	//
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//req.Header.Set("Content-Type", writer.FormDataContentType())
	//res, err := http.DefaultTransport.RoundTrip(req)
	//
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//defer res.Body.Close()
	//
	//fmt.Printf("res: %v\n", res.Header)
	//
	//body, err := ioutil.ReadAll(res.Body)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(string(body))
}
