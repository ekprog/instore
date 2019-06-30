package instore

import "testing"

func Test_Basic(t *testing.T) {
	type MyConfig struct {
		Item1 int
		Item2 string
	}
	config := MyConfig{
		Item1: 15,
		Item2: "Specific",
	}

	store := NewStore(Settings{
		Postfix: "_p",
	})
	err := store.LoadItem(config)
	if err != nil {
		t.Fatal(err)
	}

	c := new(MyConfig)
	err = store.UnloadItem(c)
	if err != nil {
		t.Fatal(err)
	}
}
