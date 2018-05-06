package sox

import (
	"fmt"
	"os/exec"	
)

func init() {
	cmd := exec.Command("which", "play")
	if err := cmd.Run(); err != nil {
		fmt.Println("没有找到play命令，发音将无法使用")
		fmt.Println("ubuntu: sudo apt-get install sox")
		fmt.Println("centos: sudo yum install sox")
		fmt.Println()
	}
}

func Play(addr string) {
	cmd := exec.Command("play", addr)
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
	}
}
