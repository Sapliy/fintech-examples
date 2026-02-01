from fastapi import FastAPI, HTTPException
from fintech import FintechClient
import os

app = FastAPI()
client = FintechClient(api_key=os.getenv("SAPLIY_API_KEY"))

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

if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="0.0.0.0", port=3000)
