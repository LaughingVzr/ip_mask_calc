package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var inputReader *bufio.Reader
var err error
var ipInterval string // IP地址段输入
var ip string         // IP地址
var mask int          // 掩码位
var parts []string

/*
   常量
*/
const (
	IPBinLen     = 32 // IP二进制化后的长度
	IPPartBinLen = 8  // IP每段的二进制固定长度
	MaskDefault  = "255"
)

func main() {
	inputReader = bufio.NewReader(os.Stdin)        // 初始化读入流
	ipInterval, err = inputReader.ReadString('\n') // 以换行符为结束标识
	if err != nil {
		fmt.Print(err)
	}
	IPMaskSplit()               // IP与掩码分割
	calcMaskIP(IPBinLen - mask) // 计算掩码IP
}

/*
IPMaskSplit IP与掩码分割
*/
func IPMaskSplit() {
	ipInterval = strings.Trim(ipInterval, "\r\n") // 去除空格与换行符
	parts = strings.Split(ipInterval, "/")        // 分割
	for index, part := range parts {
		switch index {
		case 0:
			ip = part
		case 1:
			mask, err = strconv.Atoi(part)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

/*
calcMaskIP 根据掩码地址
*/
func calcMaskIP(changeLen int) {
	var finalMaskIP string
	/*
	 * 计算掩码变更位所在的IP区
	 */
	var index = mask / IPPartBinLen
	if mask%IPPartBinLen != 0 {
		index++
	}

	fmt.Println(index)

	/*
	 * 计算填充位的掩码值
	 */
	length := changeLen
	var maskVal = 0
	for i := IPPartBinLen - 1; length > 0; i-- {
		maskVal += (1 << uint(i))
		length--
	}
	fmt.Printf("掩码值为：%d\n", maskVal)
	for i := 0; i < index-1; i++ {
		finalMaskIP += MaskDefault
		finalMaskIP += "."
	}
	finalMaskIP += strconv.Itoa(maskVal)
	fmt.Printf("掩码地址为：%s\n", finalMaskIP)
}
