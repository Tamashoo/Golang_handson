#  概要

* Go言語ハンズオンの学習記録

## 第１章

### Go言語のスタート

* 環境構築
* hello.go

### Go言語の基本

* 基本文法
  * 演算子の前後にはスペースを入れない
* 自作パッケージ

## 第２章

### 値と変数

* complex64, complex128：複素数
* 変数:=値で変数宣言と初期化を同時に行うことができる
* 全ての型変換は明示的に行う必要がある
* 全ての定義した変数は使わないとコンパイルエラーになる
* nil:値が存在しないことを表す

### 制御構文

* switchのfallthroughは次のcaseを実行する
* gotoは変数の宣言を飛び越せない

### 複雑な値

* スライスは参照型

### 関数

* パッケージとして外から呼び出す場合は関数名の先頭を大文字にする
* 返り値の変数をあらかじめ定義しておくこともできる

## 第３章

### ポインタ

* ポインタは計算できない

### 構造体

* new()で構造体を生成するとポインタが返る
* メソッドは特定の型に紐づいた関数
  * func (変数名 型) 関数名() 返り値の型 {}
    * 例：func (c Circle) Area() float64 {}
  * ex) circle.Area()みたいな使い方
* ポインタからのメソッド呼び出しは自動的に参照に変換される

### インターフェース

* 構造体に必要なメソッドを定義
* インターフェースは直接インスタンス化できない
* インターフェース型として変数を用意するときは、その値はインターフェースに用意された機能しか使えない
* 空のインターフェースはどんな型でも代入できる
* 変数 := reflect.TypeOf(値)の戻り値はreflect.Type型の値
  * さらに.Kind()とすると型を取得できる
  * reflect.Array, reflect.Slice, reflect.Map Of(Type)で配列、スライス、マップのType型と比較できる
    * Kindで比較できない場合

### 並行処理

* go func(){}()で並行処理を行う
* chanは他のスレッドに値を渡すことができる
  * make(chan 型)で宣言
  * chan <- 値で値を渡す
  * 変数 := <- chanで値を受け取る
  * チャンネルは値を受け取るまで待つ
    * mainから送る時はgo func()の後にc <- iを書く
  * 先入れ先出し
* selectで複数のチャンネルを待ち受けることができる
* mutexで排他制御を行う
  * sync.Mutexを使う
  * Lock()でロック
  * Unlock()でロック解除

## 第４章 Fyne

* ./chapter4にgo modをやった

### 基本的な使い方

## 第５章

### ネットワークアクセス

* net/httpのhttp.Get()でGETリクエストを送信できる
  * メインの内容は.Bodyで取得できる

#### goquery

* 最初にHTMLからDocument構造体を作成する
  * goquery.NewDocument(address), qoquery.NewDocumentFromResponse(Response), 
  goquery.NewDocumentFromReader(io.Reader)
* Find()でCSSセレクタを指定して要素を取得する
  * Document.Find(selector)で取得、Selectionという構造体が返る
    * Selection.Text()でテキストを取得できる
    * Selection.Attr()で属性を取得できる
    * Selection.Each()で繰り返し処理ができる
      * func(n int, sel *goquery.Selection)でnは繰り返し回数、selはSelection
* aタグはurl、href属性はリンク先のurlを指定する

#### encoding/json

* json.Unmarshal()でJSONを構造体に変換できる
  * 第一引数にJSONのbyte配列、第二引数に構造体のポインタを指定する
* jsonの各値はキーと値のペアで表される
  * m := im.(map[string]interface{})で型アサーションを行う

### データベースアクセス

* database/sqlを使う
  * sql.Open()でデータベースに接続する
    * 第一引数にドライバ名、第二引数にデータベースの情報を指定する
  * db.Query()でSQLを実行する
  * Rows.Next()で次の行があるか確認する
  * Rows.Scan()で行の値を取得する
  * 1行だけ取得する場合はQueryRow()を使う
  * db.Exec()でSQLを実行する
* ?はプレースホルダー
  * 後から値を埋め込むことができる
* %はワイルドカード
  * x%はxで始まる文字列
  * %xはxで終わる文字列
  * %x%はxを含む文字列

## 第６章

### サーバプログラムの基本

* net/httpのhttp.ListenAndServe()でサーバを起動できる
  * 第一引数にポート番号、第二引数にハンドラを指定する
* http.HandleFunc()でハンドラを登録する
  * 第一引数にパス、第二引数にハンドラを指定する

### テンプレートエンジン

* html/templateを使う
  * template.ParseFiles()でテンプレートを読み込む
  * template.Execute()でテンプレートを実行する
    * 第一引数に出力先、第二引数にテンプレートに渡す値を指定する
  * テンプレートには変数を埋め込むことができる
    * {{.変数名}}で埋め込む
  * テンプレートにはif文を埋め込むことができる
    * {{if 条件}} 条件がtrueの時に表示される {{end}}
  * テンプレートにはfor文を埋め込むことができる
    * {{range 配列}} 配列の要素が順番に表示される {{end}}
  * テンプレートでは値の有無を判定することができる
    * {{with 値}} 値がある時に表示される {{end}}
  * テンプレートには関数を埋め込むことができる
    * {{関数名 引数}}で埋め込む
    * template.FuncMapで関数を登録する
      * template.FuncMap{"関数名": 関数}
    * template.New("テンプレート名").Funcs(関数マップ).ParseFiles("テンプレートファイル名")で関数を登録しながらテンプレートを読み込む
  * テンプレートにはテンプレートを埋め込むことができる
    * {{template "テンプレート名" .}}で埋め込む
    * template.ParseFiles()でテンプレートを読み込む
    * template.ExecuteTemplate()でテンプレートを実行する
      * 第一引数に出力先、第二引数にテンプレート名、第三引数にテンプレートに渡す値を指定する

### 色々な機能

* パラメータの取得
  * http.Request.FormValue()でパラメータを取得できる
    * 第一引数にパラメータ名を指定する
