package validation

type Array []interface{}

type Object map[string]interface{}

type Patch map[string]*Config

type Payload map[string]*Config

// WithFields does something, probably deprecated.
func (payload Payload) WithFields(fields Payload) Payload {
  for k, v := range payload { fields[k] = v }
  return fields
}

type Optional map[string]*Config
