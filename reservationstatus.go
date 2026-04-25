package ocpp16types

import "fmt"

// Compile-time interface verification.
var _ fmt.Stringer = ReservationStatus("")

// ReservationStatus represents the result of a ReserveNow request.
type ReservationStatus string

// ReservationStatus enumeration values as defined in OCPP 1.6.
const (
	ReservationStatusAccepted    ReservationStatus = "Accepted"
	ReservationStatusFaulted     ReservationStatus = "Faulted"
	ReservationStatusOccupied    ReservationStatus = "Occupied"
	ReservationStatusRejected    ReservationStatus = "Rejected"
	ReservationStatusUnavailable ReservationStatus = "Unavailable"
)

// IsValid checks if the ReservationStatus value is valid per OCPP 1.6.
func (r ReservationStatus) IsValid() bool {
	switch r {
	case ReservationStatusAccepted,
		ReservationStatusFaulted,
		ReservationStatusOccupied,
		ReservationStatusRejected,
		ReservationStatusUnavailable:
		return true
	default:
		return false
	}
}

// String returns the string representation of ReservationStatus.
func (r ReservationStatus) String() string {
	return string(r)
}
