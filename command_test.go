package bobra

import "testing"

var cmd = &Command{
	Use: "test",
	Short: "test",
	Long: "test",
	Example: "test",
}

var root = &Command{
	Use: "root",
	Short: "root",
	Long: "root test",
	Example: "root test",
}

// 测试args到flags的转换
func TestCommand_ParseFlags(t *testing.T) {
	cmd.Flags().StringP("aaaa", "a", "YOUR NAME", "author name for copyright attribution")
	cmd.Flags().StringP("ddd", "d", "YOUR NAME", "author name for copyright attribution")
	cmd.Flags().StringP("c", "c", "YOUR NAME", "author name for copyright attribution")

	args := []string{"-a123", "-b 10"}
	cmd.ParseFlags(args)
	r1,_ := cmd.Flags().GetString("aaaa")
	r2 ,_:= cmd.Flags().GetString("ddd")
	e1,e2 := "123", "YOUR NAME"
	if r1 != e1 || r2 != e2 {
		t.Errorf("expected '%s', '%s' but got '%s', '%s'", e1, e2, r1, r2)
	}
}

// 测试全局flags能否在根命令中也访问到
func TestCommand_GlobalFlags(t *testing.T) {
	root.AddCommand(cmd)
	cmd.GlobalFlags().StringP("global", "g", "default", "kakakak")
	args := []string{"-ghahaha"}
	cmd.ParseFlags(args)
	r, _ := root.GlobalFlags().GetString("global")
	expected := "hahaha"
	if r != expected {
		t.Errorf("expected '%s', but got '%s'", expected,r)
	}
}

// 测试局部flags只能在当前子命令被访问
func TestCommand_LocalFlags(t *testing.T) {
	root.AddCommand(cmd)
	cmd.LocalFlags().StringP("local", "l", "default", "kakakak")
	args := []string{"-ltest"}
	cmd.ParseFlags(args)
	r1, _ := root.GlobalFlags().GetString("local")
	e1 := ""
	r2, _ := cmd.LocalFlags().GetString("local")
	e2 := "test"
	if r1 != e1 || r2 != e2 {
		t.Errorf("expected '%s', '%s' but got '%s', '%s'", e1, e2, r1, r2)
	}
}

// 测试全部flags的获取
func TestCommand_Flags(t *testing.T) {
	root.AddCommand(cmd)
	cmd.LocalFlags().StringP("local2", "t", "default", "kakakak")
	root.GlobalFlags().StringP("fff", "f", "YOUR NAME", "author name for copyright attribution")
	args := []string{"-ttestl", "-ftestg"}
	cmd.ParseFlags(args)
	root.ParseFlags(args)
	r1, _ := cmd.Flags().GetString("local2")
	e1 := "testl"
	r2, _ := cmd.Flags().GetString("fff")
	e2 := "testg"
	if r1 != e1 || r2 != e2 {
		t.Errorf("expected '%s', '%s' but got '%s', '%s'", e1, e2, r1, r2)
	}
}

// 测试命令执行路径的获取
func TestCommand_CommandPath(t *testing.T) {
	root.AddCommand(cmd)
	r := cmd.CommandPath()
	expected := "root test"
	if r != expected {
		t.Errorf("expected '%s', but got '%s'", expected,r)
	}
}

