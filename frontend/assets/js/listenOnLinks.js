// Create a WeakMap to store event listeners
const listenersMap = new WeakMap();
// Function to toggle an event listener on an element
function toggleListener(element, event, handler) {
    const listener = listenersMap.get(element);
    if (listener) {
        element.removeEventListener(listener.event, listener.handler);
    }
    element.addEventListener(event, handler);
    listenersMap.set(element, { event, handler });
}

// Example usage
const button = document.querySelector('button');

function handleClick() {
    console.log('Button clicked!');
}

// Toggle event listener on the button
toggleListener(button, 'click', handleClick);
