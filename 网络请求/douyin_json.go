package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	url := "https://www.douyin.com/aweme/v1/web/aweme/post/?device_platform=webapp&aid=6383&channel=channel_pc_web&sec_user_id=MS4wLjABAAAAlNxI61tpwf2_lkzPzboEd-s2Myl8tiAJGOnwylZMcfo&max_cursor=1685438351000&locate_item_id=7242994723852520744&locate_query=false&show_live_replay_strategy=1&count=20&publish_video_strategy_type=2&pc_client_type=1&version_code=170400&version_name=17.4.0&cookie_enabled=true&screen_width=1536&screen_height=864&browser_language=zh-CN&browser_platform=Win32&browser_name=Edge&browser_version=114.0.1823.43&browser_online=true&engine_name=Blink&engine_version=114.0.0.0&os_name=Windows&os_version=10&cpu_core_num=8&device_memory=8&platform=PC&downlink=8.85&effective_type=4g&round_trip_time=100&webid=7234137196285117985&msToken=81bCCBkGm0VNC9va2OlelVahVxoJJcLDo8Waz39Lky-K2ZEqiR0K6Pyb-R3iSlF7m2Nbyyw37iP50RJK0xot5P9db6wFPMQ8aoA3yxnRkznoKu9oBFCO4nI=&X-Bogus=DFSzswVu33sANGOVtrtP6BesEuwX"
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	// Set request headers
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36 Edg/114.0.1823.43")
	req.Header.Set("host", "www.douyin.com")
	req.Header.Set("Cookie", "csrf_session_id=0125cc1cf67b05e9c55a95b6e2c3b6cc; ttwid=1%7CxPbdTdftSvxTd9EqffgSo4XyJOZ5oE6bOGHqMW_NLoI%7C1684328837%7C77015e521ac6645a554e123486549d5a29e367aa50d91ef88ddc6cac284d9e8b; douyin.com; passport_csrf_token=db8f58ae29c8d4151338d0db8704d38c; passport_csrf_token_default=db8f58ae29c8d4151338d0db8704d38c; s_v_web_id=verify_lhrpwk91_LeIqvFm4_Mr9J_4LS6_A4Z1_jRuzZLTZ02rO; n_mh=di2ji2lzfrBFqSRT6WZJg3yamijiXDGKL5wUXulQYmU; sso_uid_tt=84f9839e4e5f36cf5935ff528cc90faf; sso_uid_tt_ss=84f9839e4e5f36cf5935ff528cc90faf; toutiao_sso_user=25664538fd9584d10e72fe544aa55180; toutiao_sso_user_ss=25664538fd9584d10e72fe544aa55180; passport_assist_user=CjxsHGESKXMr55msioDtxEpmNqXD4oLtCRjM1nmtML4_IVkbG-vqPLUO1LqmYZcBI2NPkr5Q_OuVrN1ZghAaSAo8QU99PMgaQ97Zq-U_0ya_koqMbHe7MMP0AIEpCjjPRr78G4Og8hdcb0-9azkjTaBBLoq3mUCJHoYC-WkAEPivsQ0Yia_WVCIBA0ewzBU%3D; odin_tt=280411f4f066fc71c0a3dc313b4932e59846d9725c5cdf6290bcca1a08a417ab147699bf76c217100bfebb6a9aa1a540; passport_auth_status=9b150288be05ea2e5b3b2f6b58b60b37%2C; passport_auth_status_ss=9b150288be05ea2e5b3b2f6b58b60b37%2C; uid_tt=d2829e40a62886b469a3e7f7c3f91a31; uid_tt_ss=d2829e40a62886b469a3e7f7c3f91a31; sid_tt=12966ed0a476706e7ef9ff5425c80b28; sessionid=12966ed0a476706e7ef9ff5425c80b28; sessionid_ss=12966ed0a476706e7ef9ff5425c80b28; LOGIN_STATUS=1; store-region=cn-sx; store-region-src=uid; bd_ticket_guard_server_data=; bd_ticket_guard_client_data=eyJiZC10aWNrZXQtZ3VhcmQtdmVyc2lvbiI6MiwiYmQtdGlja2V0LWd1YXJkLWl0ZXJhdGlvbi12ZXJzaW9uIjoxLCJiZC10aWNrZXQtZ3VhcmQtY2xpZW50LWNlcnQiOiItLS0tLUJFR0lOIENFUlRJRklDQVRFLS0tLS1cbk1JSUNFekNDQWJxZ0F3SUJBZ0lVYzJtRmw2YmZPeHpTZXdnbEYxL3dMenNQQ2Jvd0NnWUlLb1pJemowRUF3SXdcbk1URUxNQWtHQTFVRUJoTUNRMDR4SWpBZ0JnTlZCQU1NR1hScFkydGxkRjluZFdGeVpGOWpZVjlsWTJSellWOHlcbk5UWXdIaGNOTWpNd05URTNNVE13TnpVeFdoY05Nek13TlRFM01qRXdOelV4V2pBbk1Rc3dDUVlEVlFRR0V3SkRcblRqRVlNQllHQTFVRUF3d1BZbVJmZEdsamEyVjBYMmQxWVhKa01Ga3dFd1lIS29aSXpqMENBUVlJS29aSXpqMERcbkFRY0RRZ0FFT3ZiYW5YZVNlb3Q2TXVhVEVqRHNYRzZDaWxBcjF1R1BuNmlxQjQyeEZ3cVBJZEdieXdFWlJvNkhcbi9xclR6RHBRc2ROWHM1MWNYOVBYeEtvOTlOdWNnYU9CdVRDQnRqQU9CZ05WSFE4QkFmOEVCQU1DQmFBd01RWURcblZSMGxCQ293S0FZSUt3WUJCUVVIQXdFR0NDc0dBUVVGQndNQ0JnZ3JCZ0VGQlFjREF3WUlLd1lCQlFVSEF3UXdcbktRWURWUjBPQkNJRUlKZFhEUlhWSVQ2WmZTM2FkejJNOWF4MzR1UHV1Y2Jha3B6bkJWY1hwUzMxTUNzR0ExVWRcbkl3UWtNQ0tBSURLbForcU9aRWdTamN4T1RVQjdjeFNiUjIxVGVxVFJnTmQ1bEpkN0lrZURNQmtHQTFVZEVRUVNcbk1CQ0NEbmQzZHk1a2IzVjVhVzR1WTI5dE1Bb0dDQ3FHU000OUJBTUNBMGNBTUVRQ0lIaExLbHFWTlhHU0c3Y3lcbjRRYVhXblFjRUc5NHFzdFd6OE41eWtRZ2h2c3hBaUJWK3Z4UE5OcjU1bkVDdENQL2RVam4zVXZZZzJNZ095TjRcbnl6L3NIbk1IYnc9PVxuLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLVxuIn0=; d_ticket=d79cf78b7f1532570fcdcef76dcaaaa195732; my_rd=1; pwa2=%223%7C0%22; publish_badge_show_info=%220%2C0%2C0%2C1686369347129%22; __bd_ticket_guard_local_probe=1686369435799; SEARCH_RESULT_LIST_TYPE=%22multi%22; sid_ucp_sso_v1=1.0.0-KDRiZWNmMjBkODA2MDYyMjBkMTk1YmJkOTc2NDMyZDUwMDUxODcxODMKHQjj-oaX7AEQ5beSpAYY7zEgDDDQs5fLBTgGQPQHGgJsZiIgMjU2NjQ1MzhmZDk1ODRkMTBlNzJmZTU0NGFhNTUxODA; ssid_ucp_sso_v1=1.0.0-KDRiZWNmMjBkODA2MDYyMjBkMTk1YmJkOTc2NDMyZDUwMDUxODcxODMKHQjj-oaX7AEQ5beSpAYY7zEgDDDQs5fLBTgGQPQHGgJsZiIgMjU2NjQ1MzhmZDk1ODRkMTBlNzJmZTU0NGFhNTUxODA; sid_guard=12966ed0a476706e7ef9ff5425c80b28%7C1686412261%7C5184000%7CWed%2C+09-Aug-2023+15%3A51%3A01+GMT; sid_ucp_v1=1.0.0-KGM4NGYzYmFlZjNkNWRmMmE2MDdkMzk5MTlkNGZjYTU5ODRiODk0YzkKGQjj-oaX7AEQ5beSpAYY7zEgDDgGQPQHSAQaAmxxIiAxMjk2NmVkMGE0NzY3MDZlN2VmOWZmNTQyNWM4MGIyOA; ssid_ucp_v1=1.0.0-KGM4NGYzYmFlZjNkNWRmMmE2MDdkMzk5MTlkNGZjYTU5ODRiODk0YzkKGQjj-oaX7AEQ5beSpAYY7zEgDDgGQPQHSAQaAmxxIiAxMjk2NmVkMGE0NzY3MDZlN2VmOWZmNTQyNWM4MGIyOA; download_guide=%223%2F20230610%2F1%22; device_web_cpu_core=8; device_web_memory_size=8; webcast_local_quality=null; strategyABtestKey=%221686875160.211%22; __ac_nonce=0648bae9a008a6baaaec1; __ac_signature=_02B4Z6wo00f01PQlvkwAAIDBfXAXrdb9o9j0BbrAAFmI2zpAjOMkkhpj54k5rwl3aBNgCntbH3s4J9vpEd7u9STzQKW09G2CYyPgocJGkmwTdEnBrJrKngWwcJtESmyJk6ufqte4HJaHcM-1cf; VIDEO_FILTER_MEMO_SELECT=%7B%22expireTime%22%3A1687480605007%2C%22type%22%3A1%7D; FOLLOW_LIVE_POINT_INFO=%22MS4wLjABAAAAXfGr1LU5CLhXWe6IXBCqNme9ZKgyziz7X0Bje81kG1g%2F1686931200000%2F0%2F0%2F1686876405362%22; FOLLOW_NUMBER_YELLOW_POINT_INFO=%22MS4wLjABAAAAXfGr1LU5CLhXWe6IXBCqNme9ZKgyziz7X0Bje81kG1g%2F1686931200000%2F0%2F1686875805363%2F0%22; passport_fe_beating_status=true; home_can_add_dy_2_desktop=%221%22; msToken=cNoF-xmRgrBG-yFgk6nE8wKFFua9dXfWG9ZqRmEDfTN8-5qNYT-LVB082UNqglXqZyvZ32tFMqZNh_sJ_ADQto9CxjCsQKbB6aPTi40PUH0RD2kOMqP3xno=; tt_scid=NZUy42xod7cwoWffVcBbdPtO1ehjhH3mAXdUHvbGfAeLqD46n5ZqnzAak5kpAlD934a7; msToken=81bCCBkGm0VNC9va2OlelVahVxoJJcLDo8Waz39Lky-K2ZEqiR0K6Pyb-R3iSlF7m2Nbyyw37iP50RJK0xot5P9db6wFPMQ8aoA3yxnRkznoKu9oBFCO4nI=")
	req.Header.Set("referer", "https://www.douyin.com/")

	resp, err := client.Do(req)
	//resp, err := http.Get("https://www.baidu.com/")
	if err != nil {
		// 处理错误
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		// 处理错误
		return
	}
	if resp.StatusCode == 200 {
		println("响应码为：", resp.StatusCode)
	}
	// 解码gbk
	// decoder := simplifiedchinese.GBK.NewDecoder()
	// result, _ := decoder.Bytes(body)
	anyMap := make(map[string]interface{}, 0)
	// if err := json.Unmarshal([]byte(jsonStr), &anyMap);
	err = json.Unmarshal(body, &anyMap)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	// var responseData ResponseData
	// err = json.Unmarshal(body, &responseData)
	// if err != nil {
	// 	fmt.Println("Error decoding JSON:", err)
	// 	return
	// }

	fmt.Println(resp.StatusCode, len(body))
	// fmt.Println("body:\n", anyMap)
	fmt.Println(len(anyMap))
	for k, _ := range anyMap {
		fmt.Printf("k: %v\n", k)
	}
	// if strings.Contains(string((body)), "北京创信达科技有限公司") {
	// println("存在sb")
	// }

}
