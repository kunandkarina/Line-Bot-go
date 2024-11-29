1. 註冊 Line Bot
前往 Line Developers Console。
建立一個新的 Provider 和 Channel。
在 Messaging API 設定中，取得以下憑證資訊：
Channel Secret
Channel Access Token
2. 啟動 Ngrok
使用以下指令將本地的 3000 埠（或您程式使用的埠號）公開到公網：
ngrok http 3000
Ngrok 會提供一個 HTTPS 網址，格式如下：https://<random-string>.ngrok.io
3. 設定 Webhook URL
回到 Line Developers Console，將 Webhook URL 設定為：
https://<random-string>.ngrok.io/callback
4. 註冊 Line Notify Token
前往 Line Notify Token 註冊頁面。
建立個人或群組通知的 Token，用於讓 Bot 發送通知。
取得 Notify Token，並將其用於您的專案中。
