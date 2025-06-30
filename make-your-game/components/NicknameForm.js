import { h, Component } from '../mini-framework/framework/core.js';

class NicknameForm extends Component {
    constructor(props) {
        super(props);
        this.state = {
            nickname: ''
        };
        this.handleSubmit = this.handleSubmit.bind(this);
        this.handleInput = this.handleInput.bind(this);
    }

    handleSubmit(e) {
        e.preventDefault();
        if (this.state.nickname.trim()) {
            this.props.onSubmit(this.state.nickname);
        }
    }

    handleInput(e) {
        this.setState({ nickname: e.target.value });
    }

    render() {
        return h('div', { class: 'nickname-container' }, [
            h('h1', {}, 'Enter Your Nickname'),
            h('form', { onsubmit: this.handleSubmit }, [
                h('input', {
                    type: 'text',
                    placeholder: 'Nickname',
                    value: this.state.nickname,
                    oninput: this.handleInput,
                    autofocus: true,
                }),
                h('button', { type: 'submit' }, 'Join Game')
            ])
        ]);
    }
}

export default NicknameForm; 