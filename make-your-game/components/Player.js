import { h, Component } from '../mini-framework/framework/core.js';

class Player extends Component {
    render() {
        const { player } = this.props;
        // The player's position will be used to style the div
        const style = {
            position: 'absolute',
            top: `${player.y}px`,
            left: `${player.x}px`,
            width: '32px',
            height: '32px',
            backgroundColor: 'blue', // Placeholder color
        };

        return h('div', { class: 'player', style: style }, player.nickname);
    }
}

export default Player; 