package random_test

import (
	"study/random"
	"testing"
)

func TestDictionary_Push_New_Element_Same_Hash(t *testing.T) {
	expected := "bar"
	dictionary := random.NewDictionary()
	dictionary.Push("foo", "bar")
	dictionary.Push("bar", "foo")
	dictionary.Push("foofoofoofooo", "foo")
	v := dictionary.Get("foo")
	if v != expected {
		t.Errorf("expected %v, but found %v \n", expected, v)
		t.Fail()
	}
}

func TestDictionary_Count(t *testing.T) {
	expected := 4
	dictionary := random.NewDictionary()
	dictionary.Push("foo", "bar")
	dictionary.Push("bar", "foo")
	dictionary.Push("foofoofoofooo", "foo")
	dictionary.Push("another", "foo")
	v := dictionary.Count()
	if v != expected {
		t.Errorf("expected %v, but found %v \n", expected, v)
		t.Fail()
	}
}

func TestDictionary_Delete(t *testing.T) {
	expected := 3
	dictionary := random.NewDictionary()
	dictionary.Push("foo", "bar")
	dictionary.Push("bar", "foo")
	dictionary.Push("foofoofoofooo", "foo")
	dictionary.Push("another", "foo")
	dictionary.Delete("foo")
	v := dictionary.Count()
	if v != expected {
		t.Errorf("expected %v, but found %v \n", expected, v)
		t.Fail()
	}
}

func TestDictionary_Get(t *testing.T) {
	expected := "bar"
	dictionary := random.NewDictionary()
	dictionary.Push("foo", "bar")
	dictionary.Push("bar", "foo")
	dictionary.Push("foofoofoofooo", "foo")
	dictionary.Push("another", "foo")
	v := dictionary.Get("foo")
	if v != expected {
		t.Errorf("expected %v, but found %v \n", expected, v)
		t.Fail()
	}
}
