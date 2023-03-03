# 復習用
- [リンクはここ](https://go-tour-jp.appspot.com/list)

# Using the tour (3/3)
- formatter として，go format がある

# Basics (3/3)
## Packages, variables, and functions
- デフォルトで main package がエントリーポイントになる
- package 名についている`math/rand`では rand（最後の要素）がついているものになる（つまりpackage rand で始まるファイルということ？）
- factored import statement(`import ( "fmt" "math" )`みたいに括弧でまとめてインポートする書き方を推奨)
- 大文字で始まるモノ以外は外部参照不可能（`math.Pi`で円周率にアクセス可能）
- 関数は`func 関数名(x 型, y 型) 返り値の型{ }`の形式で書く（**型が後ろであることに注意**）
  - もし型が同じならまとめて最後に書くことも可能（`x, y int`みたいな）
  - 複数の戻り値もok（`func swap(x, y string) (string, string){return y, x}`で交換が可能）
  - **naked return**が可能（先に返り値の変数と型を書いて，最初に関数内部でアクセスした値をその値だと認識して返す方法）（ただし長い関数では可読性が悪くなるので使用しないこと）
- **変数宣言は var を使用（`var c bool`, `var i int`みたいに型は最後に指定）**
    - 初期化子を与えることも可能（`var i, j int = 1, 2`, `var c = true`などの型省略も可能となる）
    - **関数の中では**`:=`で`var`を使わない宣言もできる
- 組み込みの型は，bool, string, int(int8 ~ 64), uint(uint8 ~ 64, uintptr), byte, rune, float32, float64, complex64, complex128
- **初期値を与えず初期化すると，ゼロ値（0, false, ""）が与えられる**
- **型変換は明示的に行う必要がある**（`i := 42, f := float64(i), u:=uint(f)`）
- 初期値を与えた場合，型は省略可能だがその際の型推論は（右辺の値に基づく（元々別で与えられた型がある場合はそれに従う））
- 定数宣言が可能で，char, string, bool, numeric のみで使用可能（`const World = "世界"`，`:=`で宣言は不可能．）
    - 数値が高精度(`1<<100`も保持できる？)


## Flow control statements: for, if, else, switch and defer (3/3)
### for
- for 文は基本 C, JS と似ているが`（）`でくくる必要がない点が異なる
    - 条件式さえあれば，初期化と後処理は任意（`sum := 1` `for ; sum < 1000; { sum += sum}`みたいな）
    - **セミコロンを省略することで C における while を再現可能（`for sum < 1000 { sum += sum }`）**
    - **条件の省略で無限ループが書ける（`for{ }`）**

### if
- if も`()`がいらないだけ
  - **if 条件の前に簡単な記述が可能でこれを if スコープ内でのみ使用可能**（`if v:=math.Pow(x, n); v < lim { return v }`みたいな）
    - もちろん else の中でもこの宣言された値を用いることが可能
- 平方根を求める関数を Newton 法で実装する（面白い）
  
### switch
- **switch 文は他の swtich と似ているが，選択された case 以外を実行しないという点（実行したら break してくれる），case が定数でなくてもよい点で異なる（もちろん if と同じく短い記述を条件の前に書くことができる）**
    - 条件のない`switch {}`という記述により，if-then-else を簡潔に表現可能
- **`defer`をつけると呼び出し元の関数が return するまで遅延させるもの**
  - 複数`defer`すると，**stack されるため，LIFO で実行される**


## More types: structs, slices, and maps
### Struct
- **Go ではポインタを扱うことができ，`*T`でポインタ型を表し，ゼロ値はnil**
  - **Ｃ言語と異なり，ポインタ演算は存在しない**（`i := 42, p := &i`で p が i に対するポインタとなる）
- 構造体ももちろん宣言可能（`type Vertex struct { X int Y int}`）
  - `.`を用いてアクセスが可能（`v:= Vertex{1, 2} v.X = 4`）
  - **構造体ポインタを通してのアクセスも可能だが，正しい文法でアクセスするのは面倒なので簡単な記法が用意されている（`v := Vertex{1, 2} p := &v`とすると，本来は(*p).Xとアクセスする必要があるが，p.Xでアクセスしてもよいことになっている）**
  - **struct リテラルという方法で，初期値を割り当てられる**（`Vertex{X: 1}`とすると，`{1, 0}`と明示的でない Y は0になる．また`p = &Vertex{1, 2}`とすることでポインタが宣言可能）

### array (slice)
- `[n]T`型として宣言することで配列を宣言できる（長さまで含めて型として扱われる点に注意，例えば`var a [10]int`）
  - `[]T`とすることでスライスを宣言できる（**可変長配列であり，半開区間**）
  - **実はスライスはデータを格納しておらず，単なる配列の参照なのでデータを書き換えると元の配列も変わってしまう**
  - スライスの記法は python と同様っぽい
  - **スライスは「現在の」配列の長さを表す len と元となる配列の長さを表す cap をもつ（一度短くした配列も再拡張可能な点が面白い）**
  - **スライスのゼロ値は nil で len, cap ともに 0 で元となる配列をまったくもたない**
  - **makeを使用することでスライスを作成可能（`make([]int, 5)`, `make([]int, 0, 5)`みたいに，スライスの型，len, cap の順に指定する）**
  - 配列への追加は`append(s, 追加要素, ...)`を用いる．返ってくる値は追加後のスライスで**元の配列より cap が大きい配列になる場合は新たに配列を割り当て直す（つまり cap が変更される）** 
  - **range はスライスやマップを１つずつ処理するもので，**`for i, v := range 配列名 { // i or v でアクセス}`のように書く？（必要のない値はアンダースコアで無視できる）
- ２次元配列画像を表示する exercise

### maps
- 書式は`map[keyの型]valueの型{}`みたいにすればいい（**make を使うことでも作成できる**）（`m := make(map[string]int)`みたいに）
- **初期値は nil で配列と同じ**
- もしマップに渡す型が単純な型なら型名を省略して単純なリテラルとしてかける
    - 代入，更新(Create, Update)は`m[キー] = 値`
    - 取得(Retrieve)は`値=m[キー]`
    - 削除(Delete)は`delete(m, キー)`
    - **存在判定は`elem, ok = m[キー]`でキーが存在すれば ok が true にり，そうでなければ false となる**（もし elem, ok が未宣言なら`:=`を使う）

### function values
- 関数値として引数，戻り値として渡すことが可能（コールバック関数みたいな機能）
- Go の関数はクロージャ（fibonacci を closure で実現するエクササイズ，クロージャをちゃんと理解してないと中々厳しそう）


# Methods and interfaces (3/4)

# Concurrency (3/5)



# おまけ
- [Go言語のPrint・Println・Printfについて](https://www.flyenginer.com/low/go/go%E8%A8%80%E8%AA%9E%E3%81%AEprint%E3%83%BBprintln%E3%83%BBprintf%E3%81%AB%E3%81%A4%E3%81%84%E3%81%A6.html)
- [【Go】strings パッケージ関数まとめ](https://zenn.dev/kou_pg_0131/articles/go-strings-functions)