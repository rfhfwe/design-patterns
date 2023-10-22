package main

var (
	tv         *Tv
	onCommand  *OnCommand
	offCommand *OffCommand
	onButton   *Button
	offButton  *Button
)

func Init() {
	tv = &Tv{}
	onCommand = &OnCommand{
		device: tv,
	}
	offCommand = &OffCommand{
		device: tv,
	}
	onButton = &Button{
		command: onCommand,
	}
	offButton = &Button{
		command: offCommand,
	}
}

func main() {
	Init() // 初始化
	onButton.press()
	offButton.press()
}
