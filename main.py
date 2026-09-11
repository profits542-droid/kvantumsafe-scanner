from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware
from pydantic import BaseModel

app = FastAPI()

# Настройка CORS без конфликтов для любой версии Python
app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_credentials=False,
    allow_methods=["GET", "POST", "OPTIONS"],
    allow_headers=["Content-Type", "Authorization"],
)

# Модель данных для приема запросов с формы index.html
class ScanRequest(BaseModel):
    name: str
    company: str
    email: str
    phone: str
    message: str

@app.post("/contact/")
async def handle_scan(data: ScanRequest):
    return {
        "status": "success", 
        "message": f"Сканирование для домена {data.name} успешно запущено!"
    }
