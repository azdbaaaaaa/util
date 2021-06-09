package ximagick

import (
	"github.com/azdbaaaaaa/util/log"
	"os/exec"
)

func Convert(src, dst string) (err error) {
	_, err = exec.LookPath("convert")
	if err != nil {
		log.Errorw("Convert.LookPath", "err", err, "src", src, "dst", dst)
		return
	}
	//cmd中执行的命令：ffmpeg -i psg.flv test.mp4
	cmd := exec.Command("convert", src, dst)

	//阻塞至等待命令执行完成
	err = cmd.Run()
	if err != nil {
		log.Errorw("Convert.Run", "err", err, "src", src, "dst", dst)
	}
	return
}
