package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// キャンセル用のコンテキストを作成
	ctx, cancel := context.WithCancel(context.Background())

	// ゴールーチン内でコンテキストを使用
	go process(ctx)

	// 何らかの処理
	time.Sleep(3 * time.Second)

	// ゴールーチンをキャンセル
	cancel()

	// 後続の処理
	time.Sleep(1 * time.Second)
}

func process(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("ゴールーチンがキャンセルされました")
			return
		default:
			// 何らかの処理
			fmt.Println("ゴールーチン実行中...")
			time.Sleep(1 * time.Second)
		}
	}
}
