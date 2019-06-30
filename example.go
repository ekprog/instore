package instore

type MyConfig struct {
	Item1 int
	Item2 string
}

func Example() {
	store := NewStore(Settings{
		Postfix: "_p",
	})

	// Provide settings
	config := MyConfig{
		Item1: 15,
		Item2: "Specific",
	}
	err := store.LoadItem(config)
	if err != nil {
		panic(err)
	}

	// Unpack settings
	c := new(MyConfig)
	err = store.UnloadItem(c)
	if err != nil {
		panic(err)
	}
}
