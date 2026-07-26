cmake_minimum_required(VERSION 3.20)

project(search
    VERSION 0.0.1
    LANGUAGES CXX
)

option(BUILD_TEST ON "build unit-testes")

set(CMAKE_CXX_STANDARD 20)
set(CMAKE_CXX_STANDARD_REQUIRED ON)
set(CMAKE_POSITION_INDEPENDENT_CODE ON)


if(MSVC)
    add_compile_options(/W4)
else()
    add_compile_options(-Wall -Wextra -Wpedantic)
endif()

find_package(Catch2)

if(BUILD_TEST)

include(CTest)
enable_testing()
add_subdirectory(${CMAKE_SOURCE_DIR}/cpp_tests)

endif()
