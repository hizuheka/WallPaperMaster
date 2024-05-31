package main

import (
	"fmt"
	"os"
	"syscall"
	"unsafe"
)

// 定数の定義
const (
	SPI_SETDESKWALLPAPER = 0x0014
	SPIF_UPDATEINIFILE   = 0x01
	SPIF_SENDCHANGE      = 0x02
)

func setWallpaper(filePath string) error {
	// UTF-16に変換
	ptr, err := syscall.UTF16PtrFromString(filePath)
	if err != nil {
		return err
	}

	// SystemParametersInfo関数の呼び出し
	ret, _, err := syscall.NewLazyDLL("user32.dll").NewProc("SystemParametersInfoW").Call(
		SPI_SETDESKWALLPAPER,
		0,
		uintptr(unsafe.Pointer(ptr)),
		SPIF_UPDATEINIFILE|SPIF_SENDCHANGE,
	)

	// 成功したかどうかの確認
	if ret == 0 {
		return fmt.Errorf("failed to set wallpaper: %v", err)
	}

	return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: WallPaperMaster <image-path>")
		return
	}

	imagePath := os.Args[1]
	if err := setWallpaper(imagePath); err != nil {
		fmt.Printf("Error setting wallpaper: %v\n", err)
	} else {
		fmt.Println("Wallpaper set successfully")
	}
}
