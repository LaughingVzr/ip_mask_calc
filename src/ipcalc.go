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

/*
   常量
*/
const (
	IPBinLen     = 32    // IP二进制化后的长度
	IPPartBinLen = 8     // IP每段的二进制固定长度
	MaskDefault  = "255" // 高位IP默认值
	MaskMin      = "0"   // 低位IP默认值
	IPSep        = "."   //IP分隔符
	IPPartCount  = 4     // IP部分数(192.168.1.1,默认4段)
	ORVal        = 0xff  // 掩码计算异或位
)

func main() {
START:
	fmt.Print("请输入掩码位(0~32)：")
	inputReader = bufio.NewReader(os.Stdin)        // 初始化读入流
	ipInterval, err = inputReader.ReadString('\n') // 以换行符为结束标识
	if err != nil {
		fmt.Print(err)
	}
	if flag, _ := IPMaskBitCheck(); flag {
		calcMaskIP() // 计算掩码IP
	} else {

		goto START
	}

}

/*
IPMaskBitCheck IP与掩码分割
*/
func IPMaskBitCheck() (bool, error) {
	ipInterval = strings.Trim(ipInterval, "\r\n") // 去除空格与换行符
	mask, err = strconv.Atoi(ipInterval)
	if mask < 0 || mask > 32 {
		fmt.Println("掩码位范围错误(0~32)")
		return false, err
	}
	if err != nil {
		fmt.Println("请输入整数数字！")
		fmt.Println()
		return false, err
	}
	return true, nil
}

/*
calcMaskIP 根据掩码地址
*/
func calcMaskIP() {
	var finalMaskIP string              // 掩码IP
	var bitpat = 0xff00                 // 掩码位运算初始值
	var maskVal = 0                     // 掩码值
	var index = mask/IPPartBinLen + 1   // 掩码值所在IP段索引
	var changeBit = mask % IPPartBinLen // 掩码二进制所需要修改的位数

	/*
	 * 计算填充位的掩码值
	 */
	bitpat = bitpat >> uint(changeBit)
	maskVal = bitpat & ORVal
	// fmt.Printf("掩码值为：%d\n", maskVal)
	for i := 1; i <= IPPartCount; i++ {
		if i < index {
			finalMaskIP += MaskDefault
		}
		if i == index {
			finalMaskIP += strconv.Itoa(maskVal)
		}
		if i > index {
			finalMaskIP += MaskMin
		}
		if IPPartCount-i >= 1 {
			finalMaskIP += IPSep
		}
	}
	fmt.Printf("\n掩码地址为：%s\n", finalMaskIP)
}
