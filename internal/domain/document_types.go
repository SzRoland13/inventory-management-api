package domain

type DocumentType string

const (
	IncomingDocument    DocumentType = "INCOMING"
	OutgoingDocument    DocumentType = "OUTGOING"
	TransferInDocument  DocumentType = "TRANSFER_IN"
	TransferOutDocument DocumentType = "TRANSFER_OUT"
	SaleDocument        DocumentType = "SALE"
	ReturnDocument      DocumentType = "RETURN"
)
