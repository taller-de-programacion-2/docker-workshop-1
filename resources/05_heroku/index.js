const express = require("express");
const { Client } = require("pg");

const app = express();
const client = new Client({
  connectionString: process.env.DATABASE_URL,
  query_timeout: 1000,
  statement_timeout: 1000,
  ssl: {
    rejectUnauthorized: false,
  },
});

client.connect();

app.get("/ping", (req, res) => res.send("Pong!"));

app.get("/status", (req, res) =>
  client.query("SELECT NOW()", (err) =>
    res.send({ service: "UP", db: err ? "DOWN" : "UP" })
  )
);

app.listen(process.env.PORT, () => {
  console.log(`App running on port ${process.env.PORT}`);
});
