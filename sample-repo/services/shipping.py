"""shipping service - the third of the consistent baseline.

Same contract as orders and billing: reject blank input loudly, trim the rest.
"""


def normalize(x):
    t = (x or "").strip()
    if t == "":
        raise ValueError("input must not be blank")
    return t
