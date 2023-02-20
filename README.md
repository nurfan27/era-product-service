### Requirement:
- Buat sebuah Rest API product dengan bahasa pemrograman go dengan RDMS mysql atau postgre
- Buat 2 endpoint.
  - API Add product
  - API List product dengan sorting


### Detail Requirement
- Dengan spesifikasi product:
  - product id
  - product name
  - product price
  - product description
  - product quantity
- Rest API bisa melakukan tambah product
- Rest API bisa mengurutkan berdasarkan berikut:
  - product terbaru
  - product harga termurah
  - product harga termahal
  - sort by product name (A-Z) dan (Z-A)


### Solusi
Untuk menjawab permasalahan diatas saya membuat program sebagai berikut :
- Pertama-tama saya membuat fungsi untuk menghandle request http dari client,
- Kemudian sediakan endpoint berdasarkan requirment diatas
  - POST : /api/v1/products
  - GET : /api/v1/products
- Pada service yang akan dibuat memiliki beberapa bagian diantaranya :
  - *Transport* bagian ini memiliki fungsi sebagai tempat berkumpulnya protokol yang di sediakan oleh service ini, sehingga bila ada penambahan protokol komunikasi baru tidak perlu merubah bagian lain, hanya menambahkan protokol dan fungsi adapternya.
  - *Adapter* fungsinya untuk mapping request dari tiap protokol yang disediakan oleh service sebelum menuju ke bagian Action, tujuannya untuk menghindari redudansi kode program pada satu logika bisnis yang sama.
  - *Action* tempat dimana logika bisnis di tulis.
  - *Repository* tempat dimana logika data ditulis.
  
### Dokumentasi API

>### Add Product
>`POST /api/v1/products`

#### Request

<pre><code>curl --request POST \
  --url http://localhost:8000/api/v1/products \
  --header 'Content-Type: application/json' \
  --data '{
	"name":"tiang tenda",
	"description":"tiang portable untuk tenda dan flysheet" ,
	"price":25000,
	"quantity":75
}'</code></pre>

#### Response

<pre><code>{
	"code": 200,
	"message": "OK",
	"data": {
		"description": "tiang portable untuk tenda dan flysheet",
		"id": 1,
		"name": "tiang tenda",
		"price": 25000,
		"quantity": 75
	}
}</code></pre>

>### Get Product
>`GET /api/v1/products?sort=<sort_value>`


#### Request
>Sort value
>
>- 1 => Latest
>- 2 => Price High
>- 3 => Price Low
>- 4 => Name (A-Z)
>- 5 => Name (Z-A)

<pre><code>curl --request GET \
  --url http://localhost:8000/api/v1/products?sort=1
</code></pre>

#### Response

<pre><code>{
	"code": 200,
	"message": "OK",
	"data": {
		"products": [
			{
				"description": "tepung putih ",
				"id": 3,
				"name": "Tepung Terigu",
				"price": 32000,
				"quantity": 300000
			},
			{
				"description": "minyak ",
				"id": 2,
				"name": "Minyak Goreng",
				"price": 12000,
				"quantity": 40000
			},
			{
				"description": "tiang portable untuk tenda dan flysheet",
				"id": 1,
				"name": "tiang tenda",
				"price": 25000,
				"quantity": 75
			}
		]
	}
}</code></pre>