## How to use go
- VSCode に remote-container の extention を install
- VSCode の左下の緑の部分をクリック
- Reopen in container を選択
- コンテナが立ち上がり、接続された状態でVSCodeが開く


## localhostから コンテナ上のmysqlの特定のDBにアクセス
sampledb に port3306でアクセス
mysql -h localhost -P 3306 -u docker sampledb  -p --protocol=tcp

## docker compose で起動しているgoから mysqlにアクセス
```go
dbConn := fmt.Sprintf("%s:%s@tcp(db-for-go)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
db, err := sql.Open("mysql", dbConn)
```
