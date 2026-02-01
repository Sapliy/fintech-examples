const express = require('express');
const { FintechClient } = require('@sapliy/fintech');

const app = express();
app.use(express.json());

const client = new FintechClient(process.env.SAPLIY_API_KEY);

app.post('/charge', async (req, res) => {
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

app.listen(3000, () => console.log('Server running on port 3000'));
