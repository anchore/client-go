/*
 * Anchore Engine API Server
 *
 * This is the Anchore Engine API. Provides the primary external API for users of the service.
 *
 * API version: 0.1.16
 * Contact: nurmi@anchore.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package external

import (
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"github.com/antihax/optional"
)

// Linger please
var (
	_ _context.Context
)

// RepositoryCredentialsApiService RepositoryCredentialsApi service
type RepositoryCredentialsApiService service

// AddRepositoryOpts Optional parameters for the method 'AddRepository'
type AddRepositoryOpts struct {
    Autosubscribe optional.Bool
    Lookuptag optional.String
    Dryrun optional.Bool
    XAnchoreAccount optional.String
}

/*
AddRepository Add repository to watch
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param repository full repository to add e.g. docker.io/library/alpine
 * @param optional nil or *AddRepositoryOpts - Optional Parameters:
 * @param "Autosubscribe" (optional.Bool) -  flag to enable/disable auto tag_update activation when new images from a repo are added
 * @param "Lookuptag" (optional.String) -  use specified existing tag to perform repo scan (default is 'latest')
 * @param "Dryrun" (optional.Bool) -  flag to return tags in the repository without actually watching the repository, default is false
 * @param "XAnchoreAccount" (optional.String) -  An account name to change the resource scope of the request to that account, if permissions allow (admin only)
@return []Subscription
*/
func (a *RepositoryCredentialsApiService) AddRepository(ctx _context.Context, repository string, localVarOptionals *AddRepositoryOpts) ([]Subscription, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []Subscription
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/repositories"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	localVarQueryParams.Add("repository", parameterToString(repository, ""))
	if localVarOptionals != nil && localVarOptionals.Autosubscribe.IsSet() {
		localVarQueryParams.Add("autosubscribe", parameterToString(localVarOptionals.Autosubscribe.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Lookuptag.IsSet() {
		localVarQueryParams.Add("lookuptag", parameterToString(localVarOptionals.Lookuptag.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Dryrun.IsSet() {
		localVarQueryParams.Add("dryrun", parameterToString(localVarOptionals.Dryrun.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.XAnchoreAccount.IsSet() {
		localVarHeaderParams["x-anchore-account"] = parameterToString(localVarOptionals.XAnchoreAccount.Value(), "")
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
