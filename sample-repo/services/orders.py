"""orders service - one of three that each 'normalize' user input its own way.

Convention drift: this author assumed input is always a non-empty string.
"""


def normalize(x):
    return x.strip()          # crashes on None; lets "   " through as ""
