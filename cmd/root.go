package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode/utf8"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"

	"github.com/newton-miku/nali/internal/constant"
	"github.com/newton-miku/nali/pkg/common"
	"github.com/newton-miku/nali/pkg/entity"
)

func Isutf8(s string)bool{
	return utf8.ValidString(s)
}

func isGBK(s string) bool {
	//先检查是否是utf8字符
	if Isutf8(s){
		return false
	}
	data:=[]byte(s)
	length := len(data)
	var i int = 0
	for i < length {
		//fmt.Printf("for %x\n", data[i])
		if data[i] <= 0xff {
			//编码小于等于127,只有一个字节的编码，兼容ASCII吗
			i++
			continue
		} else {
			//大于127的使用双字节编码
			if 	data[i] >= 0x81 &&
				data[i] <= 0xfe &&
				data[i + 1] >= 0x40 &&
				data[i + 1] <= 0xfe &&
				data[i + 1] != 0xf7 {
				i += 2
				continue
			} else {
				return false
			}
		}
	}
	return true
}

var rootCmd = &cobra.Command{
	Use:   "nali",
	Short: "An offline tool for querying IP geographic information",
	Long: `An offline tool for querying IP geographic information.

Find document on: https://github.com/newton-miku/nali

#1 Query a simple IP address

	$ nali 1.2.3.4

  or use pipe

	$ echo IP 6.6.6.6 | nali

#2 Query multiple IP addresses

	$ nali 1.2.3.4 4.3.2.1 123.23.3.0

#3 Interactive query

	$ nali
	123.23.23.23
	123.23.23.23 [越南 越南邮电集团公司]
	quit

#4 Use with dig

	$ dig nali.zu1k.com +short | nali

#5 Use with nslookup

	$ nslookup nali.zu1k.com 8.8.8.8 | nali

#6 Use with any other program

	bash abc.sh | nali

#7 IPV6 support
`,
	Version: constant.Version,
	Args:    cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		// gbk, _ := cmd.Flags().GetBool("gbk")
		isJson, _ := cmd.Flags().GetBool("json")

		if len(args) == 0 {
			stdin := bufio.NewScanner(os.Stdin)
			stdin.Split(common.ScanLines)
			for stdin.Scan() {
				line := stdin.Text()
				if isGBK(line) {
					line, _, _ = transform.String(simplifiedchinese.GBK.NewDecoder(), line)
				}
				if line := strings.TrimSpace(line); line == "quit" || line == "exit" {
					return
				}
				if isJson {
					_, _ = fmt.Fprintf(color.Output, "%s", entity.ParseLine(line).Json())
				} else {
					_, _ = fmt.Fprintf(color.Output, "%s", entity.ParseLine(line).ColorString())
				}
			}
		} else {
			if isJson {
				_, _ = fmt.Fprintf(color.Output, "%s", entity.ParseLine(strings.Join(args, " ")).Json())
			} else {
				for _, line := range args {
					_, _ = fmt.Fprintf(color.Output, "%s\n", entity.ParseLine(line).ColorString())
				}
			}
		}
	},
}

// Execute parse subcommand and run
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err.Error())
	}
}

func init() {
	rootCmd.Flags().Bool("gbk", false, "Use GBK decoder")
	rootCmd.Flags().BoolP("json", "j", false, "Output in JSON format")
}
