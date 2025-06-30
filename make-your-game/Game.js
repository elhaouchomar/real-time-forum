import { h, Component } from './mini-framework/framework/core.js';
import { connect, sendMessage } from './socket.js';
import { mapArray } from "./map.js";
import NicknameForm from './components/NicknameForm.js';
import WaitingRoom from './components/WaitingRoom.js';
import Player from './components/Player.js';

export const directions = {
    up: "ArrowUp",
    down: "ArrowDown",
    left: "ArrowLeft",
    right: "ArrowRight",
};

export const keys = {
    ArrowUp: directions.up,
    ArrowDown: directions.down,
    ArrowLeft: directions.left,
    ArrowRight: directions.right
};

export const TILE_SIZE = 32;

class Game extends Component {
  constructor() {
    super();
    this.state = {
      view: 'nickname',
      nickname: '',
      players: [],
      messages: [],
      countdown: null,
    };
    this.handleNicknameSubmit = this.handleNicknameSubmit.bind(this);
    this.handleKeyDown = this.handleKeyDown.bind(this);
    this.handleKeyUp = this.handleKeyUp.bind(this);
    this.pressedDirections = [];
    this.isMoving = false;
  }

  handleNicknameSubmit(nickname) {
    this.setState({ view: 'waiting', nickname });
    connect(nickname, (message) => {
        const data = JSON.parse(message.data);
        switch (data.type) {
            case 'update-players':
                this.setState({ players: data.players });
                break;
            case 'chat-message':
                this.setState({ messages: [...this.state.messages, data] });
                break;
            case 'start-countdown':
            case 'update-countdown':
                this.setState({ countdown: data.countdown });
                break;
            case 'stop-countdown':
                this.setState({ countdown: null });
                break;
            case 'start-game':
                this.setState({ view: 'game', players: data.players });
                break;
            case 'game-state':
                this.setState({ players: data.players });
                break;
        }
    });
  }

  componentDidMount() {
      // Event listeners are added when the view changes to 'game'
  }

  componentDidUpdate(prevProps, prevState) {
      if (this.state.view === 'game' && prevState.view !== 'game') {
          window.addEventListener('keydown', this.handleKeyDown);
          window.addEventListener('keyup', this.handleKeyUp);
          this.startGameLoop();
      } else if (this.state.view !== 'game' && prevState.view === 'game') {
          window.removeEventListener('keydown', this.handleKeyDown);
          window.removeEventListener('keyup', this.handleKeyUp);
          if (this.animationFrameId) {
              window.cancelAnimationFrame(this.animationFrameId);
          }
      }
  }

  handleKeyDown(e) {
    const dir = keys[e.code];
    if (dir && this.pressedDirections.indexOf(dir) === -1) {
        this.pressedDirections.unshift(dir);
        this.tryToMove();
    }
  }

  handleKeyUp(e) {
    const dir = keys[e.code];
    const index = this.pressedDirections.indexOf(dir);
    if (index > -1) {
        this.pressedDirections.splice(index, 1);
    }
  }

  canMove(player, nextGridX, nextGridY) {
    if (nextGridX < 0 || nextGridY < 0 || nextGridX >= mapArray[0].length || nextGridY >= mapArray.length) {
        return false;
    }
    const tile = mapArray[nextGridY][nextGridX];
    return tile !== 1 && tile !== 3; // 1: block, 3: wall
  }

  tryToMove() {
      if (this.isMoving) return;
      const direction = this.pressedDirections[0];
      if (!direction) return;

      let player = this.state.players.find(p => p.nickname === this.state.nickname);
      if (!player) return;

      let currentGridX = Math.round(player.x / TILE_SIZE);
      let currentGridY = Math.round(player.y / TILE_SIZE);
      
      let nextGridX = currentGridX;
      let nextGridY = currentGridY;

      if (direction === directions.right) nextGridX++;
      else if (direction === directions.left) nextGridX--;
      else if (direction === directions.down) nextGridY++;
      else if (direction === directions.up) nextGridY--;

      if (this.canMove(player, nextGridX, nextGridY)) {
          player.nextPixelX = nextGridX * TILE_SIZE;
          player.nextPixelY = nextGridY * TILE_SIZE;
          this.isMoving = true;
      }
  }
  
  movePlayer() {
      if (!this.isMoving) return;
      
      let player = this.state.players.find(p => p.nickname === this.state.nickname);
      if (!player || player.nextPixelX === undefined) {
          this.isMoving = false;
          return;
      }

      const speed = 2;
      const diffX = player.nextPixelX - player.x;
      const diffY = player.nextPixelY - player.y;

      if (Math.abs(diffX) < speed && Math.abs(diffY) < speed) {
          player.x = player.nextPixelX;
          player.y = player.nextPixelY;
          this.isMoving = false;
          this.tryToMove(); // Check if another key is pressed
      } else {
        if (player.x < player.nextPixelX) player.x += speed;
        else if (player.x > player.nextPixelX) player.x -= speed;
        if (player.y < player.nextPixelY) player.y += speed;
        else if (player.y > player.nextPixelY) player.y -= speed;
      }
  }

  startGameLoop() {
    const gameLoop = () => {
        this.movePlayer();
        const player = this.state.players.find(p => p.nickname === this.state.nickname);
        if (player) {
            sendMessage({
                type: 'player-move',
                nickname: this.state.nickname,
                x: player.x,
                y: player.y,
            });
        }
        
        // Re-render is handled by setState in message handler
        this.animationFrameId = window.requestAnimationFrame(gameLoop);
    };
    gameLoop();
  }

  render() {
    const { view, nickname, players, messages, countdown } = this.state;

    if (view === 'nickname') {
        return h(NicknameForm, { onSubmit: this.handleNicknameSubmit });
    }
    if (view === 'waiting') {
        return h(WaitingRoom, {
            players: players,
            messages: messages,
            nickname: nickname,
            countdown: countdown
        });
    }
    if (view === 'game') {
        return h('div', { id: 'game' }, [
            h('div', { id: 'map', style: { position: 'relative' } }, 
              players.map(player => h(Player, { player: player }))
              // I will add the map rendering here later
            )
        ]);
    }
    return h('div', {}, 'Something went wrong.');
  }
}

export default Game; 