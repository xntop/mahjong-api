// Code generated by "stringer -type=Limit"; DO NOT EDIT.

package yaku

import "strconv"

const _Limit_name = "LimitNoneLimitManganLimitHanemanLimitBaimanLimitSanbaimanLimitYakuman"

var _Limit_index = [...]uint8{0, 9, 20, 32, 43, 57, 69}

func (i Limit) String() string {
	if i < 0 || i >= Limit(len(_Limit_index)-1) {
		return "Limit(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Limit_name[_Limit_index[i]:_Limit_index[i+1]]
}