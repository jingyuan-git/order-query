package models

type Customer struct {
	Id              string          `csv:"user_id" json:"id,omitempty" gorm:"primary_key"`
	Login           string          `csv:"login" json:"login,omitempty"`
	Password        string          `csv:"password" json:"password,omitempty"`
	Name            string          `csv:"name" json:"name,omitempty"`
	CompanyId       string          `csv:"company_id" json:"companyId,omitempty"`
	CustomerCompany CustomerCompany `gorm:"foreignKey:CompanyId;references:company_id" json:"customerCompany,omitempty"`
	CreditCards     string          `csv:"credit_cards" json:"creditCards,omitempty"`
}
