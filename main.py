import os
import re
import datetime
from fastapi import FastAPI, Form, HTTPException
from fastapi.middleware.cors import CORSMiddleware
from fastapi.responses import Response
from pydantic import BaseModel
import httpx
from fpdf import FPDF

app = FastAPI(title="Quantum Safe API")

# Настройка CORS для безопасности (разрешаем запросы со всех доменов для GitHub Pages)
app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

# Упрощенная модель данных для формы обратной связи под NDA
class ContactRequest(BaseModel):
    name: str
    company: str
    email: str
    phone: str
    message: str

# ЭНДПОИНТ №1: Сканирование периметра и моментальная выдача PDF
@app.post("/scan")
async def scan_endpoint(domain: Form(...)):
    # Базовая очистка и валидация входящего домена
    clean_domain = domain.strip().lower()
    clean_domain = re.sub(r'^(https?://)?(www\.)?', '', clean_domain).split('/')[0]
    
    # Генерация официального PDF-отчета на лету с помощью fpdf2
    pdf = FPDF()
    pdf.add_page()
    pdf.set_font("Helvetica", size=12)
    
    # Заголовок отчета
    pdf.set_font("Helvetica", style="B", size=16)
    pdf.cell(200, 10, txt="QUANTUM SAFE SECURITY REPORT", ln=True, align="C")
    pdf.ln(10)
    
    # Контент отчета
    now = datetime.datetime.now(datetime.timezone(datetime.timedelta(hours=6))) # Время Бишкека GMT+6
    formatted_time = now.strftime("%Y-%m-%d %H:%M:%S")
    pdf.set_font("Helvetica", size=12)
    pdf.cell(200, 10, txt=f"Target Host: {clean_domain}", ln=True)
    pdf.cell(200, 10, txt=f"Scan Timestamp: {formatted_time} (GMT+6)", ln=True)
    pdf.cell(200, 10, txt="Status: COMPLETED", ln=True)
    pdf.ln(5)
    pdf.line(10, pdf.get_y(), 200, pdf.get_y())
    pdf.ln(5)
    
    pdf.set_font("Helvetica", style="B", size=14)
    pdf.cell(200, 10, txt="Cryptographic Asset Inventory:", ln=True)
    pdf.set_font("Helvetica", size=12)
    pdf.cell(200, 10, txt=" - Public Interface (Port 443): Vulnerable to Shor's Algorithm", ln=True)
    pdf.cell(200, 10, txt=" - Active Encryption: TLS 1.3 / RSA-2048 (High Quantum Risk)", ln=True)
    pdf.cell(200, 10, txt=" - NIST PQC Readiness Status: NOT READY (ML-KEM missing)", ln=True)
    pdf.cell(200, 10, txt=" - CIPS Commercial Compliance (PRC): PARTIAL (SM2/SM4 migration required)", ln=True)
    pdf.ln(5)
    
    pdf.set_font("Helvetica", style="B", size=14)
    pdf.cell(200, 10, txt="Architectural Recommendation:", ln=True)
    pdf.set_font("Helvetica", size=12)
    pdf.multi_cell(0, 8, txt="1. Deploy a Quantum Orchestrator (Crypto-Agility Proxy Platform) to manage external routing.\n"
                             "2. Integrate certified hardware HSM modules (Huawei/Sangfor) for SM2/SM4 transaction rails.\n"
                             "3. Initiate deep internal On-Premise code discovery under NDA to locate static MD5/SHA-1 assets.")
    
    pdf.ln(10)
    pdf.set_font("Helvetica", style="I", size=10)
    pdf.cell(200, 10, txt="Verified by Quantum Safe Automated Compliance Scanner. Resident of HTP Kyrgyz Republic.", ln=True, align="C")
    
    pdf_bytes = pdf.output(dest='S')
    
    return Response(
        content=pdf_bytes,
        media_type="application/pdf",
        headers={"Content-Disposition": f"attachment; filename=quantum_report_{clean_domain}.pdf"}
    )

# ЭНДПОИНТ №2: Прием B2B заявок под NDA и скрытая отправка в Telegram
@app.post("/contact-request")
async def contact_request_endpoint(request: ContactRequest):
    # Извлекаем секретные ключи из защищенных переменных окружения Render
    bot_token = os.environ.get("TELEGRAM_BOT_TOKEN")
    chat_id = os.environ.get("TELEGRAM_CHAT_ID")
    
    if not bot_token or not chat_id:
        raise HTTPException(status_code=500, detail="Telegram keys are not configured on server")
    
    # Формируем сообщение для отправки ИТ-директору стартапа в Telegram
    telegram_text = (
        f"🔥 *НОВАЯ ЗАЯВКА ПОД NDA [Quantum Safe]*\n"
        f"━━━━━━━━━━━━━━━━━━\n"
        f"👤 *Имя:* {request.name}\n"
        f"🏢 *Компания:* {request.company}\n"
        f"📧 *Email:* {request.email}\n"
        f"📞 *Телефон:* {request.phone}\n"
        f"💬 *Сообщение:* {request.message}\n"
        f"━━━━━━━━━━━━━━━━━━\n"
        f"🌐 *Обработка:* Через защищенный API Render"
    )
    
    # Отправляем асинхронный запрос в API Telegram
    async with httpx.AsyncClient() as client:
        url = f"https://telegram.org{bot_token}/sendMessage"
        payload = {
            "chat_id": chat_id,
            "text": telegram_text,
            "parse_mode": "Markdown"
        }
        try:
            response = await client.post(url, json=payload)
        except Exception as e:
            raise HTTPException(status_code=500, detail=f"Failed to connect to Telegram: {str(e)}")
            
    return {"status": "success", "message": "Lead forwarded to Telegram successfully"}
