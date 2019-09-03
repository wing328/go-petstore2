/*
 * OpenAPI Petstore
 *
 * This is a sample server Petstore server. For this sample, you can use the api key `special-key` to test the authorization filters.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
import (
	"os"
)

type InlineObject1 struct {
	// Additional data to pass to server
	AdditionalMetadata string `json:"additionalMetadata,omitempty"`
	// file to upload
	File *os.File `json:"file,omitempty"`
}
