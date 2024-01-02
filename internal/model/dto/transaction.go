package dto

type CreateBillRequest struct {
	BillDate   string                    `json:"billDate"`
	EntryDate  string                    `json:"entryDate"`
	FinishDate string                    `json:"finishDate"`
	EmployeeId string                    `json:"employeeId"`
	CustomerId string                    `json:"customerId"`
	BillDetail []CreateBillDetailRequest `json:"billDetails"`
}

type CreateBillDetailRequest struct {
	ProductId string `json:"productId"`
	Qty       int    `json:"qty"`
}

type CreateBillResponse struct {
	Id         string                     `json:"id"`
	BillDate   string                     `json:"billDate"`
	EntryDate  string                     `json:"entryDate"`
	FinishDate string                     `json:"finishDate"`
	EmployeeId string                     `json:"employeeId"`
	CustomerId string                     `json:"customerId"`
	BillDetail []CreateBillDetailResponse `json:"billDetail"`
}

type CreateBillDetailResponse struct {
	Id           string `json:"id"`
	BillId       string `json:"billId"`
	ProductId    string `json:"productId"`
	ProductPrice int    `json:"productPrice"`
	Quantity     int    `json:"qty"`
}

type GetBillByIdResponse struct {
	Id         string                      `json:"id"`
	BillDate   string                      `json:"billDate"`
	EntryDate  string                      `json:"entryDate"`
	FinishDate string                      `json:"finishDate"`
	Employee   EmployeeResponse            `json:"employee"`
	Customer   CustomerResponse            `json:"customer"`
	BillDetail []GetBillDetailByIdResponse `json:"billDetail"`
	TotalBill  int                         `json:"totalBill"`
}

type GetBillDetailByIdResponse struct {
	Id           string          `json:"id"`
	BillId       string          `json:"billId"`
	Product      ProductResponse `json:"product"`
	ProductPrice int             `json:"productPrice"`
	Quantity     int             `json:"qty"`
}
