package utils

/*
对 json-iterator 的二次封装
*****************
import "github.com/jessun2017/gold/utils"
...
var jsoniter utils.JSONITER
jsoniter.Marshal(*, *)

*****************

*/

import (
	jsoniter "github.com/json-iterator/go"
)

var JSONITER = jsoniter.ConfigCompatibleWithStandardLibrary
