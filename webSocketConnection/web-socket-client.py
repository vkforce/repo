import asyncio
import websockets

async def run():
    uri = "ws://localhost:8080/ws"
    async with websockets.connect(uri) as websocket:
        name = input("What's your name? ")
        message = f"Hello {name}!"

        await websocket.send(message)
        print(f"You sent: {message}")

        greeting = await websocket.recv()
        print(f"The server said: {greeting}")

asyncio.run(run())
