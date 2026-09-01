"""shipping service - a third take on the same 'normalize' step.

This one happens to match the convention the team should have agreed on all along:
reject blank input loudly. The other two drifted from it.
"""


def normalize(x):
    if x is None or not x.strip():
        raise ValueError("input must not be blank")
    return x.strip()
