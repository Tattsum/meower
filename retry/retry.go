package retry

import (
	"time"
)

type RetryFunc func(int) error

func Do(n int, f RetryFunc) (err error) {
	for i := 0; i < n; i++ {
		err := f(i)
		if err == nil {
			return nil
		}
	}
	return err
}

func DoSleep(n int, d time.Duration, f RetryFunc) (err error) {
	for i := 0; i < n; i++ {
		err = f(i)
		if err == nil {
			return nil
		}
		time.Sleep(d)
	}
	return err
}

func Forever(f RetryFunc) {
	for i := 0; ; i++ {
		err := f(i)
		if err == nil {
			return
		}
	}
}

func ForeverSleep(d time.Duration, f RetryFunc) {
	for i := 0; ; i++ {
		err := f(i)
		if err == nil {
			return
		}
		time.Sleep(d)
	}
}
