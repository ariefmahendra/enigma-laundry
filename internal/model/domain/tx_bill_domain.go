package domain

import "time"

type TxBill struct {
	Id         int
	BillDate   time.Time
	EntryDate  time.Time
	FinishDate time.Time
	EmployeeId int
	CustomerId int
	TotalBill  int
}
