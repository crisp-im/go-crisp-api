// Copyright 2016 Crisp IM, Inc. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp


import (
  "errors"
  "reflect"
)


type caller struct {
  Func        reflect.Value
  Args        reflect.Type
  ArgsPresent bool
  Out         bool
}

var (
  errorCallerNotFunc     = errors.New("f is not function")
  errorCallerNot2Args    = errors.New("f should have 1 or 2 args")
  errorCallerMaxOneValue = errors.New("f should return not more than one value")
)


func newCaller(f interface{}) (*caller, error) {
  fVal := reflect.ValueOf(f)
  if fVal.Kind() != reflect.Func {
    return nil, errorCallerNotFunc
  }

  fType := fVal.Type()
  if fType.NumOut() > 1 {
    return nil, errorCallerMaxOneValue
  }

  curCaller := &caller{
    Func: fVal,
    Out:  fType.NumOut() == 1,
  }
  if fType.NumIn() == 1 {
    curCaller.Args = nil
    curCaller.ArgsPresent = false
  } else if fType.NumIn() == 2 {
    curCaller.Args = fType.In(1)
    curCaller.ArgsPresent = true
  } else {
    return nil, errorCallerNot2Args
  }

  return curCaller, nil
}


func (c *caller) callFunc(args interface{}) []reflect.Value {
  a := []reflect.Value{reflect.ValueOf(args).Elem()}
  if !c.ArgsPresent {
    a = a[0:1]
  }

  return c.Func.Call(a)
}
