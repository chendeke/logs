package main

import (
	"github.com/chendeke/logs"
	"math/rand"
	"strings"
	"time"
)

/*
RandomString is random string
*/
func RandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	rs := make([]string, length)
	for start := 0; start < length; start++ {
		t := rand.Intn(2)
		if t == 0 {
			// A-Z
			rs = append(rs, string(rand.Intn(26)+65))
		} else {
			// a-z
			rs = append(rs, string(rand.Intn(26)+97))
		}
	}
	return strings.Join(rs, "")
}

func main() {

	// logs.Errorw("msg", "err:", fmt.Errorf("fuck ok"))

	for {
		logs.Infof("connected kafka broker:%s group:%s topic:%s", "12", "34", "56")
		logs.Debug("debug message", RandomString(6), "----", RandomString(6))
		logs.Debugf("debug something happened %s %s", RandomString(6), RandomString(6))
		logs.Debugw("Debugw", "url", RandomString(6), "attempt", 3)
		time.Sleep(time.Second * 2)
	}

	// for {
	// 	logs.Debug("debug message", RandomString(6), "----", RandomString(6))
	// 	logs.Debugf("debug something happened %s %s", RandomString(6), RandomString(6))
	// 	logs.Debugw("Debugw", "url", RandomString(6), "attempt", 3)
	//
	// 	logs.Info("debug message", RandomString(6), "----", RandomString(6))
	// 	logs.Infof("debug something happened %s %s", RandomString(6), RandomString(6))
	// 	logs.Infow("Debugw", "url", RandomString(6), "attempt", 3)
	//
	// 	time.Sleep(time.Second * 2)
	// }
}
