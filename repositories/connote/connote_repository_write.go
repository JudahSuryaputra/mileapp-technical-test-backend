package connote

import (
	"context"
	"mileapp-technical-test-backend/models/db"

	"github.com/gocraft/dbr"
	"github.com/gofrs/uuid"
)

func CreateConnote(tx *dbr.Tx, createConnote db.Connote) (uuid.UUID, *dbr.Tx, error) {
	columns := []string{
		"connote_service",
		"connote_service_price",
		"connote_amount",
		"connote_code",
		"connote_booking_code",
		"connote_order",
		"connote_state",
		"connote_state_id",
		"zone_code_from",
		"zone_code_to",
		"transaction_id",
		"actual_weight",
		"chargeable_weight",
		"organization_id",
		"location_id",
		"connote_total_package",
		"connote_sla_day",
		"location_name",
		"location_type",
		"source_tariff_db",
		"id_source_tariff",
		"custom_field",
		"current_location",
	}

	var connoteID uuid.UUID

	err := tx.InsertInto(db.Connote{}.TableName()).
		Columns(columns...).
		Record(createConnote).
		Returning("connote_id").
		LoadContext(context.Background(), &connoteID)
	if err != nil {
		tx.Rollback()
		return uuid.Nil, tx, err
	}

	return connoteID, tx, nil
}
