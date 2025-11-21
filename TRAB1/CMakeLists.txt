cmake_minimum_required(VERSION 3.10)
project(TRAB1)

set(CMAKE_C_STANDARD 11)

include_directories(include)

set(SHARED_SOURCES
    
)

add_executable(servidor
    src/server/main.c
    ${SHARED_SOURCES}
)

add_executable(cliente
    src/client/main.c
    ${SHARED_SOURCES}
)
