package snowflake

// 这个模块封装了 雪花算法生成相应用户ID的方法。
import (
	sf "github.com/bwmarrin/snowflake"
	"time"
) // 定义一个节点： 通过这个全局的 node，就可以用于制造 ID了。

var node *sf.Node

func Init(startTime string, machineID int64) (err error) {
	var st time.Time
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		return err
	}
	sf.Epoch = st.UnixNano() / 1000000
	node, err = sf.NewNode(machineID)
	return
}
func GetID() int64 {
	return node.Generate().Int64()
}
