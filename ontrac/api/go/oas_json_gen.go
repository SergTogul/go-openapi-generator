// Code generated by ogen, DO NOT EDIT.

package ontracapi

import (
	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
)

// Encode encodes Error as json.
func (s Error) Encode(e *jx.Encoder) {
	unwrapped := jx.Raw(s)

	if len(unwrapped) != 0 {
		e.Raw(unwrapped)
	}
}

// Decode decodes Error from json.
func (s *Error) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Error to nil")
	}
	var unwrapped jx.Raw
	if err := func() error {
		v, err := d.RawAppend(nil)
		unwrapped = jx.Raw(v)
		if err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = Error(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s Error) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Error) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes Pet as json.
func (s Pet) Encode(e *jx.Encoder) {
	unwrapped := jx.Raw(s)

	if len(unwrapped) != 0 {
		e.Raw(unwrapped)
	}
}

// Decode decodes Pet from json.
func (s *Pet) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Pet to nil")
	}
	var unwrapped jx.Raw
	if err := func() error {
		v, err := d.RawAppend(nil)
		unwrapped = jx.Raw(v)
		if err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = Pet(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s Pet) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Pet) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes Pets as json.
func (s Pets) Encode(e *jx.Encoder) {
	unwrapped := []Pet(s)

	e.ArrStart()
	for _, elem := range unwrapped {
		elem.Encode(e)
	}
	e.ArrEnd()
}

// Decode decodes Pets from json.
func (s *Pets) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Pets to nil")
	}
	var unwrapped []Pet
	if err := func() error {
		unwrapped = make([]Pet, 0)
		if err := d.Arr(func(d *jx.Decoder) error {
			var elem Pet
			if err := elem.Decode(d); err != nil {
				return err
			}
			unwrapped = append(unwrapped, elem)
			return nil
		}); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = Pets(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s Pets) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Pets) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}
