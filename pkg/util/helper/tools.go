package helper

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

const (
	NUmStr  = "0123456789"
	CharStr = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	SpecStr = "!@#$%&"
)

func GeneratePasswd(length int, charset string) string {
	var passwd = make([]byte, length)
	var sourceStr string

	if charset == "num" { // 数字
		sourceStr = NUmStr
	} else if charset == "char" { // 字母
		sourceStr = charset
	} else if charset == "mix" { // 数字、字母混合模式
		sourceStr = fmt.Sprintf("%s%s", NUmStr, CharStr)
	} else if charset == "advance" { // 数字、字母、字符混合模式
		sourceStr = fmt.Sprintf("%s%s%s", NUmStr, CharStr, SpecStr)
	} else {
		sourceStr = NUmStr
	}
	// fmt.Println("source:", sourceStr)

	// 生成密码
	rand.NewSource(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		index := rand.Intn(len(sourceStr))
		passwd[i] = sourceStr[index]
	}
	return string(passwd)
}

func StrInArray(target string, strArray []string) bool {
	sort.Strings(strArray)
	index := sort.SearchStrings(strArray, target)
	if index < len(strArray) && strArray[index] == target {
		return true
	}
	return false
}
