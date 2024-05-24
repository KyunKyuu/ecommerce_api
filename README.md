# CHALLENGE TIM MIKTI - ECommerce

## 🎯. Fitur
Buyer: Pengguna dapat memesan produk.
Seller: Penjual dapat menawarkan dan menjual produk.
Autentikasi: Pengguna dapat melakukan pendaftaran dan login.
Otorisasi: Setiap fungsi diakses berdasarkan peran pengguna.
CRUD: Setiap pengguna dapat melakukan operasi Create, Read, Update, dan Delete sesuai dengan hak akses masing-masing.

## 🛠️. Flowchart
![flow ecommerce drawio](https://github.com/KyunKyuu/ecommerce_api/assets/57865250/81574029-76cb-4440-ab61-b449aa870799)

### Penjelasan Flowchart
Flowchart di atas menggambarkan alur kerja dari platform E-Commerce ini, mulai dari interaksi pengguna hingga pemrosesan pesanan. Berikut adalah penjelasan dari tiap bagian dalam flowchart:

**Register**
- Pengguna diwajibkan untuk melakukan registrasi terlebih dahulu jika sebelumnya belum memiliki akun. Saat registrasi, pengguna dapat memilih Rolenya apakah akan menjadi penjual (seller) atau pembeli (buyer).

**Login**
- Pengguna diwajibkan untuk login terlebih dahulu agar dapat mengakses aplikasi. Saat melakukan login, sistem akan otomatis mengenerate token JWT untuk otorisasi.

**Beli Barang**
- Pengguna dapat melakukan pembelian barang yang tersedia. Namun, hanya pengguna dengan Role Buyer saja yang dapat melakukan pembelian.

**Lihat Order**
- Setelah melakukan pembelian, pengguna dapat melihat pembelian yang telah dibuat sebelumnya

**Tambah Kategori**
- Pengguna dengan Role Seller dapat menambahkan Kategori Barang

**Tambah Produk**
- Pengguna dengan Role Seller dapat menambahkan Barang yang ingin dijual


## 💾. ERD



