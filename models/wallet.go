package models

type Wallet struct {
	ID      int  `json:"id" gorm:"primary_key;auto_increment"`
	Balance int  `json:"balance"`
	Credit  int  `json:"credit"`
	Debit   int  `json:"debit"`
	UserId  int  `json:"userId"`
	User    User `gorm:"foreignKey:UserId"`
}
