// File generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation.
package ingestion

import (
	"encoding/json"
	"fmt"
)

// SourceInput - struct for SourceInput.
type SourceInput struct {
	SourceBigCommerce       *SourceBigCommerce
	SourceBigQuery          *SourceBigQuery
	SourceCSV               *SourceCSV
	SourceCommercetools     *SourceCommercetools
	SourceDocker            *SourceDocker
	SourceGA4BigQueryExport *SourceGA4BigQueryExport
	SourceJSON              *SourceJSON
	SourceShopify           *SourceShopify
}

// SourceCommercetoolsAsSourceInput is a convenience function that returns SourceCommercetools wrapped in SourceInput.
func SourceCommercetoolsAsSourceInput(v *SourceCommercetools) *SourceInput {
	return &SourceInput{
		SourceCommercetools: v,
	}
}

// SourceBigCommerceAsSourceInput is a convenience function that returns SourceBigCommerce wrapped in SourceInput.
func SourceBigCommerceAsSourceInput(v *SourceBigCommerce) *SourceInput {
	return &SourceInput{
		SourceBigCommerce: v,
	}
}

// SourceJSONAsSourceInput is a convenience function that returns SourceJSON wrapped in SourceInput.
func SourceJSONAsSourceInput(v *SourceJSON) *SourceInput {
	return &SourceInput{
		SourceJSON: v,
	}
}

// SourceCSVAsSourceInput is a convenience function that returns SourceCSV wrapped in SourceInput.
func SourceCSVAsSourceInput(v *SourceCSV) *SourceInput {
	return &SourceInput{
		SourceCSV: v,
	}
}

// SourceBigQueryAsSourceInput is a convenience function that returns SourceBigQuery wrapped in SourceInput.
func SourceBigQueryAsSourceInput(v *SourceBigQuery) *SourceInput {
	return &SourceInput{
		SourceBigQuery: v,
	}
}

// SourceGA4BigQueryExportAsSourceInput is a convenience function that returns SourceGA4BigQueryExport wrapped in SourceInput.
func SourceGA4BigQueryExportAsSourceInput(v *SourceGA4BigQueryExport) *SourceInput {
	return &SourceInput{
		SourceGA4BigQueryExport: v,
	}
}

// SourceDockerAsSourceInput is a convenience function that returns SourceDocker wrapped in SourceInput.
func SourceDockerAsSourceInput(v *SourceDocker) *SourceInput {
	return &SourceInput{
		SourceDocker: v,
	}
}

// SourceShopifyAsSourceInput is a convenience function that returns SourceShopify wrapped in SourceInput.
func SourceShopifyAsSourceInput(v *SourceShopify) *SourceInput {
	return &SourceInput{
		SourceShopify: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct.
func (dst *SourceInput) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]any
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup (SourceGA4BigQueryExport).")
	}

	// Hold the schema validity between checks
	validSchemaForModel := true

	// If the model wasn't discriminated yet, continue checking for other discriminating properties
	if validSchemaForModel {
		// Check if the model holds a property 'projectID'
		if _, ok := jsonDict["projectID"]; !ok {
			validSchemaForModel = false
		}
	}

	// If the model wasn't discriminated yet, continue checking for other discriminating properties
	if validSchemaForModel {
		// Check if the model holds a property 'datasetID'
		if _, ok := jsonDict["datasetID"]; !ok {
			validSchemaForModel = false
		}
	}

	// If the model wasn't discriminated yet, continue checking for other discriminating properties
	if validSchemaForModel {
		// Check if the model holds a property 'tablePrefix'
		if _, ok := jsonDict["tablePrefix"]; !ok {
			validSchemaForModel = false
		}
	}

	if validSchemaForModel {
		// try to unmarshal data into SourceGA4BigQueryExport
		err = newStrictDecoder(data).Decode(&dst.SourceGA4BigQueryExport)
		if err == nil && validateStruct(dst.SourceGA4BigQueryExport) == nil {
			jsonSourceGA4BigQueryExport, _ := json.Marshal(dst.SourceGA4BigQueryExport)
			if string(jsonSourceGA4BigQueryExport) == "{}" { // empty struct
				dst.SourceGA4BigQueryExport = nil
			} else {
				return nil
			}
		} else {
			dst.SourceGA4BigQueryExport = nil
		}
	}

	// Reset the schema validity for the next class check
	validSchemaForModel = true

	// If the model wasn't discriminated yet, continue checking for other discriminating properties
	if validSchemaForModel {
		// Check if the model holds a property 'storeHash'
		if _, ok := jsonDict["storeHash"]; !ok {
			validSchemaForModel = false
		}
	}

	if validSchemaForModel {
		// try to unmarshal data into SourceBigCommerce
		err = newStrictDecoder(data).Decode(&dst.SourceBigCommerce)
		if err == nil && validateStruct(dst.SourceBigCommerce) == nil {
			jsonSourceBigCommerce, _ := json.Marshal(dst.SourceBigCommerce)
			if string(jsonSourceBigCommerce) == "{}" { // empty struct
				dst.SourceBigCommerce = nil
			} else {
				return nil
			}
		} else {
			dst.SourceBigCommerce = nil
		}
	}

	// Reset the schema validity for the next class check
	validSchemaForModel = true
	// If the model wasn't discriminated yet, continue checking for other discriminating properties
	if validSchemaForModel {
		// Check if the model holds a property 'projectID'
		if _, ok := jsonDict["projectID"]; !ok {
			validSchemaForModel = false
		}
	}

	if validSchemaForModel {
		// try to unmarshal data into SourceBigQuery
		err = newStrictDecoder(data).Decode(&dst.SourceBigQuery)
		if err == nil && validateStruct(dst.SourceBigQuery) == nil {
			jsonSourceBigQuery, _ := json.Marshal(dst.SourceBigQuery)
			if string(jsonSourceBigQuery) == "{}" { // empty struct
				dst.SourceBigQuery = nil
			} else {
				return nil
			}
		} else {
			dst.SourceBigQuery = nil
		}
	}

	// Reset the schema validity for the next class check
	validSchemaForModel = true
	// If the model wasn't discriminated yet, continue checking for other discriminating properties
	if validSchemaForModel {
		// Check if the model holds a property 'projectKey'
		if _, ok := jsonDict["projectKey"]; !ok {
			validSchemaForModel = false
		}
	}

	if validSchemaForModel {
		// try to unmarshal data into SourceCommercetools
		err = newStrictDecoder(data).Decode(&dst.SourceCommercetools)
		if err == nil && validateStruct(dst.SourceCommercetools) == nil {
			jsonSourceCommercetools, _ := json.Marshal(dst.SourceCommercetools)
			if string(jsonSourceCommercetools) == "{}" { // empty struct
				dst.SourceCommercetools = nil
			} else {
				return nil
			}
		} else {
			dst.SourceCommercetools = nil
		}
	}

	// Reset the schema validity for the next class check
	validSchemaForModel = true
	// try to unmarshal data into SourceCSV
	err = newStrictDecoder(data).Decode(&dst.SourceCSV)
	if err == nil && validateStruct(dst.SourceCSV) == nil {
		jsonSourceCSV, _ := json.Marshal(dst.SourceCSV)
		if string(jsonSourceCSV) == "{}" { // empty struct
			dst.SourceCSV = nil
		} else {
			return nil
		}
	} else {
		dst.SourceCSV = nil
	}

	// try to unmarshal data into SourceDocker
	err = newStrictDecoder(data).Decode(&dst.SourceDocker)
	if err == nil && validateStruct(dst.SourceDocker) == nil {
		jsonSourceDocker, _ := json.Marshal(dst.SourceDocker)
		if string(jsonSourceDocker) == "{}" { // empty struct
			dst.SourceDocker = nil
		} else {
			return nil
		}
	} else {
		dst.SourceDocker = nil
	}

	// try to unmarshal data into SourceJSON
	err = newStrictDecoder(data).Decode(&dst.SourceJSON)
	if err == nil && validateStruct(dst.SourceJSON) == nil {
		jsonSourceJSON, _ := json.Marshal(dst.SourceJSON)
		if string(jsonSourceJSON) == "{}" { // empty struct
			dst.SourceJSON = nil
		} else {
			return nil
		}
	} else {
		dst.SourceJSON = nil
	}

	// try to unmarshal data into SourceShopify
	err = newStrictDecoder(data).Decode(&dst.SourceShopify)
	if err == nil && validateStruct(dst.SourceShopify) == nil {
		jsonSourceShopify, _ := json.Marshal(dst.SourceShopify)
		if string(jsonSourceShopify) == "{}" { // empty struct
			dst.SourceShopify = nil
		} else {
			return nil
		}
	} else {
		dst.SourceShopify = nil
	}

	return fmt.Errorf("Data failed to match schemas in oneOf(SourceInput)")
}

// Marshal data from the first non-nil pointers in the struct to JSON.
func (src SourceInput) MarshalJSON() ([]byte, error) {
	if src.SourceBigCommerce != nil {
		serialized, err := json.Marshal(&src.SourceBigCommerce)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal one of SourceBigCommerce of SourceInput: %w", err)
		}

		return serialized, nil
	}

	if src.SourceBigQuery != nil {
		serialized, err := json.Marshal(&src.SourceBigQuery)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal one of SourceBigQuery of SourceInput: %w", err)
		}

		return serialized, nil
	}

	if src.SourceCSV != nil {
		serialized, err := json.Marshal(&src.SourceCSV)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal one of SourceCSV of SourceInput: %w", err)
		}

		return serialized, nil
	}

	if src.SourceCommercetools != nil {
		serialized, err := json.Marshal(&src.SourceCommercetools)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal one of SourceCommercetools of SourceInput: %w", err)
		}

		return serialized, nil
	}

	if src.SourceDocker != nil {
		serialized, err := json.Marshal(&src.SourceDocker)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal one of SourceDocker of SourceInput: %w", err)
		}

		return serialized, nil
	}

	if src.SourceGA4BigQueryExport != nil {
		serialized, err := json.Marshal(&src.SourceGA4BigQueryExport)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal one of SourceGA4BigQueryExport of SourceInput: %w", err)
		}

		return serialized, nil
	}

	if src.SourceJSON != nil {
		serialized, err := json.Marshal(&src.SourceJSON)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal one of SourceJSON of SourceInput: %w", err)
		}

		return serialized, nil
	}

	if src.SourceShopify != nil {
		serialized, err := json.Marshal(&src.SourceShopify)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal one of SourceShopify of SourceInput: %w", err)
		}

		return serialized, nil
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance.
func (obj SourceInput) GetActualInstance() any {
	if obj.SourceBigCommerce != nil {
		return *obj.SourceBigCommerce
	}

	if obj.SourceBigQuery != nil {
		return *obj.SourceBigQuery
	}

	if obj.SourceCSV != nil {
		return *obj.SourceCSV
	}

	if obj.SourceCommercetools != nil {
		return *obj.SourceCommercetools
	}

	if obj.SourceDocker != nil {
		return *obj.SourceDocker
	}

	if obj.SourceGA4BigQueryExport != nil {
		return *obj.SourceGA4BigQueryExport
	}

	if obj.SourceJSON != nil {
		return *obj.SourceJSON
	}

	if obj.SourceShopify != nil {
		return *obj.SourceShopify
	}

	// all schemas are nil
	return nil
}

type NullableSourceInput struct {
	value *SourceInput
	isSet bool
}

func (v NullableSourceInput) Get() *SourceInput {
	return v.value
}

func (v *NullableSourceInput) Set(val *SourceInput) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceInput) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceInput(val *SourceInput) *NullableSourceInput {
	return &NullableSourceInput{value: val, isSet: true}
}

func (v NullableSourceInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value) //nolint:wrapcheck
}

func (v *NullableSourceInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value) //nolint:wrapcheck
}
