# Tor URL Checker

Bu proje, belirlenen bir hedef listesindeki URL'lerin erişilebilirliğini Tor ağı (SOCKS5 Proxy) üzerinden kontrol eden basit bir Go aracıdır.

## Özellikler

* Tüm trafik yerel Tor SOCKS5 proxy (127.0.0.1:9050) üzerinden yönlendirilir.
* `targets.yaml` dosyasından toplu URL okuma işlemi yapar.
* Başarılı bağlantıları, zaman aşımı (timeout) hatalarını ve bağlantı hatalarını raporlar.
* Varsayılan zaman aşımı süresi 20 saniyedir.

## Gereksinimler

* Go (Golang) kurulu olmalıdır.
* Tor servisi makinenizde kurulu ve çalışıyor olmalıdır.
* Tor SOCKS5 proxy varsayılan portta (`9050`) dinliyor olmalıdır.

## Kurulum

1. Projeyi klonlayın veya indirin.
2. Gerekli Go paketini indirin:

```bash
go get golang.org/x/net/proxy
