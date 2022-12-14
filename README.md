


## 產生 proto file
`
cd project root path
protoc --go_out=plugins=grpc:. ./proto/*.proto
`

--go_out: 設定所產生的ＧＯ程式輸出的目錄。該指令會載入 protoc-gen-go外掛程式，已達到產生ＧＯ
          程式的目的。產生的檔案以 .pb.go為檔案副檔名，這裡的 “：”有分隔做用，後跟指令所需的
          參數集，這表示把產生的ＧＯ程式輸出到指向的 protoc 編譯的目前目錄

plugins=plugin1+plugin2: 指定要載入的子外掛程式列表。我們定義的 proto 檔案是有關了 RPC 服
                         務，而預設是不會產生 RPC 程式，因此需要在 go_out 列出 plugins
                         參數，將其傳遞給 protoc-gen-go 外掛程式，及告訴編譯器，請支援 ＲＰＣ
                          