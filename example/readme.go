package example

type R struct{}

func quick(a func(int) (bool, error)) (err error) {
	return nil
}

// func (r *R) MyExample(now time.Duration, i *int) (a int, err error) {
// 	return 0, fmt.Errorf("new err is right here")
// }

/*
	Default
*/
// defer func(now time.Duration, i *int) {
//     if err != nil {
//         err = fmt.Errorf("r.MyExample(%v, %v): %w", now, i ,err)
//     }
// }(now, i)

/*
	Wrap
*/
// defer dfrr.Wrap(&err, "r.MyExample(%v, %v)",now, i)
