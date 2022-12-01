package module

type Community struct {
	Community_id   int    `db:"community_id"`
	Community_name string `db:"community_name"`
	Introduction   string `db:"introduction"`
}
