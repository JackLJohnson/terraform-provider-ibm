// Code generated by go-swagger; DO NOT EDIT.

package network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteSubnetsIDTagsTagNameParams creates a new DeleteSubnetsIDTagsTagNameParams object
// with the default values initialized.
func NewDeleteSubnetsIDTagsTagNameParams() *DeleteSubnetsIDTagsTagNameParams {
	var ()
	return &DeleteSubnetsIDTagsTagNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteSubnetsIDTagsTagNameParamsWithTimeout creates a new DeleteSubnetsIDTagsTagNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteSubnetsIDTagsTagNameParamsWithTimeout(timeout time.Duration) *DeleteSubnetsIDTagsTagNameParams {
	var ()
	return &DeleteSubnetsIDTagsTagNameParams{

		timeout: timeout,
	}
}

// NewDeleteSubnetsIDTagsTagNameParamsWithContext creates a new DeleteSubnetsIDTagsTagNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteSubnetsIDTagsTagNameParamsWithContext(ctx context.Context) *DeleteSubnetsIDTagsTagNameParams {
	var ()
	return &DeleteSubnetsIDTagsTagNameParams{

		Context: ctx,
	}
}

// NewDeleteSubnetsIDTagsTagNameParamsWithHTTPClient creates a new DeleteSubnetsIDTagsTagNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteSubnetsIDTagsTagNameParamsWithHTTPClient(client *http.Client) *DeleteSubnetsIDTagsTagNameParams {
	var ()
	return &DeleteSubnetsIDTagsTagNameParams{
		HTTPClient: client,
	}
}

/*DeleteSubnetsIDTagsTagNameParams contains all the parameters to send to the API endpoint
for the delete subnets ID tags tag name operation typically these are written to a http.Request
*/
type DeleteSubnetsIDTagsTagNameParams struct {

	/*ID
	  The resource identifier

	*/
	ID string
	/*TagName
	  The name of the tag

	*/
	TagName string
	/*Version
	  Requests the version of the API as of a date in the format `YYYY-MM-DD`. Any date up to the current date may be provided. Specify the current date to request the latest version.

	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete subnets ID tags tag name params
func (o *DeleteSubnetsIDTagsTagNameParams) WithTimeout(timeout time.Duration) *DeleteSubnetsIDTagsTagNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete subnets ID tags tag name params
func (o *DeleteSubnetsIDTagsTagNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete subnets ID tags tag name params
func (o *DeleteSubnetsIDTagsTagNameParams) WithContext(ctx context.Context) *DeleteSubnetsIDTagsTagNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete subnets ID tags tag name params
func (o *DeleteSubnetsIDTagsTagNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete subnets ID tags tag name params
func (o *DeleteSubnetsIDTagsTagNameParams) WithHTTPClient(client *http.Client) *DeleteSubnetsIDTagsTagNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete subnets ID tags tag name params
func (o *DeleteSubnetsIDTagsTagNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete subnets ID tags tag name params
func (o *DeleteSubnetsIDTagsTagNameParams) WithID(id string) *DeleteSubnetsIDTagsTagNameParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete subnets ID tags tag name params
func (o *DeleteSubnetsIDTagsTagNameParams) SetID(id string) {
	o.ID = id
}

// WithTagName adds the tagName to the delete subnets ID tags tag name params
func (o *DeleteSubnetsIDTagsTagNameParams) WithTagName(tagName string) *DeleteSubnetsIDTagsTagNameParams {
	o.SetTagName(tagName)
	return o
}

// SetTagName adds the tagName to the delete subnets ID tags tag name params
func (o *DeleteSubnetsIDTagsTagNameParams) SetTagName(tagName string) {
	o.TagName = tagName
}

// WithVersion adds the version to the delete subnets ID tags tag name params
func (o *DeleteSubnetsIDTagsTagNameParams) WithVersion(version string) *DeleteSubnetsIDTagsTagNameParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the delete subnets ID tags tag name params
func (o *DeleteSubnetsIDTagsTagNameParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSubnetsIDTagsTagNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param tag_name
	if err := r.SetPathParam("tag_name", o.TagName); err != nil {
		return err
	}

	// query param version
	qrVersion := o.Version
	qVersion := qrVersion
	if qVersion != "" {
		if err := r.SetQueryParam("version", qVersion); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}