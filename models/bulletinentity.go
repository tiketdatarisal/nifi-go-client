/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0-SNAPSHOT
 * Contact: dev@nifi.apache.org
 */

package models

// BulletinEntity struct for BulletinEntity
type BulletinEntity struct {
	Id       int64  `json:"id,omitempty"`
	GroupId  string `json:"groupId,omitempty"`
	SourceId string `json:"sourceId,omitempty"`
	// When this bulletin was generated.
	Timestamp   string `json:"timestamp,omitempty"`
	NodeAddress string `json:"nodeAddress,omitempty"`
	// Indicates whether the user can read a given resource.
	CanRead  bool         `json:"canRead,omitempty"`
	Bulletin *BulletinDto `json:"bulletin,omitempty"`
}
