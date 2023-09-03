package tool

import (
	"crypto/md5"
	rr "crypto/rand"
	"fmt"
	"io"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
var nums = []rune("1234567890")

func SmsNum() string { //4位短信验证码
	return fmt.Sprintf("%04v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000))
}
func SmsNum_6() string { //4位短信验证码
	return fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
}

func Uuid(n int) string {

	b := make([]rune, n)

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]

	}
	return string(b)
}

func RandNum(n int) string { //大小写字母随机
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]

	}
	return string(b)
}

func HashUuid() string {
	t := time.Now()
	h := md5.New()
	io.WriteString(h, "crazyof.me")
	io.WriteString(h, t.String())
	uid := fmt.Sprintf("%x", h.Sum(nil))
	return uid
}

func Fanwei() int {
	rand.Seed(time.Now().UnixNano())

	x := rand.Intn(4) //生成0-99随机整数
	return x
}

// 随机字符串
func RandString(length int) string {
	rand.Seed(time.Now().UnixNano())
	if length < 100000 {
		rs := make([]string, length)
		for start := 0; start < length; start++ {
			t := rand.Intn(3)
			if t == 0 {
				rs = append(rs, strconv.Itoa(rand.Intn(10)))
			} else if t == 1 {
				rs = append(rs, string(rand.Intn(26)+65))
			} else {
				rs = append(rs, string(rand.Intn(26)+97))
			}
		}
		return strings.Join(rs, "")
	} else {
		return ""
	}

}

// 获取唯一ID
func Get_uuid() (string, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rr.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}

// 获取唯一ID
func GetUuid(l int) (string, error) {
	uuid := make([]byte, l)
	n, err := io.ReadFull(rr.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}

// 获取当前时间戳
func Get_time_stamp() int {
	ts, _ := strconv.Atoi(strconv.FormatInt(time.Now().UnixNano()/1000000000, 10))
	return ts
}

func RandInt(num int) int {
	n := 9999
	switch num {
	case 6:
		n = 999999
	case 8:
		n = 99999999
	case 12:
		n = 999999999999
	case 18:
		n = 999999999999999999
	default:
		n = 9999
	}

	rand.Seed(time.Now().UTC().UnixNano())
	myid := rand.Intn(n)
	return myid
}


