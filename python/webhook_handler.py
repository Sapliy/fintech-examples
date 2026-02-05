from flask import Flask, request, jsonify
import os
import sapliy

app = Flask(__name__)

# Initialize client
api_key = os.getenv('SAPLIY_API_KEY', 'sk_test_123')
client = sapliy.Client(api_key=api_key)

webhook_secret = os.getenv('SAPLIY_WEBHOOK_SECRET', 'whsec_test_123')

@app.route('/health', methods=['GET'])
def health():
    return jsonify({"status": "ok"})

@app.route('/webhook', methods=['POST'])
def webhook():
    payload = request.data
    sig_header = request.headers.get('X-Sapliy-Signature')
    
    try:
        event = client.webhooks.construct_event(
            payload, sig_header, webhook_secret
        )
    except ValueError as e:
        # Invalid payload
        return jsonify({"error": str(e)}), 400
    except sapliy.error.SignatureVerificationError as e:
        # Invalid signature
        return jsonify({"error": str(e)}), 400

    # Handle the event
    if event['type'] == 'payment.succeeded':
        payment = event['data']['object']
        print(f"💰 Payment {payment['id']} succeeded for {payment['amount']}")
    elif event['type'] == 'payment.failed':
        print("❌ Payment failed")
    else:
        print(f"Received event: {event['type']}")

    return jsonify({"received": True})

if __name__ == '__main__':
    print("Python Webhook Handler running on port 5000")
    app.run(port=5000)
