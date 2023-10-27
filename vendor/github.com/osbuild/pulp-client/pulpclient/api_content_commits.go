/*
Pulp 3 API

Fetch, Upload, Organize, and Distribute Software Packages

API version: v3
Contact: pulp-list@redhat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pulpclient

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"reflect"
)


// ContentCommitsAPIService ContentCommitsAPI service
type ContentCommitsAPIService service

type ContentCommitsAPIContentOstreeCommitsListRequest struct {
	ctx context.Context
	ApiService *ContentCommitsAPIService
	checksum *string
	limit *int32
	offset *int32
	ordering *[]string
	pulpHrefIn *[]string
	pulpIdIn *[]string
	repositoryVersion *string
	repositoryVersionAdded *string
	repositoryVersionRemoved *string
	fields *[]string
	excludeFields *[]string
}

// Filter results where checksum matches value
func (r ContentCommitsAPIContentOstreeCommitsListRequest) Checksum(checksum string) ContentCommitsAPIContentOstreeCommitsListRequest {
	r.checksum = &checksum
	return r
}

// Number of results to return per page.
func (r ContentCommitsAPIContentOstreeCommitsListRequest) Limit(limit int32) ContentCommitsAPIContentOstreeCommitsListRequest {
	r.limit = &limit
	return r
}

// The initial index from which to return the results.
func (r ContentCommitsAPIContentOstreeCommitsListRequest) Offset(offset int32) ContentCommitsAPIContentOstreeCommitsListRequest {
	r.offset = &offset
	return r
}

// Ordering  * &#x60;pulp_id&#x60; - Pulp id * &#x60;-pulp_id&#x60; - Pulp id (descending) * &#x60;pulp_created&#x60; - Pulp created * &#x60;-pulp_created&#x60; - Pulp created (descending) * &#x60;pulp_last_updated&#x60; - Pulp last updated * &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending) * &#x60;pulp_type&#x60; - Pulp type * &#x60;-pulp_type&#x60; - Pulp type (descending) * &#x60;upstream_id&#x60; - Upstream id * &#x60;-upstream_id&#x60; - Upstream id (descending) * &#x60;timestamp_of_interest&#x60; - Timestamp of interest * &#x60;-timestamp_of_interest&#x60; - Timestamp of interest (descending) * &#x60;checksum&#x60; - Checksum * &#x60;-checksum&#x60; - Checksum (descending) * &#x60;relative_path&#x60; - Relative path * &#x60;-relative_path&#x60; - Relative path (descending) * &#x60;pk&#x60; - Pk * &#x60;-pk&#x60; - Pk (descending)
func (r ContentCommitsAPIContentOstreeCommitsListRequest) Ordering(ordering []string) ContentCommitsAPIContentOstreeCommitsListRequest {
	r.ordering = &ordering
	return r
}

// Multiple values may be separated by commas.
func (r ContentCommitsAPIContentOstreeCommitsListRequest) PulpHrefIn(pulpHrefIn []string) ContentCommitsAPIContentOstreeCommitsListRequest {
	r.pulpHrefIn = &pulpHrefIn
	return r
}

// Multiple values may be separated by commas.
func (r ContentCommitsAPIContentOstreeCommitsListRequest) PulpIdIn(pulpIdIn []string) ContentCommitsAPIContentOstreeCommitsListRequest {
	r.pulpIdIn = &pulpIdIn
	return r
}

// Repository Version referenced by HREF
func (r ContentCommitsAPIContentOstreeCommitsListRequest) RepositoryVersion(repositoryVersion string) ContentCommitsAPIContentOstreeCommitsListRequest {
	r.repositoryVersion = &repositoryVersion
	return r
}

// Repository Version referenced by HREF
func (r ContentCommitsAPIContentOstreeCommitsListRequest) RepositoryVersionAdded(repositoryVersionAdded string) ContentCommitsAPIContentOstreeCommitsListRequest {
	r.repositoryVersionAdded = &repositoryVersionAdded
	return r
}

// Repository Version referenced by HREF
func (r ContentCommitsAPIContentOstreeCommitsListRequest) RepositoryVersionRemoved(repositoryVersionRemoved string) ContentCommitsAPIContentOstreeCommitsListRequest {
	r.repositoryVersionRemoved = &repositoryVersionRemoved
	return r
}

// A list of fields to include in the response.
func (r ContentCommitsAPIContentOstreeCommitsListRequest) Fields(fields []string) ContentCommitsAPIContentOstreeCommitsListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentCommitsAPIContentOstreeCommitsListRequest) ExcludeFields(excludeFields []string) ContentCommitsAPIContentOstreeCommitsListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentCommitsAPIContentOstreeCommitsListRequest) Execute() (*PaginatedostreeOstreeCommitResponseList, *http.Response, error) {
	return r.ApiService.ContentOstreeCommitsListExecute(r)
}

/*
ContentOstreeCommitsList List ostree commits

A ViewSet class for OSTree commits.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ContentCommitsAPIContentOstreeCommitsListRequest
*/
func (a *ContentCommitsAPIService) ContentOstreeCommitsList(ctx context.Context) ContentCommitsAPIContentOstreeCommitsListRequest {
	return ContentCommitsAPIContentOstreeCommitsListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginatedostreeOstreeCommitResponseList
func (a *ContentCommitsAPIService) ContentOstreeCommitsListExecute(r ContentCommitsAPIContentOstreeCommitsListRequest) (*PaginatedostreeOstreeCommitResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedostreeOstreeCommitResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentCommitsAPIService.ContentOstreeCommitsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/content/ostree/commits/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.checksum != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "checksum", r.checksum, "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
	}
	if r.ordering != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ordering", r.ordering, "csv")
	}
	if r.pulpHrefIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_href__in", r.pulpHrefIn, "csv")
	}
	if r.pulpIdIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_id__in", r.pulpIdIn, "csv")
	}
	if r.repositoryVersion != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "repository_version", r.repositoryVersion, "")
	}
	if r.repositoryVersionAdded != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "repository_version_added", r.repositoryVersionAdded, "")
	}
	if r.repositoryVersionRemoved != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "repository_version_removed", r.repositoryVersionRemoved, "")
	}
	if r.fields != nil {
		t := *r.fields
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "fields", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "fields", t, "multi")
		}
	}
	if r.excludeFields != nil {
		t := *r.excludeFields
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "exclude_fields", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "exclude_fields", t, "multi")
		}
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
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

type ContentCommitsAPIContentOstreeCommitsReadRequest struct {
	ctx context.Context
	ApiService *ContentCommitsAPIService
	ostreeOstreeCommitHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r ContentCommitsAPIContentOstreeCommitsReadRequest) Fields(fields []string) ContentCommitsAPIContentOstreeCommitsReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentCommitsAPIContentOstreeCommitsReadRequest) ExcludeFields(excludeFields []string) ContentCommitsAPIContentOstreeCommitsReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentCommitsAPIContentOstreeCommitsReadRequest) Execute() (*OstreeOstreeCommitResponse, *http.Response, error) {
	return r.ApiService.ContentOstreeCommitsReadExecute(r)
}

/*
ContentOstreeCommitsRead Inspect an ostree commit

A ViewSet class for OSTree commits.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ostreeOstreeCommitHref
 @return ContentCommitsAPIContentOstreeCommitsReadRequest
*/
func (a *ContentCommitsAPIService) ContentOstreeCommitsRead(ctx context.Context, ostreeOstreeCommitHref string) ContentCommitsAPIContentOstreeCommitsReadRequest {
	return ContentCommitsAPIContentOstreeCommitsReadRequest{
		ApiService: a,
		ctx: ctx,
		ostreeOstreeCommitHref: ostreeOstreeCommitHref,
	}
}

// Execute executes the request
//  @return OstreeOstreeCommitResponse
func (a *ContentCommitsAPIService) ContentOstreeCommitsReadExecute(r ContentCommitsAPIContentOstreeCommitsReadRequest) (*OstreeOstreeCommitResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OstreeOstreeCommitResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentCommitsAPIService.ContentOstreeCommitsRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "{ostree_ostree_commit_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"ostree_ostree_commit_href"+"}", parameterValueToString(r.ostreeOstreeCommitHref, "ostreeOstreeCommitHref"), -1)  // NOTE: paths aren't escaped because Pulp uses hrefs as path parameters

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fields != nil {
		t := *r.fields
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "fields", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "fields", t, "multi")
		}
	}
	if r.excludeFields != nil {
		t := *r.excludeFields
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "exclude_fields", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "exclude_fields", t, "multi")
		}
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
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