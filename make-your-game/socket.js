let socket;

export const connect = (nickname, onMessage) => {
  socket = new WebSocket('ws://localhost:8080');

  socket.onopen = () => {
    console.log('WebSocket connection established');
    const message = {
      type: 'join',
      nickname: nickname
    };
    socket.send(JSON.stringify(message));
  };

  socket.onmessage = onMessage;

  socket.onclose = () => {
    console.log('WebSocket connection closed');
  };

  socket.onerror = (error) => {
    console.error('WebSocket error:', error);
  };
};

export const sendMessage = (message) => {
  if (socket && socket.readyState === WebSocket.OPEN) {
    socket.send(JSON.stringify(message));
  }
}; 