//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package xmlgroup

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"reflect"
	"time"
)

// MarshalXML implements the xml.Marshaller interface for type AccessPolicy.
func (a AccessPolicy) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type alias AccessPolicy
	aux := &struct {
		*alias
		Expiry *timeRFC3339 `xml:"Expiry"`
		Start  *timeRFC3339 `xml:"Start"`
	}{
		alias:  (*alias)(&a),
		Expiry: (*timeRFC3339)(a.Expiry),
		Start:  (*timeRFC3339)(a.Start),
	}
	return e.EncodeElement(aux, start)
}

// UnmarshalXML implements the xml.Unmarshaller interface for type AccessPolicy.
func (a *AccessPolicy) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type alias AccessPolicy
	aux := &struct {
		*alias
		Expiry *timeRFC3339 `xml:"Expiry"`
		Start  *timeRFC3339 `xml:"Start"`
	}{
		alias: (*alias)(a),
	}
	if err := d.DecodeElement(aux, &start); err != nil {
		return err
	}
	a.Expiry = (*time.Time)(aux.Expiry)
	a.Start = (*time.Time)(aux.Start)
	return nil
}

// MarshalXML implements the xml.Marshaller interface for type AppleBarrel.
func (a AppleBarrel) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type alias AppleBarrel
	aux := &struct {
		*alias
		BadApples  *[]*string `xml:"BadApples>Apple"`
		GoodApples *[]*string `xml:"GoodApples>Apple"`
	}{
		alias: (*alias)(&a),
	}
	if a.BadApples != nil {
		aux.BadApples = &a.BadApples
	}
	if a.GoodApples != nil {
		aux.GoodApples = &a.GoodApples
	}
	return e.EncodeElement(aux, start)
}

// MarshalXML implements the xml.Marshaller interface for type Banana.
func (b Banana) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "banana"
	type alias Banana
	aux := &struct {
		*alias
		Expiration *timeRFC3339 `xml:"expiration"`
	}{
		alias:      (*alias)(&b),
		Expiration: (*timeRFC3339)(b.Expiration),
	}
	return e.EncodeElement(aux, start)
}

// UnmarshalXML implements the xml.Unmarshaller interface for type Banana.
func (b *Banana) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type alias Banana
	aux := &struct {
		*alias
		Expiration *timeRFC3339 `xml:"expiration"`
	}{
		alias: (*alias)(b),
	}
	if err := d.DecodeElement(aux, &start); err != nil {
		return err
	}
	b.Expiration = (*time.Time)(aux.Expiration)
	return nil
}

// UnmarshalXML implements the xml.Unmarshaller interface for type Blob.
func (b *Blob) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type alias Blob
	aux := &struct {
		*alias
		Metadata additionalProperties `xml:"Metadata"`
	}{
		alias: (*alias)(b),
	}
	if err := d.DecodeElement(aux, &start); err != nil {
		return err
	}
	b.Metadata = (map[string]*string)(aux.Metadata)
	return nil
}

// MarshalXML implements the xml.Marshaller interface for type BlobProperties.
func (b BlobProperties) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type alias BlobProperties
	aux := &struct {
		*alias
		CopyCompletionTime *timeRFC1123 `xml:"CopyCompletionTime"`
		DeletedTime        *timeRFC1123 `xml:"DeletedTime"`
		LastModified       *timeRFC1123 `xml:"Last-Modified"`
	}{
		alias:              (*alias)(&b),
		CopyCompletionTime: (*timeRFC1123)(b.CopyCompletionTime),
		DeletedTime:        (*timeRFC1123)(b.DeletedTime),
		LastModified:       (*timeRFC1123)(b.LastModified),
	}
	return e.EncodeElement(aux, start)
}

// UnmarshalXML implements the xml.Unmarshaller interface for type BlobProperties.
func (b *BlobProperties) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type alias BlobProperties
	aux := &struct {
		*alias
		CopyCompletionTime *timeRFC1123 `xml:"CopyCompletionTime"`
		DeletedTime        *timeRFC1123 `xml:"DeletedTime"`
		LastModified       *timeRFC1123 `xml:"Last-Modified"`
	}{
		alias: (*alias)(b),
	}
	if err := d.DecodeElement(aux, &start); err != nil {
		return err
	}
	b.CopyCompletionTime = (*time.Time)(aux.CopyCompletionTime)
	b.DeletedTime = (*time.Time)(aux.DeletedTime)
	b.LastModified = (*time.Time)(aux.LastModified)
	return nil
}

// MarshalXML implements the xml.Marshaller interface for type Blobs.
func (b Blobs) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type alias Blobs
	aux := &struct {
		*alias
		Blob       *[]*Blob       `xml:"Blob"`
		BlobPrefix *[]*BlobPrefix `xml:"BlobPrefix"`
	}{
		alias: (*alias)(&b),
	}
	if b.Blob != nil {
		aux.Blob = &b.Blob
	}
	if b.BlobPrefix != nil {
		aux.BlobPrefix = &b.BlobPrefix
	}
	return e.EncodeElement(aux, start)
}

// UnmarshalXML implements the xml.Unmarshaller interface for type Container.
func (c *Container) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type alias Container
	aux := &struct {
		*alias
		Metadata additionalProperties `xml:"Metadata"`
	}{
		alias: (*alias)(c),
	}
	if err := d.DecodeElement(aux, &start); err != nil {
		return err
	}
	c.Metadata = (map[string]*string)(aux.Metadata)
	return nil
}

// MarshalXML implements the xml.Marshaller interface for type ContainerProperties.
func (c ContainerProperties) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type alias ContainerProperties
	aux := &struct {
		*alias
		LastModified *timeRFC1123 `xml:"Last-Modified"`
	}{
		alias:        (*alias)(&c),
		LastModified: (*timeRFC1123)(c.LastModified),
	}
	return e.EncodeElement(aux, start)
}

// UnmarshalXML implements the xml.Unmarshaller interface for type ContainerProperties.
func (c *ContainerProperties) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type alias ContainerProperties
	aux := &struct {
		*alias
		LastModified *timeRFC1123 `xml:"Last-Modified"`
	}{
		alias: (*alias)(c),
	}
	if err := d.DecodeElement(aux, &start); err != nil {
		return err
	}
	c.LastModified = (*time.Time)(aux.LastModified)
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Error.
func (e Error) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "message", e.Message)
	populate(objectMap, "status", e.Status)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Error.
func (e *Error) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "message":
			err = unpopulate(val, "Message", &e.Message)
			delete(rawMsg, key)
		case "status":
			err = unpopulate(val, "Status", &e.Status)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type JSONInput.
func (j JSONInput) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", j.ID)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type JSONInput.
func (j *JSONInput) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", j, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, "ID", &j.ID)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", j, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type JSONOutput.
func (j JSONOutput) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", j.ID)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type JSONOutput.
func (j *JSONOutput) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", j, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, "ID", &j.ID)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", j, err)
		}
	}
	return nil
}

// MarshalXML implements the xml.Marshaller interface for type ListContainersResponse.
func (l ListContainersResponse) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type alias ListContainersResponse
	aux := &struct {
		*alias
		Containers *[]*Container `xml:"Containers>Container"`
	}{
		alias: (*alias)(&l),
	}
	if l.Containers != nil {
		aux.Containers = &l.Containers
	}
	return e.EncodeElement(aux, start)
}

// MarshalXML implements the xml.Marshaller interface for type ModelWithByteProperty.
func (m ModelWithByteProperty) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type alias ModelWithByteProperty
	aux := &struct {
		*alias
		Bytes *string `xml:"Bytes"`
	}{
		alias: (*alias)(&m),
	}
	if m.Bytes != nil {
		encodedBytes := runtime.EncodeByteArray(m.Bytes, runtime.Base64StdFormat)
		aux.Bytes = &encodedBytes
	}
	return e.EncodeElement(aux, start)
}

// UnmarshalXML implements the xml.Unmarshaller interface for type ModelWithByteProperty.
func (m *ModelWithByteProperty) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type alias ModelWithByteProperty
	aux := &struct {
		*alias
		Bytes *string `xml:"Bytes"`
	}{
		alias: (*alias)(m),
	}
	if err := d.DecodeElement(aux, &start); err != nil {
		return err
	}
	if aux.Bytes != nil {
		if err := runtime.DecodeByteArray(*aux.Bytes, &m.Bytes, runtime.Base64StdFormat); err != nil {
			return err
		}
	}
	return nil
}

// MarshalXML implements the xml.Marshaller interface for type Slide.
func (s Slide) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type alias Slide
	aux := &struct {
		*alias
		Items *[]*string `xml:"item"`
	}{
		alias: (*alias)(&s),
	}
	if s.Items != nil {
		aux.Items = &s.Items
	}
	return e.EncodeElement(aux, start)
}

// MarshalXML implements the xml.Marshaller interface for type Slideshow.
func (s Slideshow) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "slideshow"
	type alias Slideshow
	aux := &struct {
		*alias
		Slides *[]*Slide `xml:"slide"`
	}{
		alias: (*alias)(&s),
	}
	if s.Slides != nil {
		aux.Slides = &s.Slides
	}
	return e.EncodeElement(aux, start)
}

// MarshalXML implements the xml.Marshaller interface for type StorageServiceProperties.
func (s StorageServiceProperties) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type alias StorageServiceProperties
	aux := &struct {
		*alias
		Cors *[]*CorsRule `xml:"Cors>CorsRule"`
	}{
		alias: (*alias)(&s),
	}
	if s.Cors != nil {
		aux.Cors = &s.Cors
	}
	return e.EncodeElement(aux, start)
}

func populate(m map[string]interface{}, k string, v interface{}) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}

func unpopulate(data json.RawMessage, fn string, v interface{}) error {
	if data == nil {
		return nil
	}
	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("struct field %s: %v", fn, err)
	}
	return nil
}
