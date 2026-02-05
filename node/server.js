import express from 'express';
import Sapliyio from '@sapliyio/fintech';
import dotenv from 'dotenv';
import path from 'path';
import { fileURLToPath } from 'url';

dotenv.config();

const __dirname = path.dirname(fileURLToPath(import.meta.url));
const app = express();

// Initialize the SDK
const client = new Sapliyio.SapliyClient(process.env.SAPLIY_API_KEY || 'sk_test_123');
const WEBHOOK_SECRET = process.env.SAPLIY_WEBHOOK_SECRET || 'whsec_test_123';

app.use(express.static('public'));
app.use(express.json());

// Serve checkout page
app.get('/', (req, res) => {
  res.send(`
    <html>
      <head>
        <title>Sapliy Checkout</title>
        <style>
          body { font-family: sans-serif; display: flex; justify-content: center; align-items: center; height: 100vh; background: #f5f5f5; }
          .card { background: white; padding: 2rem; border-radius: 8px; box-shadow: 0 2px 10px rgba(0,0,0,0.1); text-align: center; }
          button { background: #6366f1; color: white; border: none; padding: 10px 20px; border-radius: 4px; cursor: pointer; font-size: 16px; }
          button:hover { background: #4f46e5; }
        </style>
      </head>
      <body>
        <div class="card">
          <h1>Premium Plan</h1>
          <p>$20.00 / month</p>
          <button id="checkout-btn">Subscribe</button>
        </div>
        <script>
          document.getElementById('checkout-btn').addEventListener('click', async () => {
            const res = await fetch('/create-payment-intent', { method: 'POST' });
            const data = await res.json();
            alert('Created Payment Intent: ' + data.id);
            // In a real app, you would redirect to payment page or use elements
          });
        </script>
      </body>
    </html>
  `);
});

// Create Payment Intent
app.post('/create-payment-intent', async (req, res) => {
  try {
    // The generated SDK uses paymentServiceCreatePaymentIntent and amount is a string
    const response = await client.payments.paymentServiceCreatePaymentIntent({
      amount: '2000', // $20.00
      currency: 'USD',
      description: 'Premium Plan Subscription'
    });

    res.json(response.data);
  } catch (error) {
    res.status(500).json({ error: error.message });
  }
});

// Webhook Handler
app.post('/webhook', express.raw({ type: 'application/json' }), (req, res) => {
  try {
    const event = JSON.parse(req.body.toString());

    console.log(`Received event: ${event.type}`);

    switch (event.type) {
      case 'payment.succeeded':
        console.log(`💰 Payment ${event.data?.object?.id} succeeded!`);
        break;
      case 'payment.failed':
        console.log(`❌ Payment failed.`);
        break;
      default:
        console.log(`Unhandled event type ${event.type}`);
    }

    res.json({ received: true });
  } catch (err) {
    console.error(`Webhook Error: ${err.message}`);
    res.status(400).send(`Webhook Error: ${err.message}`);
  }
});

const PORT = 3000;
app.listen(PORT, () => {
  console.log(`Server running on http://localhost:${PORT}`);
});
