＃環境構築  
go mod init go-gin-productManagerPro

- gin（公式 Doc： https://gin-gonic.com/ja/docs/quickstart/ ）  
  go get -u github.com/gin-gonic/gin  
  curl localhost:8080/ping // 公式の検証  

- ホットリロードの設定（公式 Doc： https://github.com/cosmtrek/air ）  
  実装を変更するたびにリロードしなくて良い  
  go install github.com/cosmtrek/air@latest  
  air init
  air  
  参考：https://zenn.dev/urakawa_jinsei/articles/a5a222f67a4fac
- うまくいかない場合は README3 を参考
