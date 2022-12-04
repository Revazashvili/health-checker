/*
Copyright (C) 2022  Levan Revazashvili <levanrevazashvili@gmail.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package main

import (
	"fmt"
  "os"
  "strconv"
	"time"
  "net/http"
)

func main() {
  duration, err := strconv.Atoi(os.Getenv("DURATION_IN_SECONDS"))
  if err != nil {
      fmt.Println("Error during duration conversion")
      os.Exit(1)
  }

  url := os.Getenv("HEALTH_CHECK_URL")
  if url == ""{
      fmt.Println("health check url can't be null or empty")
      os.Exit(1)
  }

  for {
    time.Sleep(time.Duration(duration) * time.Second)

	  res, err := http.Get(url)
    if err != nil {
      notify("error making http request");
      os.Exit(1)
    }

    if res.StatusCode == 500 {
      notify("service is down");
      os.Exit(1)
    }
    fmt.Println("service is up");
  }
}


func notify(message string) {
  fmt.Println(message);
}