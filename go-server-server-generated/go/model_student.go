/*
 * Swagger Petstore - OpenAPI 3.0
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.11
 * Contact: apiteam@swagger.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
import (
	"time"
)

type Student struct {

	STD int64 `json:"STD,omitempty"`

	Age int32 `json:"Age,omitempty"`

	Birthdate time.Time `json:"birthdate,omitempty"`
	// Order Status
	Name string `json:"name,omitempty"`

	Complete bool `json:"complete,omitempty"`
}
