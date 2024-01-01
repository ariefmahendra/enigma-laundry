package domain

type TxBill struct {
	Id         int
	BillDate   string
	EntryDate  string
	FinishDate string
	EmployeeId int
	CustomerId int
	TotalBill  int
}
