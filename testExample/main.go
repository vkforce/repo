package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type DlqResultDto struct {
	EndOffset string       `json:"endOffset"`
	Payload   []MessageDto `json:"payload"`
}

var registeredConverters map[reflect.Type]map[reflect.Type]reflect.Value = make(map[reflect.Type]map[reflect.Type]reflect.Value)

func RegisterConverter(convertFunc interface{}) {
	v := reflect.ValueOf(convertFunc)
	t := v.Type()

	if t.Kind() != reflect.Func {
		panic("Type converter must be a function")
	}
	if t.NumIn() != 1 && t.NumOut() != 1 {
		panic("Type converter must take in 1 parameter and return 1 paramater")
	}
	in := t.In(0)
	out := t.Out(0)

	if in.Kind() != reflect.Pointer || out.Kind() != reflect.Pointer {
		panic("Type convert must accept pointer and output a pointer")
	}

	if inMapValue, ok := registeredConverters[in.Elem()]; ok {
		inMapValue[out.Elem()] = v
	} else {
		registeredConverters[in.Elem()] = make(map[reflect.Type]reflect.Value)
		registeredConverters[in.Elem()][out.Elem()] = v
	}
}

func Convert(in interface{}, out interface{}) interface{} {
	inValue := reflect.ValueOf(in)

	inType := inValue.Type()
	if inType.Kind() == reflect.Array || inType.Kind() == reflect.Slice {
		return convertArray(in, out)
	}

	if inType.Kind() == reflect.Pointer {
		inType = inType.Elem()
	} else {
		inValue = inValue.Addr()
	}

	outType := reflect.TypeOf(out)
	if outType.Kind() == reflect.Pointer {
		outType = outType.Elem()
	}

	if inType == outType {
		return inValue.Interface()
	}

	if inMapValue, ok := registeredConverters[inType]; ok {
		if convertFunc, ok := inMapValue[outType]; ok {
			outValues := convertFunc.Call([]reflect.Value{inValue})
			return outValues[0].Interface()
		}
	}
	return nil
}

func convertArray(in interface{}, out interface{}) interface{} {
	inValue := reflect.ValueOf(in)

	totalElems := inValue.Len()
	output := []interface{}{}

	for i := 0; i < totalElems; i++ {
		elem := inValue.Index(i)
		output = append(output, Convert(elem.Interface(), out))
	}

	return output
}

type MessageDto struct {
	CollectionName   string         `json:"collectionName"`
	CorrelationID    string         `json:"correlationId"`
	ErrorMessage     string         `json:"errorMessage"`
	Payload          map[string]any `json:"payload"`
	PointProductCode string         `json:"pointProductCode"`
	TimeStamp        string         `json:"timeStamp"`
}

func main() {
	var endOffset int64 = 12345

	payload := []MessageDto{}
	payload = append(payload, MessageDto{
		CollectionName:   "name",
		CorrelationID:    "correlationId",
		ErrorMessage:     "no error",
		Payload:          map[string]any{"id": "1234"},
		PointProductCode: "RBI",
		TimeStamp:        "2024-09-09",
	})
	payload = append(payload, MessageDto{
		CollectionName:   "name",
		CorrelationID:    "correlationId",
		ErrorMessage:     "no error",
		Payload:          map[string]any{"id": "1234"},
		PointProductCode: "RBI",
		TimeStamp:        "2024-09-09",
	})

	result := map[string]interface{}{"payload": payload, "endOffset": strconv.FormatInt(endOffset, 10)}
	resultPtr := &result
	dtos := Convert(resultPtr, &DlqResultDto{})
	fmt.Println(dtos)
}
