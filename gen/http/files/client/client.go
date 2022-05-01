// Code generated by goa v3.7.3, DO NOT EDIT.
//
// files client HTTP transport
//
// Command:
// $ goa gen webcup/design

package client

import (
	"context"
	"mime/multipart"
	"net/http"
	files "webcup/gen/files"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the files service endpoint HTTP clients.
type Client struct {
	// ImportFile Doer is the HTTP client used to make requests to the importFile
	// endpoint.
	ImportFileDoer goahttp.Doer

	// DeleteFile Doer is the HTTP client used to make requests to the deleteFile
	// endpoint.
	DeleteFileDoer goahttp.Doer

	// CORS Doer is the HTTP client used to make requests to the  endpoint.
	CORSDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// FilesImportFileEncoderFunc is the type to encode multipart request for the
// "files" service "importFile" endpoint.
type FilesImportFileEncoderFunc func(*multipart.Writer, *files.ImportFilePayload) error

// NewClient instantiates HTTP clients for all the files service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		ImportFileDoer:      doer,
		DeleteFileDoer:      doer,
		CORSDoer:            doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// ImportFile returns an endpoint that makes HTTP requests to the files service
// importFile server.
func (c *Client) ImportFile(filesImportFileEncoderFn FilesImportFileEncoderFunc) goa.Endpoint {
	var (
		encodeRequest  = EncodeImportFileRequest(NewFilesImportFileEncoder(filesImportFileEncoderFn))
		decodeResponse = DecodeImportFileResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildImportFileRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ImportFileDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("files", "importFile", err)
		}
		return decodeResponse(resp)
	}
}

// DeleteFile returns an endpoint that makes HTTP requests to the files service
// deleteFile server.
func (c *Client) DeleteFile() goa.Endpoint {
	var (
		encodeRequest  = EncodeDeleteFileRequest(c.encoder)
		decodeResponse = DecodeDeleteFileResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildDeleteFileRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.DeleteFileDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("files", "deleteFile", err)
		}
		return decodeResponse(resp)
	}
}