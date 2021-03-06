/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0-SNAPSHOT
 * Contact: dev@nifi.apache.org
 */

package models

// TemplateDto struct for TemplateDto
type TemplateDto struct {
	// The URI for the template.
	Uri string `json:"uri,omitempty"`
	// The id of the template.
	Id string `json:"id,omitempty"`
	// The id of the Process Group that the template belongs to.
	GroupId string `json:"groupId,omitempty"`
	// The name of the template.
	Name string `json:"name,omitempty"`
	// The description of the template.
	Description string `json:"description,omitempty"`
	// The timestamp when this template was created.
	Timestamp string `json:"timestamp,omitempty"`
	// The encoding version of this template.
	EncodingVersion string         `json:"encodingVersion,omitempty"`
	Snippet         FlowSnippetDto `json:"snippet,omitempty"`
}
