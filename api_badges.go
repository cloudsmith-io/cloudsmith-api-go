/*
 * Cloudsmith API
 *
 * The API to the Cloudsmith Service
 *
 * API version: 0.51.65
 * Contact: support@cloudsmith.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package cloudsmith

import (
	_context "context"
	"github.com/antihax/optional"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// BadgesApiService BadgesApi service
type BadgesApiService service

// BadgesVersionListOpts Optional parameters for the method 'BadgesVersionList'
type BadgesVersionListOpts struct {
	BadgeToken optional.String
	Render     optional.Bool
}

/*
BadgesVersionList Get latest package version for a package or package group.
Get latest package version for a package or package group.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param owner
 * @param repo
 * @param packageFormat
 * @param packageName
 * @param packageVersion
 * @param packageIdentifiers
 * @param optional nil or *BadgesVersionListOpts - Optional Parameters:
 * @param "BadgeToken" (optional.String) -  Badge token to authenticate for private packages
 * @param "Render" (optional.Bool) -  If true, badge will be rendered
@return map[string]interface{}
*/
func (a *BadgesApiService) BadgesVersionList(ctx _context.Context, owner string, repo string, packageFormat string, packageName string, packageVersion string, packageIdentifiers string, localVarOptionals *BadgesVersionListOpts) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]interface{}
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/badges/version/{owner}/{repo}/{package_format}/{package_name}/{package_version}/{package_identifiers}/"
	localVarPath = strings.Replace(localVarPath, "{"+"owner"+"}", _neturl.QueryEscape(parameterToString(owner, "")), -1)

	localVarPath = strings.Replace(localVarPath, "{"+"repo"+"}", _neturl.QueryEscape(parameterToString(repo, "")), -1)

	localVarPath = strings.Replace(localVarPath, "{"+"package_format"+"}", _neturl.QueryEscape(parameterToString(packageFormat, "")), -1)

	localVarPath = strings.Replace(localVarPath, "{"+"package_name"+"}", _neturl.QueryEscape(parameterToString(packageName, "")), -1)

	localVarPath = strings.Replace(localVarPath, "{"+"package_version"+"}", _neturl.QueryEscape(parameterToString(packageVersion, "")), -1)

	localVarPath = strings.Replace(localVarPath, "{"+"package_identifiers"+"}", _neturl.QueryEscape(parameterToString(packageIdentifiers, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if localVarOptionals != nil && localVarOptionals.BadgeToken.IsSet() {
		localVarQueryParams.Add("badge_token", parameterToString(localVarOptionals.BadgeToken.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Render.IsSet() {
		localVarQueryParams.Add("render", parameterToString(localVarOptionals.Render.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["X-Api-Key"] = key
		}
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["X-CSRFToken"] = key
		}
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v Status
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v Status
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
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
