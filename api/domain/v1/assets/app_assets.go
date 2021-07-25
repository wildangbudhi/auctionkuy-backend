package assets

type AppAssets struct {
	CompanyBankID               *string `json:"company_bank_id"`
	CompanyBankAccountID        *string `json:"company_bank_account_id"`
	CompanyBankAccountOwnerName *string `json:"company_bank_account_owner_name"`
	SupportedBanks              []Banks `json:"supported_banks"`
}
