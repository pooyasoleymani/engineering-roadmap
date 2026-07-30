add_executable(test_search ${CMAKE_SOURCE_DIR}/cpp_tests/binary_search.cpp)

target_link_libraries(test_search PUBLIC binary_searching Catch2::Catch2WithMain)

add_test(NAME binary_search_test COMMAND test_search)