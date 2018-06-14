# go-mecab
dockerized go-mecab REST API Server  

IPA辞書をベースに[neologd](https://github.com/neologd/mecab-ipadic-neologd/blob/master/README.ja.md)と任意の単語でシステム辞書を作成します。  

## 単語の追加方法

seedディレクトリに[単語の追加方法](http://taku910.github.io/mecab/dic.html)のフォーマットで作成したCSVを格納してdocker buildを行います。

## SEE ALSO
- [go-mecab](https://github.com/shogo82148/go-mecab)
- [MeCab](http://taku910.github.io/mecab/)
- [MeCab システム辞書への単語追加（mecab-ipadic-neologd）](https://blog.apar.jp/linux/2796/)
