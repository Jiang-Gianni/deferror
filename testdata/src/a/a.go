package a

func f() (err error) { // want "deferror suggests "
	return nil
}

func a() (err error) {
	defer func() {}()
	return nil
}

func b() error {
	defer func() {}()
	return nil
}
