// Copyright 2014 Joseph Spurrier
// Author: Joseph Spurrier (http://josephspurrier.com)
// License: http://www.apache.org/licenses/LICENSE-2.0.html

// Tests for the global registry

package goregistry

import (
	"fmt"
	"testing"
)

// Setup the registry
func TestMain(t *testing.T) {
	Set("foo", "bar")
}

func TestDelete(t *testing.T) {
	Delete("foo")

	if _, ok := Get("foo"); ok {
		t.Error("Expected:", false, "|", "Actual:", ok)
	}

	// Reset
	TestMain(t)
}

func TestClear(t *testing.T) {
	Clear()

	if _, ok := Get("foo"); ok {
		t.Error("Expected:", false, "|", "Actual:", ok)
	}

	// Reset
	TestMain(t)
}

func ExampleUse() {
	Set("test", "bar")

	if v, ok := Get("test"); ok {
		fmt.Println(v)
		// Output: bar
	} else {
		fmt.Println("Value not found")
	}
}

// *****************************************************************************
// Get
// *****************************************************************************

func TestGetInterface(t *testing.T) {
	if v, ok := Get("foo"); !ok || v != "bar" {
		t.Error("Expected:", "bar", "|", "Actual:", v)
	}
}

func TestSetIntGetInterfaceInt(t *testing.T) {
	Set("bar", 4)

	if v, ok := Get("bar"); !ok || v != 4 {
		t.Error("Expected:", 4, "|", "Actual:", v)
	}
}

func TestGetFunc(t *testing.T) {
	Set("func", func() string {
		return "ok"
	}())

	if v, ok := Get("func"); !ok || v != "ok" {
		t.Error("Expected:", "ok", "|", "Actual:", v)
	}
}

// *****************************************************************************
// GetString
// *****************************************************************************

func TestSetIntGetString(t *testing.T) {
	Set("bar", 4)

	if v, ok := GetString("bar"); !ok || v != "4" {
		t.Error("Expected:", "4", "|", "Actual:", v)
	}
}

func TestSetRuneGetString(t *testing.T) {
	Set("bar", 'b')

	if v, ok := GetString("bar"); !ok || v != "b" {
		t.Error("Expected:", "b", "|", "Actual:", v)
	}
}

func TestGetString(t *testing.T) {
	if v, ok := GetString("foo"); !ok || v != "bar" {
		t.Error("Expected:", "bar", "|", "Actual:", v)
	}
}

func TestGetFuncString(t *testing.T) {
	if v, ok := GetString("func"); !ok || v != "ok" {
		t.Error("Expected:", "bar", "|", "Actual:", v)
	}
}

func TestSetBoolGetString(t *testing.T) {
	Set("bool", true)

	if v, ok := GetString("bool"); !ok || v != "true" {
		t.Error("Expected:", true, "|", "Actual:", v)
	}
}

func TestSetFloatGetString(t *testing.T) {
	Set("float", 1.5)

	if v, ok := GetString("float"); !ok || v != "1.5" {
		t.Error("Expected:", "1.5", "|", "Actual:", v)
	}
}

// *****************************************************************************
// GetBool
// *****************************************************************************

func TestSetBoolGetBoolTrue(t *testing.T) {
	Set("bool", true)

	if v, ok := GetBool("bool"); !ok || v == false {
		t.Error("Expected:", true, "|", "Actual:", v)
	}
}

func TestSetIntGetBoolTrue(t *testing.T) {
	Set("bool", 1)

	if v, ok := GetBool("bool"); !ok || v == false {
		t.Error("Expected:", true, "|", "Actual:", v)
	}
}

func TestSetStringGetBoolTrueT(t *testing.T) {
	Set("bool", "T")

	if v, ok := GetBool("bool"); !ok || v == false {
		t.Error("Expected:", true, "|", "Actual:", v)
	}
}

func TestSetRuneGetBoolTrueT(t *testing.T) {
	Set("bool", 'T')

	if v, ok := GetBool("bool"); !ok || v == false {
		t.Error("Expected:", true, "|", "Actual:", v)
	}
}

func TestSetFloatGetBoolTrue1(t *testing.T) {
	Set("bool", 1.0)

	if v, ok := GetBool("bool"); !ok || v == false {
		t.Error("Expected:", true, "|", "Actual:", v)
	}
}

func TestSetStringGetBoolFalse(t *testing.T) {
	Set("bool", false)

	if v, ok := GetBool("bool"); !ok || v == true {
		t.Error("Expected:", false, "|", "Actual:", v)
	}
}

// *****************************************************************************
// GetInt
// *****************************************************************************

func TestSetIntGetInt1(t *testing.T) {
	Set("int", 1)

	if v, ok := GetInt("int"); !ok || v != 1 {
		t.Error("Expected:", 1, "|", "Actual:", v)
	}
}

func TestSetIntGetInt0(t *testing.T) {
	Set("int", 0)

	if v, ok := GetInt("int"); !ok || v != 0 {
		t.Error("Expected:", 0, "|", "Actual:", v)
	}
}

func TestSetBoolGetIntT(t *testing.T) {
	Set("int", true)

	if v, ok := GetInt("int"); !ok || v != 1 {
		t.Error("Expected:", 1, "|", "Actual:", v)
	}
}

func TestSetBoolGetIntF(t *testing.T) {
	Set("int", false)

	if v, ok := GetInt("int"); !ok || v != 0 {
		t.Error("Expected:", 0, "|", "Actual:", v)
	}
}

func TestSetStringGetInt(t *testing.T) {
	Set("int", "1")

	if v, ok := GetInt("int"); !ok || v != 1 {
		t.Error("Expected:", 1, "|", "Actual:", v)
	}
}

func TestSetFloatGetInt(t *testing.T) {
	Set("int", 1.4)

	if v, ok := GetInt("int"); !ok || v != 1 {
		t.Error("Expected:", 1, "|", "Actual:", v)
	}
}

func TestSetRuneGetInt(t *testing.T) {
	Set("int", 'T')

	if v, ok := GetInt("int"); !ok || v != 84 {
		t.Error("Expected:", 84, "|", "Actual:", v)
	}
}
