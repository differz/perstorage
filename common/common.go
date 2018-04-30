package common

import (
	"fmt"
	"strings"
	"time"

	"../configuration"
)

const (
	configSize    = 4
	componentSize = 14
)

// ContextUpMessage print message about up status of component
func ContextUpMessage(component, message string) {
	t := time.Now()
	s := t.Format(TimeFormat()) + "\t" +
		"[" + trimToSize(configuration.Name(), configSize) + "]" +
		" " + trimToSize(component, componentSize) + " " +
		":" + message
	fmt.Println(s)
}

// TimeFormat string for project console output
// time.Now().Format(time.RFC3339) or time.Now().String()
func TimeFormat() string {
	return "2006-01-02 15:04:05.999"
}

func trimToSize(value string, size int) string {
	if len(value) > size {
		return value[0:size]
	}
	return value + strings.Repeat(" ", size-len(value))
}

// DownloadLink return download path with server name:port
func DownloadLink(link string) string {
	return configuration.ServerAddress() + "/download/" + link
}
