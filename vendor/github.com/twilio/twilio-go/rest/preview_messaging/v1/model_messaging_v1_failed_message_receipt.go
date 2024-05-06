/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Bulk Messaging and Broadcast
 * Bulk Sending is a public Twilio REST API for 1:Many Message creation up to 100 recipients. Broadcast is a public Twilio REST API for 1:Many Message creation up to 10,000 recipients via file upload.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

// MessagingV1FailedMessageReceipt struct for MessagingV1FailedMessageReceipt
type MessagingV1FailedMessageReceipt struct {
	// The recipient phone number
	To string `json:"to,omitempty"`
	// The description of the error_code
	ErrorMessage string `json:"error_message,omitempty"`
	// The error code associated with the message creation attempt
	ErrorCode int `json:"error_code,omitempty"`
}