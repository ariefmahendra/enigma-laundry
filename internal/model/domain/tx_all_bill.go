package domain

import "time"

type Bill struct {
	Id         int
	BillDate   time.Time
	EntryDate  time.Time
	FinishDate time.Time
	Employee   Employee
	Customer   Customer
	BillDetail []BillDetail
	TotalBill  int
}

type BillDetail struct {
	Id           int
	BillId       int
	Product      Product
	ProductPrice int
	Quantity     int
}
