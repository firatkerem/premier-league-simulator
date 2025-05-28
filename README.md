# Premier League Simulator

Bu proje, 4 takım arasında futbol ligi simülasyonu yapan bir web uygulamasıdır. Go programlama dili kullanılarak geliştirilmiş olup, modern web teknolojileri ve veritabanı yönetimi içerir.

## Özellikler

- 4 farklı takım ve her takımın kendine özgü güç değeri
- Haftalık veya toplu maç simülasyonu
- Canlı puan durumu takibi
- Premier Lig tarzı puanlama sistemi (3-1-0)
- Modern ve kullanıcı dostu arayüz
- RESTful API desteği

## Teknik Detaylar

- Backend: Go (Gin Framework)
- Frontend: HTML, CSS, JavaScript
- Veritabanı: SQLite
- API: RESTful endpoints

## Kurulum

1. Go'yu yükleyin (1.21 veya üstü)

2. Projeyi klonlayın:
```bash
git clone <repository-url>
cd INSIDER
```

3. Bağımlılıkları yükleyin:
```bash
go mod tidy
```

4. Uygulamayı çalıştırın:
```bash
go run main.go
```

5. Tarayıcınızda açın:
- http://localhost:8080

## Kullanım

1. Ana sayfada puan durumu tablosunu göreceksiniz
2. "Simulate Next Week" butonu ile bir sonraki haftanın maçlarını simüle edebilirsiniz
3. "Simulate All Remaining" butonu ile kalan tüm maçları simüle edebilirsiniz
4. "Reset League" butonu ile ligi sıfırlayabilirsiniz

## API Endpoints

- `GET /api/league/table` — Güncel puan durumu
- `POST /api/league/simulate/week` — Bir sonraki haftayı simüle et
- `POST /api/league/simulate/all` — Kalan tüm maçları simüle et
- `POST /api/league/reset` — Ligi sıfırla

## Proje Yapısı

```
.
├── main.go           # Ana uygulama dosyası
├── models/           # Veri modelleri
│   ├── team.go      # Takım modeli
│   └── match.go     # Maç modeli
├── simulation/       # Simülasyon mantığı
│   └── simulator.go
└── static/          # Frontend dosyaları
    ├── index.html
    ├── styles.css
    └── app.js
```

## Geliştirme

Proje, modüler bir yapıda tasarlanmıştır:
- `models/`: Veri modelleri ve veritabanı şeması
- `simulation/`: Maç simülasyonu mantığı
- `static/`: Frontend dosyaları

## Notlar

- Uygulama SQLite veritabanı kullanmaktadır
- Tüm API endpoint'leri JSON formatında yanıt verir
- Frontend tarafında modern CSS ve JavaScript kullanılmıştır
- Responsive tasarım ile mobil uyumludur 
