// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type AccountDetails struct {
	Accountnumber     *int    `json:"Accountnumber"`
	Accountid         *int    `json:"Accountid"`
	AccountHolderName *string `json:"AccountHolderName"`
	Balance           *int    `json:"Balance"`
	LastUpdated       *string `json:"LastUpdated"`
}

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

type MatrixResponse struct {
	Special int   `json:"special"`
	Matrix  []int `json:"matrix"`
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
	Ownerid   string `json:"ownerid"`
	Carnumber string `json:"carnumber"`
}

type NewTollBooth struct {
	Accountnumber string `json:"accountnumber"`
	Name          string `json:"name"`
}

type NewTollTax struct {
	Cartype *string `json:"cartype"`
	Amount  *int    `json:"amount"`
}

type Owner struct {
	Ownerid       string `json:"ownerid"`
	Accountnumber string `json:"accountnumber"`
	Name          string `json:"name"`
}

type OwnerInfoDetails struct {
	Name          *string `json:"name"`
	AccountNumber *string `json:"accountNumber"`
	Action        *string `json:"Action"`
}

type PayTollTax struct {
	Rfid    string  `json:"rfid"`
	Tollid  int     `json:"tollid"`
	Amount  int     `json:"amount"`
	Remarks *string `json:"remarks"`
}

type Relation struct {
	Owner *Owner `json:"owner"`
	Car   []*Car `json:"car"`
}

type TollBoothInfoDetails struct {
	Name          *string `json:"name"`
	AccountNumber *string `json:"accountNumber"`
	Action        *string `json:"Action"`
}

type Tollbooth struct {
	Tollboothid   string `json:"tollboothid"`
	Name          string `json:"name"`
	Accountnumber string `json:"accountnumber"`
}

type TransactionHistory struct {
	TransactionID       *int    `json:"TransactionID"`
	Timestamp           *string `json:"Timestamp"`
	DebitAcoountNumber  *int    `json:"DebitAcoountNumber"`
	CreditAccountNumber *int    `json:"CreditAccountNumber"`
	Amount              *int    `json:"Amount"`
	Remarks             *string `json:"Remarks"`
}

type ValidateRfid struct {
	Rfid  string `json:"rfid"`
	Carid int    `json:"carid"`
}
