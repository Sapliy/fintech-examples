from fastapi import FastAPI, HTTPException, Request, Header
from sapliyio_fintech import SapliyClient
import os
import json

app = FastAPI()
client = SapliyClient(api_key=os.getenv("SAPLIY_API_KEY"))
WEBHOOK_SECRET = os.getenv("SAPLIY_WEBHOOK_SECRET")

@app.post("/charge")
async def create_charge(amount: int):
    try:
        payment = client.payments.create(
            amount=amount,
            currency="USD",
            source_id="tok_visa",
            description="Example charge"
        )
        return payment
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))

@app.post("/webhook")
async def webhook(request: Request, x_sapliy_signature: str = Header(None)):
    payload = await request.body()
    
    try:
        event = client.webhooks.construct_event(
            payload=payload,
            signature=x_sapliy_signature,
            secret=WEBHOOK_SECRET
        )
        
        print(f"Received event: {event['type']}")
        
        if event['type'] == 'payment.succeeded':
            payment = event['data']['object']
            print(f"Payment for {payment['amount']} succeeded!")
            
        return {"status": "success"}
    except Exception as e:
        raise HTTPException(status_code=400, detail=str(e))

if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="0.0.0.0", port=3000)
