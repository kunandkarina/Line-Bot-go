# Line Bot 與 Line Notify 整合專案使用說明

## 簡介

本專案透過 Go 語言整合 Line Bot 和 Line Notify，實現訊息互動與通知功能。

## 設定步驟

### 1. 註冊 Line Bot

1. 前往 [Line Developers Console](https://developers.line.biz/console/)。
2. 建立一個新的 **Provider** 和 **Channel**。
3. 進入 **Messaging API** 設定頁面，並取得以下憑證資訊：
   - **Channel Secret**
   - **Channel Access Token**

   ```bash
   ChannelSecret="YOUR_CHANNEL_SECRET"
   ChannelAccessToken=YOUR_CHANNEL_ACCESS_TOKEN
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

### 3. 設定 Webhook URL

1. 前往 [Line Developers Console](https://developers.line.biz/console/)。
2. 在您的 **Channel** 中進入 **Messaging API** 設定頁面。
3. 將 **Webhook URL** 設定為：
   ```bash
   https://<random-string>.ngrok.io/callback
   ```
4. 儲存設定後，點擊 **Verify** 以測試 Webhook 連線是否成功。

###  4. 註冊 Line Notify Token

1. 前往 [Line Notify Token 註冊頁面](https://notify-bot.line.me/my/)。
2. 點擊 **生成存取權杖**（Generate Token）。
3. 選擇發送通知的對象：
   - **個人**：將通知發送至您的個人 Line 帳號。
   - **群組**：將通知發送至某個群組（需事先邀請 Line Notify 加入該群組）。
4. 為 Token 命名，選擇要授予的權限，然後點擊 **生成**。
5. 取得生成的 **Notify Token**，並將其用於專案中:
   ```bash
   LINENotifyToken="YOUR_LINE_NOTIFY_TOKEN"
   ```
