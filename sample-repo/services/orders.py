"""orders service - part of the consistent baseline the team already ships.

The convention (blank input is rejected, valid input trimmed) is followed here, and in
billing and shipping. It was never written down anywhere - it only lives in the code.
"""


def normalize(x):
    if x is None or not x.strip():
        raise ValueError("input must not be blank")
    return x.strip()
