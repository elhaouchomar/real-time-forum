import { h, Component } from '../mini-framework/framework/core.js';
import { sendMessage } from '../socket.js';

class WaitingRoom extends Component {
  constructor(props) {
    super(props);
    this.state = {
      message: ''
    };
    this.handleInput = this.handleInput.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }

  handleInput(e) {
    this.setState({ message: e.target.value });
  }

  handleSubmit(e) {
    e.preventDefault();
    if (this.state.message.trim()) {
      sendMessage({
        type: 'chat',
        nickname: this.props.nickname,
        message: this.state.message
      });
      this.setState({ message: '' });
    }
  }

  render() {
    const { players, nickname, messages, countdown } = this.props;
    return h('div', { class: 'waiting-room' }, [
      h('h1', {}, 'Waiting for Players...'),
      countdown !== null && h('h2', { class: 'countdown' }, `Game starts in ${countdown}...`),
      h('div', { class: 'players-list' }, [
        h('h2', {}, 'Players:'),
        h('ul', {}, players.map(player => h('li', {}, player.nickname)))
      ]),
      h('div', { class: 'chat' }, [
        h('h2', {}, 'Chat'),
        h('div', { class: 'messages' }, 
          messages.map(msg => h('div', {}, `${msg.nickname}: ${msg.message}`))
        ),
        h('form', { onsubmit: this.handleSubmit }, [
          h('input', {
            type: 'text',
            placeholder: 'Type a message...',
            value: this.state.message,
            oninput: this.handleInput
          }),
          h('button', { type: 'submit' }, 'Send')
        ])
      ])
    ]);
  }
}

export default WaitingRoom; 