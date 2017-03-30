/*
 * gouid
 * https://github.com/chrisenytc/gouid
 *
 * Copyright (c) 2014, Christopher EnyTC
 * Licensed under the MIT license.
 */

package gouid

/*
 * Dependencies
 */

import (
	"math/rand"
	"time"
)

type UId struct {
	Length int
}

func (uid UId) SetSeed() {
	rand.Seed(time.Now().Unix())
}

func (uid *UId) NewUId() string {
	if uid.Length > 256 {
		panic("[gouid] Maximum length of charset for NewUId is 256")
	}
	buf := make([]rune, uid.Length)
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
	charlen := len(chars)
	for i := range buf {
		buf[i] = chars[rand.Intn(charlen)]
	}
	return string(buf)
}
