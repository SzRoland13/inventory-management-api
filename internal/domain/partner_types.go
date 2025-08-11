package domain

type PartnerType string

const (
	CustomerPartner PartnerType = "CUSTOMER"
	SupplierPartner PartnerType = "SUPPLIER"
	BothPartner     PartnerType = "BOTH"
)
