//package main
//
//import (
//	"bytes"
//	"fmt"
//	"io/ioutil"
//	"mime/multipart"
//	"net/http"
//)
//
//func main() {
//
//	url := "http://acm.hdu.edu.cn/submit.php?action=submit"
//	method := "POST"
//
//	payload := &bytes.Buffer{}
//	writer := multipart.NewWriter(payload)
//	_ = writer.WriteField("check", "0")
//	_ = writer.WriteField("_usercode", "JTIzaW5jbHVkZSUzQ2JpdHMlMkZzdGRjJTJCJTJCLmglM0UlMEF1c2luZyUyMG5hbWVzcGFjZSUyMHN0ZCUzQiUwQWludCUyMG1haW4oKSUwQSU3QiUwQWludCUyMGElMkNiJTNCJTBBd2hpbGUoY2luJTIwJTNFJTNFJTIwYSUyMCUzRSUzRSUyMGIpJTBBJTdCJTBBY291dCUyMCUzQyUzQyUyMGElMjAlMkIlMjBiJTIwJTNDJTNDJTIwZW5kbCUzQiUwQSU3RCUwQXJldHVybiUyMDAlM0IlMEElN0Q=")
//	_ = writer.WriteField("problemid", "1000")
//	_ = writer.WriteField("language", "0")
//	err := writer.Close()
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	client := &http.Client{}
//	req, err := http.NewRequest(method, url, payload)
//
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	req.Header.Add("Cookie", "PHPSESSID=kr1fka674g2jb179u02pvkkru2; exesubmitlang=0")
//
//	req.Header.Set("Content-Type", writer.FormDataContentType())
//	res, err := client.Do(req)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	defer res.Body.Close()
//
//	fmt.Println(res.Header.Get("Cookie"))
//	fmt.Println(res.StatusCode)
//	_, err = ioutil.ReadAll(res.Body)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	//fmt.Println(string(body))
//}

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

func main() {

	url := "http://acm.hdu.edu.cn/userloginex.php?action=login&cid=0&notice=0"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("username", "bqxtt233")
	_ = writer.WriteField("userpass", "tcg19981108")
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	//client := &http.Client{
	//	CheckRedirect: func(req *http.Request, via []*http.Request) error {
	//		return http.ErrUseLastResponse /* 不进入重定向 */
	//	},
	//}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := http.DefaultTransport.RoundTrip(req)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	fmt.Printf("res: %v\n", res.Header)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
