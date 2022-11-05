package tools

import (
	"testing"
)

func initkit(Name string)*ToolKit{
	return &ToolKit{
		name:Name,
	}
}

//单测
//got:=kit.StringToUpper("abKJKbssh")
//want:="ABKJKBSSH"
//if got!=want{
//	t.Errorf("StringToUpper method err,because [want: %v] [got: %v]",want,got)
//	return
//}


//字符串转大写
func TestToolKit_StringToUpper(t *testing.T) {
	kit:=initkit("jq")
	//用来测试多个
	type test struct{
		got string
		want string
	}
	tests:=map[string]test{
		"abKJKbssh":{kit.StringToUpper("abKJKbssh"),"ABKJKBSSH"},
		"savavDFYU":{kit.StringToUpper("savavDFYU"),"SAVAVDFYU"},
		"abKsbfesh":{kit.StringToUpper("abKsbfesh"),"ABKSBFESH"},
	}
	for _,v:=range tests{
		if v.want!=v.got{
			t.Errorf("StringToUpper Method err: [want:%s] [got:%s]",v.want,v.got)
			return
		}
	}
}

//字符串转小写
func TestToolKit_StringToLower(t *testing.T) {
	kit:=initkit("jq")
	type test struct{
		got string
		want string
	}
	tests:=map[string]test{
		"abKJKbssh":{kit.StringToLower("abKJKbssh"),"abkjkbssh"},
		"savavDFYU":{kit.StringToLower("savavDFYU"),"savavdfyu"},
		"abKsbfesh":{kit.StringToLower("abKsbfesh"),"abksbfesh"},
	}
	for _,v:=range tests{
		if v.want!=v.got{
			t.Errorf("StringToUpper Method err: [want:%s] [got:%s]",v.want,v.got)
			return
		}
	}
}

//下划线转小驼峰
func TestToolKit_UnderlineToLowerCamelCase(t *testing.T) {
	kit:=initkit("jq")
	type test struct{
		got string
		want string
	}
	tests:=map[string]test{
		"ab_K_JKb_ssh":{kit.UnderlineToLowerCamelCase("ab_K_JKb_ssh"),"abKJKbSsh"},
		"Sa_v_avDF_YU":{kit.UnderlineToLowerCamelCase("Sa_v_avDF_YU"),"saVAvDFYU"},
		"a_bK_sb_fesh":{kit.UnderlineToLowerCamelCase("a_bK_sb_fesh"),"aBKSbFesh"},
	}
	for _,v:=range tests{
		if v.want!=v.got{
			t.Errorf("StringToUpper Method err: [want:%s] [got:%s]",v.want,v.got)
			return
		}
	}
}

//下划线转大驼峰
func TestToolKit_UnderlineToUpperCamelCase(t *testing.T) {
	kit:=initkit("jq")
	type test struct{
		got string
		want string
	}
	tests:=map[string]test{
		"ab_K_JKb_ssh":{kit.UnderlineToUpperCamelCase(("ab_K_JKb_ssh")),"AbKJKbSsh"},
		"sa_v_avDF_YU":{kit.UnderlineToUpperCamelCase("sa_v_avDF_YU"),"SaVAvDFYU"},
		"a_bK_sb_fesh":{kit.UnderlineToUpperCamelCase("a_bK_sb_fesh"),"ABKSbFesh"},
	}
	for _,v:=range tests{
		if v.want!=v.got{
			t.Errorf("StringToUpper Method err: [want:%s] [got:%s]",v.want,v.got)
			return
		}
	}
}
