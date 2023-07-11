// Code generated by go-swagger; DO NOT EDIT.

package example

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ma91n/rfc7807gen/models"
)

// HelloReader is a Reader for the Hello structure.
type HelloReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HelloReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHelloOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewHelloOK creates a HelloOK with default headers values
func NewHelloOK() *HelloOK {
	return &HelloOK{}
}

/*HelloOK handles this case with default header values.

successful operation
*/
type HelloOK struct {
	Payload *models.Hello
}

func (o *HelloOK) Error() string {
	return fmt.Sprintf("[GET /hello][%d] helloOK  %+v", 200, o.Payload)
}

func (o *HelloOK) GetPayload() *models.Hello {
	return o.Payload
}

func (o *HelloOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Hello)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}