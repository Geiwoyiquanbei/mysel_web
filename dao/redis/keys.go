package redis

//redis key
const (
	KeyPrefix          = "myself:"
	KeyPostTimeZSet    = "post:time"   //zset 帖子和发帖的时间
	KeyPostScoreZSet   = "post:score"  // zset 帖子以及投票的分数
	KeyPostVotedPrefix = "post:voted:" //zset 记录用户及投票类型；参数是post id
	KeyCommunitySetPF  = "community:"  //set ;保存每个分区下的帖子的id
)

func GetKeys(string2 string) string {
	return KeyPrefix + string2
}
