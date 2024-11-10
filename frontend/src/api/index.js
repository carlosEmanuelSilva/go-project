var socket = new WebSocket("ws://localhost:8080/ws");

let connect = cb => {
  console.log("Tentando conectar");

  socket.onopen = () => {
    console.log("Sucesso na conexão");
  };

  socket.onmessage = msg => {
    console.log(msg);
	cb(msg)
  };

  socket.onclose = event => {
    console.log("Conexão com o socket encerrada: ", event);
  };

  socket.onerror =  error => {
    console.log("Erro no socket: ", error);
  }; 
};

let sendMsg = msg => {
  console.log("Enviando mensagem");
  socket.send(msg);
};

export { connect, sendMsg};
