// SPDX-FileCopyrightText: Copyright (C) SchedMD LLC.
// SPDX-License-Identifier: Apache-2.0

package utils

import (
	"cmp"
	"errors"
	"reflect"
	"regexp"

	"k8s.io/utils/ptr"
	"k8s.io/utils/set"
)

// GetStateAsSet returns states as a set. A nil pointer is an empty set.
func GetStateAsSet[T cmp.Ordered](states *[]T) set.Set[T] {
	return set.New(ptr.Deref(states, nil)...)
}

// Clone returns a deep copy of in.
func Clone[T any](in *T) *T {
	if in == nil {
		return nil
	}
	out := deepCopyVal(reflect.ValueOf(in).Elem()).Interface().(T)
	return &out
}

func deepCopyVal(src reflect.Value) reflect.Value {
	switch src.Kind() {
	case reflect.Pointer:
		if src.IsNil() {
			return reflect.Zero(src.Type())
		}
		dst := reflect.New(src.Elem().Type())
		dst.Elem().Set(deepCopyVal(src.Elem()))
		return dst
	case reflect.Slice:
		if src.IsNil() {
			return reflect.Zero(src.Type())
		}
		dst := reflect.MakeSlice(src.Type(), src.Len(), src.Len())
		for i := range src.Len() {
			dst.Index(i).Set(deepCopyVal(src.Index(i)))
		}
		return dst
	case reflect.Map:
		if src.IsNil() {
			return reflect.Zero(src.Type())
		}
		dst := reflect.MakeMapWithSize(src.Type(), src.Len())
		iter := src.MapRange()
		for iter.Next() {
			dst.SetMapIndex(deepCopyVal(iter.Key()), deepCopyVal(iter.Value()))
		}
		return dst
	case reflect.Array:
		dst := reflect.New(src.Type()).Elem()
		for i := range src.Len() {
			dst.Index(i).Set(deepCopyVal(src.Index(i)))
		}
		return dst
	case reflect.Struct:
		dst := reflect.New(src.Type()).Elem()
		dst.Set(src)
		for i := range src.NumField() {
			f := dst.Field(i)
			if !f.CanSet() {
				continue
			}
			switch f.Kind() {
			case reflect.Pointer, reflect.Slice, reflect.Map, reflect.Interface, reflect.Array, reflect.Struct:
				f.Set(deepCopyVal(src.Field(i)))
			}
		}
		return dst
	case reflect.Interface:
		if src.IsNil() {
			return reflect.Zero(src.Type())
		}
		return deepCopyVal(src.Elem())
	default:
		return src
	}
}

func ParseNodeName(nodeConf string) (string, error) {
	re := regexp.MustCompile(`(?i)NodeName=([^\s]+)`)
	matches := re.FindStringSubmatch(nodeConf)
	if len(matches) < 2 {
		return "", errors.New("NodeName not found in node configuration string")
	}
	return matches[1], nil
}
