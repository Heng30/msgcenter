const socket = new WebSocket('ws://localhost:8001/?token=testToken&&topic=hello');
socket.addEventListener('open', event => {
    console.log('Connected to server');
    socket.send('Hello, server!');
});
socket.addEventListener('message', event => {
    console.log(`Received message: ${event.data}`);
    socket.close();
});
