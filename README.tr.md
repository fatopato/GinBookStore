# GinBookStore
Golang ve Gin FrameWork ile oluşturulmuş basit bir REST API

Oluşturulan Book nesnesi ile örnek bir kitapçı kontrol edilebilmektedir.

Bu REST API sqlite3 ile oluşturuldu. Bu nedenle sunucu ilk kez ayaklandırıldığında 'data.db' adında bir dosya oluşturulacaktır.

İlgili modülleri indirmek için:
```
$ go get .
```
Sunucuyu ayaklandırmak için:
```
$ go run .
```
Sunucu 8080 portunda koşmaktadır. API http://localhost:8080/api/v1/ adresinden erişilebilir haldedir.

CRUD operasyonları API aracılığıyla gerçekleştirilebilir.

# Örnek İstekler

Tüm Kitapları Getir:
```
$ curl -i http://localhost:8080/api/v1/books
```

Kitap Kaydet:
```
$ curl -i -X POST -H "Content-Type: application/json" -d "{ \"title\": \"How to be worse at programming\", \"author\": \"Fatopato\" }" http://localhost:8080/api/v1/books
```

ID ile Tekil Kitap Getir (id: 1):

```
curl -i http://localhost:8080/api/v1/books/1
```

ID ile Kitap Güncelle (id:1)

```
curl -i -X PUT -H "Content-Type: application/json" -d "{ \"title\": \"It's getting even worse\", \"author\": \"Fatopato\" }" http://localhost:8080/api/v1/books/1
```

ID ile Kitap Sil (id:1)
```
curl -i -X DELETE http://localhost:8080/api/v1/books/1
```
