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

type TotalsOkResponseBodyMintedTotals struct {
	Day      int64 `json:"day,omitempty"`
	Lifetime int64 `json:"lifetime,omitempty"`
	Month    int64 `json:"month,omitempty"`
	Week     int64 `json:"week,omitempty"`
}