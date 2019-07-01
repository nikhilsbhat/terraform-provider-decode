package decode

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/hashicorp/errwrap"
	"github.com/hashicorp/terraform/helper/schema"
)

// DecodeJSON will be called by provider to convert json to map
func decodeJSON() *schema.Resource {
	return &schema.Resource{
		Create: deCODE,
		Read:   deCODE,
		Delete: schema.RemoveFromState,

		Schema: map[string]*schema.Schema{
			"jsonfile": {
				Type:        schema.TypeString,
				Description: "Path to JSON file which has to be decoded to ",
				Required:    true,
				ForceNew:    true,
			},
			"json_data": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"alert_map": {
				Type:     schema.TypeMap,
				Computed: true,
			},
			"documentation": {
				Type:     schema.TypeMap,
				Computed: true,
			},
			"notificationChannels": {
				Type:     schema.TypeList,
				Computed: true,
			},
		},
	}
}

// Decode is the core function that decodes the json to map
func deCODE(d *schema.ResourceData, _ interface{}) error {

	var mapJSON map[string]interface{}

	file := d.Get("jsonfile").(string)
	//sonString := (d.Get("json_data")).(string)

	jsonData, jserr := readFile(file)
	if jserr != nil {
		return jserr
	}

	if decodneuerr := jsonDecode(jsonData, &mapJSON); decodneuerr != nil {
		return errwrap.Wrapf("Error Decoding JSON to map", decodneuerr)
	}

	//return errwrap.Wrapf(fmt.Sprintf("The map value %v", mapJSON), fmt.Errorf("Test"))
	d.Set("json_map", mapJSON)
	d.SetId("none")
	return nil
}

func setValidAttributes(value map[string]interface{}) error {

	return nil
}

func readFile(filename string) ([]byte, error) {
	if _, dirneuerr := os.Stat(filename); os.IsNotExist(dirneuerr) {
		return nil, dirneuerr
	}

	content, conterr := ioutil.ReadFile(filename)
	if conterr != nil {
		return nil, conterr
	}
	return content, nil
	//return nil, nil
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
