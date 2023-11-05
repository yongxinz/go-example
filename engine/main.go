package main

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"sync"
)

type DataContext struct {
	lockVars sync.Mutex
	lockBase sync.Mutex
	base     map[string]reflect.Value
}

func NewDataContext() *DataContext {
	dc := &DataContext{
		base: make(map[string]reflect.Value),
	}
	return dc
}

func (dc *DataContext) Add(key string, obj interface{}) {
	dc.lockBase.Lock()
	defer dc.lockBase.Unlock()
	dc.base[key] = reflect.ValueOf(obj)
}

func (dc *DataContext) Del(keys ...string) {
	if len(keys) == 0 {
		return
	}
	dc.lockBase.Lock()
	defer dc.lockBase.Unlock()

	for _, key := range keys {
		delete(dc.base, key)
	}
}

func (dc *DataContext) Get(key string) (reflect.Value, error) {
	dc.lockBase.Lock()
	v, ok := dc.base[key]
	dc.lockBase.Unlock()
	if ok {
		return v, nil
	} else {
		return reflect.ValueOf(nil), errors.New(fmt.Sprintf("NOT FOUND key :%s ", key))
	}
}

/**
execute the injected functions: a(..)
function execute supply multi return values, but simplify ,just return one value
*/
func (dc *DataContext) ExecFunc(funcName string, parameters []reflect.Value) (reflect.Value, error) {
	dc.lockBase.Lock()
	v, ok := dc.base[funcName]
	dc.lockBase.Unlock()

	if ok {
		args := ParamsTypeChange(v, parameters)
		res := v.Call(args)
		raw, e := GetRawTypeValue(res)
		if e != nil {
			return reflect.ValueOf(nil), e
		}
		return raw, nil
	}

	return reflect.ValueOf(nil), errors.New(fmt.Sprintf("NOT FOUND function \"%s(..)\"", funcName))
}

const (
	_int   = 1
	_uint  = 2
	_float = 3
)

/*
number type exchange
*/
func ParamsTypeChange(f reflect.Value, params []reflect.Value) []reflect.Value {
	tf := f.Type()
	if tf.Kind() == reflect.Ptr {
		tf = tf.Elem()
	}
	plen := tf.NumIn()
	for i := 0; i < plen; i++ {
		switch tf.In(i).Kind() {
		case reflect.Int:
			tag := getNumType(params[i])
			if tag == _int {
				params[i] = reflect.ValueOf(int(params[i].Int()))
			} else if tag == _uint {
				params[i] = reflect.ValueOf(int(params[i].Uint()))
			} else {
				params[i] = reflect.ValueOf(int(params[i].Float()))
			}
			break
		case reflect.Int8:
			tag := getNumType(params[i])
			if tag == _int {
				params[i] = reflect.ValueOf(int8(params[i].Int()))
			} else if tag == _uint {
				params[i] = reflect.ValueOf(int8(params[i].Uint()))
			} else {
				params[i] = reflect.ValueOf(int8(params[i].Float()))
			}
			break
		case reflect.Int16:
			tag := getNumType(params[i])
			if tag == _int {
				params[i] = reflect.ValueOf(int16(params[i].Int()))
			} else if tag == _uint {
				params[i] = reflect.ValueOf(int16(params[i].Uint()))
			} else {
				params[i] = reflect.ValueOf(int16(params[i].Float()))
			}
			break
		case reflect.Int32:
			tag := getNumType(params[i])
			if tag == _int {
				params[i] = reflect.ValueOf(int32(params[i].Int()))
			} else if tag == _uint {
				params[i] = reflect.ValueOf(int32(params[i].Uint()))
			} else {
				params[i] = reflect.ValueOf(int32(params[i].Float()))
			}
			break
		case reflect.Int64:
			tag := getNumType(params[i])
			if tag == _int {
				params[i] = reflect.ValueOf(params[i].Int())
			} else if tag == _uint {
				params[i] = reflect.ValueOf(int64(params[i].Uint()))
			} else {
				params[i] = reflect.ValueOf(int64(params[i].Float()))
			}
			break
		case reflect.Uint:
			tag := getNumType(params[i])
			if tag == _int {
				params[i] = reflect.ValueOf(uint(params[i].Int()))
			} else if tag == _uint {
				params[i] = reflect.ValueOf(uint(params[i].Uint()))
			} else {
				params[i] = reflect.ValueOf(uint(params[i].Float()))
			}
			break
		case reflect.Uint8:
			tag := getNumType(params[i])
			if tag == _int {
				params[i] = reflect.ValueOf(uint8(params[i].Int()))
			} else if tag == _uint {
				params[i] = reflect.ValueOf(uint8(params[i].Uint()))
			} else {
				params[i] = reflect.ValueOf(uint8(params[i].Float()))
			}
			break
		case reflect.Uint16:
			tag := getNumType(params[i])
			if tag == _int {
				params[i] = reflect.ValueOf(uint16(params[i].Int()))
			} else if tag == _uint {
				params[i] = reflect.ValueOf(uint16(params[i].Uint()))
			} else {
				params[i] = reflect.ValueOf(uint16(params[i].Float()))
			}
			break
		case reflect.Uint32:
			tag := getNumType(params[i])
			if tag == _int {
				params[i] = reflect.ValueOf(uint32(params[i].Int()))
			} else if tag == _uint {
				params[i] = reflect.ValueOf(uint32(params[i].Uint()))
			} else {
				params[i] = reflect.ValueOf(uint32(params[i].Float()))
			}
			break
		case reflect.Uint64:
			tag := getNumType(params[i])
			if tag == _int {
				params[i] = reflect.ValueOf(uint64(params[i].Int()))
			} else if tag == _uint {
				params[i] = reflect.ValueOf(params[i].Uint())
			} else {
				params[i] = reflect.ValueOf(uint64(params[i].Float()))
			}
			break
		case reflect.Float32:
			tag := getNumType(params[i])
			if tag == _int {
				params[i] = reflect.ValueOf(float32(params[i].Int()))
			} else if tag == _uint {
				params[i] = reflect.ValueOf(float32(params[i].Uint()))
			} else {
				params[i] = reflect.ValueOf(float32(params[i].Float()))
			}
			break
		case reflect.Float64:
			tag := getNumType(params[i])
			if tag == _int {
				params[i] = reflect.ValueOf(float64(params[i].Int()))
			} else if tag == _uint {
				params[i] = reflect.ValueOf(float64(params[i].Uint()))
			} else {
				params[i] = reflect.ValueOf(params[i].Float())
			}
			break
		case reflect.Ptr:
			break
		case reflect.Interface:
			if !reflect.ValueOf(params[i]).IsValid() {
				params[i] = reflect.New(tf.In(i))
			}
		default:
			continue
		}
	}
	return params
}

func getNumType(param reflect.Value) int {
	ts := param.Kind().String()
	if strings.HasPrefix(ts, "int") {
		return _int
	}

	if strings.HasPrefix(ts, "uint") {
		return _uint
	}

	if strings.HasPrefix(ts, "float") {
		return _float
	}

	panic(fmt.Sprintf("it is not number type, type is %s !", ts))
}

/**
if want to support multi return ,change this method implements
*/
func GetRawTypeValue(rs []reflect.Value) (reflect.Value, error) {
	if len(rs) == 0 {
		return reflect.ValueOf(nil), nil
	} else {
		return rs[0], nil
	}
}

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func PrintHello(user User) {
	fmt.Println(user.Name)
	fmt.Println(user.Age)
}

func main() {
	// 假设从数据库中获取函数名
	functionName := "PrintHello"

	// 使用反射获取函数
	functionValue := reflect.ValueOf(nil).MethodByName(functionName)
	if !functionValue.IsValid() {
		fmt.Printf("Function '%s' not found\n", functionName)
		return
	}

	// 创建函数参数
	user := User{
		Name: "John Doe",
		Age:  30,
	}
	args := []reflect.Value{reflect.ValueOf(user)}

	// 动态调用函数
	functionValue.Call(args)

	// dc := NewDataContext()
	// dc.Add("hello", PrintHello)

	// params := map[string]interface{}{
	// 	"name": "zhangsan",
	// 	"age":  11,
	// }

	// funcName := "hello"

	// userType := reflect.TypeOf(User{})
	// userValue := reflect.New(userType).Elem()
	// for i := 0; i < userType.NumField(); i++ {
	// 	field := userType.Field(i)
	// 	jsonTag := field.Tag.Get("json")
	// 	if val, ok := params[jsonTag]; ok {
	// 		fieldValue := userValue.Field(i)
	// 		fieldType := fieldValue.Type()
	// 		value := reflect.ValueOf(val)

	// 		if value.Type().ConvertibleTo(fieldType) {
	// 			fieldValue.Set(value.Convert(fieldType))
	// 		}
	// 	}
	// }

	// dc.ExecFunc(funcName, []reflect.Value{userValue})
}
