# メモ

## 1. Getting started
- 前提として，module > package で，Go リポジトリには１つのモジュールが直下に格納されているイメージ
- `go.mod`ファイルは，**モジュールのパス（名前）とバージョン情報（go と package）を指定しておくファイル**（これにより，再現性が生まれ，どのバージョンのモジュールを管理するかを直接指示できる）
- 以降は，hello.go ファイルの内容
- `package`は同じディレクトリにいる全てのファイルから構成される関数をグループ化したもの
- `fmt` package は，テキストのフォーマットとか，コンソールへの出力とかを担う Go のインストール時についてくる std library の１つ
- `main package`を run するとデフォルトで，main 関数が走る
- `go run`コマンドで実行
- `go help`コマンドで，go のコマンドを列挙可能
### Call code in an external package
- `go mod tidy`を実行することで，import 対象のモジュールのインストールや`go.mod` / `go.sum`の更新を行うらしい（通常はいちいち明示的にインストール・アンインストールを行う必要があるはずなので便利）


## 2. Create a Go module
- `go mod init`でモジュールパスを与えることができる（実際に公開する際には，Go からアクセスできるようなパスにする必要がある）
- 自分が所属する package を先頭に宣言する（**必ず各ソースコードに１つ以上割り当てる**）
- **Go では大文字から始まる関数名にすると外部パッケージから呼び出せるようになる**
- `:=`で，**宣言**と**初期化**を同時に行える（通常は型も明示的に宣言するが，この方法を使うことで型推論をしてくれる）
- `example.com/hello`から`example.com/greetings`を使おうとすると，go からアクセスできるところに公開されていないので探しにいけない（今回は`example.com/hello`側を調整して，`example.com/greetings`を見つけられるようにする．具体的には，`go mod edit`コマンドでローカルディレクトリにリダイレクトさせる）
- `go mod edit --replace example.com/greetings=../greetings`とすることで，replace A=BのAがBに置き換わる（`go.mod`を見ても replace があることがわかる）
- `example.com/hello`モジュールの依存関係を同期するために，`go mod tidy`を実行する（`go.mod`に require が追加される）
- `go.mod`の`require`のパスの次に続いているのは疑似バージョンらしい（意味のあるバージョン番号の代わりに使用される）
- **公開されているモジュールを使用する場合は，`replace`がなく，`require [path] [tagged version number]`という形になるはず**
### Return and handle an error
- Go 関数は複数の値を返すことが可能
- standard library の`errors`パッケージを`errors.New`関数を使うために import する
- エラー出ないことを意味する`nil`を返すと 呼び出し側が関数が成功したと判定する
- standard library の`log`パッケージを使う（SetFlags(0)でタイムスタンプとか行数などの情報を出さないようにしている）
- `err`が返ってくれば，`log.Fatal`関数でエラーの値を出力して止まるようにする

## 3. Getting started with multi-module workspaces
- まず，workspace ディレクトリを作成し，直下に`golang.org/x/example` モジュールに依存する hello モジュールを作成する（hello ディレクトリに`go mod init example.com/hello` -> `go get golang.org/x/example`でダウンロード）
- workspace ディレクトリに戻り，`go work init ./hello`と叩くことで，go が go.work file 
- `go.work`は`go.mod`に似ている
  - go directive は Go に Go のどのバージョンとして解釈するべきかを伝える
  - use directive は GO に hello ディレクトリ内の module がビルド時に main modules になることを伝えている（つまり workspace 直下にいればどこであってもそのモジュールを run させられるということ）
- `go run`をモジュールや workspace 外で実行するとどのモジュールを使っていいかわからずエラーになる
- さっき`go get`で取ってきた`golang.org/x/example`を直接 clone して workspace に　module として追加してみる
- こうすることで，stringutil の中身を書き換えることができる（今回は reverse から UpperCase に書き換えた）
- **go コマンドは，`go.work`に明記されているディレクトリ内のモジュール名を確認しに行って解決している**
- `go.work`は全てのモジュールで`replace`命令を追記する代わりに使える
- 今回の内容を正式にリリースしようとしたら，まず書き換えた`golang.org/x/example`モジュールを v0.1.0 とかでリリースした後に，`hello/go.mod`にバージョン情報を更新して追記する必要がある
- `go work`には:
  - `go work use`（use 命令を`go.work`に追記したり，存在しない場合は，`go work`から除外したりできる）
  - `go work edit`（`go mod edit`と同様にファイル編集）
  - **`go work sync`は workspace の ビルドリストから workspace modules に依存関係を同期する**（？）