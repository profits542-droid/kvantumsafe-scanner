import os
import re
import ipaddress
import socket
import tempfile
from typing import List, Optional, Tuple

import httpx
from fastapi import FastAPI, Form, HTTPException
from fastapi.middleware.cors import CORSMiddleware
from fastapi.responses import HTMLResponse, FileResponse
from pydantic import BaseModel, field_validator
from external_crypto_scanner import scan_domain_crypto, generate_pdf_report

app = FastAPI(
    title="ОсОО «Квантум Сейф»",
    description="Платформа постквантовой безопасности и Crypto-Agility для банков КР",
    version="2.3.0"
)

# Домены, с которых разрешено обращаться к API (сайт index.html хостится
# отдельно от бэкенда). Замените на реальный домен сайта перед деплоем —
# "*" НЕ используйте, если бэкенд также обслуживает NDA-заявки с персональными
# данными.
ALLOWED_ORIGINS = [
    "https://ВАШ-ДОМЕН-САЙТА.example",  # TODO: заменить на реальный домен
    "http://localhost:3000",             # для локальной разработки
]

app.add_middleware(
    CORSMiddleware,
    allow_origins=ALLOWED_ORIGINS,
    allow_methods=["POST"],
    allow_headers=["Content-Type"],
)

# Секреты Telegram-бота задаются ТОЛЬКО через переменные окружения на
# Render (Settings → Environment). Никогда не хардкодьте токен в коде и
# никогда не передавайте его в браузер — см. обсуждение в чате.
TELEGRAM_BOT_TOKEN = os.environ.get("TELEGRAM_BOT_TOKEN")
TELEGRAM_CHAT_ID = os.environ.get("TELEGRAM_CHAT_ID")

EMAIL_RE = re.compile(r"^[^@\s]+@[^@\s]+\.[^@\s]+$")
MAX_FIELD_LEN = 300
MAX_MESSAGE_LEN = 2000


class ContactRequest(BaseModel):
    name: str
    company: str
    email: str
    phone: Optional[str] = ""
    message: Optional[str] = ""
    lang: Optional[str] = "ru"

    @field_validator("name", "company", "email", "phone", "message")
    @classmethod
    def strip_and_limit(cls, v: str) -> str:
        if v is None:
            return ""
        v = v.strip()
        limit = MAX_MESSAGE_LEN if len(v) > MAX_FIELD_LEN else MAX_FIELD_LEN
        return v[:limit]

    @field_validator("email")
    @classmethod
    def validate_email(cls, v: str) -> str:
        if not EMAIL_RE.match(v):
            raise ValueError("invalid email")
        return v


def resolve_and_validate_target(hostname: str) -> Tuple[bool, List[str]]:
    """
    Полная защита от SSRF (IPv4 + IPv6, DNS Rebinding prevention):
    1. Резолвит все A и AAAA записи через socket.getaddrinfo().
    2. Проверяет каждый IP на вхождение в приватные/служебные подсети.
    3. Возвращает флаг безопасности и список разрешенных IP-адресов.

    Список валидных IP, возвращаемый здесь, должен использоваться для
    последующего соединения (IP pinning) — повторный резолв hostname
    на этапе самого сканирования сводит эту проверку на нет (DNS rebinding).
    """
    cleaned = hostname.strip().lower()

    if cleaned in ["localhost", "127.0.0.1", "0.0.0.0", "::1"]:
        return False, []

    safe_ips: List[str] = []
    try:
        addr_info = socket.getaddrinfo(cleaned, None, socket.AF_UNSPEC, socket.SOCK_STREAM)

        if not addr_info:
            return False, []

        for item in addr_info:
            ip_str = item[4][0]
            ip = ipaddress.ip_address(ip_str)

            if ip.is_private or ip.is_loopback or ip.is_link_local or ip.is_multicast or ip.is_reserved:
                return False, []

            if ip_str not in safe_ips:
                safe_ips.append(ip_str)

        return True, safe_ips
    except Exception:
        return False, []


HTML_CONTENT = """<!DOCTYPE html>
<html lang="ru">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Квантум Сейф — Постквантовая безопасность банков КР</title>
    <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/css/bootstrap.min.css" rel="stylesheet">
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap-icons@1.11.0/font/bootstrap-icons.css">
    <style>
        :root {
            --bg-dark: #0f172a;
            --card-bg: #1e293b;
            --accent-blue: #0284c7;
            --accent-green: #10b981;
            --text-light: #f8fafc;
            --text-muted: #94a3b8;
            --border-color: #334155;
        }
        body {
            background-color: var(--bg-dark);
            color: var(--text-light);
            font-family: 'Segoe UI', Roboto, Helvetica, Arial, sans-serif;
            min-height: 100vh;
        }
        .navbar-custom {
            background-color: rgba(15, 23, 42, 0.9);
            backdrop-filter: blur(10px);
            border-bottom: 1px solid var(--border-color);
        }
        .hero-section {
            padding: 60px 0 40px 0;
            background: radial-gradient(circle at top center, rgba(2, 132, 199, 0.15) 0%, transparent 70%);
        }
        .card-custom {
            background-color: var(--card-bg);
            border: 1px solid var(--border-color);
            border-radius: 12px;
            transition: transform 0.2s ease, box-shadow 0.2s ease;
            height: 100%;
        }
        .card-custom:hover {
            transform: translateY(-4px);
            box-shadow: 0 10px 25px -5px rgba(0, 0, 0, 0.5);
        }
        .badge-product {
            font-size: 0.8rem;
            padding: 6px 12px;
            border-radius: 20px;
            font-weight: 600;
        }
        .scan-box {
            background: linear-gradient(135deg, #1e293b 0%, #0f172a 100%);
            border: 1px solid var(--accent-blue);
            border-radius: 16px;
            padding: 30px;
            box-shadow: 0 0 30px rgba(2, 132, 199, 0.2);
        }
        .form-control-custom {
            background-color: #0f172a;
            border: 1px solid var(--border-color);
            color: #fff;
            padding: 12px 20px;
        }
        .form-control-custom:focus {
            background-color: #0f172a;
            color: #fff;
            border-color: var(--accent-blue);
            box-shadow: 0 0 0 0.25rem rgba(2, 132, 199, 0.25);
        }
        .btn-action {
            background: linear-gradient(135deg, #0284c7 0%, #0369a1 100%);
            color: #fff;
            font-weight: 600;
            padding: 12px 28px;
            border: none;
            border-radius: 8px;
            transition: all 0.3s;
        }
        .btn-action:hover {
            background: linear-gradient(135deg, #0369a1 0%, #075985 100%);
            color: #fff;
            box-shadow: 0 0 15px rgba(2, 132, 199, 0.4);
        }
        .legal-box {
            background-color: rgba(2, 132, 199, 0.05);
            border: 1px solid rgba(2, 132, 199, 0.3);
            border-radius: 12px;
            padding: 24px;
        }
        footer {
            border-top: 1px solid var(--border-color);
            padding: 30px 0;
            color: var(--text-muted);
            font-size: 0.9rem;
        }
    </style>
</head>
<body>

    <nav class="navbar navbar-expand-lg navbar-dark navbar-custom sticky-top">
        <div class="container">
            <a class="navbar-brand d-flex align-items-center fw-bold" href="#">
                <i class="bi bi-shield-lock-fill text-info me-2 fs-4"></i>
                <span>Квантум Сейф</span>
            </a>
            <span class="badge bg-outline-info text-info border border-info rounded-pill px-3 py-2">Кыргызская Республика</span>
        </div>
    </nav>

    <section class="hero-section text-center">
        <div class="container">
            <h1 class="display-5 fw-bold mb-3">ОсОО «Квантум Сейф»</h1>
            <p class="lead text-muted max-w-2xl mx-auto">
                Экосистема постквантовой криптографической устойчивости и Crypto-Agility для коммерческих банков и финансово-кредитных организаций КР
            </p>
        </div>
    </section>

    <section class="container my-5">
        <div class="scan-box">
            <div class="row align-items-center">
                <div class="col-lg-7 mb-4 mb-lg-0">
                    <span class="badge badge-product bg-primary mb-2">Продукт №1 — Экспресс Аудит</span>
                    <h3 class="fw-bold mb-2"><i class="bi bi-search me-2 text-info"></i>Внешний Сканер Криптостойкости</h3>
                    <p class="text-muted mb-0">
                        Введите публичный домен ДБО или веб-портала банка. Система проверит параметры TLS/RSA, оценит соответствие стандартам NIST PQC / KHP и сформирует официальный PDF-отчет.
                    </p>
                </div>
                <div class="col-lg-5">
                    <form action="/scan" method="post" target="_blank">
                        <div class="mb-3">
                            <label for="domain" class="form-label text-muted small fw-semibold">Публичный домен банка:</label>
                            <div class="input-group">
                                <span class="input-group-text bg-dark border-secondary text-muted"><i class="bi bi-globe"></i></span>
                                <input type="text" class="form-control form-control-custom" id="domain" name="domain" placeholder="например, bank.kg" required>
                            </div>
                        </div>
                        <button type="submit" class="btn btn-action w-100 d-flex align-items-center justify-content-center">
                            <i class="bi bi-file-earmark-pdf-fill me-2"></i> Экспресс-сканирование и PDF
                        </button>
                    </form>
                </div>
            </div>
        </div>
    </section>

    <section class="container my-5">
        <div class="row g-4">
            <div class="col-md-4">
                <div class="card card-custom p-4">
                    <div class="d-flex justify-content-between align-items-center mb-3">
                        <span class="badge badge-product bg-primary">Продукт №1</span>
                        <i class="bi bi-globe2 text-primary fs-3"></i>
                    </div>
                    <h4 class="fw-bold mb-3">Внешний Сканер Периметра</h4>
                    <p class="text-muted small flex-grow-1">
                        Дистанционный анализ стойкости внешних сервисов (ДБО, API). Выявление устаревших алгоритмов (RSA-2048, TLS 1.0/1.1) и оценка готовности к миграции на постквантовые стандарты.
                    </p>
                    <hr class="border-secondary my-3">
                    <div class="d-flex justify-content-between align-items-center">
                        <span class="text-info small fw-bold">Экспресс-диагностика</span>
                        <a href="#domain" class="btn btn-sm btn-outline-info">Тестировать</a>
                    </div>
                </div>
            </div>

            <div class="col-md-4">
                <div class="card card-custom p-4">
                    <div class="d-flex justify-content-between align-items-center mb-3">
                        <span class="badge badge-product bg-warning text-dark">Продукт №2</span>
                        <i class="bi bi-cpu-fill text-warning fs-3"></i>
                    </div>
                    <h4 class="fw-bold mb-3">On-Premise Инвентаризатор</h4>
                    <p class="text-muted small flex-grow-1">
                        Глубокая локальная инвентаризация крипто-активов (АБС, SWIFT, Процессинг) в режиме NDA. Формирование спецификаций для закупки сертифицированных шлюзов шифрования.
                    </p>
                    <hr class="border-secondary my-3">
                    <div class="d-flex justify-content-between align-items-center">
                        <span class="text-warning small fw-bold">Изолированный контур</span>
                        <span class="badge bg-secondary">По запросу (NDA)</span>
                    </div>
                </div>
            </div>

            <div class="col-md-4">
                <div class="card card-custom p-4">
                    <div class="d-flex justify-content-between align-items-center mb-3">
                        <span class="badge badge-product bg-success">Продукт №3</span>
                        <i class="bi bi-diagram-3-fill text-success fs-3"></i>
                    </div>
                    <h4 class="fw-bold mb-3">Квантум Оркестратор</h4>
                    <p class="text-muted small flex-grow-1">
                        Интеллектуальный API-маршрутизатор (Crypto-Agility). Автоматически перенаправляет банковский трафик на нужные аппаратные шлюзы (SWIFT &rarr; NIST PQC, CIPS &rarr; SM2/SM4) без изменения ядра АБС.
                    </p>
                    <hr class="border-secondary my-3">
                    <div class="d-flex justify-content-between align-items-center">
                        <span class="text-success small fw-bold">Crypto-Agility Engine</span>
                        <span class="badge bg-success">Модуль интеграции</span>
                    </div>
                </div>
            </div>
        </div>
    </section>

    <section class="container my-5">
        <div class="legal-box">
            <div class="row align-items-center">
                <div class="col-md-1 text-center mb-3 mb-md-0">
                    <i class="bi bi-shield-check text-info display-6"></i>
                </div>
                <div class="col-md-11">
                    <h5 class="fw-bold text-info mb-2">Архитектурный подход и нормативное соответствие (ГКНБ КР и НБКР)</h5>
                    <p class="text-muted small mb-0">
                        ПО «Квантум Оркестратор» функционирует в классе шин интеграции (API Gateway / Proxy) и осуществляет интеллектуальную маршрутизацию запросов к внешним сертифицированным аппаратным СКЗИ (HSM / VPN-шлюзам). Данный подход минимизирует затраты на доработку ядра АБС, а финальная схема встраивания проходит комплексную валидацию в рамках предэксплуатационного аудита регуляторов.
                    </p>
                </div>
            </div>
        </div>
    </section>

    <footer>
        <div class="container text-center">
            <p class="mb-1">ОсОО «Квантум Сейф» &copy; 2026. Все права защищены.</p>
            <p class="small text-muted mb-0">г. Бишкек, Кыргызская Республика | Поддержка стандартов НБКР, NIST PQC, KHP SM2/SM4</p>
        </div>
    </footer>

</body>
</html>"""


@app.get("/", response_class=HTMLResponse)
async def read_root():
    return HTMLResponse(content=HTML_CONTENT, status_code=200)


@app.post("/scan")
async def handle_scan(domain: str = Form(...)):
    cleaned_domain = domain.strip().lower().replace("https://", "").replace("http://", "").split("/")[0]

    if not cleaned_domain:
        raise HTTPException(status_code=400, detail="Укажите корректное доменное имя банка.")

    is_safe, pinned_ips = resolve_and_validate_target(cleaned_domain)
    if not is_safe or not pinned_ips:
        raise HTTPException(
            status_code=400,
            detail="Запрошенный адрес недоступен для внешнего сканирования (блокировка внутренних и служебных сетей)."
        )

    try:
        # ВАЖНО: hostname передаётся отдельно от pinned_ips, чтобы
        # scan_domain_crypto соединялся строго по уже проверенному IP
        # (для Host-заголовка / TLS SNI при этом используется исходный
        # hostname, а НЕ выполняет повторный DNS-резолв домена).
        # Это требует, чтобы сигнатура scan_domain_crypto в
        # external_crypto_scanner принимала pinned-адрес — см. заметку ниже.
        audit_data = await scan_domain_crypto(
            hostname=cleaned_domain,
            pinned_ips=pinned_ips,
        )

        temp_dir = tempfile.gettempdir()
        pdf_path = os.path.join(temp_dir, f"quantum_safe_report_{cleaned_domain}.pdf")

        generate_pdf_report(audit_data, pdf_path)

        return FileResponse(
            path=pdf_path,
            filename=f"QuantumSafe_Audit_{cleaned_domain}.pdf",
            media_type="application/pdf"
        )
    except Exception as e:
        raise HTTPException(status_code=500, detail=f"Ошибка при проведении внешнего аудита: {str(e)}")


def build_telegram_message(payload: ContactRequest) -> str:
    lang_label = "RU" if payload.lang != "en" else "EN"
    phone = payload.phone or "—"
    message = payload.message or "—"
    return (
        "🔥 НОВАЯ ЗАЯВКА ПОД NDA [Quantum Safe]\n"
        "━━━━━━━━━━━━━━━━━━\n"
        f"👤 Имя: {payload.name}\n"
        f"🏢 Компания: {payload.company}\n"
        f"📧 Email: {payload.email}\n"
        f"📞 Телефон: {phone}\n"
        f"💬 Сообщение: {message}\n"
        "━━━━━━━━━━━━━━━━━━\n"
        f"🌐 Язык интерфейса: {lang_label}"
    )


@app.post("/contact-request")
async def handle_contact_request(payload: ContactRequest):
    if not TELEGRAM_BOT_TOKEN or not TELEGRAM_CHAT_ID:
        # Секреты не настроены на сервере — не раскрываем деталей клиенту.
        raise HTTPException(status_code=500, detail="Сервис приёма заявок временно недоступен.")

    text = build_telegram_message(payload)
    telegram_url = f"https://api.telegram.org/bot{TELEGRAM_BOT_TOKEN}/sendMessage"

    try:
        async with httpx.AsyncClient(timeout=10.0) as client:
            response = await client.post(
                telegram_url,
                json={
                    "chat_id": TELEGRAM_CHAT_ID,
                    "text": text,
                    # Без parse_mode: имя/компания/сообщение приходят от
                    # пользователя и могут содержать символы, ломающие
                    # MarkdownV2/HTML-разметку Telegram.
                },
            )
        response.raise_for_status()
        data = response.json()
        if not data.get("ok"):
            raise HTTPException(status_code=502, detail="Telegram API вернул ошибку.")
    except httpx.HTTPError:
        raise HTTPException(status_code=502, detail="Не удалось отправить заявку в Telegram.")

    return {"ok": True}
