package sox

import (
	"fmt"
	"os"
	"os/exec"	
)

func init() {
	if _, err := exec.LookPath("play"); err != nil {
		info :=
`没有找到play命令，发音将无法使用
ubuntu: sudo apt-get install sox
centos: sudo yum install sox
`
		fmt.Fprintln(os.Stderr, info)
	}

}

func Play(addr string) {
	cmd := exec.Command("play", addr)
	if err := cmd.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
