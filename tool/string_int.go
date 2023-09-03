package tool

import (
	"encoding/json"
	"fmt"
	"log"
	"mayi/model"
	"strconv"
)

func String_int(val string) int {
	var a int
	if s, err := strconv.Atoi(val); err != nil {
		fmt.Print(err.Error())
	} else {
		a = s
	}
	return a
}

func String_int64(val string) int64 {
	var a int64
	if s, err := strconv.ParseInt(val, 10, 64); err != nil {
		fmt.Print(err.Error())
	} else {
		a = s
	}
	return a
}

func String_float64(val string) float64 {
	num, _ := strconv.ParseFloat(val, 64)
	return float64(num)
}

func String_float(val string) float64 {
	num, _ := strconv.ParseFloat(val, 64)
	return float64(num)
}

func Float64_string(val float64) string {

	s := strconv.FormatFloat(val, 'E', -1, 64)
	return s
}

func Int_string(val int) string {

	v := strconv.Itoa(val)

	return v
}

func Int64_string(val int64) string {

	v := strconv.FormatInt(val, 10)

	return v
}

func Float64_Decimal(value float64, prec int) float64 { //保留小数点后几位
	value, _ = strconv.ParseFloat(strconv.FormatFloat(value, 'f', prec, 64), 64)
	return value
}

func Float_String(num float64) string {
	str := strconv.FormatFloat(num, 'G', 30, 32)
	return str
}
func Interface_int(i interface{}) int64 {
	var res int64
	res = i.(int64)
	return res
}
func Interface_string(value interface{}) (string, bool) {
	switch value.(type) {
	case string:
		op, ok := value.(string)
		return op, ok

	default:
		return "unkown", false
	}

}

// json转map
func Json2map(req string) (s map[string]interface{}, err error) {
	var result map[string]interface{}
	if err := json.Unmarshal([]byte(req), &result); err != nil {
		return nil, err
	}
	return result, nil
}
func MapToSlice(m map[string]interface{}) []interface{} {
	s := make([]interface{}, 0, len(m))
	for _, v := range m {
		s = append(s, v)
	}
	return s
}
func Struct2Byte(msg *model.TestCity) []byte {
	var res []byte

	if jsonStr, err := json.Marshal(msg); err != nil {
		log.Print(err.Error())

	} else {
		res = jsonStr
	}
	return res
}


////string到int
//int,err:=strconv.Atoi(string)
////string到int64
//int64, err := strconv.ParseInt(string, 10, 64)
////int到string
//string:=strconv.Itoa(int)
////int64到string
//string:=strconv.FormatInt(int64,10)
////string到float32(float64)
//float,err := strconv.ParseFloat(string,32/64)
////float到string
//string := strconv.FormatFloat(float32, 'E', -1, 32)
//string := strconv.FormatFloat(float64, 'E', -1, 64)
//// 'b' (-ddddp±ddd，二进制指数)
//// 'e' (-d.dddde±dd，十进制指数)
//// 'E' (-d.ddddE±dd，十进制指数)
//// 'f' (-ddd.dddd，没有指数)
//// 'g' ('e':大指数，'f':其它情况)
//// 'G' ('E':大指数，'f':其它情况)

//// string到int
//int, err := strconv.Atoi(string)
//
//// string到int64
//int64, err := strconv.ParseInt(string, 10, 64)
//
//// int到string
//string := strconv.Itoa(int)
//
//// int64到string
//string := strconv.FormatInt(int64,10)
