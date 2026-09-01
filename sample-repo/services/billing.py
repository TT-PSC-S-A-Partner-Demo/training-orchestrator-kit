"""billing service - a second take on the same 'normalize' step.

Convention drift: this author chose to swallow blanks silently.
"""


def normalize(x):
    return (x or "").strip()   # never raises; returns "" for None or blank
