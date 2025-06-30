import { h, render } from './mini-framework/framework/core.js';
import { store } from './mini-framework/framework/store.js';
import Game from './Game.js';

const appContainer = document.getElementById('game');

// Initial render
render(h(Game), appContainer); 