package bot

import (
	"fmt"
	"io/ioutil"
	"net/http"
	lg "quote/pkg/utils"
	"time"
)

func Bot() {
	var timeDuration = 1
	timer := time.NewTicker(time.Second * time.Duration(timeDuration))
	defer timer.Stop()

	for {
		select {
		case <-timer.C:
			resp, err := http.Get("http://localhost:1323/bot")

			if err != nil {
				lg.Errl.Fatal(err)
			}

			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)

			if err != nil {
				lg.Errl.Fatal(err)
			}

			fmt.Println(string(body))
		}
	}
}
