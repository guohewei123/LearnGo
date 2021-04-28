package model

import "encoding/json"

// 定义用户信息结构体
type Profile struct {
	Name       string
	Gender     string
	Residence  string
	Age        int
	IncomeOrEducation  string
	Marriage   string
	Height     int
}

func FromJsonObj(o interface{}) (Profile, error) {
	var profile Profile
	s, err := json.Marshal(o)
	if err != nil {
		return profile, err
	}
	err = json.Unmarshal(s, &profile)
	return profile, err
}
