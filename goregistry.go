// Copyright (C) Joseph Spurrier. All rights reserved.
// Author: Joseph Spurrier (http://josephspurrier.com)
// License: http://www.apache.org/licenses/LICENSE-2.0.html

// Global registry that uses the singleton pattern

package goregistry

import (
	"math"
	"strconv"
)

// Package variable
var i *registry

// Registry uses map
type registry struct {
	m map[string]interface{}
}

// Get a singleton instance of the registry type
func Registry() *registry {
	if i != nil {
		return i
	} else {
		i = new(registry)
		i.m = make(map[string]interface{})
		return i
	}
}

// Set the value of a map key
func (r *registry) Set(key string, value interface{}) {
	r.m[key] = value
}

// Get the value of a map key
func (r *registry) Get(key string) (interface{}, bool) {
	s, b := r.m[key]
	return s, b
}

// Get the string value of a map key
func (r *registry) GetString(key string) (string, bool) {
	s, b := r.m[key]

	// Compare the type
	switch s.(type) {
	case int:
		return strconv.Itoa(s.(int)), b
	case string:
		return s.(string), b
	case float64:
		return strconv.FormatFloat(s.(float64), 'f', -1, 64), b
	case bool:
		return strconv.FormatBool(s.(bool)), b
	case rune:
		return string(s.(rune)), b
	default:
		return s.(string), b
	}
}

// Get the bool value of a map key
func (r *registry) GetBool(key string) (bool, bool) {
	s, b := r.m[key]

	// Compare the type
	switch s.(type) {
	case int:
		a, _ := strconv.ParseBool(strconv.Itoa(s.(int)))
		return a, b
	case string:
		a, _ := strconv.ParseBool(s.(string))
		return a, b
	case float64:
		a, _ := strconv.ParseBool(strconv.FormatFloat(s.(float64), 'f', -1, 64))
		return a, b
	case bool:
		return s.(bool), b
	case rune:
		a, _ := strconv.ParseBool(string(s.(rune)))
		return a, b
	default:
		return s.(bool), b
	}
}

// Get the bool value of a map key
func (r *registry) GetInt(key string) (int, bool) {
	s, b := r.m[key]

	// Compare the type
	switch s.(type) {
	case int:
		return s.(int), b
	case string:
		a, _ := strconv.ParseInt(s.(string), 10, 0)
		return int(a), b
	case float64:
		a, _ := strconv.ParseInt(strconv.FormatFloat(round(s.(float64), .5, 0), 'f', -1, 64), 10, 0)
		return int(a), b
	case bool:
		if s.(bool) {
			return 1, b
		} else {
			return 0, b
		}
	case rune:
		return int(s.(rune)), b
	default:
		return s.(int), b
	}
}

// Source: https://gist.github.com/pelegm/c48cff315cd223f7cf7b
func round(val float64, roundOn float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	_div := math.Copysign(div, val)
	_roundOn := math.Copysign(roundOn, val)
	if _div >= _roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}
