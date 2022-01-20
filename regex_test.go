package main

import (
	"math/rand"
	"net/url"
	"regexp"
	"testing"
	"time"
)

func TestRegex(t *testing.T) {
	var list = []string{
		"123",
		"1",
		"a",
		"1223123",
		"12344411sadfhwe213443",
		"12a你",
		"asd124,",
	}
	for _, v := range list {
		b, e := regexp.MatchString(`^[a-zA-Z0-9_-]+$`, v)
		if e != nil {
			t.Errorf("check %s error: %s\n", v, e)
			continue
		}
		if !b {
			t.Logf("%s is false\n", v)
		}
	}

}

func TestUrl(t *testing.T) {
	u := "http://127.0.0.1:9099/"
	uu, _ := url.Parse(u)
	t.Log(uu.Path)
}

// 0-9 -> 48-57 | A-Z 65-90 | a-z | 97-122
func TestASCII(t *testing.T) {
	words := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
	str := ""
	for i := 0; i < 6; i++ {
		rand.Seed(time.Now().UnixNano())
		index := rand.Intn(len(words))
		str += string(words[index])
	}
	t.Log(str)
}
