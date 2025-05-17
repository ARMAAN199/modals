package transaction

import (
	"github.com/google/uuid"
)

type Transaction struct {
	TransactionTypeID   int            `json:"transaction_type_id" validate:"required"`
	StatusID            int            `json:"status_id" validate:"required"`
	SourceEntityID      int            `json:"source_entity_id" validate:"required"`
	DestinationEntityID int            `json:"destination_entity_id" validate:"required"`
	LedgerID            uuid.UUID      `json:"ledger_id" validate:"required"`
	Amount              int            `json:"amount" validate:"required"`
	Currency            string         `json:"currency"`
	ReferenceNumber     string         `json:"reference_number"`
	Description         string         `json:"description"`
	IdempotencyKey      uuid.UUID      `json:"idempotency_key"`
	Tags                []string       `json:"tags"`
	Metadata            map[string]any `json:"metadata"`
	ChangedBy           int            `json:"changed_by"`
}

type TransactionCreateJobMessage struct {
	TransactionID int         `json:"transaction_id"`
	Transaction   Transaction `json:"transaction"`
}
