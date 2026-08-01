add_executable(sort_test ${CMAKE_SOURCE_DIR}/cpp_tests/sorting_test.cpp)

target_link_libraries(sort_test PUBLIC sorting Catch2::Catch2WithMain)

add_test(NAME sort_test COMMAND sort_test)