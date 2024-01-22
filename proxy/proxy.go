package proxy

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

var proxies []string

func Load() []string {
	file, err := os.OpenFile("./proxies.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(file)

	if !scanner.Scan() {
		return nil
	}
	for scanner.Scan() {
		proxies = append(proxies, scanner.Text())
	}
	defer file.Close()
	return proxies

}

func Get() string {
	split := strings.Split(proxies[rand.Intn(len(proxies))], ":")
	if len(split) == 4 {
		proxy := string("http://") + split[2] + ":" + split[3] + "@" + split[0] + ":" + split[1]
		return proxy
	} else if len(split) == 2 {
		proxy := split[0] + ":" + split[1]
		return proxy
	}
	return ""
}
