# ❌ No-as-a-Service

<p align="center">
  <img src="https://raw.githubusercontent.com/ItsMeSudo/no-as-a-service-go/main/assets/imgs/image.png" width="800" alt="No-as-a-Service Banner"/>
</p>


Ever needed a graceful way to say “no”?  
This tiny API returns random, generic, creative, and sometimes hilarious rejection reasons — perfectly suited for any scenario: personal, professional, student life, dev life, or just because.

Built for humans, excuses, and humor.

---

## 🚀 API Usage

**Base URL**
```
https://naas.serverhu.eu/no
```

**Method:** `GET`  
**Rate Limit:** `No rate limit`

### 🔄 Example Request
```http
GET /no
```

### ✅ Example Response
```json
{
  "reason": "This feels like something Future Me would yell at Present Me for agreeing to."
}
```

Use it in apps, bots, landing pages, Slack integrations, rejection letters, or wherever you need a polite (or witty) no.

---

## 🛠️ Self-Hosting

Want to run it yourself? It’s lightweight and simple.

### 1. Clone this repository
```bash
git clone https://github.com/ItsMeSudo/no-as-a-service-go.git
cd no-as-a-service
```

### 2. Build
```bash
go build
```

### 3. Start the server
```bash
.\naas.exe or ./naas  [Windows or Linux]
```

The API will be live at:
```
http://localhost:3000/no
```

You can also change things using an environment variable:
```bash
PORT=3000 [int]
ENABLE_RATE_LIMIT=false [bool]
```

---

## 📁 Project Structure

```
no-as-service-go/
├── main.go             # Main source
├── reasons.json        # 1000+ universal rejection reasons
├── go.mod
├── go.sum
└── README.md
```

---

## 👤 Author

Created with creative stubbornness by [hotheadhacker](https://github.com/hotheadhacker)

Ported to GO by [SUDO](https://github.com/ItsMeSudo)

---

## 📄 License

MIT — do whatever, just don’t say yes when you should say no.
