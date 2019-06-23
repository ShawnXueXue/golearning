package main

import (
	"fmt"
	"reflect"
	"time"
)

type Foo struct {
	X string
	Y int
}
type FooPtr struct {
	X *string
	Y *int
}

// 方法必须为public,才可以被Type.Method()获取
func (f Foo) Do() {
	fmt.Printf("X is %s, Y is %d\n", f.X, f.Y)
}

// reflect
func m1() {
	var i = 123
	var f float32 = 1.23
	var l = []string{"a", "b", "c"}
	fmt.Println(reflect.TypeOf(i))
	fmt.Println(reflect.TypeOf(f))
	fmt.Println(reflect.TypeOf(l))

	var foo Foo
	fmt.Println(reflect.TypeOf(foo))
}

// Type
func m2() {
	var s = "abc"
	fmt.Println(reflect.TypeOf(s).String())
	fmt.Println(reflect.TypeOf(s).Name())

	var f Foo
	typ := reflect.TypeOf(f)
	fmt.Println(typ.String())
	fmt.Println(typ.Name())
}

// StructField
func m3() {
	var f Foo
	typ := reflect.TypeOf(f)
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		fmt.Printf("%s type is: %s\n", field.Name, field.Type)
	}
	field2, _ := typ.FieldByName("X")
	fmt.Println(field2.Name)
}

// Method
func m4() {
	var f Foo
	typ := reflect.TypeOf(f)
	fmt.Println(typ.NumMethod())
	m := typ.Method(0)
	fmt.Println(m.Name)
	fmt.Println(m.Type)
	fmt.Println(m.Func)

	reflect.ValueOf(f).MethodByName(m.Name).Call([]reflect.Value{})
}

// Kind
func m5() {
	var f Foo
	typ := reflect.TypeOf(f)
	fmt.Println(typ)
	fmt.Println(typ.Kind())
	var f2 = &Foo{}
	typ2 := reflect.TypeOf(f2)
	fmt.Println(typ2)
	fmt.Println(typ2.Kind())
}

// Value
func m6() {
	var i = 123
	var f = Foo{"abc", 123}
	var s = "abc"
	fmt.Println(reflect.ValueOf(i))
	fmt.Println(reflect.ValueOf(i).String())
	fmt.Println(reflect.ValueOf(f))
	fmt.Println(reflect.ValueOf(f).String())
	fmt.Println(reflect.ValueOf(s))
	fmt.Println(reflect.ValueOf(s).String())
	fmt.Println(reflect.ValueOf(i).Interface())
	fmt.Println(f)
	fmt.Println(reflect.ValueOf(f).Interface() == f)
	fmt.Println(reflect.ValueOf(f).Interface())
}

// Value.field
func m7() {
	abc := "abc"
	num := 123
	var foo = FooPtr{&abc, &num}
	val := reflect.ValueOf(foo)
	fmt.Println(val.FieldByName("Y").Kind())
	typ := reflect.TypeOf(foo)
	fmt.Println(typ.FieldByName("Y"))
	rv := reflect.ValueOf(foo)
	rt := reflect.TypeOf(foo)
	for i := 0; i < rv.NumField(); i++ {
		fv := rv.Field(i)
		ft := rt.Field(i)
		fmt.Printf("%s type is %s, value is %v\n", ft.Name, ft.Type, fv.Interface())
	}
}

// set value
func m8() {
	var s = "abc"
	fv := reflect.ValueOf(s)
	fmt.Println(fv.CanSet())
	//fv.SetString("def")
	fv2 := reflect.ValueOf(&s)
	fmt.Println(fv2.CanSet())
	//fv2.SetString("def")

	var i = 123
	fv = reflect.ValueOf(i)
	fe := reflect.ValueOf(&i).Elem() //必须是指针的Value才能调用Elem
	fmt.Println(fv)
	fmt.Println(fe)
	fmt.Println(fv == fe)
	fmt.Println(fe.CanSet())
	fe.SetInt(456)
	fmt.Println(i)

	f := Foo{"abc", 123}
	elem := reflect.ValueOf(&f).Elem()
	elem.FieldByName("X").SetString("def")
	elem.FieldByName("Y").SetInt(456)
	fmt.Println(f)
}

func main() {
	//fmt.Println("m1-------------------")
	//m1()
	//fmt.Println("m2-------------------")
	//m2()
	//fmt.Println("m3-------------------")
	//m3()
	//fmt.Println("m4-------------------")
	//m4()
	//fmt.Println("m5-------------------")
	//m5()
	//fmt.Println("m6-------------------")
	//m6()
	//fmt.Println("m7-------------------")
	//m7()
	fmt.Println("m8-------------------")
	m8()

	fmt.Println(time.Now())
	//var size1 int64 = 50
	//var size2 int64 = 50
	//var createat1 = "2"
	//var createat2 = "2"
	//example := &CmdbDisk{Size: &size1, CreateAt: &createat1, Deleted: 0, Deadline: "d"}
	//old := CmdbDisk{Size: &size2, CreateAt: &createat2, Deleted: 0, Deadline: "d"}
	//setFields := []string{"Size","CreateAt","Deleted","Deadline"}
	//fmt.Println(Updated(example, old, setFields))

	//ar := []string{"1","2","3","4"}
	//fmt.Println(ar[2: 4])
}

func updated(example *CmdbDisk, old *CmdbDisk, setFields []string) (bool, error) {
	for _, setField := range setFields {
		exampleVal := reflect.ValueOf(*example)
		oldVal := reflect.ValueOf(*old)
		expFieldValue := exampleVal.FieldByName(setField)
		oldFieldValue := oldVal.FieldByName(setField)
		if expFieldValue.Kind().String() == "ptr" {
			expFieldValue = expFieldValue.Elem()
			oldFieldValue = oldFieldValue.Elem()
		}
		fmt.Println(expFieldValue.Kind())
		fmt.Println(oldFieldValue.Kind())
		switch expFieldValue.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if expFieldValue.Int() != oldFieldValue.Int() {
				return true, nil
			}
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			if expFieldValue.Uint() != oldFieldValue.Uint() {
				return true, nil
			}
		case reflect.String:
			if expFieldValue.String() != oldFieldValue.String() {
				return true, nil
			}
		case reflect.Float32, reflect.Float64:
			if expFieldValue.Float() != oldFieldValue.Float() {
				return true, nil
			}
		case reflect.Bool:
			if expFieldValue.Bool() != oldFieldValue.Bool() {
				return true, nil
			}
		case reflect.Complex64, reflect.Complex128:
			if expFieldValue.Complex() != oldFieldValue.Complex() {
				return true, nil
			}
		default:
			return false, fmt.Errorf("unsopprted field type: %s, field name: %s", expFieldValue.Kind(), setField)
		}
	}
	return false, nil
}

func Updated(example interface{}, old interface{}, setFields []string) (bool, error) {
	exampleVal := reflect.ValueOf(example)
	oldVal := reflect.ValueOf(old)
	if exampleVal.Kind().String() == "ptr" {
		exampleVal = exampleVal.Elem()
	}
	if oldVal.Kind().String() == "ptr" {
		oldVal = oldVal.Elem()
	}
	for _, setField := range setFields {
		expFieldValue := exampleVal.FieldByName(setField)
		oldFieldValue := oldVal.FieldByName(setField)
		if expFieldValue.Kind().String() == "ptr" {
			expFieldValue = expFieldValue.Elem()
			oldFieldValue = oldFieldValue.Elem()
		}
		switch expFieldValue.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if expFieldValue.Int() != oldFieldValue.Int() {
				return true, nil
			}
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			if expFieldValue.Uint() != oldFieldValue.Uint() {
				return true, nil
			}
		case reflect.String:
			if expFieldValue.String() != oldFieldValue.String() {
				return true, nil
			}
		case reflect.Float32, reflect.Float64:
			if expFieldValue.Float() != oldFieldValue.Float() {
				return true, nil
			}
		case reflect.Bool:
			if expFieldValue.Bool() != oldFieldValue.Bool() {
				return true, nil
			}
		case reflect.Complex64, reflect.Complex128:
			if expFieldValue.Complex() != oldFieldValue.Complex() {
				return true, nil
			}
		default:
			return false, fmt.Errorf("unsopprted field type: %s, field name: %s", expFieldValue.Kind(), setField)
		}
	}
	return false, nil
}

type CmdbDisk struct {
	Id          int64   `json:"Id,omitempty"`
	Uid         string  `json:"Uid,omitempty"`
	Name        string  `json:"Name,omitempty"`
	Size        *int64  `json:"Size,omitempty"`
	State       *int64  `json:"State,omitempty"`
	Type        *int64  `json:"Type,omitempty"`
	Usage       *int64  `json:"Usage,omitempty" orm:"column(u_usage)"`
	InstanceUid string  `json:"InstanceUid,omitempty"`
	CreateAt    *string `json:"CreateAt,omitempty"`
	Deadline    string  `json:"Deadline,omitempty"`
	Deleted     int64   `json:"Deleted,omitempty"`
}
