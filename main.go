```go
package main

import (
	"encoding/base64"
	"html"
	"net/http"
	"strings"
	"time"
)

type LicenseInfo struct {
	ClientName string
	ServerID   string
	IsValid    bool
	DaysLeft   int
}

func VerifyLicenseKey(encodedKey string) LicenseInfo {
	info := LicenseInfo{IsValid: false, DaysLeft: 0}
	decodedBytes, err := base64.StdEncoding.DecodeString(encodedKey)
	if err != nil { return info }
	parts := strings.Split(string(decodedBytes), "|")
	if len(parts) == 3 {
		info.ClientName = parts
		info.ServerID = parts
		if expTime, err := time.Parse("2006-01-02", parts); err == nil {
			if time.Now().Before(expTime) {
				info.IsValid = true
				info.DaysLeft = int(expTime.Sub(time.Now()).Hours() / 24)
			}
		}
	}
	return info
}

func main() {
	premiumLicenseKey := "T0FPIEtvbW1lcmNoZXNreWkgYmFuayBLWVJHWVpTVEFOIChNQkFOSyl8MjAyNy0wOS0yMHxDb3JlLU5vZGUtMDE="
	license := VerifyLicenseKey(premiumLicenseKey)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" { http.NotFound(w, r); return }
		lang := r.URL.Query().Get("lang")
		if lang == "" { lang = "kg" }
		
		htmlPage := GetDashboardHTML(html.EscapeString(license.ClientName), html.EscapeString(license.ServerID), lang)
		
		direction := "right: 25px;"
		if lang == "ar" { direction = "left: 25px;" }
		
		botWidget := `
	<div style='position: fixed; bottom: 25px; ` + direction + ` width: 65px; height: 65px; background: #7c3aed; color: white; border-radius: 50%; display: flex; align-items: center; justify-content: center; font-size: 28px; cursor: pointer; box-shadow: 0 4px 16px rgba(124,58,237,0.4); z-index: 1000; transition: 0.3s;' onclick='toggleChat()'>🤖</div>
	<div style='position: fixed; bottom: 100px; ` + direction + ` width: 370px; height: 480px; background: white; border-radius: 12px; box-shadow: 0 8px 32px rgba(0,0,0,0.15); display: none; flex-direction: column; z-index: 1000; overflow: hidden; border: 1px solid #e2e8f0; font-size: 14px;' id='chatWindow'>
		<div style='background: #0a2540; color: white; padding: 15px; font-weight: bold; display: flex; justify-content: space-between; align-items: center;'>
			<div>KvantumSafe AI Assistant <span style='font-size: 12px; background: #16a34a; padding: 2px 6px; border-radius: 4px; margin-left: 10px;'>Online</span></div>
			<div style='cursor:pointer;' onclick='toggleChat()'>✕</div>
		</div>
		<div style='flex: 1; padding: 15px; overflow-y: auto; background: #f8fafc; display: flex; flex-direction: column; gap: 10px;' id='chatBody'>
			<div style='max-width: 80%; padding: 10px 14px; border-radius: 8px; line-height: 1.4; background: #e2e8f0; color: #1e293b; align-self: flex-start;'>` + getTranslation("bot_welcome", lang) + `</div>
		</div>
		<div style='padding: 10px; border-top: 1px solid #e2e8f0; display: flex; background: white;'>
			<input type='text' style='flex: 1; border: none; padding: 10px; outline: none; font-size: 14px;' id='chatInput' placeholder='` + getTranslation("bot_placeholder", lang) + `' onkeypress='handleKey(event)'>
			<button style='background: #0a2540; color: white; border: none; padding: 0 20px; font-weight: bold; cursor: pointer;' onclick='sendMessage()'>&gt;</button>
		</div>
	</div>
	<script>
		function toggleChat() {
			var win = document.getElementById("chatWindow");
			win.style.display = (win.style.display === "flex") ? "none" : "flex";
		}
		function handleKey(e) { if (e.key === "Enter") sendMessage(); }
		function sendMessage() {
			var input = document.getElementById("chatInput");
			var text = input.value.trim();
			if (!text) return;
			var body = document.getElementById("chatBody");
			var userMsg = document.createElement("div");
			userMsg.style = "max-width: 80%; padding: 10px 14px; border-radius: 8px; line-height: 1.4; background: #7c3aed; color: white; align-self: flex-end;";
			userMsg.innerText = text;
			body.appendChild(userMsg);
			input.value = "";
			body.scrollTop = body.scrollHeight;
			setTimeout(function() {
				var botMsg = document.createElement("div");
				botMsg.style = "max-width: 80%; padding: 10px 14px; border-radius: 8px; line-height: 1.4; background: #e2e8f0; color: #1e293b; align-self: flex-start;";
				var lowText = text.toLowerCase();
				if (lowText.includes("nist") || lowText.includes("алгоритм") || lowText.includes("квант") || lowText.includes("algorithm")) {
					botMsg.innerHTML = "` + getTranslation("bot_ans_nist", lang) + `";
				} else if (lowText.includes("встреч") || lowText.includes("купить") || lowText.includes("цена") || lowText.includes("meet") || lowText.includes("buy")) {
					botMsg.innerHTML = "` + getTranslation("bot_ans_meet", lang) + `";
				} else {
					botMsg.innerHTML = "` + getTranslation("bot_ans_default", lang) + `";
				}
				body.appendChild(botMsg);
				body.scrollTop = body.scrollHeight;
			}, 800);
		}
	</script>
</body>
</html>`
		
		htmlPage = strings.Replace(htmlPage, "</body>\n</html>", botWidget, 1)
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, _ = w.Write([]byte(htmlPage))
	})

	http.HandleFunc("/report/mbank", func(w http.ResponseWriter, r *http.Request) {
		lang := r.URL.Query().Get("lang")
		if lang == "" { lang = "kg" }
		domainName := r.URL.Query().Get("domain")
		if domainName == "" { domainName = "unknown-node.com" }
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, _ = w.Write([]byte(GetPerimeterReportHTML(lang, html.EscapeString(domainName))))
	})

	http.HandleFunc("/specification", func(w http.ResponseWriter, r *http.Request) {
		lang := r.URL.Query().Get("lang")
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, _ = w.Write([]byte(GetSpecificationPageHTML(lang)))
	})

	http.HandleFunc("/antihacker", func(w http.ResponseWriter, r *http.Request) {
		lang := r.URL.Query().Get("lang")
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, _ = w.Write([]byte(GetAntiHackerPageHTML(lang)))
	})

	http.HandleFunc("/download", func(w http.ResponseWriter, r *http.Request) {
		lang := r.URL.Query().Get("lang")
		w.Header().Set("Content-Disposition", "attachment; filename=KvantumSafe_Specification.txt")
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		_, _ = w.Write([]byte(getTranslation("download_text", lang)))
	})

	server := &http.Server{Addr: ":8080", ReadHeaderTimeout: 3 * time.Second}
	_ = server.ListenAndServe()
}
