package domain

type DocumentStatus string

const (
	DraftDocumentStatus     DocumentStatus = "DRAFT"
	PendingDocumentStatus   DocumentStatus = "PENDING"
	CompletedDocumentStatus DocumentStatus = "COMPLETED"
	CancelledDocumentStatus DocumentStatus = "CANCELLED"
)
