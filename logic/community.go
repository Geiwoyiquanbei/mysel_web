package logic

import (
	"myself/dao/mysql"
	"myself/module"
)

func GetCommunityList(cl *[]module.Community) (err error) {
	err = mysql.GetCommunity(cl)
	if err != nil {
		return err
	}
	return nil
}
func GetCommunityByID(id int64) (data *module.Community, err error) {
	data, err = mysql.GetCommunityByID(id)
	if err != nil {
		return nil, err
	}
	return data, nil
}
