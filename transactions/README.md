# exp





## swagger generate:
`mkdir gen`

`swagger generate server --exclude-main -A exp -t gen -f ./api.yml`

## run:
`go run cmd/main.go`

## cleanup for rerun:
`rm -rf gen`

`mkdir gen`


## test:
### tx:
```
`[GET] http://127.0.0.1:8080/tx?source=bofa-chk&month=201802`
`[DELETE] http://127.0.0.1:8080/tx?source=bofa-chk&month=201802`
`[GET] http://127.0.0.1:8080/tx/2342`
```
### bofa_chk:
```
`[POST] http://localhost:8080/tx/bofa-chk/` form-data:{csvfile, <csvfile>}`
`[PUT] http://localhost:8080/tx/bofa-chk/?lastId=123&source=bofa_chk` form-data:{upfile, <csvfile>}
```