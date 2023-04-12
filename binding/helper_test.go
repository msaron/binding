package binding

import (
	"fmt"
	"net/http"
	"time"
)

type AllTypes struct {
	Uint8            uint8
	PointerToUint8   *uint8
	Uint8Slice       []uint8
	Uint16           uint16
	PointerToUint16  *uint16
	Uint16Slice      []uint16
	Uint32           uint32
	PointerToUint32  *uint32
	Uint32Slice      []uint32
	Uint64           uint64
	PointerToUint64  *uint64
	Uint64Slice      []uint64
	Int8             int8
	PointerToInt8    *int8
	Int8Slice        []int8
	Int16            int16
	PointerToInt16   *int16
	Int16Slice       []int16
	Int32            int32
	PointerToInt32   *int32
	Int32Slice       []int32
	Int64            int64
	PointerToInt64   *int64
	Int64Slice       []int64
	Float32          float32
	PointerToFloat32 *float32
	Float32Slice     []float32
	Float64          float64
	PointerToFloat64 *float64
	Float64Slice     []float64
	Uint             uint
	PointerToUint    *uint
	UintSlice        []uint
	Int              int
	PointerToInt     *int
	IntSlice         []int
	Bool             bool
	PointerToBool    *bool
	BoolSlice        []bool
	String           string
	PointerToString  *string
	StringSlice      []string
	Time             time.Time
	PointerToTime    *time.Time
	TimeSlice        []time.Time
}

func (at *AllTypes) FieldMap(req *http.Request) FieldMap {
	return FieldMap{
		&at.Uint8:            Field{Form: "uint8"},
		&at.PointerToUint8:   Field{Form: "pointerToUint8"},
		&at.Uint8Slice:       Field{Form: "uint8Slice"},
		&at.Uint16:           Field{Form: "uint16"},
		&at.PointerToUint16:  Field{Form: "pointerToUint16"},
		&at.Uint16Slice:      Field{Form: "uint16Slice"},
		&at.Uint32:           Field{Form: "uint32"},
		&at.PointerToUint32:  Field{Form: "pointerToUint32"},
		&at.Uint32Slice:      Field{Form: "uint32Slice"},
		&at.Uint64:           Field{Form: "uint64"},
		&at.PointerToUint64:  Field{Form: "pointerToUint64"},
		&at.Uint64Slice:      Field{Form: "uint64Slice"},
		&at.Int8:             Field{Form: "int8"},
		&at.PointerToInt8:    Field{Form: "pointerToInt8"},
		&at.Int8Slice:        Field{Form: "int8Slice"},
		&at.Int16:            Field{Form: "int16"},
		&at.PointerToInt16:   Field{Form: "pointerToInt16"},
		&at.Int16Slice:       Field{Form: "int16Slice"},
		&at.Int32:            Field{Form: "int32"},
		&at.PointerToInt32:   Field{Form: "pointerToInt32"},
		&at.Int32Slice:       Field{Form: "int32Slice"},
		&at.Int64:            Field{Form: "int64"},
		&at.PointerToInt64:   Field{Form: "pointerToInt64"},
		&at.Int64Slice:       Field{Form: "int64Slice"},
		&at.Float32:          Field{Form: "float32"},
		&at.PointerToFloat32: Field{Form: "pointerToFloat32"},
		&at.Float32Slice:     Field{Form: "float32Slice"},
		&at.Float64:          Field{Form: "float64"},
		&at.PointerToFloat64: Field{Form: "pointerToFloat64"},
		&at.Float64Slice:     Field{Form: "float64Slice"},
		&at.Uint:             Field{Form: "uint"},
		&at.PointerToUint:    Field{Form: "pointerToUint"},
		&at.UintSlice:        Field{Form: "uintSlice"},
		&at.Int:              Field{Form: "int"},
		&at.PointerToInt:     Field{Form: "pointerToInt"},
		&at.IntSlice:         Field{Form: "intSlice"},
		&at.Bool:             Field{Form: "bool"},
		&at.PointerToBool:    Field{Form: "pointerToBool"},
		&at.BoolSlice:        Field{Form: "boolSlice"},
		&at.String:           Field{Form: "string"},
		&at.PointerToString:  Field{Form: "pointerToString"},
		&at.StringSlice:      Field{Form: "stringSlice"},
		&at.Time:             Field{Form: "time"},
		&at.PointerToTime:    Field{Form: "pointerToTime"},
		&at.TimeSlice:        Field{Form: "timeSlice"},
	}
}

func (at *AllTypes) FormValues() map[string][]string {
	fv := make(map[string][]string)
	fm := at.FieldMap(nil)

	addField := func(key, value interface{}) {
		f, ok := fm[key].(Field)
		if ok {
			if len(fv[f.Form]) == 0 {
				fv[f.Form] = make([]string, 0, 1)
			}
			fv[f.Form] = append(fv[f.Form], fmt.Sprintf("%v", value))
		}
	}

	addField(&at.Uint8, at.Uint8)
	addField(&at.PointerToUint8, *at.PointerToUint8)
	for _, v := range at.Uint8Slice {
		addField(&at.Uint8Slice, v)
	}
	addField(&at.Uint16, at.Uint16)
	addField(&at.PointerToUint16, *at.PointerToUint16)
	for _, v := range at.Uint16Slice {
		addField(&at.Uint16Slice, v)
	}
	addField(&at.Uint32, at.Uint32)
	addField(&at.PointerToUint32, *at.PointerToUint32)
	for _, v := range at.Uint32Slice {
		addField(&at.Uint32Slice, v)
	}
	addField(&at.Uint64, at.Uint64)
	addField(&at.PointerToUint64, *at.PointerToUint64)
	for _, v := range at.Uint64Slice {
		addField(&at.Uint64Slice, v)
	}
	addField(&at.Int8, at.Int8)
	addField(&at.PointerToInt8, *at.PointerToInt8)
	for _, v := range at.Int8Slice {
		addField(&at.Int8Slice, v)
	}
	addField(&at.Int16, at.Int16)
	addField(&at.PointerToInt16, *at.PointerToInt16)
	for _, v := range at.Int16Slice {
		addField(&at.Int16Slice, v)
	}
	addField(&at.Int32, at.Int32)
	addField(&at.PointerToInt32, *at.PointerToInt32)
	for _, v := range at.Int32Slice {
		addField(&at.Int32Slice, v)
	}
	addField(&at.Int64, at.Int64)
	addField(&at.PointerToInt64, *at.PointerToInt64)
	for _, v := range at.Int64Slice {
		addField(&at.Int64Slice, v)
	}
	addField(&at.Float32, at.Float32)
	addField(&at.PointerToFloat32, *at.PointerToFloat32)
	for _, v := range at.Float32Slice {
		addField(&at.Float32Slice, v)
	}
	addField(&at.Float64, at.Float64)
	addField(&at.PointerToFloat64, *at.PointerToFloat64)
	for _, v := range at.Float64Slice {
		addField(&at.Float64Slice, v)
	}
	addField(&at.Uint, at.Uint)
	addField(&at.PointerToUint, *at.PointerToUint)
	for _, v := range at.UintSlice {
		addField(&at.UintSlice, v)
	}
	addField(&at.Int, at.Int)
	addField(&at.PointerToInt, *at.PointerToInt)
	for _, v := range at.IntSlice {
		addField(&at.IntSlice, v)
	}
	addField(&at.Bool, at.Bool)
	addField(&at.PointerToBool, *at.PointerToBool)
	for _, v := range at.BoolSlice {
		addField(&at.BoolSlice, v)
	}
	addField(&at.String, at.String)
	addField(&at.PointerToString, *at.PointerToString)
	for _, v := range at.StringSlice {
		addField(&at.StringSlice, v)
	}
	addField(&at.Time, at.Time.Format(TimeFormat))
	addField(&at.PointerToTime, (*at.PointerToTime).Format(TimeFormat))
	for _, v := range at.TimeSlice {
		addField(&at.TimeSlice, v.Format(TimeFormat))
	}

	return fv
}

func NewCompleteModel() AllTypes {
	model := AllTypes{}
	model.Uint8 = 255
	model.PointerToUint8 = &model.Uint8
	model.Uint8Slice = []uint8{model.Uint8}
	model.Uint16 = 65535
	model.PointerToUint16 = &model.Uint16
	model.Uint16Slice = []uint16{model.Uint16}
	model.Uint32 = 4294967295
	model.PointerToUint32 = &model.Uint32
	model.Uint32Slice = []uint32{model.Uint32}
	model.Uint64 = 18446744073709551615
	model.PointerToUint64 = &model.Uint64
	model.Uint64Slice = []uint64{model.Uint64}
	model.Int8 = 127
	model.PointerToInt8 = &model.Int8
	model.Int8Slice = []int8{model.Int8}
	model.Int16 = 32767
	model.PointerToInt16 = &model.Int16
	model.Int16Slice = []int16{model.Int16}
	model.Int32 = 2147483647
	model.PointerToInt32 = &model.Int32
	model.Int32Slice = []int32{model.Int32}
	model.Int64 = 9223372036854775807
	model.PointerToInt64 = &model.Int64
	model.Int64Slice = []int64{model.Int64}
	model.Float32 = 3.14
	model.PointerToFloat32 = &model.Float32
	model.Float32Slice = []float32{model.Float32}
	model.Float64 = 2.718
	model.PointerToFloat64 = &model.Float64
	model.Float64Slice = []float64{model.Float64}
	model.Uint = 4294967295
	model.PointerToUint = &model.Uint
	model.UintSlice = []uint{model.Uint}
	model.Int = 4294967295
	model.PointerToInt = &model.Int
	model.IntSlice = []int{model.Int}
	model.Bool = true
	model.PointerToBool = &model.Bool
	model.BoolSlice = []bool{model.Bool}
	model.String = "I'm a little teapot"
	model.PointerToString = &model.String
	model.StringSlice = []string{model.String}
	model.Time, _ = time.Parse(TimeFormat, time.Now().Format(TimeFormat))
	model.PointerToTime = &model.Time
	model.TimeSlice = []time.Time{model.Time}

	return model
}
