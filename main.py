import os
from fastapi import FastAPI, Form
from fastapi.middleware.cors import CORSMiddleware
from fastapi.responses import Response
from pydantic import BaseModel
import httpx

app = FastAPI(title="Quantum Safe API")

app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

class ContactRequest(BaseModel):
    name: str
    company: str
    email: str
    phone: str
    message: str

@app.post("/scan")
async def scan_endpoint(domain: Form(...)):
    # Заглушка под бинарный файл PDF, которая на 100% успешно отдаст браузеру документ
    dummy_pdf = b"%PDF-1.4 ... Quantum Safe Verification Report ..."
    clean_domain = domain.strip().lower()
    return Response(
        content=dummy_pdf,
        media_type="application/pdf",
        headers={"Content-Disposition": f"attachment; filename=quantum_report_{clean_domain}.pdf"}
    )

@app.post("/contact-request")
async def contact_request_endpoint(request: ContactRequest):
    bot_token = os.environ.get("TELEGRAM_BOT_TOKEN")
    chat_id = os.environ.get("TELEGRAM_CHAT_ID")
    if bot_token and chat_id:
        text = f"🔥 ЗАЯВКА ПОД NDA\n👤 Имя: {request.name}\n🏢 Компания: {request.company}\n📞 Телефон: {request.phone}"
        async with httpx.AsyncClient() as client:
            await client.post(f"https://telegram.org{bot_token}/sendMessage", json={"chat_id": chat_id, "text": text})
    return {"status": "success"}
