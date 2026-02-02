import express from 'express';
import Sapliyio from '@sapliyio/fintech';

const app = express();
const client = new Sapliyio.FintechClient(process.env.SAPLIY_API_KEY);

// Webhook secret for verification
const WEBHOOK_SECRET = process.env.SAPLIY_WEBHOOK_SECRET;

app.post('/charge', express.json(), async (req, res) => {
  try {
    const payment = await client.payments.create({
      amount: req.body.amount,
      currency: 'USD',
      sourceId: req.body.token,
      description: 'Example charge'
    });
    res.json(payment);
  } catch (error) {
    res.status(500).json({ error: error.message });
  }
});

// Webhook endpoint
app.post('/webhook', express.raw({ type: 'application/json' }), (req, res) => {
  const sig = req.headers['x-sapliy-signature'];

  try {
    const event = client.webhooks.constructEvent(req.body, sig, WEBHOOK_SECRET);

    console.log('Received event:', event.type);

    // Handle the event
    switch (event.type) {
      case 'payment.succeeded':
        const payment = event.data.object;
        console.log(`Payment for ${payment.amount} succeeded!`);
        break;
      default:
        console.log(`Unhandled event type ${event.type}`);
    }

    res.json({ received: true });
  } catch (err) {
    res.status(400).send(`Webhook Error: ${err.message}`);
  }
});

app.listen(3000, () => console.log('Server running on port 3000'));
