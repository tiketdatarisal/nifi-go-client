/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0-SNAPSHOT
 * Contact: dev@nifi.apache.org
 */

package models

// InstantiateTemplateRequestEntity struct for InstantiateTemplateRequestEntity
type InstantiateTemplateRequestEntity struct {
	// The x coordinate of the origin of the bounding box where the new components will be placed.
	OriginX float64 `json:"originX,omitempty"`
	// The y coordinate of the origin of the bounding box where the new components will be placed.
	OriginY float64 `json:"originY,omitempty"`
	// The identifier of the template.
	TemplateId string `json:"templateId,omitempty"`
	// The encoding version of the flow snippet. If not specified, this is automatically populated by the node receiving the user request. If the snippet is specified, the version will be the latest. If the snippet is not specified, the version will come from the underlying template. These details need to be replicated throughout the cluster to ensure consistency.
	EncodingVersion string         `json:"encodingVersion,omitempty"`
	Snippet         *FlowSnippetDto `json:"snippet,omitempty"`
	// Acknowledges that this node is disconnected to allow for mutable requests to proceed.
	DisconnectedNodeAcknowledged bool `json:"disconnectedNodeAcknowledged,omitempty"`
}
