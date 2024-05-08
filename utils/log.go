package utils

import "fmt"

func PrintServerInfo(localIp, port string) {
	fmt.Printf("\n 👁 Orus Media Server\n\n"+
		"\033[90mOpen in your browser:\033[0m\n\n"+
		"	➜ \033[1mLocal:\033[0m   \033[36mhttp://localhost%s\033[0m\n"+
		"	➜ \033[1mNetwork:\033[0m \033[36mhttp://%s%s\033[0m\n\n", port, localIp, port)
}
