package teams

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/luizalabs/tapi/models"
)

// UpdateTeamReader is a Reader for the UpdateTeam structure.
type UpdateTeamReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *UpdateTeamReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateTeamOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewUpdateTeamDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewUpdateTeamOK creates a UpdateTeamOK with default headers values
func NewUpdateTeamOK() *UpdateTeamOK {
	return &UpdateTeamOK{}
}

/*UpdateTeamOK handles this case with default header values.

Team details
*/
type UpdateTeamOK struct {
	Payload *models.Team
}

func (o *UpdateTeamOK) Error() string {
	return fmt.Sprintf("[PUT /teams/{team_id}][%d] updateTeamOK  %+v", 200, o.Payload)
}

func (o *UpdateTeamOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Team)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTeamDefault creates a UpdateTeamDefault with default headers values
func NewUpdateTeamDefault(code int) *UpdateTeamDefault {
	return &UpdateTeamDefault{
		_statusCode: code,
	}
}

/*UpdateTeamDefault handles this case with default header values.

Error
*/
type UpdateTeamDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the update team default response
func (o *UpdateTeamDefault) Code() int {
	return o._statusCode
}

func (o *UpdateTeamDefault) Error() string {
	return fmt.Sprintf("[PUT /teams/{team_id}][%d] updateTeam default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateTeamDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}