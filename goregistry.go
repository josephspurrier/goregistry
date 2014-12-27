// Copyright 2014 Joseph Spurrier
// Author: Joseph Spurrier (http://josephspurrier.com)
// License: http://www.apache.org/licenses/LICENSE-2.0.html

// Generic container to use as a global registry with type conversion.
package goregistry

import (
	"fmt"
	"math"
	"strconv"
	"sync"
)

// Package variable
var reg *registry

// Registry uses map
type registry struct {
	sync.RWMutex
	interfaceMap map[string]interface{}
}

// Initialize the instance
func init() {
	reg = new(registry)
	reg.interfaceMap = make(map[string]interface{})
}

// Set the value of a map key
func Set(key string, value interface{}) {
	reg.Lock()
	reg.interfaceMap[key] = value
	reg.Unlock()
}

// Get the value of a map key
func Get(key string) (interface{}, bool) {
	reg.RLock()
	s, b := reg.interfaceMap[key]
	reg.RUnlock()
	return s, b
}

// Delete the value of a map key
func Delete(key string) {
	reg.Lock()
	delete(reg.interfaceMap, key)
	reg.Unlock()
}

// Clear all the values
func Clear() {
	reg.Lock()
	reg.interfaceMap = make(map[string]interface{})
	reg.Unlock()
}

// Get the string value of a map key
func GetString(key string) (string, bool) {
	reg.RLock()
	s, b := reg.interfaceMap[key]
	reg.RUnlock()

	if !b {
		return "", b
	}

	// Compare the type
	switch v := s.(type) {
	case rune:
		return string(v), b
	}

	return fmt.Sprintf("%v", s), b
}

// Get the bool value of a map key
func GetBool(key string) (bool, bool) {
	reg.RLock()
	s, b := reg.interfaceMap[key]
	reg.RUnlock()

	if !b {
		return false, b
	}

	// Compare the type
	switch v := s.(type) {
	case int:
		a, _ := strconv.ParseBool(strconv.Itoa(v))
		return a, b
	case string:
		a, _ := strconv.ParseBool(v)
		return a, b
	case float64:
		a, _ := strconv.ParseBool(strconv.FormatFloat(v, 'f', -1, 64))
		return a, b
	case bool:
		return v, b
	case rune:
		a, _ := strconv.ParseBool(string(v))
		return a, b
	}

	return false, b
}

// Get the int value of a map key
func GetInt(key string) (int, bool) {
	reg.RLock()
	s, b := reg.interfaceMap[key]
	reg.RUnlock()

	if !b {
		return 0, b
	}

	// Compare the type
	switch v := s.(type) {
	case int:
		return v, b
	case string:
		a, _ := strconv.ParseInt(v, 10, 0)
		return int(a), b
	case float64:
		a, _ := strconv.ParseInt(strconv.FormatFloat(round(v, .5, 0), 'f', -1, 64), 10, 0)
		return int(a), b
	case bool:
		if v {
			return 1, b
		}
		return 0, b
	case rune:
		return int(v), b
	}

	return 0, b
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
