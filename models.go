package main

type User struct {
	Id       string `bson:"Id" json:"Id"`
	Name     string `bson:"Name" json:"Name"`
	Email    string `bson:"Email" json:"Email"`
	Password string `bson:"Password" json:"Password"`
}

type Post struct {
	Id        string `bson:"Id" json:"Id"`
	UserId    string `bson:"UserId" json:"UserId"`
	Caption   string `bson:"Caption" json:"Caption"`
	ImageURL  string `bson:"ImageURL" json:"ImageURL"`
	Timestamp string `bson:"Timestamp" json:"Timestamp"`
}
