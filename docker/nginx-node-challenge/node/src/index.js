const express = require("express");
const mysql = require("mysql");
const { faker } = require("@faker-js/faker");

const port = 3000;

const config = {
  host: "db",
  user: "root",
  password: "root",
  database: "nodedb",
};

const app = express();

let connection;

app.get("/", async (req, res) => {
  try {
    connection = mysql.createConnection(config);

    await insertPerson();

    const people = await getPeople();

    let peopleHtmlList = "";
    if (people && people.length) {
      peopleHtmlList = people.map((person) => `<li>${person.name}</li>`);
    }

    await disconnectDb(connection);

    let result = `
            <h1>Full Cycle Rocks!</h1>
            <ul>
                ${peopleHtmlList}
            </ul>
        `;

    res.send(result);
  } catch (error) {
    res.status(500).send("Deu ruim: " + error);
  }
});

async function disconnectDb() {
  return new Promise((resolve) => {
    connection.end(resolve);
  });
}

async function insertPerson() {
  return new Promise((resolve) => {
    const insertSql = `INSERT INTO people(name) values('${faker.name.findName()}')`;
    connection.query(insertSql, resolve);
  });
}

async function getPeople() {
  return new Promise((resolve) => {
    const selectSql = "SELECT * from people";
    connection.query(selectSql, (_, people) => resolve(people));
  });
}

app.listen(port, () => {
  console.log(`Rodando na porta ${port}`);
});
