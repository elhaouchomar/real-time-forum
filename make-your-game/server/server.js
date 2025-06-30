import { WebSocketServer } from 'ws';

const wss = new WebSocketServer({ port: 8080 });

console.log('WebSocket server started on port 8080');

let players = [];
let gameState = 'waiting'; // waiting, countdown, in-game
let waitingTimeout = null;
let countdownInterval = null;
let countdown = 10;

const broadcast = (message) => {
    const data = JSON.stringify(message);
    wss.clients.forEach(client => {
        if (client.readyState === client.OPEN) {
            client.send(data);
        }
    });
};

const broadcastPlayers = () => {
    broadcast({ type: 'update-players', players });
};

const startGame = () => {
    console.log('Starting game...');
    gameState = 'in-game';

    // Assign starting positions
    const startPositions = [
        { x: 32, y: 32 },   // Top-left
        { x: 608, y: 32 },  // Top-right
        { x: 32, y: 416 },  // Bottom-left
        { x: 608, y: 416 }  // Bottom-right
    ];

    players.forEach((player, index) => {
        if (startPositions[index]) {
            player.x = startPositions[index].x;
            player.y = startPositions[index].y;
            player.lives = 3;
        }
    });

    broadcast({ type: 'start-game', players });
    if (waitingTimeout) clearTimeout(waitingTimeout);
    if (countdownInterval) clearInterval(countdownInterval);
};

const startCountdown = () => {
    if (gameState === 'countdown') return;

    if (waitingTimeout) {
        clearTimeout(waitingTimeout);
        waitingTimeout = null;
    }
    
    console.log('Starting 10 second countdown...');
    gameState = 'countdown';
    countdown = 10;
    
    broadcast({ type: 'start-countdown', countdown });

    countdownInterval = setInterval(() => {
        countdown--;
        broadcast({ type: 'update-countdown', countdown });
        if (countdown <= 0) {
            clearInterval(countdownInterval);
            countdownInterval = null;
            startGame();
        }
    }, 1000);
};

const checkGameState = () => {
    if (players.length >= 4) {
        startCountdown();
    } else if (players.length >= 2) {
        if (!waitingTimeout && gameState === 'waiting') {
            console.log('Two players joined. Starting 20 second timer...');
            waitingTimeout = setTimeout(startCountdown, 20000);
        }
    } else { // < 2 players
        if (waitingTimeout) {
            console.log('Not enough players, stopping timer.');
            clearTimeout(waitingTimeout);
            waitingTimeout = null;
        }
        if (countdownInterval) {
            console.log('Not enough players, stopping countdown.');
            clearInterval(countdownInterval);
            countdownInterval = null;
            gameState = 'waiting';
            broadcast({ type: 'stop-countdown' });
        }
    }
};

wss.on('connection', function connection(ws) {
  console.log('A new client connected');
  
  ws.on('message', function message(data) {
    const messageData = JSON.parse(data);
    
    if (messageData.type === 'join') {
      const newPlayer = {
        nickname: messageData.nickname,
        id: ws.id
      };
      players.push(newPlayer);
      ws.id = newPlayer.id;
      broadcastPlayers();
      checkGameState();
    } else if (messageData.type === 'chat') {
        broadcast({
            type: 'chat-message',
            nickname: messageData.nickname,
            message: messageData.message
        });
    } else if (messageData.type === 'player-move') {
        const player = players.find(p => p.nickname === messageData.nickname);
        if (player) {
            player.x = messageData.x;
            player.y = messageData.y;
        }
        broadcast({ type: 'game-state', players: players });
    }
  });

  ws.on('close', () => {
    console.log('Client disconnected');
    players = players.filter(player => player.id !== ws.id);
    broadcastPlayers();
    checkGameState();
  });

  ws.send(JSON.stringify({ message: 'Welcome to the Bomberman server!'}));
}); 