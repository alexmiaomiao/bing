package sox

import (
	"fmt"
	"os/exec"	
)

func init() {
	cmd := exec.Command("which", "play")
	if err := cmd.Run(); err != nil {
		str := "没有找到play命令，发音将无法使用\n"
		str += "ubuntu: sudo apt-get install sox\n"
		str += "centos: sudo yum install sox\n"
		fmt.Println(str)
	}
}

func Play(addr string) {
	cmd := exec.Command("play", addr)
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
	}
}
