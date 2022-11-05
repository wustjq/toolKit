package tools

import "strings"

type AsciiType int32

//工具类结构体
type ToolKit struct{
	name string
}

//提供初始化的方法
func NewToolKit(Name string)*ToolKit{
	return &ToolKit{
		name:Name,
	}
}

//字符串转大写
func (t*ToolKit)StringToUpper(str string)string{
	return strings.ToUpper(str)
}

//字符串转小写
func (t*ToolKit)StringToLower(str string)string{
	return strings.ToLower(str)
}

//下划线转大驼峰  eg:LastTime HelloWorld
func (t*ToolKit)UnderlineToUpperCamelCase(str string)string{
	s:=strings.Replace(str,"_"," ",-1)
	s=strings.Title(s)
	return strings.Replace(s," ","",-1)
}

//下划线转小驼峰  eg:lastTime helloWorld
func (t*ToolKit)UnderlineToLowerCamelCase(str string)string{
	s:=t.UnderlineToUpperCamelCase(str)
	return t.StringToLower(string(s[0]))+s[1:]
}

//字符求ascii
func (t *ToolKit)ByteToAscii(ch byte)AsciiType{
	return AsciiType(ch)
}

//ascii求字符
func (t *ToolKit)AsciiToByte(asciibyte AsciiType)byte{
	return byte(asciibyte)
}

