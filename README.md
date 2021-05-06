# project-individu-go-react

:tada: :tada: :tada:

Kita akan membuat tugas project pribadi dengan membuat back end menggunakan GO dan package GIN-GONIC dan juga package GORM. Kemudian kita juga akan membuat frond end menggunakan REACT JS

- Silahkan tentukan mau membuat project seperti apa, untuk tema bebas. Tetapi khusus untuk back end harus memenuhi kriteria dibawah ini.

# Project - Day 1
## Release 1
Buatlah sebuah tabel diagram sesuai keinginan kalian dengan catatan :
1. terdapat minimal 3 tabel di dalam database.
2. memiliki tabel user, dengan kolom sebagai berikut : 
  * id / user_id : primarykey, not null, auto increment
  * first_name  : varchar(255)
  * last_name  : varchar(255)
  * email  : varchar(255), unique
  * password  : varchar(255)
  * created_at : date(3) 
  * updated_at : date(3)

3. table user memiliki relasi dengan tabel lain, setiap tabel selain tabel user juga memiliki relasi

> NB : Silahkan laporkan ketika release 1 sudah selesai di github dalam bentuk foto atau pdf.

## release 2
Buatlah rest API yang dibutuhkan untuk menggunakan GIN - GONIC dan GORM.
- REST API yang dibuat harus memenuhi basic CRUD ( create, read, update, delete)
- Didalam rest API tersebut harus memiliki :
   * POST "/users" ket : routingan untuk membuat user baru
   * GET "/users" , ket : menampilkan seluruh user yang ada di database
   * GET "/users/:user_id" , ket : menampilkan 1 user sesuai dengan parameter yang diberikan
   * PUT "/users/:user_id", ket : melakukan update user sesuai dengan parameter yang diberikan
   * DELETE "/users/:user_id" ket : routingan untuk melakukan delete user sesuai parameter yang diberikan"

# Project - Day 2
## Release 3

- Silahkan perbaiki code REST API yang sudah di buat dengan menerapkan arsitektur yang sudah dijelaskan dengan menerapkan layer - layer :
1. repository , ket : package ini adalah package yang digunakan untuk berkomunikasi dengan database
2. service , ket : package ini adalah package yang digunakan untuk melakukan bussiness logic di back end
3. handler / controller : package ini digunakan untuk berkomunikasi dengan user / pengguna eksternal baik itu menerima data atau menampilkan data


## release 4
- Gunakan hashing untuk data yang bersifat pribadi dengan membuat package helper. Seperti password user atau data lain di tabel yang datanya perlu di enksipsi.
- Aplikasikan "env" atau environment variabel pada project ini.


# Project Day 3
## release 5
### Saatnya menerapkan authorization menggunakan JWT "json web token" dan juga middleware di project ini. 

- Terdapat beberapa code yang wajib dibuat di project ini sebagai berikut: 
1. buatlah routing untuk register user dan juga login user
  * POST "/users/register" , ket : routingan register digunakan untuk melakukan register user
  * POST "/users/login" , ket : routingan login digunakan untuk mendapatkan akses authorization dan juga mendapatkan data user yang sedang melakukan login 

2. Silahkan menerapkan middleware untuk routingan :
  * GET "users/:user_id"
  * PUT "users/:user_id"
  * DELETE "/users/:user_id"

3. Gunakan middleware untuk akses routing yang bersifat privat atau memerlukan akses user login melalui routingan POST "/users/login"

# Project Day 4
## release 6
buatlah 1 routing khusus yang didapat dari 3rd party API ( bebas apa saja ) dan diolah kemudian tampilkan menggunakan routing GET.

## release 7
- lakukan deployment server menggunakan heroku dan juga mengaplikasikan env "environment variable" untuk menginput data connection database mysql sebagai berikut :
1. MYSQL_DATABASE :	mysql (or your preferred database name)
2. MYSQL_USER	:	mysql (or your preferred database username)
3. MYSQL_PASSWORD :	A secure password for MYSQL_USER
4. MYSQL_HOST	:	tcp or protocol deployment MYSQL

- kemudian buatlah API documentation dengan menggunakan swagger (boleh membuat dokumentasi manual menggunakan API-DOC.md)

# project Day 5 - 7
## release 8
Buatlah front end tampilan yang dibutuhkan menggunakan framework REACT JS. dengan page yang wajib dibuat adalah :
1. page login 
2. page register
3. page home (untuk menampilkan home page)
4. page user detail
5. memiliki fitur log out

## release 9
Gunakan API yang sudah dideploy di heroku untuk kemudian digabungkan dengan front end ( wiring ) kemudian menerapkan beberapa fitur berikut :
1. wajib menggunakan REDUX
2. gunakan private route dipage user detail dan juga page yang dirasa memerlukan login user terlebih dahulu

## release 10
Lakukan deployment frond end yang sudah dibuat menggunakan netlify.

> NB : Segera push data terbaru jika sudah selesai tiap release.
