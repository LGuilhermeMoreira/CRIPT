# Projeto Cliente-Servidor

Este projeto contém o código-fonte de uma aplicação cliente-servidor, que utiliza **CMake** e **Make** para a compilação.

## Pré-requisitos

Para compilar e executar o projeto, você deve ter os seguintes utilitários instalados:

* **CMake**
* **Make**
* **Um compilador C/C++** (como GCC ou Clang)
* **Shell script compatível com Bash** (para rodar o `run.sh`)

## ⚙️ Como Compilar e Executar

A maneira mais simples de compilar, iniciar o servidor e rodar o cliente é utilizando o *script* de automação fornecido.

1.  **Execute o script de inicialização:**

    ```bash
    ./run.sh
    ```

    O *script* irá:
    * Criar uma pasta de *build* (`build/`).
    * Configurar a compilação com **CMake**.
    * Compilar os executáveis (**servidor** e **cliente**) com **Make**.
    * Iniciar o `./servidor` em segundo plano.
    * Executar o `./cliente`.
    * Ao finalizar o cliente, o servidor será encerrado automaticamente.

### 🔨 Etapas Manuais (Opcional)

Se preferir compilar manualmente sem rodar o cliente em seguida:

```bash
# 1. Cria e entra na pasta de build
mkdir -p build
cd build

# 2. Configura o projeto com CMake
cmake ..

# 3. Compila os executáveis
make