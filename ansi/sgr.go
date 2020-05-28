// SPDX-License-Identifier: MIT

package ansi

// SGR 控制码
const (
	SGRReset           = "0"
	SGRBold            = "1"
	SGRUnderline       = "4"
	SGRBlink           = "5" // 闪烁
	SGRReverseVideo    = "7" // 反显
	SGRConceal         = "8"
	SGRBoldOff         = "22"
	SGRUnderlineOff    = "24"
	SGRBlinkOff        = "25"
	SGRReverseVideoOff = "27"
	SGRConcealOff      = "28"

	SGRFBlack   = "30"
	SGRFRed     = "31"
	SGRFGreen   = "32"
	SGRFYellow  = "33"
	SGRFBlue    = "34"
	SGRFMagenta = "35"
	SGRFCyan    = "36"
	SGRFWhite   = "37"
	SGRFDefault = "39" // 默认前景色

	SGRBBlack   = "40"
	SGRBRed     = "41"
	SGRBGreen   = "42"
	SGRBYellow  = "43"
	SGRBBlue    = "44"
	SGRBMagenta = "45"
	SGRBCyan    = "46"
	SGRBWhite   = "47"
	SGRBDefault = "49" // 默认背景色
)

// SGR 将几个 SGR 控制符合成一个 ansi 控制符
//
//  "30", "31", "32"
//  // 以上参数将会被转换成以下内容返回
//  "\033[30;31;32m"
func SGR(args ...string) string {
	if len(args) == 0 {
		return "\033[" + SGRReset + "m"
	}

	ret := ""
	for _, v := range args {
		ret += v + ";"
	}

	return "\033[" + ret[0:len(ret)-1] + "m"
}
