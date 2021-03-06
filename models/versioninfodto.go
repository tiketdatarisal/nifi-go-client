/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0-SNAPSHOT
 * Contact: dev@nifi.apache.org
 */

package models

import (
	"time"
)

// VersionInfoDto struct for VersionInfoDto
type VersionInfoDto struct {
	// The version of this NiFi.
	NiFiVersion string `json:"niFiVersion,omitempty"`
	// Java JVM vendor
	JavaVendor string `json:"javaVendor,omitempty"`
	// Java version
	JavaVersion string `json:"javaVersion,omitempty"`
	// Host operating system name
	OsName string `json:"osName,omitempty"`
	// Host operating system version
	OsVersion string `json:"osVersion,omitempty"`
	// Host operating system architecture
	OsArchitecture string `json:"osArchitecture,omitempty"`
	// Build tag
	BuildTag string `json:"buildTag,omitempty"`
	// Build revision or commit hash
	BuildRevision string `json:"buildRevision,omitempty"`
	// Build branch
	BuildBranch string `json:"buildBranch,omitempty"`
	// Build timestamp
	BuildTimestamp time.Time `json:"buildTimestamp,omitempty"`
}
