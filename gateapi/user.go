/*
 * Spinnaker API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type User struct {

	Email string `json:"email,omitempty"`

	AccountNonExpired bool `json:"accountNonExpired,omitempty"`

	AccountNonLocked bool `json:"accountNonLocked,omitempty"`

	Username string `json:"username,omitempty"`

	CredentialsNonExpired bool `json:"credentialsNonExpired,omitempty"`

	LastName string `json:"lastName,omitempty"`

	FirstName string `json:"firstName,omitempty"`

	Authorities []GrantedAuthority `json:"authorities,omitempty"`

	AllowedAccounts []string `json:"allowedAccounts,omitempty"`

	Enabled bool `json:"enabled,omitempty"`

	Roles []string `json:"roles,omitempty"`
}
