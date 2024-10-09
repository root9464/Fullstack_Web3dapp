package model

type Body struct {
	Id                int      `json:"id" gorm:"primaryKey;autoIncrement"`
	UserWalletAddress string   `json:"user_wallet_address" gorm:"type:string;not null"`
	DepositeDate      string   `json:"deposite_date" gorm:"type:string;not null"`
	ReceivingDate     string   `json:"receiving_date" gorm:"type:string;not null"`
	Amount            string   `json:"amount" gorm:"type:string;not null"`
	Rewards           string   `json:"rewards" gorm:"type:string;not null"`
	Headers           []Header `json:"headers" gorm:"foreignKey:BodyID;constraint:OnDelete:CASCADE;"` // связь "один ко многим" с каскадным удалением
}
