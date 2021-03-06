/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0-SNAPSHOT
 * Contact: dev@nifi.apache.org
 */

package models

// CreateTemplateRequestEntity struct for CreateTemplateRequestEntity
type CreateTemplateRequestEntity struct {
	// The name of the template.
	Name string `json:"name,omitempty"`
	// The description of the template.
	Description string `json:"description,omitempty"`
	// The identifier of the snippet.
	SnippetId string `json:"snippetId,omitempty"`
	// Acknowledges that this node is disconnected to allow for mutable requests to proceed.
	DisconnectedNodeAcknowledged bool `json:"disconnectedNodeAcknowledged,omitempty"`
}
