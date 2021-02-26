package db

// Post ... Post Model for DB
type Post struct {
	ID        string `bson:"_id" json:"id"`
	Name      string `bson:"name,omitempty" json:"name"`
	Timestamp string `bson:"timestamp" json:"timestamp"`
	Contents  string `bson:"contents" json:"contents"`
	Account   string `bson:"account" json:"account"`
}
