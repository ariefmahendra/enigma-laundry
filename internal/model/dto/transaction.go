package dto

type CreateBillRequest struct {
	BillDate   string                    `json:"billDate"`
	EntryDate  string                    `json:"entryDate"`
	FinishDate string                    `json:"finishDate"`
	EmployeeId int                       `json:"employeeId"`
	CustomerId int                       `json:"customerId"`
	BillDetail []CreateBillDetailRequest `json:"billDetails"`
}

type CreateBillDetailRequest struct {
	ProductId int `json:"productId"`
	Qty       int `json:"qty"`
}

type CreateBillResponse struct {
	Id         int                        `json:"id"`
	BillDate   string                     `json:"billDate"`
	EntryDate  string                     `json:"entryDate"`
	FinishDate string                     `json:"finishDate"`
	Employee   int                        `json:"employee"`
	Customer   int                        `json:"customer"`
	BillDetail []CreateBillDetailResponse `json:"billDetail"`
	TotalBill  int                        `json:"totalBill"`
}

type CreateBillDetailResponse struct {
	Id           int `json:"id"`
	BillId       int `json:"billId"`
	ProductId    int `json:"productId"`
	ProductPrice int `json:"productPrice"`
	Quantity     int `json:"qty"`
}

type GetBillByIdResponse struct {
	Id         int                         `json:"id"`
	BillDate   string                      `json:"billDate"`
	EntryDate  string                      `json:"entryDate"`
	FinishDate string                      `json:"finishDate"`
	Employee   EmployeeResponse            `json:"employee"`
	Customer   CustomerResponse            `json:"customer"`
	BillDetail []GetBillDetailByIdResponse `json:"billDetail"`
	TotalBill  int                         `json:"totalBill"`
}

type GetBillDetailByIdResponse struct {
	Id           int             `json:"id"`
	BillId       int             `json:"billId"`
	Product      ProductResponse `json:"product"`
	ProductPrice int             `json:"productPrice"`
	Quantity     int             `json:"qty"`
}
