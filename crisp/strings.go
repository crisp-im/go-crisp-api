// Copyright 2016 Crisp IM, Inc. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp


import (
  "bytes"
  "fmt"
  "io"
  "reflect"
)


// Stringify creates a string representation for given interface instance
func Stringify(instance interface{}) string {
  var buf bytes.Buffer

  v := reflect.ValueOf(instance)
  stringifyValue(&buf, v)

  return buf.String()
}


// stringifyValue prints the correct value
func stringifyValue(w io.Writer, val reflect.Value) {
  if val.Kind() == reflect.Ptr && val.IsNil() {
    w.Write([]byte("<nil>"))
    return
  }

  v := reflect.Indirect(val)

  switch v.Kind() {
      case reflect.String:
      fmt.Fprintf(w, `"%s"`, v)

    case reflect.Slice:
      w.Write([]byte{'['})
      for i := 0; i < v.Len(); i++ {
        if i > 0 {
          w.Write([]byte{' '})
        }

        stringifyValue(w, v.Index(i))
      }

      w.Write([]byte{']'})
      return

    case reflect.Struct:
      if v.Type().Name() != "" {
        w.Write([]byte(v.Type().String()))
      }

      w.Write([]byte{'{'})

      var sep bool

      for i := 0; i < v.NumField(); i++ {
        fv := v.Field(i)
        if fv.Kind() == reflect.Ptr && fv.IsNil() {
          continue
        }
        if fv.Kind() == reflect.Slice && fv.IsNil() {
          continue
        }

        if sep {
          w.Write([]byte(", "))
        } else {
          sep = true
        }

        w.Write([]byte(v.Type().Field(i).Name))
        w.Write([]byte{':'})
        stringifyValue(w, fv)
      }

      w.Write([]byte{'}'})

    default:
      if v.CanInterface() {
        fmt.Fprint(w, v.Interface())
      }
  }
}
