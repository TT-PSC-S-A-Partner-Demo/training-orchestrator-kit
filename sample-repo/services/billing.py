"""billing service - same convention as orders, written in a different style.

Consistent behaviour, different code: that is fine. The convention is the behaviour,
not the exact lines.
"""


def normalize(x):
    text = (x or "").strip()
    if not text:
        raise ValueError("input must not be blank")
    return text
