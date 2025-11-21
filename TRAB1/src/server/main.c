#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <arpa/inet.h>

#define PORT 8080
#define BUFFER_SIZE 1024

int main() {
    int server_fd, client_fd;
    struct sockaddr_in server_addr, client_addr;
    char buffer[BUFFER_SIZE];
    socklen_t addr_len = sizeof(client_addr);

    server_fd = socket(AF_INET, SOCK_STREAM, 0);
    if (server_fd < 0) {
        perror("socket");
        exit(1);
    }

    server_addr.sin_family = AF_INET;
    server_addr.sin_addr.s_addr = INADDR_ANY;
    server_addr.sin_port = htons(PORT);

    if (bind(server_fd, (struct sockaddr*)&server_addr, sizeof(server_addr)) < 0) {
        perror("bind");
        exit(1);
    }

    if (listen(server_fd, 1) < 0) {
        perror("listen");
        exit(1);
    }

    printf("Servidor aguardando conexão na porta %d...\n", PORT);

    client_fd = accept(server_fd, (struct sockaddr*)&client_addr, &addr_len);
    if (client_fd < 0) {
        perror("accept");
        exit(1);
    }

    printf("Cliente conectado!\n");

    while (1) {
        memset(buffer, 0, BUFFER_SIZE);

        int bytes = recv(client_fd, buffer, BUFFER_SIZE, 0);
        if (bytes <= 0) {
            printf("Cliente desconectado.\n");
            break;
        }

        printf("Cliente disse: %s\n", buffer);

        send(client_fd, "Mensagem recebida!", 19, 0);
    }

    close(client_fd);
    close(server_fd);

    return 0;
}
