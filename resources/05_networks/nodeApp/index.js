const express = require("express");
const dotenv = require("dotenv");


dotenv.config({ path: process.env.ENV_PATH || '.env' });

const app = express();
app.get("/ping", (req, res) => res.send("Pong!"));

app.get("/isConnected", async (req, res) =>{
  const hostname = process.env.HOSTNAME;
  const port = process.env.PORT;
  const url = `${hostname}${port}/isAlive`;
  try {
    await fetch(url);
    res.status(200).send('yep');
  } catch (e) {
    res.status(400).send('nop')
  }
});

const server = app.listen(process.env.PORT, () => {
  console.log(`App running on port ${process.env.PORT}`);
});

process.on('SIGTERM', () => {server.close()});
