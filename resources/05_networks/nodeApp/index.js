const express = require("express");
const dotenv = require("dotenv");
const axios = require("axios");


dotenv.config({ path: process.env.ENV_PATH || '.env' });

const app = express();
app.get("/ping", (req, res) => res.send("Pong!"));

app.get("/isConnected", async (req, res) =>{
  const hostname = process.env.HOSTNAME;
  const port = process.env.GO_PORT;
  const url = `http://${hostname}:${port}/isAlive`;
  console.log('sending request to see if it is alive to:', url);
  try {
    await axios.get(url);
    res.status(200).send('yep');
  } catch (e) {
    console.log('error happened:', e.message);
    res.status(400).send('nop')
  }
});

const server = app.listen(process.env.PORT, () => {
  console.log(`App running on port ${process.env.PORT}`);
});

process.on('SIGTERM', () => {server.close()});
