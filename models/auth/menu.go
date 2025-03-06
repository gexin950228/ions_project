package auth

type Tree struct {
	Id       int     `orm:"id,primary" json:"id"`
	AuthName string  `orm:"name" json:"name"`
	UrlFor   string  `orm:"url_for" json:"url_for"`
	Weight   int     `orm:"weight" json:"weight"`
	Children []*Tree `orm:"-" json:"children"`
}
