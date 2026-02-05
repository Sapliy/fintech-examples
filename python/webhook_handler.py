import os
import json
from flask import Flask, request, jsonify
from sapliyio_fintech import SapliyClient

app = Flask(__name__)

# Initialize client
api_key = os.getenv('SAPLIY_API_KEY', 'sk_test_123')
# The SDK uses SapliyClient
client = SapliyClient(api_key=api_key)

webhook_secret = os.getenv('SAPLIY_WEBHOOK_SECRET', 'whsec_test_123')

@app.route('/health', methods=['GET'])
def health():
    return jsonify({"status": "ok"})

@app.route('/webhook', methods=['POST'])
def webhook():
    try:
        # For now, we bypass signature verification as it's not in the base SDK
        payload = request.get_data().decode('utf-8')
        event = json.loads(payload)
        
        print(f"Received event: {event.get('type')}")

        # Handle the event
        event_type = event.get('type')
        if event_type == 'payment.succeeded':
            payment = event.get('data', {}).get('object', {})
            print(f"💰 Payment {payment.get('id')} succeeded for {payment.get('amount')}")
        elif event_type == 'payment.failed':
            print("❌ Payment failed")
        else:
            print(f"Unhandled event type: {event_type}")

        return jsonify({"received": True})
    except Exception as e:
        print(f"Webhook Error: {str(e)}")
        return jsonify({"error": str(e)}), 400

if __name__ == '__main__':
    print("Python Webhook Handler running on port 5001")
    app.run(port=5001)
