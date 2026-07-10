cmake_minimum_required(VERSION 3.20)

project(Array
    VERSION 0.1.0
    LANGUAGES CXX
)

set(CMAKE_CXX_STANDARD 20)
set(CMAKE_CXX_STANDARD_REQUIRED ON)
set(CMAKE_POSITION_INDEPENDENT_CODE ON)

# Better warnings (cross-platform)
if(MSVC)
    add_compile_options(/W4)
else()
    add_compile_options(-Wall -Wextra -Wpedantic)
endif()

# Enable folders in IDEs
set_property(GLOBAL PROPERTY USE_FOLDERS ON)

add_executable(array)

target_sources(array
    PRIVATE
    main.cpp
)

include(CTest)

