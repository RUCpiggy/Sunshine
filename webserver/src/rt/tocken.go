package rt

import (
	"crypto/md5"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"
)

const salt = "sunshineweb"

func genTocken(userID int) (string, string, string) {
	t := time.Now()
	userIDString := strconv.Itoa(userID)
	now := strconv.FormatInt(t.Unix(), 10)
	rowData := userIDString + now
	h := md5.New()
	io.WriteString(h, rowData)
	pmd5 := fmt.Sprintf("%x", h.Sum(nil))
	io.WriteString(h, salt)
	io.WriteString(h, pmd5)
	tocken := fmt.Sprintf("%x", h.Sum(nil))
	return userIDString, tocken, now
}

func testTocken(userID string, t string) string {
	rowData := string(userID) + string(t)
	h := md5.New()
	io.WriteString(h, rowData)
	pmd5 := fmt.Sprintf("%x", h.Sum(nil))
	io.WriteString(h, salt)
	io.WriteString(h, pmd5)
	tocken := fmt.Sprintf("%x", h.Sum(nil))
	return tocken
}

func isLogin(tocken string) (string, bool) {
	tockenArray := strings.Split(tocken, "#")
	t := time.Now().Unix()
	before, err := strconv.ParseInt(tockenArray[2], 10, 64)
	if err != nil {
		print(err)
	}
	if tockenArray[1] == testTocken(tockenArray[0], tockenArray[2]) && t-before < 3600 {
		return tockenArray[0], true
	} else {
		return "", false
	}
}
