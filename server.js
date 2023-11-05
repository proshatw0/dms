const express = require('express');
const app = express();
const net = require('net');

app.get('/:value', (req, res) => {
  const value = req.params.value;
  if (value === "favicon.ico") {
    res.status(204).end();
    return; 
  }

  const client = new net.Socket();
  client.connect(6379, '10.241.125.222', () => {
    const request = `links.data\nget\nlocalhost/${value}\n\n\n\n"`;
    client.write(request);
  });

  let responseData = ''; 

  client.on('data', (data) => {
    responseData += data.toString();
    responseData = responseData.slice(4);
  });

  client.on('end', () => {
    res.redirect(responseData);

    client.end();
  });

  client.on('error', (error) => {
    console.error(`Ошибка при подключении к серверу: ${error}`);
    res.status(500).send('Internal Server Error');
  });
});

app.listen(80, () => {
  console.log('Сервер запущен на порту 80');
});