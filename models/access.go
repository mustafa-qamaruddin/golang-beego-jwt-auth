package models

import (
	"strconv"
	"time"
)

var (
	AccessList map[string]*Access
)

func init() {
	AccessList = make(map[string]*Access)
	u := Access{"access_11111", "1234", "9876", "1495982544", "true"}
	AccessList["user_11111"] = &u
}

type Access struct {
	Id       string
	APID string `valid:"Required;Numeric";json:"ap_id"`
	PassID string `valid:"Required;Numeric";json:"pass_id"`
	Timestamp string `valid:"Required;Numeric;Length(10)";json:"timestamp"`
	Access string `valid:"Required;Match(/^true|false$/)";json:"access"`
}

func AddAccess(u Access) string {
	u.Id = "log_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	AccessList[u.Id] = &u
	return u.Id
}

func GetAllAccess() map[string]*Access {
	return AccessList
}