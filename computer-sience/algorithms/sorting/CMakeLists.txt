cmake_minimum_required(VERSION 3.20)


project(sorting
    VERSION 0.0.1
    LANGUAGES CXX
)

option(BUILD_TESTS ON "build unittests")

set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)
set(CMAKE_POSITION_INDEPENDENT_CODE ON)


if(MSVC)
    add_compile_options(/W4)
else()
    add_compile_options(-Wall -Wextra -Wpedantic)
endif()


add_library(sorting SHARED ${CMAKE_CURRENT_SOURCE_DIR}/sorting.cpp)

if (BUILD_TESTS)

    include(CTest)
    find_package(Catch2)
    enable_testing()
    add_subdirectory(cpp_tests)

endif()