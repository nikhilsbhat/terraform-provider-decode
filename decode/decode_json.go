package decode

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/errwrap"
	"github.com/hashicorp/terraform/helper/schema"
)

func DecodeJSON() *schema.Resource {
	return &schema.Resource{
		Create: Decode,
		Delete: schema.RemoveFromState,

		Schema: map[string]*schema.Schema{
			"json_data": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"json_map": {
				Type:     schema.TypeMap,
				Computed: true,
			},
		},
	}
}

// Test
func Decode(d *schema.ResourceData, meta interface{}) error {

	var mapJSON map[string]interface{}

	jsonString := (d.Get("json_data")).(string)
	if decodneuerr := jsonDecode([]byte(jsonString), &mapJSON); decodneuerr != nil {
		return errwrap.Wrapf("Error Decoding JSON to map", decodneuerr)
	} else {
		d.Set("json_map", mapJSON)
	}
	return nil
}

func jsonDecode(data []byte, i interface{}) error {
	err := json.Unmarshal(data, i)
	if err != nil {

		switch err.(type) {
		case *json.UnmarshalTypeError:
			return unknownTypeError(data, err)
		case *json.SyntaxError:
			return syntaxError(data, err)
		}
	}

	return nil
}

func syntaxError(data []byte, err error) error {
	syntaxErr, ok := err.(*json.SyntaxError)
	if !ok {
		return err
	}

	newline := []byte{'\x0a'}

	start := bytes.LastIndex(data[:syntaxErr.Offset], newline) + 1
	end := len(data)
	if index := bytes.Index(data[start:], newline); index >= 0 {
		end = start + index
	}

	line := bytes.Count(data[:start], newline) + 1

	err = fmt.Errorf("error occurred at line %d, %s\n%s",
		line, syntaxErr, data[start:end])
	return err
}

func unknownTypeError(data []byte, err error) error {
	unknownTypeErr, ok := err.(*json.UnmarshalTypeError)
	if !ok {
		return err
	}

	newline := []byte{'\x0a'}

	fmt.Println(bytes.LastIndex(data[:unknownTypeErr.Offset], newline))
	start := bytes.LastIndex(data[:unknownTypeErr.Offset], newline) + 1
	end := len(data)
	if index := bytes.Index(data[start:], newline); index >= 0 {
		end = start + index
	}

	line := bytes.Count(data[:start], newline) + 1

	err = fmt.Errorf("error occurred at line %d, %s\n%s\nThe data type you entered for the value is wrong",
		line, unknownTypeErr, data[start:end])
	return err
}
