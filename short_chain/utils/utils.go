package utils

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	uuid "github.com/satori/go.uuid"
	"time"
)

//GetUuid 获取uuid
func GetUuid() string {
	u := uuid.NewV4()
	return Get16MD5Encode(u.String())
}

func Get16MD5Encode(data string) string {
	return GetMD5Encode(data)[8:24]
}

// GetMD5Encode 返回一个32位md5加密后的字符串
func GetMD5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

// ToMap 通用转map功能
func ToMap(data interface{}, removeKey []string) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	j, _ := json.Marshal(data)
	err := json.Unmarshal(j, &m)
	createTime, _ := time.Parse(time.RFC3339, m["CreatedAt"].(string))
	updateTime, _ := time.Parse(time.RFC3339, m["UpdatedAt"].(string))
	m["create_time"] = createTime.Format("2006-01-02 15:04:05")
	m["update_time"] = updateTime.Format("2006-01-02 15:04:05")
	m["id"] = m["ID"]
	delete(m, "ID")
	delete(m, "CreatedAt")
	delete(m, "UpdatedAt")
	delete(m, "DeletedAt")
	if removeKey != nil && len(removeKey) > 0 {
		for _, v := range removeKey {
			delete(m, v)
		}
	}
	return m, err
}
