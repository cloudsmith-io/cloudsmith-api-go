/*
Cloudsmith API (v1)

The API to the Cloudsmith Service

API version: 1.202.1
Contact: support@cloudsmith.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudsmith

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// BadgesApiService BadgesApi service
type BadgesApiService service

type ApiBadgesVersionListRequest struct {
	ctx                context.Context
	ApiService         *BadgesApiService
	owner              string
	repo               string
	packageFormat      string
	packageName        string
	packageVersion     string
	packageIdentifiers string
	badgeToken         *string
	cacheSeconds       *string
	color              *string
	label              *string
	labelColor         *string
	logoColor          *string
	logoWidth          *string
	render             *bool
	shields            *bool
	showLatest         *bool
	style              *string
}

// Badge token to authenticate for private packages
func (r ApiBadgesVersionListRequest) BadgeToken(badgeToken string) ApiBadgesVersionListRequest {
	r.badgeToken = &badgeToken
	return r
}

// Override the shields.io badge cacheSeconds value.
func (r ApiBadgesVersionListRequest) CacheSeconds(cacheSeconds string) ApiBadgesVersionListRequest {
	r.cacheSeconds = &cacheSeconds
	return r
}

// Override the shields.io badge color value.
func (r ApiBadgesVersionListRequest) Color(color string) ApiBadgesVersionListRequest {
	r.color = &color
	return r
}

// Override the shields.io badge label value.
func (r ApiBadgesVersionListRequest) Label(label string) ApiBadgesVersionListRequest {
	r.label = &label
	return r
}

// Override the shields.io badge labelColor value.
func (r ApiBadgesVersionListRequest) LabelColor(labelColor string) ApiBadgesVersionListRequest {
	r.labelColor = &labelColor
	return r
}

// Override the shields.io badge logoColor value.
func (r ApiBadgesVersionListRequest) LogoColor(logoColor string) ApiBadgesVersionListRequest {
	r.logoColor = &logoColor
	return r
}

// Override the shields.io badge logoWidth value.
func (r ApiBadgesVersionListRequest) LogoWidth(logoWidth string) ApiBadgesVersionListRequest {
	r.logoWidth = &logoWidth
	return r
}

// If true, badge will be rendered
func (r ApiBadgesVersionListRequest) Render(render bool) ApiBadgesVersionListRequest {
	r.render = &render
	return r
}

// If true, a shields response will be generated
func (r ApiBadgesVersionListRequest) Shields(shields bool) ApiBadgesVersionListRequest {
	r.shields = &shields
	return r
}

// If true, for latest version badges a &#39;(latest)&#39; suffix is added
func (r ApiBadgesVersionListRequest) ShowLatest(showLatest bool) ApiBadgesVersionListRequest {
	r.showLatest = &showLatest
	return r
}

// Override the shields.io badge style value.
func (r ApiBadgesVersionListRequest) Style(style string) ApiBadgesVersionListRequest {
	r.style = &style
	return r
}

func (r ApiBadgesVersionListRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.BadgesVersionListExecute(r)
}

/*
BadgesVersionList Get latest package version for a package or package group.

Get latest package version for a package or package group.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param owner
 @param repo
 @param packageFormat
 @param packageName
 @param packageVersion
 @param packageIdentifiers
 @return ApiBadgesVersionListRequest
*/
func (a *BadgesApiService) BadgesVersionList(ctx context.Context, owner string, repo string, packageFormat string, packageName string, packageVersion string, packageIdentifiers string) ApiBadgesVersionListRequest {
	return ApiBadgesVersionListRequest{
		ApiService:         a,
		ctx:                ctx,
		owner:              owner,
		repo:               repo,
		packageFormat:      packageFormat,
		packageName:        packageName,
		packageVersion:     packageVersion,
		packageIdentifiers: packageIdentifiers,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *BadgesApiService) BadgesVersionListExecute(r ApiBadgesVersionListRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BadgesApiService.BadgesVersionList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/badges/version/{owner}/{repo}/{package_format}/{package_name}/{package_version}/{package_identifiers}/"
	localVarPath = strings.Replace(localVarPath, "{"+"owner"+"}", url.PathEscape(parameterToString(r.owner, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"repo"+"}", url.PathEscape(parameterToString(r.repo, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"package_format"+"}", url.PathEscape(parameterToString(r.packageFormat, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"package_name"+"}", url.PathEscape(parameterToString(r.packageName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"package_version"+"}", url.PathEscape(parameterToString(r.packageVersion, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"package_identifiers"+"}", url.PathEscape(parameterToString(r.packageIdentifiers, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.badgeToken != nil {
		localVarQueryParams.Add("badge_token", parameterToString(*r.badgeToken, ""))
	}
	if r.cacheSeconds != nil {
		localVarQueryParams.Add("cacheSeconds", parameterToString(*r.cacheSeconds, ""))
	}
	if r.color != nil {
		localVarQueryParams.Add("color", parameterToString(*r.color, ""))
	}
	if r.label != nil {
		localVarQueryParams.Add("label", parameterToString(*r.label, ""))
	}
	if r.labelColor != nil {
		localVarQueryParams.Add("labelColor", parameterToString(*r.labelColor, ""))
	}
	if r.logoColor != nil {
		localVarQueryParams.Add("logoColor", parameterToString(*r.logoColor, ""))
	}
	if r.logoWidth != nil {
		localVarQueryParams.Add("logoWidth", parameterToString(*r.logoWidth, ""))
	}
	if r.render != nil {
		localVarQueryParams.Add("render", parameterToString(*r.render, ""))
	}
	if r.shields != nil {
		localVarQueryParams.Add("shields", parameterToString(*r.shields, ""))
	}
	if r.showLatest != nil {
		localVarQueryParams.Add("show_latest", parameterToString(*r.showLatest, ""))
	}
	if r.style != nil {
		localVarQueryParams.Add("style", parameterToString(*r.style, ""))
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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apikey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Api-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorDetail
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v ErrorDetail
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
