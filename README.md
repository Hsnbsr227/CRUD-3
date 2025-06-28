# Go CRUD Uygulaması

Bu proje, PostgreSQL veritabanı kullanarak temel CRUD (Create, Read, Update, Delete) işlemlerini gerçekleştiren bir Go uygulamasıdır.

## 🚀 Özellikler

- ✅ Kullanıcı ekleme (Create)
- ✅ Kullanıcı listeleme (Read)
- ✅ Kullanıcı güncelleme (Update)
- ✅ Kullanıcı silme (Delete)
- 🔒 Transaction desteği
- 🛡️ Güvenli veritabanı bağlantısı
- 📝 Kullanıcı dostu konsol arayüzü

## 📋 Gereksinimler

- Go 1.16 veya üzeri
- PostgreSQL veritabanı
- Git

## 🛠️ Kurulum

### 1. Projeyi klonlayın
```bash
git clone <repository-url>
cd CRUD
```

### 2. Bağımlılıkları yükleyin
```bash
go mod init crud-app
go get github.com/lib/pq
go get github.com/joho/godotenv
```

### 3. Veritabanını hazırlayın
PostgreSQL'de aşağıdaki tabloyu oluşturun:

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(100) NOT NULL,
    password VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### 4. Environment değişkenlerini ayarlayın
`.env` dosyası oluşturun:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_username
DB_PASSWORD=your_password
DB_NAME=your_database_name
DB_SSLMODE=disable
```

### 5. Uygulamayı çalıştırın
```bash
go run .
```

## 🎯 Kullanım

Uygulama başlatıldığında aşağıdaki menü görünecektir:

```
--- Main Menu ---
1. Add User
2. List Users
3. Delete User
4. Update User
0. Exit
```

### Kullanıcı Ekleme
- Menüden `1` seçin
- Kullanıcı adı ve şifre girin

### Kullanıcı Listeleme
- Menüden `2` seçin
- Tüm kullanıcılar listelenecektir

### Kullanıcı Silme
- Menüden `3` seçin
- Silinecek kullanıcının ID'sini girin

### Kullanıcı Güncelleme
- Menüden `4` seçin
- Güncellenecek kullanıcının ID'sini girin
- Yeni kullanıcı adı ve şifre girin

## 📁 Proje Yapısı

```
CRUD/
├── main.go          # Ana uygulama ve menü sistemi
├── db.go            # Veritabanı bağlantısı ve transaction yönetimi
├── create.go        # Kullanıcı ekleme işlemleri
├── read.go          # Kullanıcı listeleme işlemleri
├── update.go        # Kullanıcı güncelleme işlemleri
├── delete.go        # Kullanıcı silme işlemleri
├── .env.example     # Örnek environment değişkenleri
├── .gitignore       # Git ignore dosyası
└── README.md        # Bu dosya
```

## 🔧 Teknik Detaylar

- **Veritabanı:** PostgreSQL
- **ORM:** Standart database/sql paketi
- **Transaction Yönetimi:** Manuel transaction kontrolü
- **Environment:** godotenv ile environment değişkenleri
- **Hata Yönetimi:** Kapsamlı hata yakalama ve loglama

## 🛡️ Güvenlik

- Veritabanı şifreleri `.env` dosyasında saklanır
- `.env` dosyası `.gitignore` ile korunur
- Transaction kullanılarak veri tutarlılığı sağlanır

## 🤝 Katkıda Bulunma

1. Bu repository'yi fork edin
2. Yeni bir branch oluşturun (`git checkout -b feature/yeni-ozellik`)
3. Değişikliklerinizi commit edin (`git commit -am 'Yeni özellik eklendi'`)
4. Branch'inizi push edin (`git push origin feature/yeni-ozellik`)
5. Pull Request oluşturun

## 📝 Lisans

Bu proje MIT lisansı altında lisanslanmıştır.

## 👨‍💻 Geliştirici

Bu proje Go öğrenme sürecinde geliştirilmiştir.

---

⭐ Bu projeyi beğendiyseniz yıldız vermeyi unutmayın! 
