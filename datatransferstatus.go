package ocpp16types

import "fmt"

// Compile-time interface verification.
var _ fmt.Stringer = DataTransferStatus("")

// DataTransferStatus represents the result of a DataTransfer request
// as defined in OCPP 1.6.
type DataTransferStatus string

// DataTransferStatus enumeration values as defined in OCPP 1.6.
const (
	DataTransferStatusAccepted         DataTransferStatus = "Accepted"
	DataTransferStatusRejected         DataTransferStatus = "Rejected"
	DataTransferStatusUnknownMessageID DataTransferStatus = "UnknownMessageId"
	DataTransferStatusUnknownVendor    DataTransferStatus = "UnknownVendor"
)

// IsValid checks if the DataTransferStatus value is valid per OCPP 1.6.
func (d DataTransferStatus) IsValid() bool {
	switch d {
	case DataTransferStatusAccepted,
		DataTransferStatusRejected,
		DataTransferStatusUnknownMessageID,
		DataTransferStatusUnknownVendor:
		return true
	default:
		return false
	}
}

// String returns the string representation of DataTransferStatus.
func (d DataTransferStatus) String() string {
	return string(d)
}
