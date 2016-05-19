// generated with github.com/Bridgevine/wsdlgo; DO NOT EDIT

package types

import (
	"encoding/xml"
	"time"
)

// against "unused imports"
var _ time.Time
var _ xml.Name

type intReqNil struct {
	*int32
}

// MarshalXML satisfies the XML Marshaler interface for type intReqNil.
func (t intReqNil) MarshalXML(e *xml.Encoder, s xml.StartElement) error {
	if t.int32 == nil {
		return e.EncodeElement("", s)
	}

	return e.EncodeElement(t, s)
}

type stringReqNil struct {
	*string
}

// MarshalXML satisfies the XML Marshaler interface for type stringReqNil.
func (t stringReqNil) MarshalXML(e *xml.Encoder, s xml.StartElement) error {
	if t.string == nil {
		return e.EncodeElement("", s)
	}

	return e.EncodeElement(t, s)
}

type dateTimeReqNil struct {
	*time.Time
}

// MarshalXML satisfies the XML Marshaler interface for type dateTimeReqNil.
func (t dateTimeReqNil) MarshalXML(e *xml.Encoder, s xml.StartElement) error {
	if t.Time == nil {
		return e.EncodeElement("", s)
	}

	tt := time.Time(*t.Time)
	if tt.IsZero() {
		return e.EncodeElement("", s)
	}

	return e.EncodeElement(t, s)
}

type myelements struct {
	Nonboth     *string        `xml:"nonboth"`
	Minzero     *int32         `xml:"minzero"`
	Nilint      intReqNil      `xml:"nilint"`
	Nilstring   stringReqNil   `xml:"nilstring"`
	Minzeronil  *string        `xml:"minzeronil"`
	DateOfBirth dateTimeReqNil `xml:"DateOfBirth"`
}
