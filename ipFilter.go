package main

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

func ipFilter(w http.ResponseWriter, r *http.Request) {
	var err error
	fmt.Println(r.RemoteAddr)
	reqIP := strings.Split(r.RemoteAddr, ":")[0]
	for _, ip := range config.Server.BlockedIPs {
		if reqIP == ip {
			err = errors.New("ip not allowed to post images")
		}
	}

	for _, ip := range config.Server.AllowedIPs {
		if reqIP != ip {
			err = errors.New("ip not allowed to post images")
		}
	}
	//return err
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}
}
