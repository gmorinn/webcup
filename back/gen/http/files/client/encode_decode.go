// Code generated by goa v3.7.2, DO NOT EDIT.
//
// files HTTP client encoders and decoders
//
// Command:
// $ goa gen webcup/back/design

package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"
	files "webcup/back/gen/files"

	goahttp "goa.design/goa/v3/http"
)

// BuildImportFileRequest instantiates a HTTP request object with method and
// path set to call the "files" service "importFile" endpoint
func (c *Client) BuildImportFileRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ImportFileFilesPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("files", "importFile", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeImportFileRequest returns an encoder for requests sent to the files
// importFile server.
func EncodeImportFileRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*files.ImportFilePayload)
		if !ok {
			return goahttp.ErrInvalidType("files", "importFile", "*files.ImportFilePayload", v)
		}
		if p.Oauth != nil {
			head := *p.Oauth
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		if p.JWTToken != nil {
			head := *p.JWTToken
			req.Header.Set("jwtToken", head)
		}
		if err := encoder(req).Encode(p); err != nil {
			return goahttp.ErrEncodingError("files", "importFile", err)
		}
		return nil
	}
}

// NewFilesImportFileEncoder returns an encoder to encode the multipart request
// for the "files" service "importFile" endpoint.
func NewFilesImportFileEncoder(encoderFn FilesImportFileEncoderFunc) func(r *http.Request) goahttp.Encoder {
	return func(r *http.Request) goahttp.Encoder {
		body := &bytes.Buffer{}
		mw := multipart.NewWriter(body)
		return goahttp.EncodingFunc(func(v interface{}) error {
			p := v.(*files.ImportFilePayload)
			if err := encoderFn(mw, p); err != nil {
				return err
			}
			r.Body = ioutil.NopCloser(body)
			r.Header.Set("Content-Type", mw.FormDataContentType())
			return mw.Close()
		})
	}
}

// DecodeImportFileResponse returns a decoder for responses returned by the
// files importFile endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeImportFileResponse may return the following errors:
//	- "unknown_error" (type *files.UnknownError): http.StatusInternalServerError
//	- error: internal error
func DecodeImportFileResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusCreated:
			var (
				body ImportFileResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("files", "importFile", err)
			}
			err = ValidateImportFileResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("files", "importFile", err)
			}
			res := NewImportFileResultCreated(&body)
			return res, nil
		case http.StatusInternalServerError:
			var (
				body ImportFileUnknownErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("files", "importFile", err)
			}
			err = ValidateImportFileUnknownErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("files", "importFile", err)
			}
			return nil, NewImportFileUnknownError(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("files", "importFile", resp.StatusCode, string(body))
		}
	}
}

// BuildDeleteFileRequest instantiates a HTTP request object with method and
// path set to call the "files" service "deleteFile" endpoint
func (c *Client) BuildDeleteFileRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: DeleteFileFilesPath()}
	req, err := http.NewRequest("PATCH", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("files", "deleteFile", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeDeleteFileRequest returns an encoder for requests sent to the files
// deleteFile server.
func EncodeDeleteFileRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*files.DeleteFilePayload)
		if !ok {
			return goahttp.ErrInvalidType("files", "deleteFile", "*files.DeleteFilePayload", v)
		}
		if p.Oauth != nil {
			head := *p.Oauth
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		if p.JWTToken != nil {
			head := *p.JWTToken
			req.Header.Set("jwtToken", head)
		}
		body := NewDeleteFileRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("files", "deleteFile", err)
		}
		return nil
	}
}

// DecodeDeleteFileResponse returns a decoder for responses returned by the
// files deleteFile endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeDeleteFileResponse may return the following errors:
//	- "unknown_error" (type *files.UnknownError): http.StatusInternalServerError
//	- error: internal error
func DecodeDeleteFileResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body DeleteFileResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("files", "deleteFile", err)
			}
			err = ValidateDeleteFileResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("files", "deleteFile", err)
			}
			res := NewDeleteFileResultOK(&body)
			return res, nil
		case http.StatusInternalServerError:
			var (
				body DeleteFileUnknownErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("files", "deleteFile", err)
			}
			err = ValidateDeleteFileUnknownErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("files", "deleteFile", err)
			}
			return nil, NewDeleteFileUnknownError(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("files", "deleteFile", resp.StatusCode, string(body))
		}
	}
}

// unmarshalResFileResponseBodyToFilesResFile builds a value of type
// *files.ResFile from a value of type *ResFileResponseBody.
func unmarshalResFileResponseBodyToFilesResFile(v *ResFileResponseBody) *files.ResFile {
	res := &files.ResFile{
		ID:   *v.ID,
		Name: *v.Name,
		URL:  *v.URL,
		Mime: v.Mime,
		Size: v.Size,
	}

	return res
}
