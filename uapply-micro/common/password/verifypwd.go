package password

import (
	"unicode"
)

const (
	PwdOptNumber  uint16 = 1 << iota //数字 	 0001
	PwdOptLower                      //小写 	 0010
	PwdOptUpper                      //大写 	 0100
	PwdOptSpecial                    //特殊符号 1000
)

// VerifyPwd 密码验证
func VerifyPwd(pwd string, options uint16) bool {
	if options == 0 {
		options = PwdOptNumber | PwdOptLower | PwdOptUpper | PwdOptSpecial
	}

	if len(pwd) < 12 || len(pwd) > 18 {
		return false
	}
	// 用于记录验证结果
	var result uint16
	for _, r := range pwd {
		switch {
		case unicode.IsNumber(r):
			result = result | PwdOptNumber
		case unicode.IsLower(r):
			result = result | PwdOptLower
		case unicode.IsUpper(r):
			result = result | PwdOptUpper
		case (unicode.IsPunct(r) || unicode.IsSymbol(r)) && (r == '@' || r == '.'): //标点符号 和 字符
			result = result | PwdOptSpecial
		default:
			return false
		}
		// 比较结果和设置项
		// 当 options与result != result 表示密码字符串超出 options 范围
		if options&result != result {
			return false
		}
	}
	return true
}
