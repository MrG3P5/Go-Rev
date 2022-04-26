package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"strings"
)

const rethrow_panic = "_____rethrow"

var colorReset = "\033[0m"
var colorGreen = "\033[32m"
var colorRed = "\033[31m"
var colorCyan = "\033[36m"

type (
	E         interface{}
	exception struct {
		finally func()
		Error   E
	}
)

func Throw() {
	panic(rethrow_panic)
}

func This(f func()) (e exception) {
	e = exception{nil, nil}
	defer func() {
		e.Error = recover()
	}()
	f()
	return
}

func (e exception) Catch(f func(err E)) {
	if e.Error != nil {
		defer func() {
			if e.finally != nil {
				e.finally()
			}

			if err := recover(); err != nil {
				if err == rethrow_panic {
					err = e.Error
				}
				panic(err)
			}
		}()
		f(e.Error)
	} else if e.finally != nil {
		e.finally()
	}
}

func banner() {
	__banner__ := string(colorRed) + `
  ____             ____            
 / ___| ___       |  _ \ _____   __
| |  _ / _ \ _____| |_) / _ \ \ / /
| |_| | (_) |_____|  _ <  __/\ V / 
 \____|\___/      |_| \_\___| \_/  ` + string(colorGreen) + `Created By X - MrG3P5` + string(colorReset)
	fmt.Printf("\x1bc")
	fmt.Println(__banner__)
	fmt.Println()
}

func ReverseIP() {
	banner()
	var ipList string
	client := &http.Client{}

	fmt.Print(string(colorCyan) + "[" + string(colorGreen) + "?" + string(colorCyan) + "] " + string(colorReset) + "IP List : ")
	fmt.Scanf("%s", &ipList)

	bytesRead, _ := ioutil.ReadFile(ipList)
	file_content := string(bytesRead)
	lines := strings.Split(file_content, "\n")
	output, _ := os.Create("result/rev-ip.txt")
	defer output.Close()

	for i := 0; i < len(lines); i++ {
		This(func() {
			res, _ := client.Get("https://sonar.omnisint.io/reverse/" + lines[i])
			response, _ := ioutil.ReadAll(res.Body)
			data := string(response)
			data_split := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(data, "[", ""), "]", ""), `"`, "")
			data_chunk := strings.Split(data_split, ",")
			for x := 0; x < len(data_chunk); x++ {
				output.WriteString(data_chunk[x] + "\n")
			}
			fmt.Println(string(colorCyan)+"["+string(colorGreen)+"*"+string(colorCyan)+"]"+string(colorReset), lines[i], "~>", string(colorGreen), len(data_chunk), string(colorReset)+"Domain")
		}).Catch(func(E) {
			return
		})
	}
	fmt.Println(string(colorCyan) + "[" + string(colorGreen) + "*" + string(colorCyan) + "] " + string(colorReset) + "Done")
}

func DomainToIP() {
	banner()
	var domainList string

	fmt.Print(string(colorCyan) + "[" + string(colorGreen) + "?" + string(colorCyan) + "] " + string(colorReset) + "Domain List : ")
	fmt.Scanf("%s", &domainList)

	byteReadDomain, _ := ioutil.ReadFile(domainList)
	file_content_domain := string(byteReadDomain)
	line_domain := strings.Split(file_content_domain, "\n")
	output_domain, _ := os.Create("result/dom-to-ip.txt")
	defer output_domain.Close()

	for y := 0; y < len(line_domain); y++ {
		This(func() {
			ips, _ := net.LookupIP(line_domain[y])
			for _, ip := range ips {
				if ipv4 := ip.To4(); ipv4 != nil {
					fmt.Println(string(colorCyan)+"["+string(colorGreen)+"*"+string(colorCyan)+"]"+string(colorReset), line_domain[y], "~>", ipv4)
					output_domain.WriteString(ipv4.String() + "\n")
				}
			}
		}).Catch(func(err E) {
			return
		})
	}
	fmt.Println(string(colorCyan) + "[" + string(colorGreen) + "?" + string(colorCyan) + "] " + string(colorReset) + "Done")
}

func main() {
	var pilih string
	banner()
	menu := string(colorCyan) + `
[` + string(colorGreen) + `1` + string(colorCyan) + `] ` + string(colorReset) + `Domain To IP
` + string(colorCyan) + `[` + string(colorGreen) + `2` + string(colorCyan) + `] ` + string(colorReset) + `Rev IP`
	fmt.Println(menu)
	fmt.Println()
	fmt.Print(string(colorCyan) + "[" + string(colorGreen) + "?" + string(colorCyan) + "]" + string(colorReset) + " Choice : ")
	fmt.Scanf("%s", &pilih)

	switch pilih {
	case "1":
		DomainToIP()
	case "2":
		ReverseIP()
	default:
		os.Exit(0)
	}
}
