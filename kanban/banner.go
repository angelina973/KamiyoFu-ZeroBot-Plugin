package kanban

import (
	"fmt"
	"strings"

	"github.com/fumiama/go-registry"
)

var (
	info = [...]string{
		"* OneBot + ZeroBot + Golang + Angelina973mod",
		"* Version 1.5.0-b5.rc1 - 2022-07-28 12:37:17 +0800 CST",
		"* Copyright © 2020 - 2022 FloatTech. All Rights Reserved.",
		"* Project: https://github.com/FloatTech/ZeroBot-Plugin",
	}
	// Banner ...
	Banner = strings.Join(info[:], "\n")
	reg    = registry.NewRegReader("reilia.fumiama.top:32664", "fumiama")
)

// PrintBanner ...
func PrintBanner() {
	fmt.Print(
		"\n============================================",
		"\n", Banner, "\n",
		"--------------------------------------------",
		"\n", Kanban(), "\n",
		"============================================================\n\n",
	)
}

// Kanban ...
func Kanban() string {
	err := reg.Connect()
	if err != nil {
		return err.Error()
	}
	defer reg.Close()
	text, err := reg.Get("ZeroBot-Plugin/kanban")
	if err != nil {
		return err.Error()
	}
	return text
}
