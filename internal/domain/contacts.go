package domain

type ContactInfo string

const (
	BillingContactInfo  ContactInfo = "BILLING"
	ShippingContactInfo ContactInfo = "SHIPPING"
	BothContactInfo     ContactInfo = "BOTH"
	OtherContactInfo    ContactInfo = "OTHER"
)
