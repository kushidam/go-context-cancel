# go-context-cancel


`process`関数をゴールーチン内で非同期に実行

ゴールーチンをキャンセルするために、`cancel `関数を呼び出す。


`process `関数: 以下の処理行う。

無限ループ内で`select`ステートメントを使用して、コンテキストの`ctx.Done()`チャネルを監視する。
これにより、コンテキストがキャンセルされたかどうかを確認する。

* `ctx.Done()`チャネルが閉じられた場合、終了時の処理を行い、ゴールーチンは終了する

* `ctx.Done()`チャネルが閉じられていない場合、デフォルトのケースで 何らかの処理を行う。


メインゴルーチンが `cancel`関数を使用してコンテキストをキャンセルすることにより、ゴールーチン内の処理を制御。
これによりキャンセルが発生した場合に正しく終了するような処理が設計ができる
