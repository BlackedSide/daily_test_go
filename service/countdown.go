package service

import (
	"fmt"
	"io"
)

func CountDown(out io.Writer, sleeper Sleeper, base int, finalWord string) {

	for i := base; i > 0; i-- {
		sleeper.Sleep()
		_, _ = fmt.Fprintln(out, i)
	}

	sleeper.Sleep()
	_, _ = fmt.Fprint(out, finalWord)

}
