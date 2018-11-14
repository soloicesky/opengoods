package entity

type Category struct {
	Id    string `gorm:"primary_key" json:"id"` //category id unique auto generate
	Name  string `gorm:"name" json:"name"`      //category name
	Extra string `gorm:"extra" json:"extra"`    //user define extra details json format string
}

type Origin struct {
	Id          string `gorm:"primary_key" json:"id"`        //orgin id unique
	Name        string `gorm:"name" json:"name"`             //name
	Country     string `gorm:"country" json:"country"`       //country
	Address     string `gorm:"address" json:"country"`       //address's details
	Postal_code string `gorm:"postalcode" json:"postalcode"` // postal code
	Start_time  string `gorm:"starttime" json:"starttime"`   // start time
	Extra       string `gorm:"extra" json:"extra"`           //extra statement of the factory, also in json format
}

type Goods struct {
	Id         string `gorm:"primary_key" json:"id"`        //goods id unique
	Name       string `gorm:"name" json:"name"`             // goods name
	CategoryId string `gorm:"categoryid" json:"categoryid"` //witch category belongs to
	OriginId   string `gorm:"originid" json"originid"`      // produced by witch factory
	Unit       string `gorm:"unit" json:"unit"`             //unit
	Price      string `gorm:"price" json:"price"`           // price for each unit
	Extra      string `gorm:"extra" json:"extra"`           //extra statement of the factory, also in json format
}
