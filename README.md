# 使用說明：Line Bot 與 Line Notify 整合專案

## 簡介

本專案旨在整合 Line Bot 和 Line Notify，透過 Go 語言實現訊息互動與通知功能。以下步驟將引導您完成環境設置與專案執行。

---

## 設定步驟

### 1. 註冊 Line Bot

1. 前往 [Line Developers Console](https://developers.line.biz/console/)。
2. 建立一個新的 **Provider** 和 **Channel**。
3. 進入 **Messaging API** 設定頁面，並取得以下憑證資訊：
   - **Channel Secret**
   - **Channel Access Token**

> 將上述憑證資訊設定為環境變數，或直接寫入程式碼：
```bash
CHANNEL_SECRET="your_channel_secret"
CHANNEL_ACCESS_TOKEN="your_channel_access_token"
```

### 2. 設定 Ngrok

1. **啟動 Ngrok**  
   使用以下指令將本地 `3000` 埠公開到公網：  
   ```bash
   ngrok http 3000
   ```
2. Ngrok 啟動後會提供一個 HTTPS 網址，格式如下：
```bash
https://<random-string>.ngrok.io
```
