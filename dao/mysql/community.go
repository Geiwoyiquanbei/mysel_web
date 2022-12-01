package mysql

import "myself/module"

func GetCommunity(cl *[]module.Community) (err error) {
	sqlStr := `select community_id,community_name,introduction from community`
	err = db.Select(cl, sqlStr)
	if err != nil {
		return err
	}
	return nil
}
func GetCommunityByID(id int64) (*module.Community, error) {
	sqlStr := `select community_id,community_name,introduction from community where communtiy_id=?`
	var data = new(module.Community)
	err := db.Get(data, sqlStr, id)
	if err != nil {
		return nil, err
	}
	return data, nil
}
