#!/usr/bin/env python3
"""
Pytest configuration and fixtures for the tools package.
"""
import pytest
import tempfile
import shutil
import os
from pathlib import Path


@pytest.fixture
def temp_dir():
    """Create a temporary directory that gets cleaned up after the test."""
    temp_dir = tempfile.mkdtemp()
    yield temp_dir
    shutil.rmtree(temp_dir, ignore_errors=True)


@pytest.fixture
def temp_c_file():
    """Create a temporary C file with sample content."""
    content = """
    #include <stdio.h>
    
    int sample_function(int param1, char* param2) {
        printf("param1: %d, param2: %s\\n", param1, param2);
        return param1 * 2;
    }
    
    void another_function() {
        // Empty function
    }
    """
    
    with tempfile.NamedTemporaryFile(mode='w', suffix='.c', delete=False) as f:
        f.write(content)
        temp_file = f.name
    
    yield temp_file
    
    # Cleanup
    if os.path.exists(temp_file):
        os.unlink(temp_file)


@pytest.fixture
def temp_h_file():
    """Create a temporary header file with sample content."""
    content = """
    #ifndef SAMPLE_H
    #define SAMPLE_H
    
    extern int global_variable;
    extern char* global_string;
    
    int declared_function(int param);
    void utility_function(const char* message);
    
    #endif // SAMPLE_H
    """
    
    with tempfile.NamedTemporaryFile(mode='w', suffix='.h', delete=False) as f:
        f.write(content)
        temp_file = f.name
    
    yield temp_file
    
    # Cleanup
    if os.path.exists(temp_file):
        os.unlink(temp_file)


@pytest.fixture
def sample_project_dir():
    """Create a temporary project directory with multiple C files."""
    temp_dir = tempfile.mkdtemp()
    
    # Create main.c
    main_c = os.path.join(temp_dir, "main.c")
    with open(main_c, 'w') as f:
        f.write("""
        #include "utils.h"
        
        int main(int argc, char* argv[]) {
            init_utils();
            return 0;
        }
        """)
    
    # Create utils.h
    utils_h = os.path.join(temp_dir, "utils.h")
    with open(utils_h, 'w') as f:
        f.write("""
        #ifndef UTILS_H
        #define UTILS_H
        
        extern int config_value;
        void init_utils();
        int process_data(int* data, size_t count);
        
        #endif
        """)
    
    # Create utils.c
    utils_c = os.path.join(temp_dir, "utils.c")
    with open(utils_c, 'w') as f:
        f.write("""
        #include "utils.h"
        
        int config_value = 42;
        
        void init_utils() {
            // Initialize utilities
        }
        
        int process_data(int* data, size_t count) {
            return count > 0 ? data[0] : 0;
        }
        """)
    
    # Create subdirectory with more files
    subdir = os.path.join(temp_dir, "modules")
    os.makedirs(subdir)
    
    module_c = os.path.join(subdir, "module.c")
    with open(module_c, 'w') as f:
        f.write("""
        void module_function(void) {
            // Module function
        }
        """)
    
    yield temp_dir
    
    # Cleanup
    shutil.rmtree(temp_dir, ignore_errors=True)


def pytest_configure(config):
    """Configure pytest with custom markers."""
    config.addinivalue_line(
        "markers", "slow: mark test as slow running"
    )
    config.addinivalue_line(
        "markers", "integration: mark test as integration test"
    )


def pytest_collection_modifyitems(config, items):
    """Automatically mark slow tests."""
    for item in items:
        # Mark integration tests that use sample_project_dir as slow
        if "sample_project_dir" in item.fixturenames:
            item.add_marker(pytest.mark.slow)
        
        # Mark tests with "integration" in the name
        if "integration" in item.name.lower():
            item.add_marker(pytest.mark.integration)