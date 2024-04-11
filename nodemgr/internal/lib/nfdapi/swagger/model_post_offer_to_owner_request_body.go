/*
 * NFD Management Service
 *
 * Service for querying and managing NFDs
 *
 * API version: 1.0
 * Contact: feedback@txnlab.dev
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type PostOfferToOwnerRequestBody struct {
	// Note to pass along to the NFD owner.  Must be provided but can be blank
	Note string `json:"note"`
	// Amount in microAlgo being offered to the NFD owner
	Offer  int64  `json:"offer"`
	Sender string `json:"sender"`
}