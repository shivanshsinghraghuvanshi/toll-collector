// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Car struct {
	Carid     string `json:"carid"`
	Make      string `json:"make"`
	Cartype   string `json:"cartype"`
	Carnumber string `json:"carnumber"`
}

type Deductible struct {
	ID      string `json:"id"`
	Cartype string `json:"cartype"`
	Amount  int    `json:"amount"`
}

type Netc struct {
	Netcid  string `json:"netcid"`
	Ownerid string `json:"ownerid"`
	Carid   string `json:"carid"`
	Rfid    string `json:"rfid"`
}

type NewCar struct {
	Make      string `json:"make"`
	Cartype   string `json:"cartype"`
	Carnumber string `json:"carnumber"`
}

type NewOwner struct {
	Accountnumber string `json:"accountnumber"`
	Name          string `json:"name"`
}

type NewRfid struct {
	Ownerid string `json:"ownerid"`
	Carid   string `json:"carid"`
}

type NewTollBooth struct {
	Accountnumber string `json:"accountnumber"`
	Name          string `json:"name"`
}

type Owner struct {
	Ownerid       string `json:"ownerid"`
	Accountnumber string `json:"accountnumber"`
	Name          string `json:"name"`
}

type Relation struct {
	Owner *Owner `json:"owner"`
	Car   []*Car `json:"car"`
}

type Tollbooth struct {
	Tollboothid   string `json:"tollboothid"`
	Name          string `json:"name"`
	Accountnumber string `json:"accountnumber"`
}
