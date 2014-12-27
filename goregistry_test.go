// Copyright 2014 Joseph Spurrier
// Author: Joseph Spurrier (http://josephspurrier.com)
// License: http://www.apache.org/licenses/LICENSE-2.0.html

// Tests for the global registry
package goregistry_test

import (
	"fmt"
	"github.com/josephspurrier/goregistry"
	"testing"
)

// goregistry.Setup the registry
func TestMain(t *testing.T) {
	goregistry.Set("foo", "bar")
}

func TestDelete(t *testing.T) {
	goregistry.Delete("foo")

	if _, ok := goregistry.Get("foo"); ok {
		t.Error("Expected:", false, "|", "Actual:", ok)
	}

	// Reset
	TestMain(t)
}

func TestClear(t *testing.T) {
	goregistry.Clear()

	if _, ok := goregistry.Get("foo"); ok {
		t.Error("Expected:", false, "|", "Actual:", ok)
	}

	// Reset
	TestMain(t)
}

func ExampleUse() {
	goregistry.Set("test", "bar")

	if v, ok := goregistry.Get("test"); ok {
		fmt.Println(v)
		// Output: bar
	} else {
		fmt.Println("Value not found")
	}
}

// *****************************************************************************
// goregistry.Get
// *****************************************************************************

func TestGetInterface(t *testing.T) {
	if v, ok := goregistry.Get("foo"); !ok || v != "bar" {
		t.Error("Expected:", "bar", "|", "Actual:", v)
	}
}

func TestSetIntGetInterfaceInt(t *testing.T) {
	goregistry.Set("bar", 4)

	if v, ok := goregistry.Get("bar"); !ok || v != 4 {
		t.Error("Expected:", 4, "|", "Actual:", v)
	}
}

func TestGetFunc(t *testing.T) {
	goregistry.Set("func", func() string {
		return "ok"
	}())

	if v, ok := goregistry.Get("func"); !ok || v != "ok" {
		t.Error("Expected:", "ok", "|", "Actual:", v)
	}
}

// *****************************************************************************
// goregistry.GetString
// *****************************************************************************

func TestSetIntGetString(t *testing.T) {
	goregistry.Set("bar", 4)

	if v, ok := goregistry.GetString("bar"); !ok || v != "4" {
		t.Error("Expected:", "4", "|", "Actual:", v)
	}
}

func TestSetRuneGetString(t *testing.T) {
	goregistry.Set("bar", 'b')

	if v, ok := goregistry.GetString("bar"); !ok || v != "b" {
		t.Error("Expected:", "b", "|", "Actual:", v)
	}
}

func TestGetString(t *testing.T) {
	if v, ok := goregistry.GetString("foo"); !ok || v != "bar" {
		t.Error("Expected:", "bar", "|", "Actual:", v)
	}
}

func TestGetFuncString(t *testing.T) {
	if v, ok := goregistry.GetString("func"); !ok || v != "ok" {
		t.Error("Expected:", "bar", "|", "Actual:", v)
	}
}

func TestSetBoolGetString(t *testing.T) {
	goregistry.Set("bool", true)

	if v, ok := goregistry.GetString("bool"); !ok || v != "true" {
		t.Error("Expected:", true, "|", "Actual:", v)
	}
}

func TestSetFloatGetString(t *testing.T) {
	goregistry.Set("float", 1.5)

	if v, ok := goregistry.GetString("float"); !ok || v != "1.5" {
		t.Error("Expected:", "1.5", "|", "Actual:", v)
	}
}

func TestNotFoundGetString(t *testing.T) {
	if v, ok := goregistry.GetString("notwhere"); ok || v != "" {
		t.Error("Expected:", "", "|", "Actual:", v)
	}
}

// *****************************************************************************
// goregistry.GetBool
// *****************************************************************************

func TestSetBoolGetBoolTrue(t *testing.T) {
	goregistry.Set("bool", true)

	if v, ok := goregistry.GetBool("bool"); !ok || v == false {
		t.Error("Expected:", true, "|", "Actual:", v)
	}
}

func TestSetIntGetBoolTrue(t *testing.T) {
	goregistry.Set("bool", 1)

	if v, ok := goregistry.GetBool("bool"); !ok || v == false {
		t.Error("Expected:", true, "|", "Actual:", v)
	}
}

func TestSetStringGetBoolTrueT(t *testing.T) {
	goregistry.Set("bool", "T")

	if v, ok := goregistry.GetBool("bool"); !ok || v == false {
		t.Error("Expected:", true, "|", "Actual:", v)
	}
}

func TestSetRuneGetBoolTrueT(t *testing.T) {
	goregistry.Set("bool", 'T')

	if v, ok := goregistry.GetBool("bool"); !ok || v == false {
		t.Error("Expected:", true, "|", "Actual:", v)
	}
}

func TestSetFloatGetBoolTrue1(t *testing.T) {
	goregistry.Set("bool", 1.0)

	if v, ok := goregistry.GetBool("bool"); !ok || v == false {
		t.Error("Expected:", true, "|", "Actual:", v)
	}
}

func TestSetStringGetBoolFalse(t *testing.T) {
	goregistry.Set("bool", false)

	if v, ok := goregistry.GetBool("bool"); !ok || v == true {
		t.Error("Expected:", false, "|", "Actual:", v)
	}
}

func TestNotFoundGetBool(t *testing.T) {
	if v, ok := goregistry.GetBool("notwhere"); ok || v != false {
		t.Error("Expected:", false, "|", "Actual:", v)
	}
}

func TestSetFuncGetBool(t *testing.T) {
	goregistry.Set("boolfunc", func() func() { return func() {} })

	if v, ok := goregistry.GetBool("boolfunc"); !ok || v != false {
		t.Error("Expected:", false, "|", "Actual:", v)
	}
}

// *****************************************************************************
// goregistry.GetInt
// *****************************************************************************

func TestSetIntGetInt1(t *testing.T) {
	goregistry.Set("int", 1)

	if v, ok := goregistry.GetInt("int"); !ok || v != 1 {
		t.Error("Expected:", 1, "|", "Actual:", v)
	}
}

func TestSetIntGetInt0(t *testing.T) {
	goregistry.Set("int", 0)

	if v, ok := goregistry.GetInt("int"); !ok || v != 0 {
		t.Error("Expected:", 0, "|", "Actual:", v)
	}
}

func TestSetBoolGetIntT(t *testing.T) {
	goregistry.Set("int", true)

	if v, ok := goregistry.GetInt("int"); !ok || v != 1 {
		t.Error("Expected:", 1, "|", "Actual:", v)
	}
}

func TestSetBoolGetIntF(t *testing.T) {
	goregistry.Set("int", false)

	if v, ok := goregistry.GetInt("int"); !ok || v != 0 {
		t.Error("Expected:", 0, "|", "Actual:", v)
	}
}

func TestSetStringGetInt(t *testing.T) {
	goregistry.Set("int", "1")

	if v, ok := goregistry.GetInt("int"); !ok || v != 1 {
		t.Error("Expected:", 1, "|", "Actual:", v)
	}
}

func TestSetFloatGetIntRoundDown(t *testing.T) {
	goregistry.Set("int", 1.4)

	if v, ok := goregistry.GetInt("int"); !ok || v != 1 {
		t.Error("Expected:", 1, "|", "Actual:", v)
	}
}

func TestSetFloatGetIntRoundUp(t *testing.T) {
	goregistry.Set("int", 1.6)

	if v, ok := goregistry.GetInt("int"); !ok || v != 2 {
		t.Error("Expected:", 2, "|", "Actual:", v)
	}
}

func TestSetRuneGetInt(t *testing.T) {
	goregistry.Set("int", 'T')

	if v, ok := goregistry.GetInt("int"); !ok || v != 84 {
		t.Error("Expected:", 84, "|", "Actual:", v)
	}
}

func TestNotFoundGetInt(t *testing.T) {
	if v, ok := goregistry.GetInt("notwhere"); ok || v != 0 {
		t.Error("Expected:", 0, "|", "Actual:", v)
	}
}

func TestSetFuncGetInt(t *testing.T) {
	goregistry.Set("intfunc", func() func() { return func() {} })

	if v, ok := goregistry.GetInt("intfunc"); !ok || v != 0 {
		t.Error("Expected:", 0, "|", "Actual:", v)
	}
}

// *****************************************************************************
// Race Condition Test
// *****************************************************************************

func TestGoGet(t *testing.T) {
	go goregistry.Set("foo", "bar")
	go goregistry.Get("foo")
	go goregistry.GetString("foo")
	go goregistry.GetBool("foo")
	go goregistry.GetInt("foo")
	go goregistry.Delete("foo")
	go goregistry.Clear()
}

// *****************************************************************************
// Examples
// *****************************************************************************

func ExampleClear() {
	goregistry.Clear()
}

func ExampleDelete() {
	goregistry.Delete("foo")
}

func ExampleGet() {
	if value, ok := goregistry.Get("foo"); ok {
		fmt.Println(value)
	}
}

func ExampleGetBool() {
	if value, ok := goregistry.GetBool("foo"); ok {
		fmt.Println(value)
	}
}

func ExampleGetInt() {
	if value, ok := goregistry.GetInt("foo"); ok {
		fmt.Println(value)
	}
}

func ExampleGetString() {
	if value, ok := goregistry.GetString("foo"); ok {
		fmt.Println(value)
	}
}

func ExampleSet() {
	goregistry.Set("foo", "bar")
}
