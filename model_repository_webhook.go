/*
 * Cloudsmith API
 *
 * The API to the Cloudsmith Service
 *
 * API version: 0.51.38
 * Contact: support@cloudsmith.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package cloudsmith

// RepositoryWebhook struct for RepositoryWebhook
type RepositoryWebhook struct {
	//
	CreatedAt string `json:"created_at,omitempty"`
	//
	CreatedBy string `json:"created_by,omitempty"`
	//
	CreatedByUrl string `json:"created_by_url,omitempty"`
	//
	DisableReason string `json:"disable_reason,omitempty"`
	//
	DisableReasonStr string `json:"disable_reason_str,omitempty"`
	//
	Events []string `json:"events"`
	//
	Identifier int32 `json:"identifier,omitempty"`
	// If enabled, the webhook will trigger on events and send payloads to the configured target URL.
	IsActive bool `json:"is_active,omitempty"`
	//
	IsLastResponseBad bool `json:"is_last_response_bad,omitempty"`
	//
	LastResponseStatus int32 `json:"last_response_status,omitempty"`
	//
	LastResponseStatusStr string `json:"last_response_status_str,omitempty"`
	//
	NumSent int32 `json:"num_sent,omitempty"`
	// The format of the payloads for webhook requests.
	RequestBodyFormat string `json:"request_body_format,omitempty"`
	//
	RequestBodyFormatStr string `json:"request_body_format_str,omitempty"`
	// The format of the payloads for webhook requests.
	RequestBodyTemplateFormat string `json:"request_body_template_format,omitempty"`
	//
	RequestBodyTemplateFormatStr string `json:"request_body_template_format_str,omitempty"`
	// The value that will be sent for the 'Content Type' header.
	RequestContentType string `json:"request_content_type,omitempty"`
	// The header to send the predefined secret in. This must be unique from existing headers or it won't be sent. You can use this as a form of authentication on the endpoint side.
	SecretHeader string `json:"secret_header,omitempty"`
	//
	SelfUrl string `json:"self_url,omitempty"`
	//
	SlugPerm string `json:"slug_perm,omitempty"`
	// The destination URL that webhook payloads will be POST'ed to.
	TargetUrl string `json:"target_url"`
	//
	Templates []WebhooksOwnerRepoTemplates `json:"templates"`
	//
	UpdatedAt string `json:"updated_at,omitempty"`
	//
	UpdatedBy string `json:"updated_by,omitempty"`
	//
	UpdatedByUrl string `json:"updated_by_url,omitempty"`
	// If enabled, SSL certificates is verified when webhooks are sent. It's recommended to leave this enabled as not verifying the integrity of SSL certificates leaves you susceptible to Man-in-the-Middle (MITM) attacks.
	VerifySsl bool `json:"verify_ssl,omitempty"`
}
