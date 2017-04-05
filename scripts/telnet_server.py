from miniboa import TelnetServer

CLIENTS = []

def on_connect(client):
    client.send("login:")
    CLIENTS.append(client)

def on_disconnect(client):
    CLIENTS.remove(client)

server = TelnetServer(port = 3333, on_connect = on_connect, on_disconnect = on_disconnect)

while True:
    server.poll()
    for client in CLIENTS:
            if client.active and client.cmd_ready:
                    client.get_command()
                    client.send("password:")
