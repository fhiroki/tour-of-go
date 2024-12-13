# Exercise: Equivalent Binary Trees

1. Walk 関数を実装してください。

2. Walk 関数をテストしてください。

関数 tree.New(k) は、値( k, 2k, 3k, ..., 10k )をもつ、ランダムに構造化 (しかし常にソートされています) した二分木を生成します。

新しいチャネル ch を生成し、 Walk を動かしましょう:

go Walk(tree.New(1), ch)
そして、そのチャネルから値を読み出し、10個の値を表示してください。 それは、 1, 2, 3, ..., 10 という表示になるでしょう。

3. Same 関数を実装してください。 t1 と t2 が同じ値を保存しているどうかを判断するため、 Walk を使ってください。

4. Same 関数をテストしてください。

Same(tree.New(1), tree.New(1)) は、 true を返すように、 Same(tree.New(1), tree.New(2)) は、 false を返すようにします。

Tree のドキュメントは [こちら](https://pkg.go.dev/golang.org/x/tour/tree#Tree) です。
