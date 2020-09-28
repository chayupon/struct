package main

import (
	"encoding/json"
	"fmt"
)

//PersonalHistory struct name
type PersonalHistory struct {
	AccountName                string  `json:"account_name"`
	AccountNumber              int64   `json:"account_number"`
	ContractDate               string  `json:"contract_date"`
	ContractDatetime           string  `json:"contract_datetime"`
	ContractNumber             int64   `json:"contract_number"`
	CreditLimit                int64   `json:"credit_limit"`
	CustomerNumber             int64   `json:"customer_number"`
	MaximumCustomerCreditLimit int64   `json:"maximum_customer_credit_limit"`
	Phu                        string  `json:"phu"`
	ProductName                string  `json:"product_name"`
	SalakAccountNumber         int64   `json:"salak_account_number"`
	SalakBalance               int64   `json:"salak_balance"`
	SalakList                  []salak `json:"salaklist"`
	SuccessArr                 Success `json:"successarr"`
}

type salak struct {
	Amount         int64  `json:"amount"`
	CertificateID  string `json:"certificate_id"`
	Group          string `json:"group"`
	HighNumber     int64  `json:"high_number"`
	LowNumber      int64  `json:"low_number"`
	MaturityDate   string `json:"maturity_date"`
	Period         int64  `json:"period"`
	RecordNumber   int64  `json:"record_number"`
	RemainingPrize int64  `json:"remaining_prize"`
	Status         int64  `json:"status"`
	Term           string `json:"term"`
}



//Success struct
type Success struct {
	Ok string `json:"ok"`
}

func main() {
	
	
	pesrsonalhistory := PersonalHistory{
		ContractNumber:121000000121,
		AccountNumber:200003500005,
		SalakAccountNumber:490000000005,
		ContractDate:"2020-04-04",
		ContractDatetime:"2020-04-04T00:00:00+07:00",
		ProductName:"GSBMyCreditOD",
		CreditLimit:60000,
		MaximumCustomerCreditLimit:117000,
		CustomerNumber:4038,
		AccountName:"Gaew Noi",
		SalakBalance:130000,
		SalakList: []salak{
			salak{
				RecordNumber:1,
				CertificateID:"122221",
				Period:120,
				Group:"ก",
				LowNumber:0,
				HighNumber:1359,
				Amount:68000.00,
				RemainingPrize:100.00,
				Status:0,
				Term:"3Y",
				MaturityDate:"2021-01-01",
				},
				{
				RecordNumber:2,
				CertificateID:"122222",
				Period:120,
				Group:"ข",
				LowNumber:0,
				HighNumber:239,
				Amount:12000.00,
				RemainingPrize:100.00,
				Status:0,
				Term:"3Y",
				MaturityDate:"2022-04-15",
				},
			},
			SuccessArr: Success{
				Ok: "1",
			},
		}
	

	
	output,_ :=json.Marshal(&pesrsonalhistory)
	fmt.Println("My name is :",string(output))
}

