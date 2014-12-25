// Copyright (C) Joseph Spurrier. All rights reserved.
// Author: Joseph Spurrier (http://josephspurrier.com)
// License: http://www.apache.org/licenses/LICENSE-2.0.html

// Tests for the global registry

package goregistry

import (
	"testing"
)

// Setup the registry
func TestMain(t *testing.T) {
	reg := Registry()
	reg.Set("foo", "bar")
}

func TestDelete(t *testing.T) {
	reg := Registry()

	reg.Delete("foo")

	if _, ok := reg.Get("foo"); ok {
		t.Error("Expected:", false, "|", "Actual:", ok)
	}

	// Reset
	TestMain(t)
}

func TestClear(t *testing.T) {
	reg := Registry()

	reg.Clear()

	if _, ok := reg.Get("foo"); ok {
		t.Error("Expected:", false, "|", "Actual:", ok)
	}

	// Reset
	TestMain(t)
}

// *****************************************************************************
// Get
// *****************************************************************************

func TestGetInterface(t *testing.T) {
	reg := Registry()

	if v, ok := reg.Get("foo"); !ok || v != "bar" {
		t.Error("Expected:", "bar", "|", "Actual:", v)
	}
}

func TestSetIntGetInterfaceInt(t *testing.T) {
	reg := Registry()
	reg.Set("bar", 4)

	if v, ok := reg.Get("bar"); !ok || v != 4 {
		t.Error("Expected:", 4, "|", "Actual:", v)
	}
}

func TestGetFunc(t *testing.T) {
	reg := Registry()

	reg.Set("func", func() string {
		return "ok"
	}())

	if v, ok := reg.Get("func"); !ok || v != "ok" {
		t.Error("Expected:", "ok", "|", "Actual:", v)
	}
}

// *****************************************************************************
// GetString
// *****************************************************************************

func TestSetIntGetString(t *testing.T) {
	reg := Registry()
	reg.Set("bar", 4)

	if v, ok := reg.GetString("bar"); !ok || v != "4" {
		t.Error("Expected:", "4", "|", "Actual:", v)
	}
}

func TestSetRuneGetString(t *testing.T) {
	reg := Registry()
	reg.Set("bar", 'b')

	if v, ok := reg.GetString("bar"); !ok || v != "b" {
		t.Error("Expected:", "b", "|", "Actual:", v)
	}
}

func TestGetString(t *testing.T) {
	reg := Registry()

	if v, ok := reg.GetString("foo"); !ok || v != "bar" {
		t.Error("Expected:", "bar", "|", "Actual:", v)
	}
}

func TestGetFuncString(t *testing.T) {
	reg := Registry()

	if v, ok := reg.GetString("func"); !ok || v != "ok" {
		t.Error("Expected:", "bar", "|", "Actual:", v)
	}
}

func TestSetBoolGetString(t *testing.T) {
	reg := Registry()
	reg.Set("bool", true)

	if v, ok := reg.GetString("bool"); !ok || v != "true" {
		t.Error("Expected:", true, "|", "Actual:", v)
	}
}

// *****************************************************************************
// GetBool
// *****************************************************************************

func TestSetBoolGetBoolTrue(t *testing.T) {
	reg := Registry()
	reg.Set("bool", true)

	if v, ok := reg.GetBool("bool"); !ok || v == false {
		t.Error("Expected:", true, "|", "Actual:", v)
	}
}

func TestSetIntGetBoolTrue(t *testing.T) {
	reg := Registry()
	reg.Set("bool", 1)

	if v, ok := reg.GetBool("bool"); !ok || v == false {
		t.Error("Expected:", true, "|", "Actual:", v)
	}
}

func TestSetStringGetBoolTrueT(t *testing.T) {
	reg := Registry()
	reg.Set("bool", "T")

	if v, ok := reg.GetBool("bool"); !ok || v == false {
		t.Error("Expected:", true, "|", "Actual:", v)
	}
}

func TestSetRuneGetBoolTrueT(t *testing.T) {
	reg := Registry()
	reg.Set("bool", 'T')

	if v, ok := reg.GetBool("bool"); !ok || v == false {
		t.Error("Expected:", true, "|", "Actual:", v)
	}
}

func TestSetFloatGetBoolTrue1(t *testing.T) {
	reg := Registry()
	reg.Set("bool", 1.0)

	if v, ok := reg.GetBool("bool"); !ok || v == false {
		t.Error("Expected:", true, "|", "Actual:", v)
	}
}

func TestSetStringGetBoolFalse(t *testing.T) {
	reg := Registry()
	reg.Set("bool", false)

	if v, ok := reg.GetBool("bool"); !ok || v == true {
		t.Error("Expected:", false, "|", "Actual:", v)
	}
}

// *****************************************************************************
// GetInt
// *****************************************************************************

func TestSetIntGetInt1(t *testing.T) {
	reg := Registry()
	reg.Set("int", 1)

	if v, ok := reg.GetInt("int"); !ok || v != 1 {
		t.Error("Expected:", 1, "|", "Actual:", v)
	}
}

func TestSetIntGetInt0(t *testing.T) {
	reg := Registry()
	reg.Set("int", 0)

	if v, ok := reg.GetInt("int"); !ok || v != 0 {
		t.Error("Expected:", 0, "|", "Actual:", v)
	}
}

func TestSetBoolGetIntT(t *testing.T) {
	reg := Registry()
	reg.Set("int", true)

	if v, ok := reg.GetInt("int"); !ok || v != 1 {
		t.Error("Expected:", 1, "|", "Actual:", v)
	}
}

func TestSetBoolGetIntF(t *testing.T) {
	reg := Registry()
	reg.Set("int", false)

	if v, ok := reg.GetInt("int"); !ok || v != 0 {
		t.Error("Expected:", 0, "|", "Actual:", v)
	}
}

func TestSetStringGetInt(t *testing.T) {
	reg := Registry()
	reg.Set("int", "1")

	if v, ok := reg.GetInt("int"); !ok || v != 1 {
		t.Error("Expected:", 1, "|", "Actual:", v)
	}
}

func TestSetFloatGetInt(t *testing.T) {
	reg := Registry()
	reg.Set("int", 1.4)

	if v, ok := reg.GetInt("int"); !ok || v != 1 {
		t.Error("Expected:", 1, "|", "Actual:", v)
	}
}

func TestSetRuneGetInt(t *testing.T) {
	reg := Registry()
	reg.Set("int", 'T')

	if v, ok := reg.GetInt("int"); !ok || v != 84 {
		t.Error("Expected:", 84, "|", "Actual:", v)
	}
}
