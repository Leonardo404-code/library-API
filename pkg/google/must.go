package google

func Must() Google {
	storage, err := New()
	if err != nil {
		panic(err)
	}

	return storage
}
