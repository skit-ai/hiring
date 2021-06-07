import pytest
from testing_and_reporting import complex_arithmetic


@pytest.mark.parametrize("a, b, c", [
    (1, 1, 2),
    (1, 2, 3),
    (1, -1, 0)
])
def test_complex_arithmetic(a, b, c):
    assert complex_arithmetic(a, b) == c
